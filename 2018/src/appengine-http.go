package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", handler)
	appengine.Main()
}
