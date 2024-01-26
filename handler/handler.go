package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/verissimo-sn/available-positions/config"
	"github.com/verissimo-sn/available-positions/schemas"
	"gorm.io/gorm"
)

type PositionHandler struct {
	logger        *config.Logger
	db            *gorm.DB
	httpFormatter *HttpFormatter
}

func NewPositionHandler() *PositionHandler {
	return &PositionHandler{
		logger:        config.GetLogger("position handler"),
		db:            config.GetSQLite(),
		httpFormatter: NewHttpFormatter(),
	}
}

func (h *PositionHandler) Create(ctx *gin.Context) {
	request := CretePositionDto{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		h.logger.Errorf("Validation error: %v", err)
		h.httpFormatter.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ap := schemas.AvailablePosition{
		Role:     request.Role,
		Tech:     request.Tech,
		Level:    request.Level,
		Company:  request.Company,
		Location: request.Location,
		Salary:   request.Salary,
		Link:     request.Link,
	}

	if err := h.db.Create(&ap).Error; err != nil {
		h.logger.Errorf("Error on create position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}
	h.httpFormatter.Response(ctx, http.StatusCreated, nil)
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
