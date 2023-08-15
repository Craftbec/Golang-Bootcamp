package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"jwt/types"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golang-jwt/jwt"
)

type errorPage struct {
	Text string `json:"error"`
}

type Tok struct {
	Text string `json:"token"`
}

type Res struct {
	Count  int
	Places []types.Place
	Prev   int
	Next   int
	Last   int
}

type ResJ struct {
	Name   string        `json:"name"`
	Count  int           `json:"total"`
	Places []types.Place `json:"places"`
	Prev   int           `json:"prev_page"`
	Next   int           `json:"next_page"`
	Last   int           `json:"last_page"`
}

type Recommend struct {
	Name   string        `json:"name"`
	Places []types.Place `json:"places"`
}

type Str struct {
	Id     string          `json:"_id"`
	Shards json.RawMessage `json:"_source"`
}

func CreateEsClient() *ES {
	client, err := elasticsearch.NewClient(elasticsearch.Config{})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	r := ES{
		es: client,
	}
	return &r
}

type ES struct {
	es *elasticsearch.Client
}

type Store interface {
	GetPlaces(limit int, offset int) ([]types.Place, int, error)
	GetPlacesJSON(limit int, offset int) ([]types.Place, int, error)
	GetRecommend(lat, lon float64) []byte
}

func Count(es *elasticsearch.Client) int {
	var total int
	count, err := es.Count(es.Count.WithBody(nil))
	if err != nil {
		log.Fatal(err)
	}
	var rr map[string]interface{}
	if err := json.NewDecoder(count.Body).Decode(&rr); err == nil {
		total = int(rr["count"].(float64))
	} else {
		log.Printf("Error parsing the response body of count documents: %s", err)
	}
	return total
}

func (store *ES) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	var places = []types.Place{}
	var rr map[string]interface{}
	total := Count(store.es)
	if offset < 0 || offset > total {
		return places, 0, errors.New("Incorrect page")
	}
	res, _ := store.es.Search(
		store.es.Search.WithBody(strings.NewReader(fmt.Sprintf(`{
		  "from": %d,
		  "size": %d,
		  "query": {
			"match": {
			  "_index": "places"
			}
		  },
		  "sort": [
			{"_score": "desc"},
			{"id": "asc"}
		  ]

		}`, offset, limit))),
		store.es.Search.WithPretty(),
	)
	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	} else {
		for _, hit := range rr["hits"].(map[string]interface{})["hits"].([]interface{}) {
			source := hit.(map[string]interface{})["_source"]
			v := reflect.ValueOf(source).MapRange()
			place := types.Place{}
			for v.Next() {
				switch key := v.Key().String(); key {
				case "id":
					place.Id, _ = strconv.Atoi((fmt.Sprintf("%s", v.Value())))
				case "name":
					place.Name = fmt.Sprintf("%s", v.Value())
				case "address":
					place.Address = fmt.Sprintf("%s", v.Value())
				case "phone":
					place.Phone = fmt.Sprintf("%s", v.Value())
				case "location":
					loc := v.Value().Interface()
					place.Location.Latitude = loc.(map[string]interface{})["lat"].(float64)
					place.Location.Longitude = loc.(map[string]interface{})["lon"].(float64)
				}
			}
			places = append(places, place)
		}
	}
	return places, total, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	num := 1
	var err error
	var lst int
	var places = []types.Place{}
	st := CreateEsClient()
	tmp, _ := r.URL.Query()["page"]
	if len(tmp) > 0 {
		num, err = strconv.Atoi(tmp[0])
		if err != nil {
			num = -1
		}
	}
	places, lst, err = st.GetPlaces(10, num*10-10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'page' value: '%s'", tmp[0])

	} else {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "index", Res{lst, places, num - 1, num + 1, int(math.Ceil(float64(lst) / 10))})
	}

}

func (store *ES) GetPlacesJSON(limit int, offset int) ([]byte, error) {
	places, total, err := store.GetPlaces(limit, offset)
	if err != nil {
		return nil, errors.New("incorrect page")
	}
	json, err := json.Marshal(ResJ{"Places", total, places, (offset+10)/10 - 1, (offset+10)/10 + 1, int(math.Ceil(float64(total) / 10))})
	if err != nil {
		log.Printf("%v", err)
	}
	return json, nil
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		err  error
		buff bytes.Buffer
	)
	num := 1
	tmp, _ := r.URL.Query()["page"]
	if len(tmp) > 0 {
		num, err = strconv.Atoi(tmp[0])
		if err != nil {
			num = -1
		}
	}
	st := CreateEsClient()
	js, err := st.GetPlacesJSON(10, num*10-10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jso, err := json.Marshal(errorPage{fmt.Sprintf("Invalid 'page' value: '%s'", tmp[0])})
		json.Indent(&buff, jso, "", "    ")
		if err != nil {
			log.Printf("%v", err)
		}
		_, err = w.Write(buff.Bytes())
		if err != nil {
			log.Println(err)
		}
	} else {
		json.Indent(&buff, js, "", "    ")
		_, err = w.Write(buff.Bytes())
		if err != nil {
			log.Println(err)
		}
	}
}

func (store *ES) GetRecommend(lat, lon float64) []byte {
	var (
		places = []types.Place{}
		rr     map[string]interface{}
	)
	res, _ := store.es.Search(
		store.es.Search.WithBody(strings.NewReader(fmt.Sprintf(`{
		  "size": 3,
		  "query": {
			"match": {
			  "_index": "places"
			}
		  },
		  "sort": [
			{
				"_geo_distance": {
				  "location": {
					"lat": %f,
					"lon": %f
				  },
				  "order": "asc",
				  "unit": "km",
				  "mode": "min",
				  "distance_type": "arc",
				  "ignore_unmapped": true
				}
			  }
		  ]
		}`, lat, lon))),
		store.es.Search.WithPretty(),
	)
	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		log.Printf("Error parsing the response body: %s", err)
	} else {
		for _, hit := range rr["hits"].(map[string]interface{})["hits"].([]interface{}) {
			source := hit.(map[string]interface{})["_source"]
			v := reflect.ValueOf(source).MapRange()
			place := types.Place{}
			for v.Next() {
				switch key := v.Key().String(); key {
				case "id":
					place.Id, _ = strconv.Atoi((fmt.Sprintf("%s", v.Value())))
				case "name":
					place.Name = fmt.Sprintf("%s", v.Value())
				case "address":
					place.Address = fmt.Sprintf("%s", v.Value())
				case "phone":
					place.Phone = fmt.Sprintf("%s", v.Value())
				case "location":
					loc := v.Value().Interface()
					place.Location.Latitude = loc.(map[string]interface{})["lat"].(float64)
					place.Location.Longitude = loc.(map[string]interface{})["lon"].(float64)
				}
			}
			places = append(places, place)
		}
	}
	json, err := json.Marshal(Recommend{"Recommendation", places})
	if err != nil {
		log.Printf("%v", err)
	}
	return json
}

func recom(w http.ResponseWriter, r *http.Request) {
	err := validateToken(w, r)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		var (
			lat, lon float64
			err      error
			buff     bytes.Buffer
		)
		slat, _ := r.URL.Query()["lat"]
		if len(slat) > 0 {
			lat, err = strconv.ParseFloat(slat[0], 64)
			if err != nil {
				lat = -10000
			}
		}
		slon, _ := r.URL.Query()["lon"]
		if len(slon) > 0 {
			lon, err = strconv.ParseFloat(slon[0], 64)
			if err != nil {
				lon = -10000
			}
		}
		if lat < 91 && lat > -91 && lon < 181 && lon > -181 {
			st := CreateEsClient()
			jso := st.GetRecommend(lat, lon)
			json.Indent(&buff, jso, "", "    ")
			_, err = w.Write(buff.Bytes())
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			jso, err := json.Marshal(errorPage{"Bad request"})
			json.Indent(&buff, jso, "", "    ")
			if err != nil {
				log.Printf("%v", err)
			}
			_, err = w.Write(buff.Bytes())
			if err != nil {
				log.Println(err)
			}

		}
	}
}

var sampleSecretKey = []byte("My secret key")

func generateJWT() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = "craftbec"
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		log.Printf("%v", err)
		return ""
	}
	return tokenString
}

func token(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buff bytes.Buffer
	str := generateJWT()
	jso, err := json.Marshal(Tok{str})
	json.Indent(&buff, jso, "", "    ")
	if err != nil {
		log.Printf("%v", err)
	}
	_, err = w.Write(buff.Bytes())
	if err != nil {
		log.Println(err)
	}

}

func validateToken(w http.ResponseWriter, r *http.Request) (err error) {
	if r.Header["Authorization"] == nil {
		fmt.Fprintf(w, "Can not find token")
		return errors.New("token error")
	}
	str := r.Header["Authorization"][0][7:]
	token, _ := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return sampleSecretKey, nil
	})
	if token == nil {
		fmt.Fprintf(w, "invalid token")
		return errors.New("token error")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return errors.New("token error")
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		fmt.Fprintf(w, "token expired")
		return errors.New("token error")
	}
	return nil
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/places/", api)
	http.HandleFunc("/api/recommend", recom)
	http.HandleFunc("/api/get_token", token)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
