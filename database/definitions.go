package database

import (
	"context"
	"errors"
	"time"

	"github.com/victor-lima-142/simple-financial-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BaseRepository[M any, W models.WriteConvertible[M], R any, P models.ReadConvertible[R]] struct {
	Collection *mongo.Collection
}

func (r *BaseRepository[M, W, R, P]) Create(ctx context.Context, dto W) (*R, error) {
	model := (dto).ToModel()

	switch m := any(&model).(type) {
	case models.HasTimestamps:
		m.SetTimestamps(true)
	default:
		if t, ok := any(&model).(interface{ SetTimestamps(bool) }); ok {
			t.SetTimestamps(true)
		}
	}

	_, err := r.Collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	return r.FindLastInserted(ctx, *model)
}

func (r *BaseRepository[M, W, R, P]) FindByID(ctx context.Context, id primitive.ObjectID) (*R, error) {
	var model M
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&model)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	readable := any(&model).(P)
	read := readable.ToRead()
	return read, nil
}

func (r *BaseRepository[M, W, R, P]) FindAll(ctx context.Context) ([]*R, error) {
	cur, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var modelsList []M
	if err := cur.All(ctx, &modelsList); err != nil {
		return nil, err
	}

	results := make([]*R, 0, len(modelsList))
	for _, m := range modelsList {
		read := any(&m).(P).ToRead()
		results = append(results, read)
	}

	return results, nil
}

func (r *BaseRepository[M, W, R, P]) Update(ctx context.Context, id primitive.ObjectID, update bson.M) (*R, error) {
	update["updatedAt"] = time.Now().UTC()
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	result, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepository[M, W, R, P]) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *BaseRepository[M, W, R, P]) FindLastInserted(ctx context.Context, model M) (*R, error) {
	idField, ok := any(model).(interface{ GetID() primitive.ObjectID })
	if !ok {
		return nil, errors.New("model does not expose GetID()")
	}

	id := idField.GetID()
	return r.FindByID(ctx, id)
}

func (r *BaseRepository[M, W, R, P]) FindOneBy(ctx context.Context, filter bson.M) (*R, error) {
	var model M
	err := r.Collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	readable, ok := any(&model).(P)
	if !ok {
		return nil, errors.New("model does not implement ReadConvertible")
	}

	read := readable.ToRead()
	return read, nil
}

func (r *BaseRepository[M, W, R, P]) FindAllBy(ctx context.Context, filter bson.M) ([]*R, error) {
	cur, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var modelsList []M
	if err := cur.All(ctx, &modelsList); err != nil {
		return nil, err
	}

	results := make([]*R, 0, len(modelsList))
	for _, m := range modelsList {
		readable, ok := any(&m).(P)
		if !ok {
			return nil, errors.New("model does not implement ReadConvertible")
		}
		read := readable.ToRead()
		results = append(results, read)
	}

	return results, nil
}
