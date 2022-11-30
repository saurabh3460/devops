package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBInstance() *mongo.Client {
	url := "mongodb://mongo:example@localhost:27017/"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// var dbs []string
	// if dbs, err = client.ListDatabaseNames(context.TODO(), bson.D{{}}); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("List of dataabse present:", dbs)
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	fmt.Println("Successfully connected and pinged.")
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(c *mongo.Client, cn string) *mongo.Collection {
	coll := c.Database("new").Collection(cn)
	return coll
}
