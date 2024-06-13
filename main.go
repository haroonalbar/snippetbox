package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippetbox home page."))
}

func handleSnippetView(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Display a specific snippet..."))
}

func handleSnippetCreate(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Create a snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/snippet/view", handleSnippetView)
	mux.HandleFunc("/snippet/create", handleSnippetCreate)

  fmt.Println("Start server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
