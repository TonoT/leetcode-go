package main

import (
	"fmt"
	"net/http"
)

type test111 struct {
}

func (t test111) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Print("12345")
}

func main() {
	s := test111{}
	http.Handle("/ab/rew/:123", s)
	//http.Handle("/ab/get/rew", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Print("12321") }))

	http.ListenAndServe(":8080", nil)
}
