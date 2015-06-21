# CurrencyFair test #

## cf-test-collector ##

This is an HTTP server that will receive requests with POST data in JSON.

The JSON data will look like : 

```
{"userId": "134256",
	"currencyFrom": "EUR", 
	"currencyTo": "GBP", 
	"amountSell": 1000, 
	"amountBuy": 747.10, 
	"rate": 0.7471, 
	"timePlaced" : "24-JAN-15 10:27:44", 
	"originatingCountry" : "FR"}
```

This is kept as simple as possible and focus on performances.

The interesting file to start is `main.go`

In `main` function :
* we get the flag data
* we start a new redis connection pool with login and password from flag

on "/transaction" we will call ReceiveTransaction

In `ReceiveTransaction` function :
* We create a new Transaction Record
* We load it from the JSON of the request
* We save it to redis

We could add validation but we prefer to proceed with validation (and other) later in [cf-test-alertor](https://github.com/mathvdh/cf-test-alertor) 