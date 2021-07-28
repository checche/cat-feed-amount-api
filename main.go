// +build ignore

package main

import (
	"cat-feed-amount-api/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.CatHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
