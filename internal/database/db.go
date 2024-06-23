package database

import (
	"context"

	"github.com/aperezgdev/trello-clone-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoClient() (*mongo.Client, error) {
	envApp := env.GetEnvApp(".env")

	option := options.Client().ApplyURI(envApp.MONGO_URL + envApp.MONGO_DB_NAME)
	option.SetAuth(options.Credential{
		Username: envApp.MONGO_USERNAME,
		Password: envApp.MONGO_PASSWORD,
	})

	client, err := mongo.Connect(context.TODO(), option)

	if err != nil {
		return client, nil
	}

	err = client.Ping(context.TODO(), readpref.Nearest())

	return client, err
}

func CancelMongoClient(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		panic(err)
	}
}
