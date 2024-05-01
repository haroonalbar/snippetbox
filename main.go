package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter ,r *http.Request ) {
  w.Write([]byte("Boi"))
}

func main(){
  mux:= http.NewServeMux()
  mux.HandleFunc("/",home)
  fmt.Println("Hello boss")

  log.Print("Starting server on port 4000")
  err:=  http.ListenAndServe(":4000",mux)
  log.Fatal(err)
}
