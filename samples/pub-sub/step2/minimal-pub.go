package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {

	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatal(err)
	} else {
		defer nc.Close()
		publish(nc)
		nc.Flush()
	}
}
func publish(nc *nats.Conn) {

	subj := "go-meetup"
	msg := "Hello Gophers"

	nc.Publish(subj, []byte(msg))

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
