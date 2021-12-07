package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
func home(rw http.ResponseWriter,r *http.Request){
	fmt.Println("serving home page ....")
	fmt.Fprintf(rw,"%s","welcome to home Page")
}