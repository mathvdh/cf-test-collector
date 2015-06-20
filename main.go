package main

import (
    "log"
    "net/http"
    "flag"
)

//now I have to write another go program that will :
//every X millisec check the redis data store
//for all new transactions
//will compute running stats
//last hour average amounts
//per origin country amounts
//save all this into redis back
//push into pubnub

//frontend will receive stat messages and display them in a "nice way"
//needs more thinking
//maybe a map on top with bubbles per country (like big bubbles if multiple transaction or small if only one)
//on the bottom a histogram that will display per minute #transactions or transaction amounts
//and then a counter on top right 
//#transactions (total / last hour / last minute)
//#transaction #top user #top country #top currency to #top currency from


func ReceiveTransaction(rw http.ResponseWriter, req *http.Request) {
    
    t:=new(TransactionRecord)

    t.FromJsonRequest(req.Body)

    // t.Log()

    t.ToRedis()
}

func main() {
    flag.Parse()
    pool = newPool(*redisServer, *redisPassword)

    http.HandleFunc("/transaction", ReceiveTransaction)
    log.Fatal(http.ListenAndServe(":80", nil))
}