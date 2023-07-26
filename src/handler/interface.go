package handler

import (
	"context"
	"user-crud-service/model"
)

type UserService interface {
	AddUsers(ctx context.Context, user []*model.User) error
	GetUser(ctx context.Context, userID string) (*model.User, error)
	GetUsersList(ctx context.Context) ([]model.UserInfo, error)
	UpdateUser(ctx context.Context, userId string, userNewInfo *model.UpdateUser) (*model.User, error)
	DeleteUser(ctx context.Context, userId string) error
}
