package handler

import (
    "context"
    "user-crud-service/internal/model"
)

type UserService interface {
    GetUser(ctx context.Context, userID string) (*model.User, error)
    GetUsersList(ctx context.Context, limit, offset int32) ([]*model.User, error)
}
