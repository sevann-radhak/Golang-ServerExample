package main

import "net/http"

type Metadata interface{}

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Email string
	Name  string
	Phone string
}
