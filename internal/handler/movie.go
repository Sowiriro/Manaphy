package main

import (
	"github.com/gin-gonic/gin"
)

type MovieHandlerI interface {
	Index(c *gin.context)
	Show(c *gin.context)
	Create(c *gin.context)
	Update(c *gin.context)
	Delete(c *gin.context)
}

type MovieHandler struct {
	movieUseCase usecase.MovieUseCase
}

func NewMovieHandler(usecase usecase.MovieUseCase) MovieHandler {
	return &MovieHandlerI{movieUseCase}
}

func ()Index(c *gin.Context) {
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
