package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type ContributorHandler interface {
	Initialize(engine *gin.Engine)
}

type contributorHandlerImpl struct {
	controller controllers.ContributorController
}

func NewContributorHandler() ContributorHandler {
	controller := controllers.NewContributorController()
	return &contributorHandlerImpl{
		controller: controller,
	}
}

func (h *contributorHandlerImpl) Initialize(engine *gin.Engine) {
	contributorGroup := engine.Group("/contributors")
	{
		contributorGroup.POST("", h.controller.CreateContributor)
		contributorGroup.PUT("/:id", h.controller.UpdateContributor)
		contributorGroup.DELETE("/:id", h.controller.DeleteContributor)
		contributorGroup.GET("/:id", h.controller.GetContributorByID)
	}
}
