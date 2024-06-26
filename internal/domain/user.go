package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           primitive.ObjectID
	Username     string
	PasswordHash string
}

type UserRepository interface {
	Store(context.Context, *User) error
	GetByUsername(context.Context, string) (*User, error)
}

func NewUser(username string, password string) (*User, error) {
	passwordHash, err := hashPassword(password)
	if err == nil {
		return nil, err
	}

	return &User{
		Id:           primitive.NewObjectID(),
		Username:     username,
		PasswordHash: passwordHash,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
