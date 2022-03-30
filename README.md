# Go nats step by step
This repository contains golang source code for all the samples for nats implementation as presented on [Go Israel Meetup at JFrog TLV](https://www.meetup.com/Go-Israel/events/284585914/)

- [samples](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples) folder (follow the [presentation](https://docs.google.com/presentation/d/1DBwhDyXLQ-lUEekAshG9H8bWd6YtZWvkYhkdgN5gSDY) for instructions how to run each step) :
  - [pub-sub](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/pub-sub) - Step by step implemntatiom for NATS (https://nats.io/) pub/sub using golang
  - [chat](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/chat) - Mini console based chat using NATS pub/sub
  - [req-resp](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/req-resp) - Step by step implemntatiom for NATS requset/reply in go
  - [smash](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/smash) - Implementation of a mini service mash NATS requset/reply pattern.

- [raffle](https://github.com/guybrand/go-nats-step-by-step/tree/main/raffle) - A small raffle console tool that was used in the meetup

- [benchmark-grpc-protobuf-vs-http-json](https://github.com/guybrand/go-nats-step-by-step/tree/main/benchmark-grpc-protobuf-vs-http-json) - Benchmark comparison tool between 
  - http/grpc (https://github.com/plutov/benchmark-grpc-protobuf-vs-http-json clone tweaked to allow concurrent http/grpc clients) 
    running go run . -n 50000 -h 10 -g 20
  - this can be compared to [req-resp](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/req-resp) pattern
    running 
    
    `$ go run step4/many-reply.go`

    and 
       
    `$ go run step5/many-async-req.go  -n 50000`


       In 2 seperate terminal windows
       
  - or, if you wish to limit nats concurrency (to 20 concurrent clients):

    `$ go run step4/many-reply.go`

       and
       
       `$ go run step6/many-clients-req.go  -n 50000 -c 20`
    


# Presntation link:
[Go NATS](https://docs.google.com/presentation/d/1DBwhDyXLQ-lUEekAshG9H8bWd6YtZWvkYhkdgN5gSDY) - 
The presentation also includes much of code built step by step, and instructions how to run each step(s).

# Requirements
Go 1.14

# Meetup video link:
Comming soon :)
