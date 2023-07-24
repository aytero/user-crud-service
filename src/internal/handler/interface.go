package handler

import (
	"context"
	"user-crud-service/model"
)

type UserService interface {
	GetUser(ctx context.Context, userID string) (*model.User, error)
	GetUsersList(ctx context.Context) ([]model.User, error)
	AddUsers(ctx context.Context, user []*model.User) ([]*model.User, error)
	UpdateUser(ctx context.Context, userId string, userNewInfo *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, userId string) error
}
