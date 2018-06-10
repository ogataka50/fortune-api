package main

import (
	"net/http"

	"github.com/ogataka50/fortune-api/fortune"
)

func main() {
	http.HandleFunc("/", fortune.Handler)
	http.ListenAndServe(":8080", nil)
}
