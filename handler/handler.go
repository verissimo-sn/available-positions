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
	id := ctx.Param("id")
	request := UpdatePositionDto{}
	ctx.BindJSON(&request)
	position := schemas.AvailablePosition{}
	if err := h.db.First(&position, id).Error; err != nil {
		h.logger.Errorf("Error on update position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusNotFound, "Position not found")
		return
	}

	if request.Role != "" {
		position.Role = request.Role
	}
	if request.Tech != "" {
		position.Tech = request.Tech
	}
	if request.Level != "" {
		position.Level = request.Level
	}
	if request.Company != "" {
		position.Company = request.Company
	}
	if request.Location != "" {
		position.Location = request.Location
	}
	if request.Salary != 0 {
		position.Salary = request.Salary
	}
	if request.Link != "" {
		position.Link = request.Link
	}

	if err := h.db.Save(&position).Error; err != nil {
		h.logger.Errorf("Error on update position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}
	h.httpFormatter.Response(ctx, http.StatusNoContent, nil)
}

func (h *PositionHandler) Get(ctx *gin.Context) {
	positions := []schemas.AvailablePosition{}
	if err := h.db.Find(&positions).Error; err != nil {
		h.logger.Errorf("Error on get position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.httpFormatter.Response(ctx, http.StatusOK, mapAvailablePositionList(positions))
}

func (h *PositionHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	position := schemas.AvailablePosition{}
	if err := h.db.First(&position, id).Error; err != nil {
		h.logger.Errorf("Error on get position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusNotFound, "Position not found")
		return
	}
	h.httpFormatter.Response(ctx, http.StatusOK, mapAvailablePosition(position))
}

func (h *PositionHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := h.db.First(&schemas.AvailablePosition{}, id).Error; err != nil {
		h.logger.Errorf("Error on delete position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusNotFound, "Position not found")
		return
	}

	if err := h.db.Delete(&schemas.AvailablePosition{}, id).Error; err != nil {
		h.logger.Errorf("Error on delete position: %v", err)
		h.httpFormatter.Response(ctx, http.StatusBadRequest, err.Error())
		return
	}
	h.httpFormatter.Response(ctx, http.StatusNoContent, nil)
}

func mapAvailablePosition(data schemas.AvailablePosition) schemas.AvailablePositionResponse {
	return schemas.AvailablePositionResponse{
		Id:        data.ID,
		Role:      data.Role,
		Tech:      data.Tech,
		Level:     data.Level,
		Company:   data.Company,
		Location:  data.Location,
		Salary:    data.Salary,
		Link:      data.Link,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func mapAvailablePositionList(data []schemas.AvailablePosition) []schemas.AvailablePositionResponse {
	var apResponse []schemas.AvailablePositionResponse
	for _, item := range data {
		apResponse = append(apResponse, mapAvailablePosition(item))
	}
	return apResponse
}
