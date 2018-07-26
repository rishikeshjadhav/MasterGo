package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "\n Hello World from web service for home url \n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "\n Hello World from web service \n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "\n Hello World from web service for about url \n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

func main() {
	fmt.Println("Welcome to web service")

	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/about", about)

	fmt.Println("Listening on port 8000...")
	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}

}
