package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"github.com/victor-lima-142/simple-financial-api/internal/services"
	"github.com/victor-lima-142/simple-financial-api/internal/utils"
)

type CostController interface {
	CreateCost(ctx *gin.Context)
	UpdateCost(ctx *gin.Context)
	DeleteCost(ctx *gin.Context)
	GetCostByID(ctx *gin.Context)
}

type costControllerImpl struct {
	service services.CostService
}

func NewCostController() CostController {
	service := services.NewCostService()
	return &costControllerImpl{
		service: service,
	}
}

func (c *costControllerImpl) CreateCost(ctx *gin.Context) {
	var write models.CostWrite
	if err := ctx.ShouldBindJSON(&write); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := c.service.CreateCost(ctx, &write)
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

func (c *costControllerImpl) UpdateCost(ctx *gin.Context) {
	var write models.CostWrite
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of cost should be provided"),
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
	c.service.UpdateCost(ctx, &write)
}

func (c *costControllerImpl) DeleteCost(ctx *gin.Context) {
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
	err = c.service.DeleteCost(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *costControllerImpl) GetCostByID(ctx *gin.Context) {
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
	cost, err := c.service.GetCostByID(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": cost,
	})
}
