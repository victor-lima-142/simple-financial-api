package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CostWrite struct {
	ID          *primitive.ObjectID `json:"_id,omitempty"`
	Name        string              `json:"name" validate:"required,min=1"`
	Description string              `json:"description,omitempty"`
	Value       float64             `json:"value" validate:"required,gte=0"`
	Scenario    primitive.ObjectID  `json:"scenario" validate:"required"`
}

type CostRead struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	Value       float64            `json:"value"`
	Scenario    primitive.ObjectID `json:"scenario"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

type Cost struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Value       float64            `bson:"value" json:"value"`
	Scenario    primitive.ObjectID `bson:"scenario" json:"scenario"`
	BaseModel   `bson:",inline"`
}

func (w *CostWrite) ToModel() *Cost {
	if w.ID != nil {
		return &Cost{
			ID:          *w.ID,
			Name:        w.Name,
			Description: w.Description,
			Value:       w.Value,
			Scenario:    w.Scenario,
		}
	}
	return &Cost{
		Name:        w.Name,
		Description: w.Description,
		Value:       w.Value,
		Scenario:    w.Scenario,
	}
}

func (w *CostWrite) ToBson() *bson.M {
	return &bson.M{
		"name":        w.Name,
		"description": w.Description,
		"value":       w.Value,
		"scenario":    w.Scenario,
	}
}

func (c *Cost) ToRead() *CostRead {
	return &CostRead{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		Value:       c.Value,
		Scenario:    c.Scenario,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c *Cost) GetID() primitive.ObjectID {
	return c.ID
}
