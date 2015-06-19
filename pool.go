package main

import (
    "github.com/garyburd/redigo/redis"
    "time"
    "flag"
    )

var (
    pool *redis.Pool
    redisServer = flag.String("redisServer", "transactionscf.ndvlwb.0001.euw1.cache.amazonaws.com:6379", "")
    redisPassword = flag.String("redisPassword", "", "")
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
            
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}

