package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
)

var msgCount int

func main() {

	opts := []nats.Option{nats.Name("NATS Sample Replier")}
	opts = setupConnOptions(opts)
	log.SetFlags(log.Lmicroseconds)

	nc, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	listenAndServe(nc)
}

func handleMsg(nc *nats.Conn) nats.MsgHandler {
	return func(msg *nats.Msg) {
		msgCount += 1
		msg.Respond([]byte("I hear you " + fmt.Sprint(msgCount)))
		m := string(msg.Data)
		if m == "start" || m == "end" {
			nc.Flush()
			log.Printf("[#%d] Received on [%s]", msgCount, m)
		}
	}
}

func listenAndServe(nc *nats.Conn) {

	subj := "go-meetup"

	nc.QueueSubscribe(subj, subj, handleMsg(nc))
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subj)

	// Setup the interrupt handler to drain so we don't miss requests when scaling down.
	drainBeforeExit(nc)
}

func drainBeforeExit(nc *nats.Conn) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatalf("Exiting")
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
