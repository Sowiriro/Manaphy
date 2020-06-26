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
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signUpAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/movie/:id", show)
	//mux.HandleFunc("/movie/create", create)
	mux.HandleFunc("/movie/:id/update", update)
	log.Printf("updateの後ろまで来た！")
	mux.HandleFunc("/movie/:id/delete", delete)
	log.Printf("delete完了あとはlistenのみ")

	http.ListenAndServe(":8000", mux)
	log.Printf("listen終わったよ")
}
