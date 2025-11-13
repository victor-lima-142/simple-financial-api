package services

import (
	"context"

	"github.com/victor-lima-142/simple-financial-api/database"
	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ContributorService interface {
	CreateContributor(ctx context.Context, write *models.ContributorWrite) (*models.ContributorRead, error)
	UpdateContributor(ctx context.Context, write *models.ContributorWrite) (*models.ContributorRead, error)
	DeleteContributor(ctx context.Context, id primitive.ObjectID) error
	GetContributorFromScenario(ctx context.Context, scenarioID primitive.ObjectID) ([]*models.ContributorRead, error)
	GetContributorByID(ctx context.Context, contributorID primitive.ObjectID) (*models.ContributorRead, error)
}

type contributorServiceImpl struct{}

func NewContributorService() ContributorService {
	return &contributorServiceImpl{}
}

func (s *contributorServiceImpl) CreateContributor(ctx context.Context, write *models.ContributorWrite) (*models.ContributorRead, error) {
	return database.ContributorRepo.Create(ctx, write)
}

func (s *contributorServiceImpl) UpdateContributor(ctx context.Context, write *models.ContributorWrite) (*models.ContributorRead, error) {
	return database.ContributorRepo.Update(ctx, write.ToModel().GetID(), *write.ToBson())
}

func (s *contributorServiceImpl) DeleteContributor(ctx context.Context, id primitive.ObjectID) error {
	return database.ContributorRepo.Delete(ctx, id)
}

func (s *contributorServiceImpl) GetContributorFromScenario(ctx context.Context, scenarioID primitive.ObjectID) ([]*models.ContributorRead, error) {
	return database.ContributorRepo.FindAllBy(ctx, bson.M{"scenario": scenarioID})
}

func (s *contributorServiceImpl) GetContributorByID(ctx context.Context, contributorID primitive.ObjectID) (*models.ContributorRead, error) {
	return database.ContributorRepo.FindByID(ctx, contributorID)
}
