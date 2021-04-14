package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		message := []byte("Hello, World!")
		res.Write(message)
	})
	http.ListenAndServe(":8080", nil)
}
