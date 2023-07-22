package service

import (
    "context"
    "user-crud-service/internal/model"
)

type Repository interface {
    GetById(ctx context.Context, id string) (*model.User, error)
    GetUsersList(ctx context.Context, limit, offset int32) ([]*model.User, error)
}
