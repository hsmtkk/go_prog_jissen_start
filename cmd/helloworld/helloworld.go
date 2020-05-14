package main

import (
	"net/http"

	"github.com/hsmtkk/go_prog_jissen_start/pkg/helloworld"
)

func main() {
	http.HandleFunc("/", helloworld.Handler)
	http.ListenAndServe(":8080", nil)
}
