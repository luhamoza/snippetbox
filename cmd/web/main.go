package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet/create", snippetCreateHandler)
	mux.HandleFunc("/snippet/view", snippetViewHandler)
	log.Print("Starting at server 4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
