package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func subscribe(nc *nats.Conn, subj string, printMsg func([]byte, string)) (*nats.Subscription, error) {
	//this is not concurrency safe nor checks if subscription with similar name already exist !
	if sub, err := nc.Subscribe(subj, func(msg *nats.Msg) {
		printMsg(msg.Data, msg.Subject)
	}); err != nil {
		return nil, err
	} else {
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatalf("subscribe error:%s", err.Error())
		}

		fmt.Printf("Welcome to [%s] group\n", subj)

		return sub, nil
	}
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
