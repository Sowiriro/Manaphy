package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/auth0/go-jwt-middleware"
	"github.com/urfave/negroni"
)

func main() {

	r := mux.NewRouter()

	jwtMiddleWare := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token)(interface{}, error){
			return []byte("Mysecret"), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})

	r.Handle("/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleWare.HandlerWithNext),
		negroni.Wrap(index),
	))

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static/", files))

	r.Handle("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login",login)
	r.Handle("/logout", jwtMiddleWare.Handler(http.HandlerFunc(logout)))
	mux.HandleFunc("/signUp",signUp)
	r.Handle("/authenticate",jwtMiddleWare.Handler(http.HandlerFunc(authenticate)))

	r.Handle("/movie", jwtMiddleWare.Handler(show))
	r.Handle("/movie/create", jwtMiddleWare.Handler(movieCreate))
	r.Handle("/movie/update", jwtMiddleWare.Handler(http.HandlerFunc(update)))
	r.Handle("/movie/delete",jwtMiddleWare.Handler(http.HandlerFunc(delete)))
	log.Printf("呼んでるしん")

	http.ListenAndServe(":8000", mux)
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
