package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"github.com/victor-lima-142/simple-financial-api/internal/services"
	"github.com/victor-lima-142/simple-financial-api/internal/utils"
)

type ScenarioController interface {
	CreateScenario(ctx *gin.Context)
	UpdateScenario(ctx *gin.Context)
	DeleteScenario(ctx *gin.Context)
	GetScenarios(ctx *gin.Context)
	GetContributorFromScenario(ctx *gin.Context)
	GetCostFromScenario(ctx *gin.Context)
	GetScenarioByID(ctx *gin.Context)
}

type scenarioControllerImpl struct {
	service services.ScenarioService
}

func NewScenarioController() ScenarioController {
	service := services.NewScenarioService()
	return &scenarioControllerImpl{
		service: service,
	}
}

func (c *scenarioControllerImpl) CreateScenario(ctx *gin.Context) {
	var write models.ScenarioWrite
	if err := ctx.ShouldBindJSON(&write); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := c.service.CreateScenario(ctx, &write)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": created,
	})
}

func (c *scenarioControllerImpl) UpdateScenario(ctx *gin.Context) {
	var write models.ScenarioWrite
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of scenario should be provided"),
		})
		return
	}
	if err := ctx.ShouldBindJSON(&write); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	objectID, err := utils.StringToObjectID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	write.ID = &objectID
	c.service.UpdateScenario(ctx, &write)
}

func (c *scenarioControllerImpl) DeleteScenario(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of scenario should be provided"),
		})
		return
	}
	objectID, err := utils.StringToObjectID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.service.DeleteScenario(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *scenarioControllerImpl) GetScenarios(ctx *gin.Context) {
	scenarios, err := c.service.GetScenarios(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": scenarios,
	})
}

func (c *scenarioControllerImpl) GetContributorFromScenario(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of contributor should be provided"),
		})
		return
	}
	objectID, err := utils.StringToObjectID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	contributors, err := c.service.GetContributorFromScenario(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": contributors,
	})
}

func (c *scenarioControllerImpl) GetCostFromScenario(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of cost should be provided"),
		})
		return
	}
	objectID, err := utils.StringToObjectID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	costs, err := c.service.GetCostFromScenario(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": costs,
	})
}

func (c *scenarioControllerImpl) GetScenarioByID(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of scenario should be provided"),
		})
		return
	}
	objectID, err := utils.StringToObjectID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	scenario, err := c.service.GetScenarioByID(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": scenario,
	})
}
