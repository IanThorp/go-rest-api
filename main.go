package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  // instantiate gorilla/mux router
  r := mux.NewRouter()

  // serve static index page by default
  r.Handle("/", http.FileServer(http.Dir("./views/")))
  // allow server to serve static assets
  r.Path.Prefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

  // declare port, pass in router
  http.ListenAndServe(":3000", r)
}
