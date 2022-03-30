package main

import (
	"flag"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

var numberOfMessagesToSend int

func usage() {
	log.Printf("Usage: [-n ]\n")
	flag.PrintDefaults()
}

func parseFlags() {
	n := flag.Int("n", 1, "the number of requests to send")

	log.SetFlags(log.Lmicroseconds)
	flag.Usage = usage
	flag.Parse()
	numberOfMessagesToSend = *n
}

func main() {

	opts := []nats.Option{nats.Name("NATS Sample Requestor")}
	parseFlags()

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatal(err)
	} else {
		defer nc.Close()

		request(nc, "start", true)
		for i := 0; i < numberOfMessagesToSend; i++ {
			request(nc, "Hello Gophers", false)
		}
		request(nc, "end", true)
	}
}

func request(nc *nats.Conn, msg string, withLog bool) {
	subj := "go-meetup"
	req := []byte(msg)

	if _, err := nc.Request(subj, req, 2*time.Second); err != nil {
		if nc.LastError() != nil {
			log.Fatalf("%v for request", nc.LastError())
		}
		log.Fatalf("%v for request", err)
	} else if withLog {
		log.Printf("[#%d] %s", numberOfMessagesToSend, req)
	}
}
