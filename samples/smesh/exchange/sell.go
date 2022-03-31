package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func init() {
	handlers["sell"] = processSell
}

func processSell(data []byte) []byte {
	type req struct {
		AccountId  string  `json:"accountId"`
		CurrencyId string  `json:"currencyId"`
		Amount     float64 `json:"amount"`
	}

	type rep struct {
		TrsId string  `json:"tsrId"`
		Rate  float64 `json:"rate"`
		Total float64 `json:"total"`
		Error string  `json:"error,omitempty"`
	}

	var rq req
	var rp rep

	if err := json.Unmarshal(data, &rq); err != nil {
		log.Printf("[#sell] cant Unmarshal [%s] to request", string(data))
	} else if rate, err := rateAdapter(rq.CurrencyId); err != nil {
		log.Printf("[#sell] error getting rate for [%+v] : %s", rq, err.Error())
		return svcError
	} else {
		//check balance, insert trs into db etc
		rp = rep{TrsId: fmt.Sprint(20000 + msgCount), Rate: rate, Total: rate * rq.Amount}
	}
	if reply, err := json.Marshal(rp); err != nil {
		log.Printf("[#sell] cant Marshal [%+v] to response: %s", rp, err.Error())
		return svcError
	} else {
		return reply
	}
}
