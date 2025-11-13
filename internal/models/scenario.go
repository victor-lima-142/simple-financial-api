package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ScenarioWrite struct {
	ID           *primitive.ObjectID  `json:"_id,omitempty"`
	Name         string               `json:"name" validate:"required,min=1"`
	Description  *string              `json:"description,omitempty" validate:"omitempty,max=500"`
	Contributors []primitive.ObjectID `json:"contributors,omitempty"`
	Costs        []primitive.ObjectID `json:"costs,omitempty"`
}

type ScenarioRead struct {
	ID           primitive.ObjectID   `json:"_id"`
	Name         string               `json:"name"`
	Description  *string              `json:"description,omitempty"`
	Contributors []primitive.ObjectID `json:"contributors,omitempty"`
	Costs        []primitive.ObjectID `json:"costs,omitempty"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
}

type Scenario struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string               `bson:"name" json:"name"`
	Description  *string              `bson:"description,omitempty" json:"description,omitempty"`
	Contributors []primitive.ObjectID `bson:"contributors,omitempty" json:"contributors,omitempty"`
	Costs        []primitive.ObjectID `bson:"costs,omitempty" json:"costs,omitempty"`
	BaseModel    `bson:",inline"`
}

func (w *ScenarioWrite) ToModel() *Scenario {
	if w.ID != nil {
		return &Scenario{
			ID:           *w.ID,
			Name:         w.Name,
			Description:  w.Description,
			Contributors: w.Contributors,
			Costs:        w.Costs,
		}
	}
	return &Scenario{
		Name:         w.Name,
		Description:  w.Description,
		Contributors: w.Contributors,
		Costs:        w.Costs,
	}
}

func (w *ScenarioWrite) ToBson() *bson.M {
	return &bson.M{
		"name":         w.Name,
		"description":  w.Description,
		"contributors": w.Contributors,
		"costs":        w.Costs,
	}
}

func (s *Scenario) ToRead() *ScenarioRead {
	return &ScenarioRead{
		ID:           s.ID,
		Name:         s.Name,
		Description:  s.Description,
		Contributors: s.Contributors,
		Costs:        s.Costs,
		CreatedAt:    s.CreatedAt,
		UpdatedAt:    s.UpdatedAt,
	}
}

func (s *Scenario) GetID() primitive.ObjectID {
	return s.ID
}
