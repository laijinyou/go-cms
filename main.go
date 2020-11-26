package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello World!This is a go website.</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "Welcome to the about page."+"<a href=\"https://trip123.com\">Trip123.com</a>")
	} else {
		fmt.Fprint(w, "<h1>The request page is not found.</h1>"+"<p>if you have any questions,please contact me.</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
