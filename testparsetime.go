package main

import (
	"fmt"
)


func test() {
	p := fmt.Println

	testtime := "24-JAN-15 10:27:44"

    tplaced := ParseDate(testtime)

    p(tplaced)
    p(tplaced.Unix())
}