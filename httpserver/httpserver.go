package main

import (
	"httpserver/handler"
	"log"
	"net/http"
)

func main() {
	h := handler.New()

	log.Println("httpserver start on localhost:8000")

	err := http.ListenAndServe(":8000", h)
	if err != nil {
		panic(err)
	}
}
