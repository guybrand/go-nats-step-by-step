### gRPC+Protobuf or JSON+HTTP?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

 - Go 1.11


### Running 
`go run . -n 50000 -h 10 -g 20`

-n: number of calls
-h: number of concurrent http clients
-g: number og concurrent gRPC clients

Comments:
- high number of http clients (> ~10-20) is inefficient and at some oint will hung up your CPU 
- high number of gRPC client can improve preformance up to a degree (> ~150-300 will actually lower your preformance)

- these benchmarks can be compared to [req-resp](/samples/req-resp) pattern
    running 
    
    `$ go run step4/many-reply.go`

    and 
       
    `$ go run step5/many-async-req.go  -n 50000`

     In 2 seperate terminal windows
       
   or, if you wish to limit nats concurrency (to 20 concurrent clients):

    `$ go run step4/many-reply.go`

     and
       
    `$ go run step6/many-clients-req.go  -n 50000 -c 20`
    
    In 2 seperate terminal windows as well
    
