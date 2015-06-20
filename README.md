# CurrencyFair #

## test for Mathieu Van der Haegen ##

This is the code for the collector part of the test for CurrencyFair

It simply collect the transaction records through POST'ed json data, parse it and store it into AWS Elasticache (Redis).

*Why is it coded in Go?* Because I've been wanting to play a bit with Go for quite a time and as it is depicted as having very good performances between C and Java. I read some articles about Go before but never code with it. It's quite clean actually.

You can run some tests by doing : `go test` in this folder.

After that there is another component here : which whill inspect REDIS do some repetitive stats calculations and push them in a realtime bus. For that I used Pubnub which is very scalable. It could be replace by other like socketio or else for cost or other reasons.



