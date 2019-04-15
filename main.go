package main

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r)
}

func main() {
	http.HandleFunc("/", Handle)
	s := http.Server{Addr: "0.0.0.0:12345"}
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("error: %s", err)
	}
}
