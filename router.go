package main

import (
	"fmt"
	"net/http"
)

func HelloMMG(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world")
}

func newRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", HelloMMG)
	return router
}
