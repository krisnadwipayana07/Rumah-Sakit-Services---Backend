package mongodb

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var global connect db

func Connect(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(viper.GetString(`mongodb.uri`))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(viper.GetString(`mongodb.database`)), nil
}
