package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Here is a simple HTTP server in Go!")

	//	create a default route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to learning Go web!")
	})

	//serving static assets
	//To serve static assets like JavaScript, CSS and images, we use the inbuilt http.FileServer and point it to a url path.
	fs := http.FileServer(http.Dir("static"))
	fmt.Println("Serving static files from directory at", fs)

	//point the url path
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//listen to the port
	http.ListenAndServe(":9300", nil)
}
