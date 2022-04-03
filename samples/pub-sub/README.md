# Creating basic publisher(s) and subscriber(s)

# Requirements
Go 1.14 or later version

NATS server running on your local machine, easiest way to run it:
```
$ docker pull nats
$ docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
```

# Step1:
Open a terminal window

`$ cd samples/pub-sub`

`$ go run step1/minimal-sub.go`

You may open run many subscribers on multiple terminal windows to see they will all receive same messages

# Step2:
Open another terminal window

`$ cd samples/pub-sub`

`$ go run step2/minimal-pub.go`

You should now see each subscriber has recieved "Hello Gophers" message.

# Step3:
Use the terminal window from step 2, but instead type:

`$ go run step3/many-pub.go -n 10`

or

`$ go run step3/many-pub.go -n 1000`

You will now see each subscriber has received many (10/1000) messages


# Step4:
Stop all subscribers (step1 windows) and replace then with a subscriber that only print first and last message to the console:

`$ go run step4/print-start-and-end-sub.go`


# Step5:
On the publisher terminal window type:

`$ go run step5/many-with-start-and-end-pub.go -n 1000000`

or

`$ go run step5/many-with-start-and-end-pub.go -n 10000000`

And see how much time does it take for 1M or 10M messages to be sent
