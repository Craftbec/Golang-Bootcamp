package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Money      int
	CandyType  string
	CandyCount int
}

type Str_Created struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

func main() {
	var struc Data
	k := flag.String("k", "", "Candy type")
	m := flag.Int("m", 0, "Count of candy to buy")
	c := flag.Int("c", 0, "Amount of money")
	flag.Parse()
	if !isFlagPassed("k") || !isFlagPassed("m") || !isFlagPassed("c") {
		log.Fatal("Wrong arguments")
	}

	struc.Money = *m
	struc.CandyCount = *c
	struc.CandyType = *k

	json_str, err := json.Marshal(struc)
	if err != nil {
		log.Fatalln("Error encoding data: ", err)
	}

	rootCA, err := ioutil.ReadFile("../cert/minica.pem")
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool := x509.NewCertPool()
	rootCAPool.AppendCertsFromPEM(rootCA)
	cc, err := tls.LoadX509KeyPair("../cert/client/cert.pem", "../cert/client/key.pem")
	if err != nil {
		log.Fatalf("Error loading key pair: %v\n", err)
	}
	client := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAPool,
				GetClientCertificate: func(info *tls.CertificateRequestInfo) (*tls.Certificate, error) {
					return &cc, nil
				},
			},
		}}

	req, err := http.NewRequest("POST", "https://localhost:3333/buy_candy", bytes.NewBuffer(json_str))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 201 {
		screat := Str_Created{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(body, &screat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("%s Your change is %d", screat.Thanks, screat.Change))
	} else {
		io.Copy(os.Stdout, resp.Body)
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
