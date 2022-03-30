package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

var msgCount int

func main() {

	opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
	opts = setupConnOptions(opts)

	if nc, err := nats.Connect(nats.DefaultURL, opts...); err != nil {
		log.Fatalf("connect error:#s", err.Error())
	} else {
		subscribe(nc)
	}
}

func subscribe(nc *nats.Conn) {
	subj := "go-meetup"
	nc.Subscribe(subj, handleMsg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatalf("subscribe error:%s", err.Error())
	}

	log.Printf("Listening on [%s]", subj)

	runtime.Goexit()
}

func handleMsg(msg *nats.Msg) {
	msgCount += 1
	log.Printf("[#%d] Received on [%s]: '%s'", msgCount, msg.Subject, string(msg.Data))
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to error %s: will attempt reconnects for %.0fm", err.Error(), totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
