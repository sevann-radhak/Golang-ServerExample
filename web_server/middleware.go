package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")

			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Log Step!")
			start := time.Now()
			time.Sleep(1 * time.Second)

			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			f(w, r)
		}
	}
}
