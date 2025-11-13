package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type CostHandler interface {
	Initialize(engine *gin.Engine)
}

type costHandlerImpl struct {
	controller controllers.CostController
}

func NewCostHandler() CostHandler {
	controller := controllers.NewCostController()
	return &costHandlerImpl{
		controller: controller,
	}
}

func (h *costHandlerImpl) Initialize(engine *gin.Engine) {
	costGroup := engine.Group("/costs")
	{
		costGroup.POST("", h.controller.CreateCost)
		costGroup.PUT("/:id", h.controller.UpdateCost)
		costGroup.DELETE("/:id", h.controller.DeleteCost)
		costGroup.GET("/:id", h.controller.GetCostByID)
	}
}
