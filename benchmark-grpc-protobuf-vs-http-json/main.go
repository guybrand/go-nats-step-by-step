package main

import (
	"flag"
	"fmt"
	"log"
)

func usage() {
	fmt.Printf("Usage: [-n ] [-h ] [-g ]\n")
	flag.PrintDefaults()
}

func parseFlags() (int, int, int) {
	n := flag.Int("n", 50000, "the number of requests to send")
	h := flag.Int("h", 10, "the number of concurrent http clients")
	g := flag.Int("g", 20, "the number of concurrent gRpc clients")

	log.SetFlags(log.Lmicroseconds)
	flag.Usage = usage
	flag.Parse()
	return *n, *h, *g
}

func main() {
	n, h, g := parseFlags()
	msgToSend := n
	BenchmarkHTTPJSON(msgToSend, h)
	BenchmarkGRPCProtobuf(msgToSend, g)
}
