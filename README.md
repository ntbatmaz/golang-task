# golang-task
 
A REST-API service that works as an in memory key-value store.

For the run program please use this command;
->go run main.go

Usage For Set Method(Using for adding key-value pair);
-> curl --location --request POST "localhost:8080/set?key=&val="

Usage For Get Method(Using for return the value of a key);
-> curl --location --request GET "localhost:8080/get/<key>"
 
Usage For Flush Method
->curl --location --request PUT "localhost:8080/flush"
 
It runs on port 8080 by default.
