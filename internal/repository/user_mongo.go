package repository

import (
    "context"
    "user-crud-service/internal/database"
    "user-crud-service/internal/model"
)

type UserRepository struct {
    db *database.Database
}

func NewUserRepository(dbConn *database.Database) *UserRepository {
    return &UserRepository{
        db: dbConn,
    }
}

func (r *UserRepository) GetById(ctx context.Context, id string) (*model.User, error) {
    return &model.User{}, nil
}

func (r *UserRepository) GetUsersList(ctx context.Context, limit, offset int32) ([]*model.User, error) {
    users := make([]*model.User, 0, limit)
    return users, nil
}
