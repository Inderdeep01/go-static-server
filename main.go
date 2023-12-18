package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/value", valueHandler)
	log.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	if req.Method != "POST" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := req.ParseForm(); err != nil {
		http.Error(res, "Something went wrong", http.StatusInternalServerError)
		return
	}
	name := req.FormValue("name")
	fmt.Fprintf(res, "Hello, %s!", name)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(res, "Hello, World!")
}

func valueHandler(res http.ResponseWriter, req *http.Request) {
	un := req.URL.Query()
	val1 := req.URL.Query().Get("key")
	val2 := req.URL.Query()["key2"]
	fmt.Fprintf(res, "val1: %s\n", val1)
	fmt.Fprintf(res, "val2: %s\n", val2)
	fmt.Fprintf(res, "un: %s\n", un)
}
