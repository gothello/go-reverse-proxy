package main

import (

	"fmt"
	"log"
	"net/http"
)

func Client(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Path[1:])
	log.Println(r.URL)
}

func main() {

	http.HandleFunc("/", Client)

	log.Println("init client server port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}