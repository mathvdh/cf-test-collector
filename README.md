# CurrencyFair #

## test for Mathieu Van der Haegen ##

This is the code for the collector part of the test for CurrencyFair

It simply collect the transaction records through POST'ed json data, parse it and store it into AWS Elasticache (Redis).

*Why is it coded in Go?* Because I've been wanting to play a bit with Go for quite a time and as it is depicted as having very good performances, like between C and Java. I thought it might be a good fit for this part of the project. I read some articles about Go before but never coded with it. It's quite clean actually :)

You can run the tests by doing : `go test` in this folder.
You can build the project : `go build` in this folder and then `./cf-test-collector` to launch it.

After that there is another component here : which whill inspect REDIS do some repetitive stats calculations and push them in a realtime bus. For that I used Pubnub which is very scalable. It could be replace by other like socketio or else for cost or other reasons. For the frontend part it's located in a AWS S3 Bucket with theorically infinite scaling (You could say "enough")



