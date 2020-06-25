package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	//mux.HandleFunc("/", Index)
	//mux.HandleFunc("/err", err)
	//
	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signUpAccount)
	//mux.HandleFunc("/authenticate", authenticate)
	//
	//mux.HandleFunc("/movie/:id", show)
	log.Printf("hoihoi")
	mux.HandleFunc("/movie/create", create)
	//mux.HandleFunc("/movie/:id/update", update)
	//mux.HandleFunc("/movie/:id/delete", delete)
	//f
	http.ListenAndServe(":8000", mux)
}
