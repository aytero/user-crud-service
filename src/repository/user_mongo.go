package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user-crud-service/database"
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

func (r *UserRepository) AddUser(ctx context.Context, user *model.User) error {
	res, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.InsertedId = res.InsertedID.(primitive.ObjectID).String()
	return nil

}

func (r *UserRepository) AddUsers(ctx context.Context, users []*model.User) error {
	var usersDocs []interface{}
	for _, u := range users {
		usersDocs = append(usersDocs, u)
	}
	res, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).InsertMany(ctx, usersDocs)
	if err != nil {
		return err
	}
	i := 0
	for _, r := range res.InsertedIDs {
		users[i].InsertedId = r.(primitive.ObjectID).String()
		i++
	}
	return nil
}

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

func (r *UserRepository) GetUsersList(ctx context.Context) ([]model.UserInfo, error) {
	cur, err := r.db.Conn.Collection(r.db.Cfg.CollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var users []model.UserInfo
	for cur.Next(ctx) {
		var us model.UserInfo
		if err = cur.Decode(&us); err != nil {
			return nil, err
		}
		users = append(users, us)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, user *model.UpdateUser) (*model.User, error) {
	uByte, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}
	var in bson.M
	err = bson.Unmarshal(uByte, &in)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": id}
	fields := bson.D{{"$set", in}}

	var res *model.User

	err = r.db.Conn.Collection(r.db.Cfg.CollectionName).FindOneAndUpdate(
		ctx,
		filter,
		fields,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&res)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res, nil
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
