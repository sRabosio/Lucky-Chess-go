package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const devPort = "8080"
const prodPort = ""

func main() {
	fmt.Println("starting server")

	//todo: replace w/ port from args
	const currentPort = devPort

	index := func(writer http.ResponseWriter, req *http.Request) {
		templ := template.Must(template.ParseFiles("./pages/master.html"))
		templ.Execute(writer, nil)
	}

	http.HandleFunc("/", index)

	//assets server
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//start server
	log.Fatal(http.ListenAndServe(":"+currentPort, nil))
}
