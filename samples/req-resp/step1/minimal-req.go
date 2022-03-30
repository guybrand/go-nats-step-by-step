package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {

	opts := []nats.Option{nats.Name("NATS Sample Requestor")}

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatal(err)
	} else {
		defer nc.Close()
		request(nc, "Hello Gophers!")
	}
}

func request(nc *nats.Conn, msg string) {
	subj := "go-meetup"
	req := []byte(msg)

	if reply, err := nc.Request(subj, req, 2*time.Second); err != nil {
		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		log.Fatalf("%v for request", err)
	} else {
		log.Printf("Published [%s] : '%s'", subj, req)
		log.Printf("Received  [%v] : '%s'", reply.Subject, string(reply.Data))
	}
}
