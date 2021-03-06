package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("START")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>\n<body>\n")
	fmt.Fprintf(w, "<h1>Success!!</h1>\n")
	fmt.Fprintf(w, "</body>\n<html>\n")
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
}
