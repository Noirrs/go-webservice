package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(url string, dbName string) *mongo.Collection {

	ctx := context.TODO()
	var mongourl string

	if url != "" {
		mongourl = url
	} else {
		mongourl = "mongodb://localhost:27017"
	}

	mongoconn := options.Client().ApplyURI(mongourl)
	client, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal("unknown err occurred", err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal("unknown err occurred", err)
	}

	fmt.Println("connected to mongodb")

	if (dbName == "") {
		dbName = "leaw"
	}

	categoryCollection := client.Database(dbName).Collection("categories")

	return categoryCollection
}
