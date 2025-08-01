package main

import (
	"log"
	"net/http"

	"github.com/dramax8569/CI-CD-groovy-golang-docker/handler"
)

func main() {
	http.HandleFunc("/ping", handler.PingHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
