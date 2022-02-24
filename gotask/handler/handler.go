//Package provide handler functions implementations.
package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var cache = make(map[string]string)

//It checks which api method is called.
//It writes the data in the cache to the file every 10 seconds
func Handlers(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		Get(w, r)
	case "POST":
		Set(w, r)
	case "PUT":
		FlushData(w)
	}

	defer func() {
		go writeFile()
	}()

}


func Get(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.Path, "/")[2]

	e := cache[key]
	if len(e) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Record not found for this key!"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(e))
	}

	log.Println("Get Key = " + key + "Get Value = " + e)

}

func Set(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("val")

	cache[key] = value
	log.Println("Key = " + key + " " +"Value = " + value )

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))

}

//It ensures that the data in the cache is saved to the file every 10 seconds.
func writeFile(){
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		formatData, _ := json.Marshal(cache)
		os.WriteFile("temp/data.json", formatData, os.ModeAppend)
	}
}

//Function that inserts the records in the file into the cache when the project is first run.
func OriginFunc(){
	dataFile, err := os.Open("temp/data.json")

	if err != nil {
		fmt.Println(err)
	}

	file, _ := ioutil.ReadAll(dataFile)

	json.Unmarshal([]byte(file), &cache)
}

func FlushData(w http.ResponseWriter) {
	for cac, _ := range cache {
		delete(cache, cac)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cache is empty!"))
}