package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"user-crud-service/config"
	"user-crud-service/model"
	"user-crud-service/repository"
	mock_service "user-crud-service/repository/mocks"
)

var cfg = config.SRVC{
	FileLocation: "temp/",
}

func TestGetUsersList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()

	var age model.StringInt
	age = 27
	mockRepo := []model.UserInfo{
		{
			Id: "1",
			//Password:   "12345",
			IsActive:   true,
			Balance:    "$2,547.50",
			Age:        age,
			Name:       "Antoine",
			Gender:     "make",
			Company:    "ANI",
			Email:      "antoine@ani.com",
			Phone:      "+1 (868) 439-2675",
			Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
			About:      "Co about about",
			Registered: "2014-10-12T04:14:20 -02:00",
			Latitude:   -10.647121,
			Longitude:  126.04006,
			Tags:       []string{"Consectetur", "Ipsum"},
			Friends: []model.Friends{
				{
					Id:   0,
					Name: "Petty Warren",
				},
			},
			Data: "data data data",
		},
	}

	expected := []model.UserInfo{
		{
			Id:         "1",
			IsActive:   true,
			Balance:    "$2,547.50",
			Age:        age,
			Name:       "Antoine",
			Gender:     "make",
			Company:    "ANI",
			Email:      "antoine@ani.com",
			Phone:      "+1 (868) 439-2675",
			Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
			About:      "Co about about",
			Registered: "2014-10-12T04:14:20 -02:00",
			Latitude:   -10.647121,
			Longitude:  126.04006,
			Tags:       []string{"Consectetur", "Ipsum"},
			Friends: []model.Friends{
				{
					Id:   0,
					Name: "Petty Warren",
				},
			},
			Data: "data data data",
		},
	}

	// Validity check
	repo.EXPECT().GetUsersList(ctx).Return(mockRepo, nil).Times(1)

	service := NewUserService(cfg, repo)
	orders, err := service.GetUsersList(ctx)
	require.NoError(t, err)
	require.Equal(t, expected, orders)

	// Invalidity check
	errDb := errors.New("db is down")
	repo.EXPECT().GetUsersList(ctx).Return(nil, errDb).Times(1)

	_, err = service.GetUsersList(ctx)
	require.Error(t, err)

}

func TestAddUsers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()

	mockRepo := &model.User{
		Id: "1",
		//Password:   service,
		IsActive:   true,
		Balance:    "$2,547.50",
		Age:        model.StringInt(27),
		Name:       "Antoine",
		Gender:     "make",
		Company:    "ANI",
		Email:      "antoine@ani.com",
		Phone:      "+1 (868) 439-2675",
		Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:      "Co about about",
		Registered: "2014-10-12T04:14:20 -02:00",
		Latitude:   -10.647121,
		Longitude:  126.04006,
		Tags:       []string{"Consectetur", "Ipsum"},
		Friends: []model.Friends{
			{
				Id:   0,
				Name: "Petty Warren",
			},
		},
		Data: "data data data",
	}
	mockRequest := []*model.User{
		{
			Id: "1",
			//Password:   "1234",
			IsActive:   true,
			Balance:    "$2,547.50",
			Age:        model.StringInt(27),
			Name:       "Antoine",
			Gender:     "make",
			Company:    "ANI",
			Email:      "antoine@ani.com",
			Phone:      "+1 (868) 439-2675",
			Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
			About:      "Co about about",
			Registered: "2014-10-12T04:14:20 -02:00",
			Latitude:   -10.647121,
			Longitude:  126.04006,
			Tags:       []string{"Consectetur", "Ipsum"},
			Friends: []model.Friends{
				{
					Id:   0,
					Name: "Petty Warren",
				},
			},
			Data: "data data data",
		},
	}
	// Validity check
	repo.EXPECT().GetById(ctx, "1").Return(nil, repository.ErrNotFound).Times(1)
	repo.EXPECT().AddUser(ctx, mockRepo).Return(nil).Times(1)

	service := NewUserService(cfg, repo)
	err := service.AddUsers(ctx, mockRequest)
	require.NoError(t, err)

	// Invalidity check
	errDb := errors.New("db is down")
	repo.EXPECT().GetById(ctx, "1").Return(nil, repository.ErrNotFound).Times(1)
	repo.EXPECT().AddUser(ctx, mockRepo).Return(errDb).Times(1)

	err = service.AddUsers(ctx, mockRequest)
	require.Error(t, err)

}

func TestAddUsersInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()

	mockRepo := &model.User{
		Id:   "1",
		Name: "Antoine",
	}
	mockRequest := []*model.User{
		{
			Id:   "1",
			Name: "Antoine",
		},
	}

	repo.EXPECT().GetById(ctx, "1").Return(nil, repository.ErrNotFound).Times(1)
	repo.EXPECT().AddUser(ctx, mockRepo).Return(errors.New("")).Times(1)

	service := NewUserService(cfg, repo)
	err := service.AddUsers(ctx, mockRequest)

	require.Error(t, err)

}

func TestGetUserById(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()

	retur := &model.User{
		Id:         "1",
		Password:   "1234",
		IsActive:   true,
		Balance:    "$2,547.50",
		Age:        model.StringInt(27),
		Name:       "Antoine",
		Gender:     "make",
		Company:    "ANI",
		Email:      "antoine@ani.com",
		Phone:      "+1 (868) 439-2675",
		Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:      "Co about about",
		Registered: "2014-10-12T04:14:20 -02:00",
		Latitude:   -10.647121,
		Longitude:  126.04006,
		Tags:       []string{"Consectetur", "Ipsum"},
		Friends: []model.Friends{
			{
				Id:   0,
				Name: "Petty Warren",
			},
		},
		Data: "data data data",
	}

	expected := &model.User{
		Id:         "1",
		Password:   "1234",
		IsActive:   true,
		Balance:    "$2,547.50",
		Age:        model.StringInt(27),
		Name:       "Antoine",
		Gender:     "make",
		Company:    "ANI",
		Email:      "antoine@ani.com",
		Phone:      "+1 (868) 439-2675",
		Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:      "Co about about",
		Registered: "2014-10-12T04:14:20 -02:00",
		Latitude:   -10.647121,
		Longitude:  126.04006,
		Tags:       []string{"Consectetur", "Ipsum"},
		Friends: []model.Friends{
			{
				Id:   0,
				Name: "Petty Warren",
			},
		},
		Data: "data data data",
	}

	repo.EXPECT().GetById(ctx, "1").Return(retur, nil).Times(1)

	service := NewUserService(cfg, repo)
	user, err := service.GetUser(ctx, "1")
	require.NoError(t, err)
	require.Equal(t, expected, user)

	// Db error test
	errDb := errors.New("db is down")
	repo.EXPECT().GetById(ctx, "1").Return(nil, errDb).Times(1)

	_, err = service.GetUser(ctx, "1")
	require.Error(t, err)
}

func TestGetUserById_NotFound(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()

	repo.EXPECT().GetById(ctx, "1").Return(nil, repository.ErrNotFound).Times(1)
	service := NewUserService(cfg, repo)

	_, err := service.GetUser(ctx, "1")
	require.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()

	repo.EXPECT().DeleteUser(ctx, "1").Return(nil).Times(1)
	service := NewUserService(cfg, repo)
	err := service.DeleteUser(ctx, "1")
	require.NoError(t, err)

	// Db error test
	errDb := errors.New("db is down")
	repo.EXPECT().DeleteUser(ctx, "1").Return(errDb).Times(1)
	err = service.DeleteUser(ctx, "1")
	require.Error(t, err)
}

func TestUpdateUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()
	retur := &model.User{
		Id:         "1",
		Password:   "1234",
		IsActive:   true,
		Balance:    "$2,547.50",
		Age:        model.StringInt(27),
		Name:       "Antoine",
		Gender:     "make",
		Company:    "ANI",
		Email:      "antoine@ani.com",
		Phone:      "+1 (868) 439-2675",
		Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:      "Co about about",
		Registered: "2014-10-12T04:14:20 -02:00",
		Latitude:   -10.647121,
		Longitude:  126.04006,
		Tags:       []string{"Consectetur", "Ipsum"},
		Friends: []model.Friends{
			{
				Id:   0,
				Name: "Petty Warren",
			},
		},
		Data: "data data data",
	}

	updReq := &model.UpdateUser{
		Id:      "1",
		Age:     model.StringInt(15),
		Name:    "Nana Yo",
		Gender:  "female",
		Company: "ANI",
		Email:   "antoine@ani.com",
		Phone:   "+1 (868) 439-2675",
		Address: "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:   "Co about about",
		Data:    "upd test",
	}
	expected := &model.User{
		Id:         "1",
		IsActive:   true,
		Balance:    "$2,547.50",
		Age:        model.StringInt(15),
		Name:       "Nana Yo",
		Gender:     "female",
		Company:    "ANI",
		Email:      "antoine@ani.com",
		Phone:      "+1 (868) 439-2675",
		Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:      "Co about about",
		Registered: "2014-10-12T04:14:20 -02:00",
		Latitude:   -10.647121,
		Longitude:  126.04006,
		Tags:       []string{"Consectetur", "Ipsum"},
		Friends: []model.Friends{
			{
				Id:   0,
				Name: "Petty Warren",
			},
		},
		Data: "upd test",
	}

	repo.EXPECT().GetById(ctx, "1").Return(retur, nil).Times(1)
	repo.EXPECT().UpdateUser(ctx, "1", updReq).Return(expected, nil).Times(1)
	service := NewUserService(cfg, repo)
	newInfo, err := service.UpdateUser(ctx, "1", updReq)
	require.NoError(t, err)
	require.Equal(t, expected, newInfo)

	// Db error test
	//errDb := errors.New("db is down")
	//repo.EXPECT().GetById(ctx, "1").Return(nil, repository.ErrNotFound).Times(1)
	//repo.EXPECT().UpdateUser(ctx, "1", updReq).Return(nil, errDb).Times(1)
	//repo.EXPECT().DeleteUser(ctx, "1").Return(errDb).Times(1)
	//err = service.DeleteUser(ctx, "1")
	//_, err = service.UpdateUser(ctx, "1", updReq)
	//require.Error(t, err)
}

func TestUpdateUserInvalid(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_service.NewMockRepository(ctl)

	ctx := context.Background()
	retur := &model.User{
		Id:         "1",
		Password:   "1234",
		IsActive:   true,
		Balance:    "$2,547.50",
		Age:        model.StringInt(27),
		Name:       "Antoine",
		Gender:     "make",
		Company:    "ANI",
		Email:      "antoine@ani.com",
		Phone:      "+1 (868) 439-2675",
		Address:    "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:      "Co about about",
		Registered: "2014-10-12T04:14:20 -02:00",
		Latitude:   -10.647121,
		Longitude:  126.04006,
		Tags:       []string{"Consectetur", "Ipsum"},
		Friends: []model.Friends{
			{
				Id:   0,
				Name: "Petty Warren",
			},
		},
		Data: "data data data",
	}

	updReq := &model.UpdateUser{
		Id:      "1",
		Age:     model.StringInt(15),
		Name:    "Nana Yo",
		Gender:  "female",
		Company: "ANI",
		Email:   "antoine@ani.com",
		Phone:   "+1 (868) 439-2675",
		Address: "588 Schaefer Street, Falconaire, Missouri, 9457",
		About:   "Co about about",
		Data:    "upd test",
	}

	err := errors.New("")
	repo.EXPECT().GetById(ctx, "1").Return(retur, nil).Times(1)
	repo.EXPECT().UpdateUser(ctx, "1", updReq).Return(nil, err).Times(1)
	service := NewUserService(cfg, repo)
	_, err = service.UpdateUser(ctx, "1", updReq)
	require.Error(t, err)
}
