package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Snippetbox home page."))
}

func handleSnippetView(w http.ResponseWriter, r *http.Request) {
	// extract id parameter and check it's postive number
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

  // formatted http.ResponseWriter
	fmt.Fprintf(w, "View snipped of id: %d", id)
}

func handleSnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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
