package main

import (
	"testing"
)


func TestParseDate(t *testing.T) {
	testtime := "24-JAN-15 10:27:44"

    tplaced := ParseDate(testtime)

    testtime_back := tplaced.Format("02-Jan-06 15:04:05")
    testtime_back_expected :=  "24-Jan-15 10:27:44"

    if testtime_back != testtime_back_expected {
    	t.Errorf("ParseDate(%s) returned %s, expected %s", testtime, testtime_back, testtime_back_expected)
	} 

}

var expectedResults = map[string]string {
	"JAN":  "Jan",
	"FEB":  "Feb",
	"24-JAN-15 10:27:44":  "24-Jan-15 10:27:44",
}

func TestSanitizeDate(t *testing.T) {
	for k, v := range expectedResults {
		res := SanitizeDate(k)
		if res != v {
    		t.Errorf("SanitizeDate(%s) returned %s, expected %s", k, res, v)
		} 
	}
}