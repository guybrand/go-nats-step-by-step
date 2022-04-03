# Creating a small service mesh using nats and http API gateway

# Requirements
Go 1.14 or later version

NATS server running on your local machine, easiest way to run it:
```
$ docker pull nats
$ docker run --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats --http_port 8222
```



# Use case : 
Lets assume we need a Cryptocurrency exchange with an API gateway and 2 services:
- When we call /rate handler we expect a quote for the currency (rate service)
- When we call /purchase or /sell handler we expect exchange service to:
- Get the latest quote for the purchased currency (using the rate service)
- Commit the transaction, return TrsID and total amount converted

![smesh](/samples/smesh/smesh.png)


# Step1 (run the http API gateway):
Open a terminal windows

`$ cd samples/smesh`

`$ go run .`


Open another terminal window

`curl '127.0.0.1:8090/rate' -H 'content-type: application/json' --data-raw '{"id":"btc"}'`

An you will recieve an error - as rate service is not running !


# Step2 (run the rate service):
Open another terminal window

`$ cd samples/smesh`

`$ go run rate/*.go`


in the previous curl terminal window and type:

```
#Get the rate for btc
curl '127.0.0.1:8090/rate' -H 'content-type: application/json' --data-raw '{"id":"btc"}'

#Get the rate for eth
curl '127.0.0.1:8090/rate' -H 'content-type: application/json' --data-raw '{"id":"eth"}'

#Get the rate for unknown currency xxx (error)
curl '127.0.0.1:8090/rate' -H 'content-type: application/json' --data-raw '{"id":"xxx"}'
```



# Step3 (try to purchase without an exchange service):
in the previous curl terminal window and type:

`curl '127.0.0.1:8090/purchase' -H 'content-type: application/json' --data-raw '{"accountId":"123","currencyId":"btc","amount":100}'`

An you will recieve an error - as purchase service is not running !

Open another terminal window


# Step4 (run the exchange service):
Open another terminal window

`$ cd samples/smesh`

`$ go run exchange/*.go`

in the previous curl terminal window and type:

```
#purchase btc
curl '127.0.0.1:8090/purchase' -H 'content-type: application/json' --data-raw '{"accountId":"123","currencyId":"btc","amount":100}'

#purchase eth
curl '127.0.0.1:8090/purchase' -H 'content-type: application/json' --data-raw '{"accountId":"123","currencyId":"eth","amount":15}'

#sell eth
curl '127.0.0.1:8090/sell' -H 'content-type: application/json' --data-raw '{"accountId":"123","currencyId":"eth","amount":20}'

#sell unknown currency (get an error)
curl '127.0.0.1:8090/sell' -H 'content-type: application/json' --data-raw '{"accountId":"123","currencyId":"xxx","amount":20}'
```

