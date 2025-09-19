package repository

import (
	"context"
	"time"

	"github.com/Vis7044/GinCrud2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestRepository struct {
	collection *mongo.Collection
}

func NewTestRepository(db *mongo.Database) *TestRepository {
	return &TestRepository{
		collection: db.Collection("tests"),
	}
}

func (r *TestRepository) Create(test models.Test) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()
	return r.collection.InsertOne(ctx,test)
}

func (r *TestRepository) GetAll(limit , skip int) (*[]models.Test, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64(skip))

	cursor, err := r.collection.Find(ctx,bson.M{},findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tests[] models.Test
	if err = cursor.All(ctx, &tests); err != nil {
		return nil, err
	}
	return &tests , nil
}

func (r *TestRepository) GetOne(id primitive.ObjectID) (*models.Test, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var test models.Test
	err := r.collection.FindOne(ctx,bson.M{"_id": id}).Decode(&test)
	if err != nil {
		return nil, err
	}
	return &test , nil
}

func (r *TestRepository) UpdateOne(id primitive.ObjectID, updateTest models.Test) (*int64,error) {
	
	update := bson.M{"$set": bson.M{
		"title":  updateTest.Title,
		"description": updateTest.Description,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.UpdateByID(ctx, id, update)
	if err != nil || result.MatchedCount == 0 {
		return nil, err
	}
	return &result.MatchedCount, err
}

func (r *TestRepository) DeleteTest(id primitive.ObjectID) (*int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		return nil, err
	}
	return &result.DeletedCount,nil
}



