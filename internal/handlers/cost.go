package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type CostHandler interface {
	Initialize()
}

type costHandlerImpl struct {
	controller controllers.CostController
	engine     *gin.Engine
}

func NewCostHandler(engine *gin.Engine) CostHandler {
	controller := controllers.NewCostController()
	return &costHandlerImpl{
		controller: controller,
		engine:     engine,
	}
}

func (h *costHandlerImpl) Initialize() {
	costGroup := h.engine.Group("/costs")
	{
		costGroup.POST("", h.controller.CreateCost)
		costGroup.PUT("/:id", h.controller.UpdateCost)
		costGroup.DELETE("/:id", h.controller.DeleteCost)
		costGroup.GET("/:id", h.controller.GetCostByID)
	}
}
