//Package main provide a key-value pair rest api main function
package main

import (
	"gotask/handler"
	"log"
	"net/http"
)

func main() {
	Start()
}

//All api handlers are here.
func Start() {
	handler.OriginFunc()
	http.HandleFunc("/set", handler.Handlers)
	http.HandleFunc("/get/", handler.Handlers)
	http.HandleFunc("/flush", handler.Handlers)

	http.ListenAndServe(":8080", nil)
	log.Println("Api Started Port: 8080")

}

