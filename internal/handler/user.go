package main

import "github.com/gin-gonic/gin"

func NewUserHandler(UserUseCase usecase.UserUseCase) UserHandler {
	return &userHandler{userUseCase}
}

func Index(c *gin.Context) {
	return
}

func Show(c *gin.Context) {
	return
}
func Create(c *gin.Context) {
	return
}
func Update(c *gin.Context) {
	return
}
func Delete(c *gin.Context) {
	return
}
