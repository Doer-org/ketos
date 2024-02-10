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
	fmt.Println(`

       ###  ##  #######  ######    #####    #####
	##  ##   ##   #  # ## #   ##   ##  ##   ##
	## ##    ## #      ##     ##   ##  #
	####     ####      ##     ##   ##   #####
	## ##    ## #      ##     ##   ##       ##
	##  ##   ##   #    ##     ##   ##  ##   ##
       ###  ##  #######   ####     #####    #####

	`)
	fmt.Println("Server started on http://localhost:8090")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
