package main

import (
    "encoding/json"
    "log"
    "reflect"
    "io"
    "fmt"
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

func (record *TransactionRecord) Log() {
    s := reflect.ValueOf(&record).Elem()

    log.Println("***transaction***")
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        log.Println(f.Interface())
    }
    log.Println("*****************")
}

func (record *TransactionRecord) FromJsonRequest(jsondata io.ReadCloser) {

	decoder := json.NewDecoder(jsondata)
    err := decoder.Decode(&record)
    if err != nil {
        panic(err)
    }
}

func (record *TransactionRecord) ToJson() string {
	encoded, err := json.Marshal(record)
    if err != nil {
        panic(err)
    }
    encoded_str := string(encoded)
    return encoded_str
}

func (record *TransactionRecord) ToRedis() {
    conn := pool.Get()
    defer conn.Close()

    tplaced := ParseDate(record.TimePlaced)

    encoded_str := record.ToJson()

    tid, err := conn.Do("INCR","transaction_id")
    if err != nil {
        panic(err)
    }

    n1, err1 := conn.Do("SET", fmt.Sprint("transaction:",tid), encoded_str )
    if err1 != nil {
        panic(err1)
    }

    log.Println(n1)

    n2, err2 := conn.Do("ZADD", fmt.Sprint("transactions"), tplaced.Unix(), tid )
    if err2 != nil {
        panic(err2)
    }
    
    log.Println(n2)
}