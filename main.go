package main

import (
	"net/http"
	"time"

	"github.com/ogataka50/fortune-api/fortune"
)

func main() {
	http.Handle("/", &fortune.Handler{Time: time.Now()})
	http.ListenAndServe(":8080", nil)
}
