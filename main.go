package main

import (
    "log"
    "net/http"
    "flag"
)



func ReceiveTransaction(rw http.ResponseWriter, req *http.Request) {
    
    var t TransactionRecord

    t.FromJsonRequest(req.Body)

    t.Log()

    t.ToRedis()

    t.Log()
}

func main() {
    flag.Parse()
    pool = newPool(*redisServer, *redisPassword)

    http.HandleFunc("/transaction", ReceiveTransaction)
    log.Fatal(http.ListenAndServe(":80", nil))
}