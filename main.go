package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/signup_account", signUpAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/movie", show)
	mux.HandleFunc("/movie/create", create)
	mux.HandleFunc("/movie/update", update)
	mux.HandleFunc("/movie/delete", delete)
	log.Printf("呼んでるしん")

	http.ListenAndServe(":8000", mux)
}
