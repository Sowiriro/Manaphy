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

func (h *MovieHandler)Index(c *gin.Context) {
	return
}

func (h *MovieHandler) Show(c *gin.Context) {
	return
}

func (h *MovieHandler) Create(c *gin.Context) {
	return
}

func  (h *MovieHandler)Update(c *gin.Context) {
	return
}

func  (h *MovieHandler)Delete(c *gin.Context) {
	return
}
