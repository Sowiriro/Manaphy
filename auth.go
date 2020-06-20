package main

import (
	"fmt"
	"net/http"
)

func authenticate (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ちゃんと承認をしたしん")
}

func signUpAccount(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "登録をしてくれだしん")
}

func signup(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "signupを表示しました")
}

func logout(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "logoutをしただしん")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "ログインをしただしん")
}

