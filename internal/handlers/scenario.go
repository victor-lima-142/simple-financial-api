package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type ScenarioHandler interface {
	Initialize(engine *gin.Engine)
}

type scenarioHandlerImpl struct {
	controller controllers.ScenarioController
}

func NewScenarioHandler() ScenarioHandler {
	controller := controllers.NewScenarioController()
	return &scenarioHandlerImpl{
		controller: controller,
	}
}

func (h *scenarioHandlerImpl) Initialize(engine *gin.Engine) {
	scenarioGroup := engine.Group("/scenarios")
	{
		scenarioGroup.POST("", h.controller.CreateScenario)
		scenarioGroup.PUT("/:id", h.controller.UpdateScenario)
		scenarioGroup.DELETE("/:id", h.controller.DeleteScenario)
		scenarioGroup.GET("", h.controller.GetScenarios)
		scenarioGroup.GET("/:id", h.controller.GetScenarioByID)
		scenarioGroup.GET("/costs/:id", h.controller.GetCostFromScenario)
		scenarioGroup.GET("/contributors/:id", h.controller.GetContributorFromScenario)
	}
}
