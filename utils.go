package main

import (
    "time"
    "strings"
    "unicode"
)



func SanitizeDate(s string) string {
    cnt := 0

    mtlr := func(r rune) rune {
        retval := r
        if (unicode.IsUpper(r) && cnt >= 1) {
            retval=unicode.ToLower(r)
        }

        if (unicode.IsLetter(r)) {
            cnt = cnt + 1
        }
        
        return retval
    }

    return strings.Map(mtlr, s)
}



func ParseDate(datestr string) time.Time {

    newdatestr := SanitizeDate(datestr)

    tobj, err := time.Parse("02-Jan-06 15:04:05", newdatestr)
    if err != nil {
        panic(err)
    }

    return tobj
}