package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handleEcho(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("text")
	fmt.Fprintf(w, val)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method != http.MethodGet {
			http.Error(w, "405 | method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "")
		//tmpl.Execute(w, nil)
	case "/echo":
		if r.Method != http.MethodPost {
			http.Error(w, "405 | method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handleEcho(w, r)
	default:
		http.NotFound(w, r)
	}
}

var tmpl *template.Template
var err error

func main() {
	// tmpl, err = template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	log.Fatal("error", err)
	// }
	fmt.Println("http://localhost:8080:")
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
