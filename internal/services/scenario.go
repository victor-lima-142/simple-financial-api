package services

import (
	"context"

	"github.com/victor-lima-142/simple-financial-api/database"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScenarioService interface {
	CreateScenario(ctx context.Context, write *models.ScenarioWrite) (*models.ScenarioRead, error)
	UpdateScenario(ctx context.Context, write *models.ScenarioWrite) (*models.ScenarioRead, error)
	DeleteScenario(ctx context.Context, id primitive.ObjectID) error
	GetScenarios(ctx context.Context, scenarioID primitive.ObjectID) ([]*models.ScenarioRead, error)
	GetScenarioByID(ctx context.Context, scenarioID primitive.ObjectID) (*models.ScenarioRead, error)
}

type scenarioServiceImpl struct{}

func NewScenarioService() ScenarioService {
	return &scenarioServiceImpl{}
}

func (s *scenarioServiceImpl) CreateScenario(ctx context.Context, write *models.ScenarioWrite) (*models.ScenarioRead, error) {
	return database.ScenarioRepo.Create(ctx, write)
}

func (s *scenarioServiceImpl) UpdateScenario(ctx context.Context, write *models.ScenarioWrite) (*models.ScenarioRead, error) {
	return database.ScenarioRepo.Update(ctx, write.ToModel().GetID(), *write.ToBson())
}

func (s *scenarioServiceImpl) DeleteScenario(ctx context.Context, id primitive.ObjectID) error {
	return database.ScenarioRepo.Delete(ctx, id)
}

func (s *scenarioServiceImpl) GetScenarios(ctx context.Context, scenarioID primitive.ObjectID) ([]*models.ScenarioRead, error) {
	return database.ScenarioRepo.FindAll(ctx)
}

func (s *scenarioServiceImpl) GetScenarioByID(ctx context.Context, scenarioID primitive.ObjectID) (*models.ScenarioRead, error) {
	return database.ScenarioRepo.FindByID(ctx, scenarioID)
}
