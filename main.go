package main

import (
	"log"
	"net/http"
	"os/exec"
)

type Router struct{}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("Method: %s | URL: %s | Header: %s\n", r.Method, r.URL.Path, r.Header)
}

func main() {
	http.ListenAndServe(":3902", &Router{})
}
