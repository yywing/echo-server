package main

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil{
		log.Printf("error: %v", err)
	}
	err = r.ParseMultipartForm(512)
	if err != nil{
		log.Printf("error: %v", err)
	}
	log.Printf("[Method]: %s\n", r.Method)
	log.Printf("[RemoteAddr]: %s\n", r.RemoteAddr)
	log.Printf("[REQUEST_URI]: %s\n", r.RequestURI)
	log.Printf("[UserAgent]: %s\n", r.UserAgent())
	log.Printf("[Cookies]: %s\n", r.Cookies())
	log.Printf("[Formdata]: %v\n", r.Form)
	log.Printf("[Headers]: %s\n", r.Header)
}

func main() {
	http.HandleFunc("/", Handle)
	s := http.Server{Addr: "0.0.0.0:12345"}
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("error: %s", err)
	}
}
