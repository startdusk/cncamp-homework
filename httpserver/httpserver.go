package main

import (
	"httpserver/handler"
	"net/http"
)

func main() {
	h := handler.New()
	err := http.ListenAndServe(":8000", h)
	if err != nil {
		panic(err)
	}
}
