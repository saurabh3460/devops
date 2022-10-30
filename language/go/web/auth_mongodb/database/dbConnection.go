package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBInstance() *mongo.Client {
	url := "mongodb://mongo:example@mongo:27017/"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	defer func() {
		client.Disconnect(context.TODO())
	}()

	fmt.Println("Successfully connected and pinged.")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(c *mongo.Client, cn string) *mongo.Collection {
	coll := c.Database("cluster").Collection(cn)
	return coll
}
