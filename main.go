package main

import (
	"log"
	"net/http"
	"proxy/application"
)

func main() {
	port := ":8000"

	log.Printf("ProxyHandler started on port %s\n", port)

	s := application.NewProxyHandler()

	http.HandleFunc("/", s.ServeHTTP)
	log.Fatal(http.ListenAndServe(port, nil))
}
