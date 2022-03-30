package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

var numberOfMessagesToSend int
var concurrentRequests int

func usage() {
	log.Printf("Usage: [-n ]\n")
	flag.PrintDefaults()
}

func parseFlags() {
	n := flag.Int("n", 1, "the number of requests to send")
	c := flag.Int("c", 1, "the number of concurrent requests")

	log.SetFlags(log.Lmicroseconds)
	flag.Usage = usage
	flag.Parse()
	numberOfMessagesToSend = *n
	concurrentRequests = *c
}

func main() {

	opts := []nats.Option{nats.Name("NATS Sample Requestor")}
	parseFlags()

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatal(err)
	} else {
		defer nc.Close()

		t := time.Now()

		ch := make(chan bool, concurrentRequests)

		for i := 0; i < numberOfMessagesToSend; i++ {
			ch <- true
			go request(nc, fmt.Sprint("Hello Gophers", i), false, ch)
		}
		for len(ch) > 0 {
		}
		d := time.Since(t)
		fmt.Printf("ran %d nats calls on %s time\n", numberOfMessagesToSend, d.String())

	}
}

func request(nc *nats.Conn, msg string, withLog bool, ch chan bool) {
	defer func() {
		<-ch
	}()

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
