package main

import (
	"fmt"
	"log"
	"time"

	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf"
	"github.com/plutov/benchmark-grpc-protobuf-vs-http-json/grpc-protobuf/proto"
	"golang.org/x/net/context"
	g "google.golang.org/grpc"
)

func init() {
	go grpcprotobuf.Start()
	fmt.Println("grpc warming up")
	time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(msgToSend int, concurrentClients int) {
	conn, err := g.Dial("127.0.0.1:60000", g.WithInsecure())
	if err != nil {
		log.Fatalf("grpc connection failed: %v", err)
	}

	client := proto.NewAPIClient(conn)
	t := time.Now()

	ch := make(chan bool, concurrentClients)
	for n := 0; n < msgToSend; n++ {
		ch <- true
		go doGRPC(client, ch)
	}

	d := time.Now().Sub(t)
	fmt.Printf("ran %d grpc calls on %s time\n", msgToSend, d.String())
}

func doGRPC(client proto.APIClient, ch chan bool) {
	defer func() {
		<-ch
	}()
	resp, err := client.CreateUser(context.Background(), &proto.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
	})

	if err != nil {
		log.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		log.Fatalf("grpc response is wrong: %v", resp)
	}
}
