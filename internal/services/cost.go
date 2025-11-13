package services

import (
	"context"

	"github.com/victor-lima-142/simple-financial-api/database"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CostService interface {
	CreateCost(ctx context.Context, write *models.CostWrite) (*models.CostRead, error)
	UpdateCost(ctx context.Context, write *models.CostWrite) (*models.CostRead, error)
	DeleteCost(ctx context.Context, id primitive.ObjectID) error
	GetCostByID(ctx context.Context, costID primitive.ObjectID) (*models.CostRead, error)
}

type costServiceImpl struct{}

func NewCostService() CostService {
	return &costServiceImpl{}
}

func (s *costServiceImpl) CreateCost(ctx context.Context, write *models.CostWrite) (*models.CostRead, error) {
	return database.CostRepo.Create(ctx, write)
}

func (s *costServiceImpl) UpdateCost(ctx context.Context, write *models.CostWrite) (*models.CostRead, error) {
	return database.CostRepo.Update(ctx, write.ToModel().GetID(), *write.ToBson())
}

func (s *costServiceImpl) DeleteCost(ctx context.Context, id primitive.ObjectID) error {
	return database.CostRepo.Delete(ctx, id)
}

func (s *costServiceImpl) GetCostByID(ctx context.Context, costID primitive.ObjectID) (*models.CostRead, error) {
	return database.CostRepo.FindByID(ctx, costID)
}
