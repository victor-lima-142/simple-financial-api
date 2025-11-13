package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/controllers"
)

type ScenarioHandler interface {
	Initialize()
}

type scenarioHandlerImpl struct {
	controller controllers.ScenarioController
	engine     *gin.Engine
}

func NewScenarioHandler(engine *gin.Engine) ScenarioHandler {
	controller := controllers.NewScenarioController()
	return &scenarioHandlerImpl{
		controller: controller,
		engine:     engine,
	}
}

func (h *scenarioHandlerImpl) Initialize() {
	scenarioGroup := h.engine.Group("/scenarios")
	{
		scenarioGroup.POST("", h.controller.CreateScenario)
		scenarioGroup.PUT("/:id", h.controller.UpdateScenario)
		scenarioGroup.DELETE("/:id", h.controller.DeleteScenario)
		scenarioGroup.GET("", h.controller.GetScenarios)
		scenarioGroup.GET("/:id", h.controller.GetScenarioByID)
		scenarioGroup.GET("/:id/costs", h.controller.GetCostFromScenario)
		scenarioGroup.GET("/:id/contributors", h.controller.GetContributorFromScenario)
	}
}
