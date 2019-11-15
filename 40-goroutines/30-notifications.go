package main

import (
	"fmt"
	"net/http"
)

func logger(requests chan string) {
	for request := range requests {
		fmt.Println(request)
	}
}

func main() {
	notifications := make(chan string)
	go logger(notifications)

	http.HandleFunc("/", request(notifications))
	http.ListenAndServe(":8080", nil)
}

func request(requests chan string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requests <- "hello world"
	}
}
