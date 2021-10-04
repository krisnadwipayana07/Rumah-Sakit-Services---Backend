package mongodb

import (
	"context"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var global connect db

func Connect(ctx context.Context) *mongo.Database {
	clientOptions := options.Client()
	clientOptions.ApplyURI(viper.GetString(`mongodb.uri`))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(viper.GetString(`mongodb.database`))
}
