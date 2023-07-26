package service

import (
	"context"
	"user-crud-service/model"
)

type Repository interface {
	AddUsers(ctx context.Context, users []*model.User) error
	AddUser(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, id string) (*model.User, error)
	GetUsersList(ctx context.Context) ([]model.UserInfo, error)
	UpdateUser(ctx context.Context, id string, userNewInfo *model.UpdateUser) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}
