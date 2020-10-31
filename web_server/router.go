package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, urlExists := r.rules[path]
	handler, methodExists := r.rules[path][method]

	return handler, methodExists, urlExists
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(w, "Hellow World")
	handler, methodExists, urlRxists := r.FindHandler(request.URL.Path, request.Method)

	if !urlRxists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
