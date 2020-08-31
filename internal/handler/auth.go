package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	fmt.Fprintf("ちゃんと承認をしたしん")
	return
}


func Logout(c *gin.Context) {
	fmt.Fprintf("logoutをしただしん")
	return
}

func Login(c *gin.Context) {
	fmt.Fprintf("ログインをしただしん")
	return
}
