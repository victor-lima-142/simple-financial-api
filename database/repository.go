package database

import "github.com/victor-lima-142/simple-financial-api/internal/models"

var (
	ProjectionRepo  *BaseRepository[models.Projection, *models.ProjectionWrite, models.ProjectionRead, *models.Projection]
	CostRepo        *BaseRepository[models.Cost, *models.CostWrite, models.CostRead, *models.Cost]
	ContributorRepo *BaseRepository[models.Contributor, *models.ContributorWrite, models.ContributorRead, *models.Contributor]
	ScenarioRepo    *BaseRepository[models.Scenario, *models.ScenarioWrite, models.ScenarioRead, *models.Scenario]
)

func InitRepositories() {
	ProjectionRepo = &BaseRepository[models.Projection, *models.ProjectionWrite, models.ProjectionRead, *models.Projection]{
		Collection: ProjectionCollection,
	}
	CostRepo = &BaseRepository[models.Cost, *models.CostWrite, models.CostRead, *models.Cost]{
		Collection: CostCollection,
	}
	ContributorRepo = &BaseRepository[models.Contributor, *models.ContributorWrite, models.ContributorRead, *models.Contributor]{
		Collection: ContributorCollection,
	}
	ScenarioRepo = &BaseRepository[models.Scenario, *models.ScenarioWrite, models.ScenarioRead, *models.Scenario]{
		Collection: ScenarioCollection,
	}
}
