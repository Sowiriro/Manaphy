package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"database/sql"
)

var (
	sqlDB *sql.DB
	mainCfg *config.Config
)

var (
	userHandler handler.User
)
var (
	userUsecase usercase.User
)
var (
	userRepository repository.User
)

func main() {


	// initAppmodule()
	initHandler()
	initUsecase()
	initRepository()
	run()
}

// func initAppmodule(){

// }

func initRepository(){
	userRepository data.Repository
}

func initUsecase(){
	userUsecase
}


func initHandler(){
	us
}

func run(){
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