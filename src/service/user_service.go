package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"user-crud-service/model"
	"user-crud-service/repository"
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
	//if entry == nil {
	//	return nil, nil
	//}
	return entry, nil
}

func (uc *UserService) UpdateUser(ctx context.Context, userId string, userNewInfo model.User) (model.User, error) {
	entry, err := uc.userRepo.UpdateUser(ctx, userId, userNewInfo)
	if err != nil {
		return model.User{}, err
	}
	return entry, nil
}

func (uc *UserService) DeleteUser(ctx context.Context, userId string) error {
	err := uc.userRepo.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	// todo delete file with the data
	return nil
}

func (uc *UserService) GetUsersList(ctx context.Context) ([]model.User, error) {
	//entry, err := uc.userRepo.GetUsersList(ctx, limit, offset)
	entry, err := uc.userRepo.GetUsersList(ctx)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *UserService) AddUser(ctx context.Context, user *model.User) (*model.User, error) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hpass)

	entry, err := uc.userRepo.AddUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *UserService) AddUsers(ctx context.Context, users []*model.User) ([]*model.User, error) {
	for _, user := range users {
		hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hpass)
		_, err = uc.userRepo.GetById(ctx, user.Id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				_, err = uc.userRepo.AddUser(ctx, user)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	// todo
	return nil, nil
}
