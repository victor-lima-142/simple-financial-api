package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type ProjectionHandler interface {
	Initialize(engine *gin.Engine)
}

type projectionHandlerImpl struct {
	controller controllers.ProjectionController
}

func NewProjectionHandler() ProjectionHandler {
	controller := controllers.NewProjectionController()
	return &projectionHandlerImpl{
		controller: controller,
	}
}

func (h *projectionHandlerImpl) Initialize(engine *gin.Engine) {
	projectionGroup := engine.Group("/projections")
	{
		projectionGroup.POST("", h.controller.CreateProjection)
		projectionGroup.PUT("/:id", h.controller.UpdateProjection)
		projectionGroup.DELETE("/:id", h.controller.DeleteProjection)
		projectionGroup.GET("", h.controller.GetProjections)
		projectionGroup.GET("/:id", h.controller.GetProjectionByID)
	}
}
