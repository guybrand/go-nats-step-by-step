# Use what we learned about pub/sub to write a small console chat

# Requirements
Go 1.14 or later version

NATS server running on your local machine, easiest way to run it:
```
$ docker pull nats
$ docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
```

Open multiple terminal windows  

`$ cd samples/chat`
`$ go run .`
