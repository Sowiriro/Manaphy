package main

import (
	"log"

	"github.com/Sowiriro/Manaphy/internal/domain/repository"
	"github.com/Sowiriro/Manaphy/internal/handler"
	data "github.com/Sowiriro/Manaphy/internal/interface/repository"
	"github.com/Sowiriro/Manaphy/internal/pkg/config"
	"github.com/Sowiriro/Manaphy/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/hinzhu/gorm"
)

var (
	DB      *gorm.DB
	mainCfg *config.Config
)

var (
	userHandler handler.UserHandler
)
var (
	userUseCase usecase.UserUseCase
)
var (
	userRepo repository.UserRepository
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

func initRepository() {
	userRepo = data.NewUserRepository(DB)
}

func initUsecase() {
	userUseCase = usecase.NewUserRepository(userRepo)
}

func initHandler() {
	userHandler = handler.NewUserHandler(userUseCase)
}

func run() {
	e := gin.Default()

	//routing for user
	e.GET("/user", userHandler.Index)
	e.GET("/user/:id", userHandler.Show)
	e.POST("/user", userHandler.Create)
	e.POST("/user/:id", userHandler.Update)
	e.DELETE("/user/:id", userHandlerDelete)

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
