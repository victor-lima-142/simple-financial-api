package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type ProjectionHandler interface {
	Initialize()
}

type projectionHandlerImpl struct {
	controller controllers.ProjectionController
	engine     *gin.Engine
}

func NewProjectionHandler(engine *gin.Engine) ProjectionHandler {
	controller := controllers.NewProjectionController()
	return &projectionHandlerImpl{
		controller: controller,
		engine:     engine,
	}
}

func (h *projectionHandlerImpl) Initialize() {
	projectionGroup := h.engine.Group("/projections")
	{
		projectionGroup.POST("", h.controller.CreateProjection)
		projectionGroup.PUT("/:id", h.controller.UpdateProjection)
		projectionGroup.DELETE("/:id", h.controller.DeleteProjection)
		projectionGroup.GET("", h.controller.GetProjections)
		projectionGroup.GET("/:id", h.controller.GetProjectionByID)
	}
}
