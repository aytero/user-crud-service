package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"os"
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

func (uc UserService) CreateUserFile(userId string, data string) error {
	f, err := os.Create(userId)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func (uc UserService) DeleteUserFile(userId string) error {
	err := os.Remove(userId)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserService) UpdateUser(ctx context.Context, userId string, userNewInfo *model.User) (*model.User, error) {
	old, err := uc.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	entry, err := uc.userRepo.UpdateUser(ctx, userId, userNewInfo)
	if err != nil {
		return nil, err
	}

	// todo if data changed
	// todo goroutine
	if userNewInfo.Data != "" && userNewInfo.Data != old.Data {
		go uc.CreateUserFile(userId, userNewInfo.Data)
	}
	return entry, nil
}

func (uc *UserService) DeleteUser(ctx context.Context, userId string) error {
	err := uc.userRepo.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	// todo delete file with the data
	go uc.DeleteUserFile(userId)
	return nil
}

func (uc *UserService) GetUsersList(ctx context.Context) ([]model.User, error) {
	entry, err := uc.userRepo.GetUsersList(ctx)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (uc *UserService) AddUsers(ctx context.Context, users []*model.User) ([]*model.User, error) {
	for _, user := range users {
		us := user
		go func(user *model.User) (*model.User, error) {

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
					// todo transaction
					go uc.CreateUserFile(user.Id, user.Data)
					//err = os.WriteFile(user.Id, []byte(user.Data), 0644)
				}
			}
			return user, nil
		}(us)
		//if err != nil {
		//	log.Error().Msgf("%v\n", err)
		//}
	}
	// todo
	return users, nil
}
