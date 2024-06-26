package repository

import (
	"context"
	"github.com/taranovegor/jurnalo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type userRepository struct {
	domain.UserRepository
	col *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) domain.UserRepository {
	return userRepository{
		col: col,
	}
}

func (r userRepository) Store(ctx context.Context, user *domain.User) error {
	_, err := r.col.InsertOne(ctx, user)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return err
}

func (r userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.col.FindOne(ctx, bson.D{{"username", username}}).Decode(&user)

	return &user, err
}
