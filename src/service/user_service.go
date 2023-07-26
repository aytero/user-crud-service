package service

import (
	"context"
	"errors"
	"sync"
	"user-crud-service/config"
	"user-crud-service/model"
	"user-crud-service/repository"
)

type UserService struct {
	userRepo Repository
	cfg      config.SRVC
}

func NewUserService(cfg config.SRVC, repo Repository) *UserService {
	return &UserService{
		userRepo: repo,
		cfg:      cfg,
	}
}

func (uc *UserService) GetUser(ctx context.Context, userId string) (*model.User, error) {
	entry, err := uc.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *UserService) GetUsersList(ctx context.Context) ([]model.UserInfo, error) {
	entry, err := uc.userRepo.GetUsersList(ctx)
	if err != nil {
		return []model.UserInfo{}, err
	}
	return entry, nil
}

func (uc *UserService) AddUsers(ctx context.Context, users []*model.User) error {
	for _, user := range users {
		if user.Password != "" {
			pwd, err := HashPassword(user.Password)
			if err != nil {
				return err
			}
			user.Password = pwd
		}
	}

	var wg sync.WaitGroup
	for _, user := range users {
		_, err := uc.userRepo.GetById(ctx, user.Id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				wg.Add(1)
				go func() {
					defer wg.Done()
					CreateFile(uc.cfg.FileLocation+user.Id, user.Data)
				}()
				err = uc.userRepo.AddUser(ctx, user)
				if err != nil {
					return err
				}
			}
		}
	}
	wg.Wait()
	return nil
}

func (uc *UserService) UpdateUser(ctx context.Context, userId string, userNewInfo *model.UpdateUser) (*model.User, error) {
	old, err := uc.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if userNewInfo.Password != "" && !ComparePasswords(userNewInfo.Password, []byte(old.Password)) {
		pwd, err := HashPassword(userNewInfo.Password)
		if err != nil {
			return nil, err
		}
		userNewInfo.Password = pwd
	}

	var wg sync.WaitGroup

	if userNewInfo.Data != "" && userNewInfo.Data != old.Data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CreateFile(uc.cfg.FileLocation+userId, userNewInfo.Data)

		}()
	}

	entry, err := uc.userRepo.UpdateUser(ctx, userId, userNewInfo)
	if err != nil {
		return nil, err
	}
	wg.Wait()
	return entry, nil
}

func (uc *UserService) DeleteUser(ctx context.Context, userId string) error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		DeleteFile(uc.cfg.FileLocation + userId)
	}()

	err := uc.userRepo.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}
