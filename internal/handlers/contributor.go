package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type ContributorHandler interface {
	Initialize()
}

type contributorHandlerImpl struct {
	controller controllers.ContributorController
	engine     *gin.Engine
}

func NewContributorHandler(engine *gin.Engine) ContributorHandler {
	controller := controllers.NewContributorController()
	return &contributorHandlerImpl{
		controller: controller,
		engine:     engine,
	}
}

func (h *contributorHandlerImpl) Initialize() {
	contributorGroup := h.engine.Group("/contributors")
	{
		contributorGroup.POST("", h.controller.CreateContributor)
		contributorGroup.PUT("/:id", h.controller.UpdateContributor)
		contributorGroup.DELETE("/:id", h.controller.DeleteContributor)
		contributorGroup.GET("/:id", h.controller.GetContributorByID)
	}
}
