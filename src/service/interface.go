package service

import (
	"context"
	"user-crud-service/model"
)

type Repository interface {
	GetById(ctx context.Context, id string) (*model.User, error)
	GetUsersList(ctx context.Context) ([]model.User, error)
	AddUsers(ctx context.Context, users []*model.User) ([]*model.User, error)
	AddUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}
