package services

import (
	"context"

	"github.com/Subasri-V/long-authentication/dal/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	ctx        context.Context
	collection *mongo.Collection
	client     *mongo.Client
}

func InitUserService(ctx context.Context, collection *mongo.Collection, client *mongo.Client) interfaces.IUser {
	return &UserService{ctx, collection, client}
}


// SignInUser implements interfaces.IUser.
func (*UserService) SignInUser() {
	panic("unimplemented")
}

// SignUpUser implements interfaces.IUser.
func (*UserService) SignUpUser() {
	panic("unimplemented")
}

// VerifyEmail implements interfaces.IUser.
func (*UserService) VerifyEmail() {
	panic("unimplemented")
}
