package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func connectNats() *nats.Conn {
	log.Printf("connecting to nats")
	opts := []nats.Option{nats.Name("NATS Sample Requestor")}

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatal(err)
		return nil
	} else {
		log.Printf("connected to nats")
		return nc
	}
}

func request(nc *nats.Conn, subj string, req []byte) ([]byte, error) {

	if reply, err := nc.Request(subj, req, 2*time.Second); err != nil {
		if e := nc.LastError(); e != nil {
			return nil, e
		}
		return nil, err
	} else {
		return reply.Data, nil
	}
}
