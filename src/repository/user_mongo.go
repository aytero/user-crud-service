package repository

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"user-crud-service/internal/database"
	"user-crud-service/model"
)

var (
	ErrNotFound = errors.New("not found")
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(dbConn *database.Database) *UserRepository {
	return &UserRepository{
		db: dbConn,
	}
}

// todo rename to getUserById

func (r *UserRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	var us model.User

	err := r.db.Conn.Collection(r.db.Cfg.CollectionName).FindOne(ctx, bson.M{"id": id}).Decode(&us)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &us, nil
}

func (r *UserRepository) AddUser(ctx context.Context, user *model.User) (*model.User, error) {
	res, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	// todo
	//user.InsertedId = res.InsertedID.(primitive.ObjectID)
	user.InsertedId = res.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (r *UserRepository) AddUsers(ctx context.Context, users []*model.User) ([]*model.User, error) {
	var usersDocs []interface{}
	for _, u := range users {
		usersDocs = append(usersDocs, u)
	}
	res, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).InsertMany(ctx, usersDocs)
	if err != nil {
		return nil, err
	}

	i := 0
	for _, r := range res.InsertedIDs {
		users[i].InsertedId = r.(primitive.ObjectID).String()
		i++
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, user model.User) (model.User, error) {
	in := bson.M{}
	if user.Name != "" {
		in["name"] = user.Name
	}
	if user.Password != "" && len(user.Password) > 0 {
		in["password"] = user.Password
	}
	// todo etc model fields
	filter := bson.M{"id": user.Id}
	fields := bson.M{"$set": user}

	out, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).UpdateOne(ctx, filter, fields)
	if err != nil {
		return model.User{}, err
	}
	if out.MatchedCount == 0 {
		return model.User{}, ErrNotFound
	}
	return user, nil
}

func (r UserRepository) DeleteUser(ctx context.Context, id string) error {
	filter := bson.M{"id": id}
	out, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *UserRepository) GetUsersList(ctx context.Context) ([]model.User, error) {
	cur, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []model.User
	for cur.Next(ctx) {
		//var us bson.M
		var us model.User
		if err = cur.Decode(&us); err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, us)
	}

	return users, nil
}
