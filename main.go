package main

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Method]: %s\n[RemoteAddr]: %s\n[REQUEST_URI]: %s\n[UserAgent]: %s\n[Cookies]: %s\n[Headers]: %s\n", r.Method, r.RemoteAddr, r.RequestURI, r.UserAgent(), r.Cookies(), r.Header)
}

func main() {
	http.HandleFunc("/", Handle)
	s := http.Server{Addr: "0.0.0.0:12345"}
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("error: %s", err)
	}
}
