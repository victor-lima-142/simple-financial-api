package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"github.com/victor-lima-142/simple-financial-api/internal/services"
	"github.com/victor-lima-142/simple-financial-api/internal/utils"
)

type ProjectionController interface {
	CreateProjection(ctx *gin.Context)
	UpdateProjection(ctx *gin.Context)
	DeleteProjection(ctx *gin.Context)
	GetProjections(ctx *gin.Context)
	GetProjectionByID(ctx *gin.Context)
}

type projectionControllerImpl struct {
	service services.ProjectionService
}

func NewProjectionController() ProjectionController {
	service := services.NewProjectionService()
	return &projectionControllerImpl{
		service: service,
	}
}

func (c *projectionControllerImpl) CreateProjection(ctx *gin.Context) {
	var write models.ProjectionWrite
	if err := ctx.ShouldBindJSON(&write); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := c.service.CreateProjection(ctx, &write)
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

func (c *projectionControllerImpl) UpdateProjection(ctx *gin.Context) {
	var write models.ProjectionWrite
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of projection should be provided"),
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
	c.service.UpdateProjection(ctx, &write)
}

func (c *projectionControllerImpl) DeleteProjection(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of projection should be provided"),
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
	err = c.service.DeleteProjection(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *projectionControllerImpl) GetProjections(ctx *gin.Context) {
	projections, err := c.service.GetProjections(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": projections,
	})
}

func (c *projectionControllerImpl) GetProjectionByID(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of projection should be provided"),
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
	projection, err := c.service.GetProjectionByID(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": projection,
	})
}
