package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PositionHandler struct{}

func NewPositionHandler() *PositionHandler {
	return &PositionHandler{}
}

func (h *PositionHandler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success POST available positions",
	})
}

func (h *PositionHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success PUT available positions",
	})
}

func (h *PositionHandler) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success GET available positions",
	})
}

func (h *PositionHandler) GetById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success GET available positions",
	})
}

func (h *PositionHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"msg": "success DELETE available positions",
	})
}
