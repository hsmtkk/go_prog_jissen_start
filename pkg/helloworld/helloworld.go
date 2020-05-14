package helloworld

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Hello World, %s!", path)
}
