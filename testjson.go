package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type TransactionRecord struct {
    UserId string
    CurrencyFrom string 
    CurrencyTo string 
    AmountSell json.Number
    AmountBuy json.Number
    Rate json.Number
    TimePlaced string
    OriginatingCountry string
}

func ReceiveTransaction(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t TransactionRecord   
    err := decoder.Decode(&t)
    if err != nil {
        panic("Error while decoding")
    }

    log.Println("***transaction***")
    log.Println(t.UserId)
    log.Println(t.CurrencyFrom)
    log.Println(t.CurrencyTo)
    log.Println(t.AmountSell)
    log.Println(t.AmountBuy)
    log.Println(t.TimePlaced)
    log.Println(t.OriginatingCountry)
    log.Println("*****************")
}

func main() {
    http.HandleFunc("/transaction", ReceiveTransaction)
    log.Fatal(http.ListenAndServe(":80", nil))
}