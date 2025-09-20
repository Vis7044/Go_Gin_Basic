package repository

import (
	"context"
	"time"

	"github.com/Vis7044/GinCrud2/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	collection *mongo.Collection
}

func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{
		collection: db.Collection("User"),
	}
}

func (r *AuthRepository) Register(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *AuthRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"email":email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
