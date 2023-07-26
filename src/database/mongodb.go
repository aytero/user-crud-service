package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"user-crud-service/config"
)

type Database struct {
	Client *mongo.Client
	Conn   *mongo.Database
	Cfg    config.MG
}

func New(ctx context.Context, cfg config.MG) (*Database, error) {

	// local Mongo URI
	//mongoURI := "mongodb://" + cfg.User + ":" + cfg.Password + "@localhost:27017"
	mongoURI := "mongodb://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":27017"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	// todo indexes
	//coll := client.Database(cfg.DbName).Collection(cfg.CollectionName)
	//index := []mongo.IndexModel{
	//	{
	//		Keys: bsonx.Doc{{Key: "id", Value: bsonx.String("text")}},
	//	},
	//}
	//opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	//_, errIndex = coll.Indexes().CreateMany(context, index, opts)
	//if err != nil {
	//	panic(errIndex)
	//}
	//c := client.Database(cfg.DbName).Collection(cfg.CollectionName)
	//_, err = c.Indexes().CreateOne(
	//	context.Background(),
	//	mongo.IndexModel{
	//		Keys: bson.M{
	//			"id": "",
	//		},
	//		Options: options.Index().SetUnique(true),
	//	},
	//)

	return &Database{
		Client: client,
		Conn:   client.Database(cfg.DbName),
		Cfg:    cfg,
	}, nil
}

func (db *Database) Close(ctx context.Context) {

	defer func(client *mongo.Client, ctx context.Context) {
		client.Disconnect(ctx)
	}(db.Client, ctx)
}
