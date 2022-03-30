package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func publish(nc *nats.Conn, subj, msg string) {

	nc.Publish(subj, []byte(msg))

	if err := nc.LastError(); err != nil {
		log.Printf("error publishing %s on %s group : %s", msg, subj, err.Error())
	} else {
		//log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
