package main

import (
	// the dot (.) make the use of fmt when using its functions unecessary
	. "fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		Fprintf(w, "Parseform() err: %v", err)
		return
	}
	Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	Fprintf(w, "name = %s\n", name)
	Fprintf(w, "address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "methos not found", http.StatusNotFound)
		return
	}
	Fprintf(w, "Hello!")
}

func main() {
	// the ":=" operator is used to declare and assign the value
	// int a = 10 is the same as a := 10
	// http.Dir(“xxx”) : returns a filesystem.
	// In this case, the collection of files will be the list of files present in the folder we passed as an input to http.Dir()
	// http.FileServer(): It takes a filesystem(return value http.Dir()) and creates an handler.
	// The handler will in turn return all the files listed by http.Dir to the client, that is why to access the form you need to go to the form.html url
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// Because of the index file the root "/" will return a static page (index)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	Printf("Server Started!")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
