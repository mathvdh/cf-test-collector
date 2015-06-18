package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/garyburd/redigo/redis"
    "time"
    "flag"
    "fmt"
)

func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        Dial: func () (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}

var (
    pool *redis.Pool
    redisServer = flag.String("redisServer", "transactionscf.ndvlwb.0001.euw1.cache.amazonaws.com:6379", "")
    redisPassword = flag.String("redisPassword", "", "")
)

func ReceiveTransaction(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t TransactionRecord   
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }

    conn := pool.Get()
    defer conn.Close()


    tplaced := ParseDate(t.TimePlaced)

    n1, err1 := conn.Do("ZADD", "transactions", tplaced.Unix(), req.Body )
    if err1 != nil {
        panic(err1)
    }

    fmt.Println(n1)

    n2, err2 := conn.Do("ZADD", fmt.Sprint("transactions:",t.UserId), tplaced.Unix(), req.Body )
    if err != nil {
        panic(err2)
    }
    
    fmt.Println(n2)

    

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
    flag.Parse()
    pool = newPool(*redisServer, *redisPassword)

    http.HandleFunc("/transaction", ReceiveTransaction)
    log.Fatal(http.ListenAndServe(":80", nil))
}