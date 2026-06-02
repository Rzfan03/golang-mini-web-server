package main

import (
	"fmt"
	"log"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
if err := r.ParseForm(); err != nil {
	fmt.Fprintf(w, "ParseForm() err: %v", err)
	return
}
fmt.Fprintf(w, "POST Request Succesfully\n")
nama := r.FormValue("name")
alamat := r.FormValue("address")

fmt.Fprintf(w, "Name = %s\n", nama)
fmt.Fprintf(w, "Address = %s\n", alamat)
}


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not support", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello Nga!")
}

func main() {
	FileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", FileServer)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.Println("Starting at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}	
}