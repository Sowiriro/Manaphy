package handler

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{} `json:"data, omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteData(ctx *gin.Context, httpStatus int, data interface{}){
	write(ctx, httpStatus, Response{Data: data})
}

func WriteError(ctx *gin.Context, httpStatus int, err error){
	write(ctx, httpStatus, Response{Error: err.Error()})
}

func write(ctx *gin.Context, httpStatus int, res Response){
	ctx.Json(httpStatu, resp)
}