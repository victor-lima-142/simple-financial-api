package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProjectionWrite struct {
	ID                  *primitive.ObjectID `json:"_id,omitempty"`
	Name                string              `json:"name" validate:"required,min=1,max=100"`
	Description         string              `json:"description,omitempty" validate:"max=500"`
	InitialAmount       float64             `json:"initialAmount" validate:"required,gte=0"`
	AnnualInterestRate  float64             `json:"annualInterestRate" validate:"required,gte=0"`
	Months              int                 `json:"months" validate:"required,gt=0"`
	MonthlyContribution float64             `json:"monthlyContribution" validate:"required,gte=0"`
}

type ProjectionRead struct {
	ID                  primitive.ObjectID `json:"_id"`
	Name                string             `json:"name"`
	Description         string             `json:"description,omitempty"`
	InitialAmount       float64            `json:"initialAmount"`
	AnnualInterestRate  float64            `json:"annualInterestRate"`
	Months              int                `json:"months"`
	MonthlyContribution float64            `json:"monthlyContribution"`
	FinalAmount         float64            `json:"finalAmount"`
	TotalContributions  float64            `json:"totalContributions"`
	TotalInterestEarned float64            `json:"totalInterestEarned"`
	CreatedAt           time.Time          `json:"createdAt"`
	UpdatedAt           time.Time          `json:"updatedAt"`
}

type Projection struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name                string             `bson:"name" json:"name"`
	Description         string             `bson:"description,omitempty" json:"description,omitempty"`
	InitialAmount       float64            `bson:"initialAmount" json:"initialAmount"`
	AnnualInterestRate  float64            `bson:"annualInterestRate" json:"annualInterestRate"`
	Months              int                `bson:"months" json:"months"`
	MonthlyContribution float64            `bson:"monthlyContribution" json:"monthlyContribution"`
	FinalAmount         float64            `bson:"finalAmount" json:"finalAmount"`
	TotalContributions  float64            `bson:"totalContributions" json:"totalContributions"`
	TotalInterestEarned float64            `bson:"totalInterestEarned" json:"totalInterestEarned"`
	BaseModel           `bson:",inline"`
}

func (w *ProjectionWrite) ToModel() *Projection {
	if w.ID != nil {
		return &Projection{
			ID:                  *w.ID,
			Name:                w.Name,
			Description:         w.Description,
			InitialAmount:       w.InitialAmount,
			AnnualInterestRate:  w.AnnualInterestRate,
			Months:              w.Months,
			MonthlyContribution: w.MonthlyContribution,
		}
	}
	return &Projection{
		Name:                w.Name,
		Description:         w.Description,
		InitialAmount:       w.InitialAmount,
		AnnualInterestRate:  w.AnnualInterestRate,
		Months:              w.Months,
		MonthlyContribution: w.MonthlyContribution,
	}
}

func (w *ProjectionWrite) ToBson() *bson.M {
	return &bson.M{
		"name":                w.Name,
		"description":         w.Description,
		"initialAmount":       w.InitialAmount,
		"annualInterestRate":  w.AnnualInterestRate,
		"months":              w.Months,
		"monthlyContribution": w.MonthlyContribution,
	}
}

func (p *Projection) ToRead() *ProjectionRead {
	return &ProjectionRead{
		ID:                  p.ID,
		Name:                p.Name,
		Description:         p.Description,
		InitialAmount:       p.InitialAmount,
		AnnualInterestRate:  p.AnnualInterestRate,
		Months:              p.Months,
		MonthlyContribution: p.MonthlyContribution,
		FinalAmount:         p.FinalAmount,
		TotalContributions:  p.TotalContributions,
		TotalInterestEarned: p.TotalInterestEarned,
		CreatedAt:           p.CreatedAt,
		UpdatedAt:           p.UpdatedAt,
	}
}

func (p *Projection) GetID() primitive.ObjectID {
	return p.ID
}
