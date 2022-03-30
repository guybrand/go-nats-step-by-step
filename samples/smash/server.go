package main

import (
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"net/http"
)

var msgCount int

func createHandlerWithNats(nc *nats.Conn, subj string) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		msgCount += 1

		if bt, err := ioutil.ReadAll(req.Body); err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		} else {
			log.Printf("Sending a request [%s]", subj)
			if reply, err := request(nc, subj, bt); err != nil {
				log.Printf("Service reply error [%s] : '%s'", subj, err.Error())
				http.Error(w, "cant process request", http.StatusInternalServerError)
			} else {
				log.Printf("Replied [%s] '%s'", subj, string(reply))
				w.Write(reply)
			}
		}
	}
}

func main() {

	nc := connectNats()

	http.HandleFunc("/rate", createHandlerWithNats(nc, "rate"))

	http.HandleFunc("/purchase", createHandlerWithNats(nc, "purchase"))
	http.HandleFunc("/sell", createHandlerWithNats(nc, "sell"))

	log.Printf("serving on :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Printf("cant serve on :8090 - %s", err.Error())
	}
}
