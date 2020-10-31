package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exists := r.rules[path]

	return handler, exists
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(w, "Hellow World")
	handler, exists := r.FindHandler(request.URL.Path)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}
