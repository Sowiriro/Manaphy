package main

import "github.com/gin-gonic/gin"

type UserHandlerI interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &UserHandlerI{userUseCase}
}

func (h *UserHandler) Index(c *gin.Context) {
	return
}

func (h *UserHandler) Show(c *gin.Context) {
	return
}
func (h *UserHandler) Create(c *gin.Context) {
	return
}
func (h *UserHandler) Update(c *gin.Context) {
	return
}
func (h *UserHandler) Delete(c *gin.Context) {
	return
}
