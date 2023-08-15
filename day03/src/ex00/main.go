package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type Properties struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone,omitempty"`
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	} `json:"location"`
}

const mapping = `{
	"mappings": {
		"properties": {
			"id": {
				"type":"long"
			},
			"name": {
				"type":"text"
			},
			"phone": {
				"type":"text"
			},
			"address": {
				"type":"text"
			},
			"location": {
				"type":"geo_point"
			}
		}
	}
}`

func createRestaurants(data [][]string) []Properties {
	var restaurants []Properties
	for i, line := range data {
		var tmp Properties
		if i == 0 {
			continue
		}
		tmp.Name = line[1]
		tmp.Address = line[2]
		tmp.Phone = line[3]
		id, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		tmp.Id = id + 1
		tmp.Location.Longitude, err = strconv.ParseFloat(line[4], 64)
		if err != nil {
			log.Fatal(err)
		}
		tmp.Location.Latitude, err = strconv.ParseFloat(line[5], 64)
		if err != nil {
			log.Fatal(err)
		}
		restaurants = append(restaurants, tmp)
	}
	return restaurants
}

func CreateEsClient() *elasticsearch.Client {
	retryBackoff := backoff.NewExponentialBackOff()
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		MaxRetries: 5,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return es
}

func CreateBulkIndexer(es *elasticsearch.Client) esutil.BulkIndexer {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         "places",
		Client:        es,
		NumWorkers:    runtime.NumCPU(),
		FlushBytes:    5e+6,
		FlushInterval: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}
	return bi
}

func Indexator(restaurants []Properties, bi esutil.BulkIndexer, es *elasticsearch.Client) {
	var (
		countSuccessful uint64
		res             *esapi.Response
		err             error
	)
	if res, err = es.Indices.Delete([]string{"places"}, es.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || res.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}
	res.Body.Close()
	res, err = es.Indices.Create("places", es.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Cannot create index: %s", res)
	}
	res.Body.Close()

	for _, a := range restaurants {
		data, err := json.Marshal(a)
		if err != nil {
			log.Fatal(err)
		}
		err = bi.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(a.Id),
				Body:       bytes.NewReader(data),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Printf("ERROR: %s", err)
					} else {
						log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
					}
				},
			},
		)
		if err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}
	}
	if err := bi.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}
}

func main() {
	f, err := os.Open("../../materials/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	csvReader.Comma = '\t'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	restaurants := createRestaurants(data)
	es := CreateEsClient()
	bi := CreateBulkIndexer(es)
	Indexator(restaurants, bi, es)
}
