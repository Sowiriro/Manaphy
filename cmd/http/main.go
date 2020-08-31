package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	e := gin.Default()

	//routing for user
	e.GET("/user", Index)
	e.GET("/user/:id", Show)
	e.POST("/user", Create)
	e.POST("/user/:id", Update)
	e.DELETE("/user/:id", Delete)

	//routing for movie
	e.GET("/movie", Index)
	e.GET("/movie/:id", Show)
	e.POST("/movie", Create)
	e.POST("/movie/'id", Update)
	e.DELETE("/movie/:id", Delete)

	//routing for login
	e.GET("/login", Login)
	e.GET("/logout", Logout)
	e.POST("/authenticate", Signup)

	log.Printf("呼んでるしん")

	e.run(":8080")
}
