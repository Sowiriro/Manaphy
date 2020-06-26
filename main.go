package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", Index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signUpAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/movie/:id", Show)
	mux.HandleFunc("/movie/create", Create)
	mux.HandleFunc("/movie/:id/update", Update)
	mux.HandleFunc("/movie/:id/delete", Delete)
	log.Printf("deleteまでだいじょうぶだった")

	http.ListenAndServe(":8080", mux)
}
