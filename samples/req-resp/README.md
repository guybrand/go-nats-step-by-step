# Creating basic requester(s) and responser(s)

# Requirements
Go 1.14 or later version

NATS server running on your local machine, easiest way to run it:
```
$ docker pull nats
$ docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
```

# Step1 (running single request-reply):
Open a terminal windows

`$ cd samples/req-resp`

`$ go run step2/minimal-reply.go`


Open another terminal window

`$ cd samples/req-resp`

`$ go run step1/minimal-req.go`

You should now see the replier gets "Hello Gophers" message and replies: "I hear you" (+msg number) .


# Step2 (running single request vs multi repliers):
You can open more windows with the replier (step2/minimal-reply.go) and see that when you run the requester (step1/minimal-req.go), each time a different replier will reply, note: there that the selected replier is random



# Step3:
Stop all repliers, and instead type:

`$ go run step4/many-reply.go`

This replier only prints "start" and "end" messages to the console


# Step4 (running multiple requests):
On the requester terminal window, type:

`$ go run step3/many-req.go -n 15000`


# Step5 (running multiple requests with concurrency):
On the requester terminal window, type:

`$ go run step5/many-async-req.go  -n 120000`


