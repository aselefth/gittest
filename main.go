package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	s := []byte("hello_world")
	w.Write(s)
	w.WriteHeader(201)

}

func main() {
	fmt.Println("running server")
	http.HandleFunc("/", helloWorld)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		os.Exit(1)
	}

}
