package main

import (
	"fmt"
	"net/http"
)

func health() string { return "ok" }

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { fmt.Fprint(w, health()) })
	_ = http.ListenAndServe(":8080", nil)
}
