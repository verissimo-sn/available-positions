package router

import (
	"github.com/gin-gonic/gin"
	"github.com/verissimo-sn/available-positions/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	availablePositionsRoutes(v1)
}

func availablePositionsRoutes(routerGroup *gin.RouterGroup) {
	apHandler := handler.NewPositionHandler()
	routerGroup.POST("/available-positions", apHandler.Create)
	routerGroup.GET("/available-positions", apHandler.Get)
	routerGroup.GET("/available-positions/:id", apHandler.GetById)
	routerGroup.PUT("/available-positions/:id", apHandler.Update)
	routerGroup.DELETE("/available-positions/:id", apHandler.Delete)
}
