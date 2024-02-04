package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
