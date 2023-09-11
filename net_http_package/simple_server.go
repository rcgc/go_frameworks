package main

import (
	"fmt"
	"net/http"
)

/*
	This is the simplest web server you
	can implement in Go
*/
func main(){
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}