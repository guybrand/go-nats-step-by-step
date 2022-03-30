package main

import (
	"flag"
	"github.com/nats-io/nats.go"
	"log"
)

var numberOfMessagesToSend int

func usage() {
	log.Printf("Usage: [-n ]\n")
	flag.PrintDefaults()
}

func parseFlags() {
	n := flag.Int("n", 1, "the number of messages to publish")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	numberOfMessagesToSend = *n
}

func main() {

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}
	parseFlags()
	log.SetFlags(log.Lmicroseconds)

	// Connect to NATS
	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatal(err)
	} else {
		defer nc.Close()
		publish(nc, "start", true)
		for i := 0; i < numberOfMessagesToSend; i++ {
			publish(nc, "Hello Gophers", false)
		}
		publish(nc, "end", true)
		nc.Flush()
	}
}

func publish(nc *nats.Conn, msg string, withLog bool) {

	subj := "go-meetup"
	nc.Publish(subj, []byte(msg))

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else if withLog {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
