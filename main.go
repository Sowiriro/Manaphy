package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()


	//r.Handle("/ping", negroni.New(
	//	negroni.HandlerFunc(jwtMiddleWare.HandlerWithNext),
	//	negroni.Wrap(index),
	//))

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	//ここでroutingをしている　mux.HandleFunc()
	mux.Handle("/static", http.StripPrefix("/static/", files))

	r.Handle("/", http.HandlerFunc(index))
	r.HandleFunc("/err", err)

	r.HandleFunc("/login",login)
	r.Handle("/logout", jwtMiddleWare.Handler(http.HandlerFunc(logout)))
	r.HandleFunc("/signUp", signUp)
	r.Handle("/authenticate",jwtMiddleWare.Handler(http.HandlerFunc(authenticate)))

	r.Handle("/movie", jwtMiddleWare.Handler(http.HandlerFunc(show)))
	r.Handle("/movie/create", jwtMiddleWare.Handler(movieCreate)).Methods("POST")
	r.Handle("/movie/update", jwtMiddleWare.Handler(http.HandlerFunc(update)))
	r.Handle("/movie/delete", jwtMiddleWare.Handler(http.HandlerFunc(delete)))
	log.Printf("呼んでるしん")

	http.ListenAndServe(":8000", r)
}

//func main() {
//	var array = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//	//このarrayから二部探索木をつくって、9をみつける。
//	num := len(array)
//	println(num)
//
//	search := 9
//
//	left := 0
//
//	right := num
//}
