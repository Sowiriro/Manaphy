package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ちゃんと承認をしたしん")
	return
}

func signUpAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "登録をしてくれだしん")
	return
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "signupを表示しました")
	return
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logoutをしただしん")
	return
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ログインをしただしん")
	return
}

func Logout(c *gin.Context) {
	return
}
