package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Gorilla MUX routing!")

	//	create a new router
	router := mux.NewRouter()

	//	Registering a Request Handler
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MUX rout handler")
	})

	//URL Parameters
	//The biggest strength of the gorilla/mux Router is the ability to extract segments from the request URL. As an example, this is a URL in your application:
	///books/go-programming-blueprint/page/10

	//The last thing is to get the data from these segments. The package comes with the function mux.Vars(r) which takes the http.Request as parameter and returns a map of the segments.
	//func(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	vars["title"] // the book title slug
	//	vars["page"]  // the page
	//}
	//listen to port
	err := http.ListenAndServe(":9500", router)
	if err != nil {
		return
	}
}
