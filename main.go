package main

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", r)
}

func main() {
	http.HandleFunc("/", Handle)
	s := http.Server{Addr: "0.0.0.0"}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
