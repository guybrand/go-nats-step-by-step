package main

import (
	"encoding/json"
	"errors"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func request(nc *nats.Conn, subj string, req []byte) ([]byte, error) {

	if reply, err := nc.Request(subj, req, 2*time.Second); err != nil {
		if e := nc.LastError(); e != nil {
			return nil, e
		}
		return nil, err
	} else {
		return reply.Data, nil
	}
}

//a better solution would be an abstracted function, not for this demo :)
var rateAdapter func(currencyId string) (float64, error)

func registerAdapters(nc *nats.Conn) {
	rateAdapter = func(currencyId string) (float64, error) {

		type req struct {
			ID string `json:"id"`
		}

		type rep struct {
			Id    string  `json:"id"`
			Rate  float64 `json:"rate"`
			Error string  `json:"error,omitempty"`
		}
		rq := req{ID: currencyId}
		var rp rep

		if bt, err := json.Marshal(&rq); err != nil {
			log.Printf("rate adapter cant Unmarshal [%s] to request", string(bt))
			return 0, err
		} else if ret, err := request(nc, "rate", bt); err != nil { //sorry for the shortcuts :)
			return 0, err
		} else if err := json.Unmarshal(ret, &rp); err != nil {
			log.Printf("[#rate] cant Unmarshal [%s] to rate", string(ret))
			return 0, err
		} else if rp.Error != "" {
			return 0, errors.New(rp.Error)
		} else {
			return rp.Rate, nil
		}
	}
}
