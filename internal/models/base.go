package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type WriteConvertible[T any] interface {
	ToModel() *T
	ToBson() *bson.M
}

type ReadConvertible[R any] interface {
	ToRead() *R
}

type BaseModel struct {
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type HasTimestamps interface {
	SetTimestamps(isNew bool)
	GetID() primitive.ObjectID
}

func (b *BaseModel) SetTimestamps(isNew bool) {
	now := time.Now().UTC()
	if isNew {
		b.CreatedAt = now
	}
	b.UpdatedAt = now
}
