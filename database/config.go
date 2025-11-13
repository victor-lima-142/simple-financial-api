package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	ProjectionCollection  *mongo.Collection
	CostCollection        *mongo.Collection
	ContributorCollection *mongo.Collection
	ScenarioCollection    *mongo.Collection
)

func InitMongoDB(uri, dbName string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	ProjectionCollection = db.Collection("projections")
	CostCollection = db.Collection("costs")
	ContributorCollection = db.Collection("contributors")
	ScenarioCollection = db.Collection("scenarios")
	InitRepositories()
	return client, nil
}
