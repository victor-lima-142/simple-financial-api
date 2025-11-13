package services

import (
	"context"

	"github.com/victor-lima-142/simple-financial-api/database"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectionService interface {
	CreateProjection(ctx context.Context, write *models.ProjectionWrite) (*models.ProjectionRead, error)
	UpdateProjection(ctx context.Context, write *models.ProjectionWrite) (*models.ProjectionRead, error)
	DeleteProjection(ctx context.Context, id primitive.ObjectID) error
	GetProjections(ctx context.Context) ([]*models.ProjectionRead, error)
	GetProjectionByID(ctx context.Context, projectionID primitive.ObjectID) (*models.ProjectionRead, error)
}

type projectionServiceImpl struct{}

func NewProjectionService() ProjectionService {
	return &projectionServiceImpl{}
}

func (s *projectionServiceImpl) CreateProjection(ctx context.Context, write *models.ProjectionWrite) (*models.ProjectionRead, error) {
	return database.ProjectionRepo.Create(ctx, write)
}

func (s *projectionServiceImpl) UpdateProjection(ctx context.Context, write *models.ProjectionWrite) (*models.ProjectionRead, error) {
	return database.ProjectionRepo.Update(ctx, write.ToModel().GetID(), *write.ToBson())
}

func (s *projectionServiceImpl) DeleteProjection(ctx context.Context, id primitive.ObjectID) error {
	return database.ProjectionRepo.Delete(ctx, id)
}

func (s *projectionServiceImpl) GetProjections(ctx context.Context) ([]*models.ProjectionRead, error) {
	return database.ProjectionRepo.FindAll(ctx)
}

func (s *projectionServiceImpl) GetProjectionByID(ctx context.Context, projectionID primitive.ObjectID) (*models.ProjectionRead, error) {
	return database.ProjectionRepo.FindByID(ctx, projectionID)
}
