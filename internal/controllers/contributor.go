package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"github.com/victor-lima-142/simple-financial-api/internal/services"
	"github.com/victor-lima-142/simple-financial-api/internal/utils"
)

type ContributorController interface {
	CreateContributor(ctx *gin.Context)
	UpdateContributor(ctx *gin.Context)
	DeleteContributor(ctx *gin.Context)
	GetContributorByID(ctx *gin.Context)
}

type contributorControllerImpl struct {
	service services.ContributorService
}

func NewContributorController() ContributorController {
	service := services.NewContributorService()
	return &contributorControllerImpl{
		service: service,
	}
}

func (c *contributorControllerImpl) CreateContributor(ctx *gin.Context) {
	var write models.ContributorWrite
	if err := ctx.ShouldBindJSON(&write); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := c.service.CreateContributor(ctx, &write)
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

func (c *contributorControllerImpl) UpdateContributor(ctx *gin.Context) {
	var write models.ContributorWrite
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("id of contributor should be provided"),
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
	c.service.UpdateContributor(ctx, &write)
}

func (c *contributorControllerImpl) DeleteContributor(ctx *gin.Context) {
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
	err = c.service.DeleteContributor(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (c *contributorControllerImpl) GetContributorByID(ctx *gin.Context) {
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
	contributor, err := c.service.GetContributorByID(ctx, objectID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": contributor,
	})
}
