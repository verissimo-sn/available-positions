package handler

import (
	"github.com/gin-gonic/gin"
)

type HttpFormatter struct {
}

func NewHttpFormatter() *HttpFormatter {
	return &HttpFormatter{}
}

func (formatter *HttpFormatter) Response(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.Header("content-type", "application/json")
	if data == nil {
		ctx.Status(statusCode)
		return
	}

	if message, ok := data.(string); ok {
		ctx.JSON(statusCode, gin.H{
			"message": message,
		})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}
