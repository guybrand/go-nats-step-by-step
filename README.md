# Step by step NATS inplementation in golang
This repository contains golang source code for all the samples for nats implementation as presented on [Go Israel Meetup at JFrog TLV](https://www.meetup.com/Go-Israel/events/284585914/)

- [samples](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples) folder (follow the [presentation](https://docs.google.com/presentation/d/1DBwhDyXLQ-lUEekAshG9H8bWd6YtZWvkYhkdgN5gSDY) for instructions how to run each step) :
  - [pub-sub](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/pub-sub) - Step by step implemntatiom for NATS (https://nats.io/) pub/sub using golang
  - [chat](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/chat) - Mini console based chat using NATS pub/sub
  - [req-resp](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/req-resp) - Step by step implemntatiom for NATS requset/reply in go
  - [smesh](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/smesh) - Implementation of a mini service mesh NATS requset/reply pattern.

- [raffle](https://github.com/guybrand/go-nats-step-by-step/tree/main/raffle) - A small raffle console tool that was used in the meetup

- [benchmark-grpc-protobuf-vs-http-json](https://github.com/guybrand/go-nats-step-by-step/tree/main/benchmark-grpc-protobuf-vs-http-json) - Benchmark comparison tool between 
  - http/grpc (https://github.com/plutov/benchmark-grpc-protobuf-vs-http-json clone tweaked to allow concurrent http/grpc clients) 
    and a request-reply pattern as coded in [req-resp](https://github.com/guybrand/go-nats-step-by-step/tree/main/samples/req-resp)
    

### Presntation link:
[Go NATS](https://docs.google.com/presentation/d/1DBwhDyXLQ-lUEekAshG9H8bWd6YtZWvkYhkdgN5gSDY) - 
The presentation also includes much of code built step by step, and instructions how to run each step(s).

### Requirements
Go 1.14 or later version

NATS server running on your local machine, easiest way to run it:
```
$ docker pull nats
$ docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
```

### Meetup video link:
Comming soon :)
