package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ContributorWrite struct {
	ID          *primitive.ObjectID `json:"_id,omitempty"`
	Name        string              `json:"name" validate:"required,min=1"`
	Description string              `json:"description,omitempty"`
	Salary      float64             `json:"salary" validate:"required,gte=0"`
	Scenario    primitive.ObjectID  `json:"scenario" validate:"required"`
}

type ContributorRead struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Salary      float64            `json:"salary"`
	Scenario    primitive.ObjectID `json:"scenario"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

type Contributor struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Salary      float64            `bson:"salary" json:"salary"`
	Scenario    primitive.ObjectID `bson:"scenario" json:"scenario"`
	BaseModel   `bson:",inline"`
}

func (w *ContributorWrite) ToModel() *Contributor {
	if w.ID != nil {
		return &Contributor{
			ID:          *w.ID,
			Name:        w.Name,
			Description: w.Description,
			Salary:      w.Salary,
			Scenario:    w.Scenario,
		}
	}
	return &Contributor{
		Name:        w.Name,
		Description: w.Description,
		Salary:      w.Salary,
		Scenario:    w.Scenario,
	}
}

func (w *ContributorWrite) ToBson() *bson.M {
	return &bson.M{
		"name":        w.Name,
		"description": w.Description,
		"salary":      w.Salary,
		"scenario":    w.Scenario,
	}
}

func (c *Contributor) ToRead() *ContributorRead {
	return &ContributorRead{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Salary:      c.Salary,
		Scenario:    c.Scenario,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c *Contributor) GetID() primitive.ObjectID {
	return c.ID
}
