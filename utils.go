package main

import (
    "fmt"
    "time"
    "strings"
    "unicode"
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

var cnt int = 0

func monthToLower(r rune) rune {
    retval := r
    if (unicode.IsUpper(r) && cnt >= 1) {
        retval=unicode.ToLower(r)
    }

    if (unicode.IsLetter(r)) {
        cnt = cnt + 1
    }
    
    return retval
}

func ParseDate(datestr string) time.Time {

    cnt=0

    newdatestr := strings.Map(monthToLower, datestr)
    
    fmt.Println(newdatestr)

    tobj, err := time.Parse("02-Jan-06 15:04:05", newdatestr)
    if err != nil {
        panic(err)
    }

    return tobj
    
}