package service

import (
    "context"
    "user-crud-service/internal/model"
)

type UserService struct {
    userRepo Repository
}

func NewUserService(repo Repository) *UserService {
    return &UserService{
        userRepo: repo,
    }
}

func (uc *UserService) GetUser(ctx context.Context, userId string) (*model.User, error) {
    entry, err := uc.userRepo.GetById(ctx, userId)
    if err != nil {
        return nil, err
    }
    if entry == nil {
        return nil, nil
    }
    return entry, nil
}

func (uc *UserService) GetUsersList(ctx context.Context, limit, offset int32) ([]*model.User, error) {
    entry, err := uc.userRepo.GetUsersList(ctx, limit, offset)
    if err != nil {
        return nil, err
    }
    return entry, nil
}
