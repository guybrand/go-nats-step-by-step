package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

type msg struct {
	User    string
	Message string
	Time    string
}

func main() {

	opts := []nats.Option{nats.Name("Simple chat")}
	opts = setupConnOptions(opts)

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatalf("connect error:#s", err.Error())
	} else {
		defer nc.Close()

		showChatOnConsole(nc)

		nc.Flush()
	}

}
