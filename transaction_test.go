package main

import (
	"testing"
	"io"
	"reflect"
	"bytes"
	// "log"
)


type nopCloser struct { 
    io.Reader 
}

func (nopCloser) Close() error { return nil } 

func TestTransactionRecordFromJsonRequest(t *testing.T) {

	transaction_ok := TransactionRecord{
		UserId:"134256",
		CurrencyFrom: "EUR", 
		CurrencyTo: "GBP", 
		AmountSell: "1000", 
		AmountBuy: "747.10", 
		Rate: "0.7471", 
		TimePlaced : "24-JAN-15 10:27:44", 
		OriginatingCountry : "FR",
	}

	json_ok := `{"userId": "134256",
					"currencyFrom": "EUR", 
					"currencyTo": "GBP", 
					"amountSell": 1000, 
					"amountBuy": 747.10, 
					"rate": 0.7471, 
					"timePlaced" : "24-JAN-15 10:27:44", 
					"originatingCountry" : "FR"}`

	// log.Println(json_ok)


	body := nopCloser{bytes.NewBufferString(json_ok)}

	// log.Println(body)
	
	transaction_json:=new(TransactionRecord)
	transaction_json.FromJsonRequest(body)

	// log.Println(&transaction_ok)
	// log.Println(transaction_json)

	if !reflect.DeepEqual(&transaction_ok,transaction_json) {
		t.Errorf("TestTransactionRecordFromJsonRequest failed")
	}

}