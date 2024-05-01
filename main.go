package main

import (
	"log"
	"net/http"
)

// fucnction to handle home path
func home(w http.ResponseWriter ,r *http.Request ) {
  w.Write([]byte("Boi"))
}

func bro(w http.ResponseWriter ,r *http.Request ) {
  w.Write([]byte("Sup Bro!"))
}

func main(){
  // create a new mux
  // mux is essentionaly a router or http live multiplexer/server

  mux:= http.NewServeMux()

  // to handle the root route with home function
  mux.HandleFunc("/",home)
  // handle for /bro route using bro function
  mux.HandleFunc("/bro",bro)
  

  log.Print("Starting server on port 4000")

  // listen to listen to http live server/multiplexer aka mux
  err:=  http.ListenAndServe(":4000",mux)
  // if there is err log it
  log.Fatal(err)
}
