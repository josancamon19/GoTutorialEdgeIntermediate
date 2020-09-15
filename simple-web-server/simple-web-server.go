package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func home(w http.ResponseWriter, r *http.Request){
	file := r.URL.Path[1:]
	fmt.Println("File requested is " + file)
	http.ServeFile(w, r,file)
}

func increment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./simple-web-server/static")))
	http.HandleFunc("/increment",increment)

	log.Fatal(http.ListenAndServe(":8001", nil))
	// HTTPS
	// log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
