package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Sowiriro/internal/usecase"
	"github.com/Sowiriro/internal/domain/entity"
	"net/http"
)

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
	var req entity.UserPostRequest

	err := c.ShouldBindJSON(&req)
	if  err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := h.userUseCase.All(ctx, req)
	if err != nil{
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}
	WriteData(ctx, http.StatusOK,  response)
}

func (h *UserHandler) Show(c *gin.Context) {
	var req entity.User
	id := ctx.Param("id")

	err := ctx.ShouldBindJSON(&req)
	if  err != nil {
		WriteError(ctx, http.StatusOK, err)
	}

	user, err := h.userUseCase.ByID(id)
	if err != nil {
		return
	}

	WriteData(ctx, http.StatusOK, user)
}
}
func (h *UserHandler) Create(c *gin.Context) {
	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
	}

	resp, err := h.userUseCase.Create(req)
	if err != nil {
		return
	}

	WriteData(ctx, http.StatusOK, resp)
}
func (h *UserHandler) Update(c *gin.Context) {
	var req entity.User
	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
	}

	err = h.userUseCase.Update(req, userId)

	WriteData(ctx, http.StatusOK, err)

}
func (h *UserHandler) Delete(c *gin.Context) {
	return
}
