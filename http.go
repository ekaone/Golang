package main

import (
  “fmt”
  “log”
  “net/http”
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, “Hi, this is Docker Dewiweb!”)
}

func about(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, “Hi, this is Docker Dewiweb!”)
}
 
func main() {
  http.HandleFunc(“/”, index)
  http.HandleFunc(“/about”, about)

  log.Println(“application started at port :8000”)
  log.Fatal(http.ListenAndServe(“:8000”, nil))
}