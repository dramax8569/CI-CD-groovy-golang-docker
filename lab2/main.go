package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, Ping())
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Ping devuelve un string "pong"
func Ping() string {
	return "pong"
}
