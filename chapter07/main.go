package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "Hello %q", html.EscapeString(r.URL.Path)) //File writer
		//fmt.Sprintf() / string writer
		//fmt.Println() //console writer
	})
	http.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"hi")
	})
	log.Fatal(http.ListenAndServe(":8081",nil))
}

