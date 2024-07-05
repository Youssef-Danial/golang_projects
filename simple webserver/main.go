package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name     string
	location string
}

func (self User) save(w http.ResponseWriter) {
	fmt.Fprintf(w, "Name = %s\n", self.name)
	fmt.Fprintf(w, "Location = %s\n", self.location)
	fmt.Println(self.name)
	fmt.Println(self.location)
}

func hello_handler(w http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}
func form_handler(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	if request.Method != "POST" {
		http.Error(w, "this method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := request.FormValue("name")
	location := request.FormValue("location")
	user := User{name, location}
	user.save(w)
}

func main() {
	file_server := http.FileServer(http.Dir("./static"))
	http.Handle("/", file_server)
	http.HandleFunc("/form", form_handler)
	http.HandleFunc("/hello", hello_handler)
	fmt.Println("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
