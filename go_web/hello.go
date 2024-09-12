package main

import (
	"fmt"
	"net/http"
)

type Todo struct {
	name     string
	done     bool
	assigned bool
}

func main() {
	fmt.Println("hello world")

	//	register a simple route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("My first Go route handler.")
		fmt.Println("You requested", r.URL.Path)
	})

	//todos handler
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := []Todo{{name: "pack the clothes", done: false, assigned: true}, {name: "clean the house", done: true, assigned: true}}
		fmt.Println(todos)
	})

	//	listen to the handler on port 9200
	err := http.ListenAndServe(":9200", nil)
	if err != nil {
		return
	}
}
