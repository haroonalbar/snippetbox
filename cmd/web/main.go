package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// server files in ui/static
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/snippet/view", handleSnippetView)
	mux.HandleFunc("/snippet/create", handleSnippetCreate)

	// handle file server also strip the prefix to get the correct path
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Start server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
