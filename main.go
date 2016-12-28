package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func middlewareFirst(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Before MiddlewareFirst")
	next(w, r)
	log.Println("After MiddlewareFirst")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Before MiddlewareSecond")
	if r.URL.Path == "/message" {
		if r.URL.Query().Get("password") == "password" {
			next(w, r)
			log.Println("Correct password.")
		} else {
			log.Println("Wrong password.")
			return
		}
	} else {
		next(w, r)
	}
	log.Println("After MiddlewareSecond")
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index Handler")
	fmt.Fprintf(w, "Welcome")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing message handler")
	fmt.Fprintf(w, "Secret message : gophers are awesome")
}

func main() {
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(middlewareFirst),
		negroni.HandlerFunc(middlewareSecond),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/message", message)

	n.UseHandler(r)
	n.Run(":8080")
}
