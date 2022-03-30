package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/http-json"
)

func init() {
	go httpjson.Start()
	fmt.Println("http warming up")
	time.Sleep(time.Second)
}

func BenchmarkHTTPJSON(msgToSend int, concurrentClients int) {
	client := &http.Client{}
	t := time.Now()

	ch := make(chan bool, concurrentClients)
	for n := 0; n < msgToSend; n++ {
		ch <- true
		go doPost(client, ch)
	}

	d := time.Now().Sub(t)
	fmt.Printf("ran %d http calls on %s time\n", msgToSend, d.String())
}

func doPost(client *http.Client, ch chan bool) {
	defer func() {
		<-ch
	}()
	u := &httpjson.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(u)

	resp, err := client.Post("http://127.0.0.1:60001/", "application/json", buf)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target httpjson.Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		log.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.User == nil || target.User.ID != "1000000" {
		log.Fatalf("http response is wrong: %v", resp)
	}
}
