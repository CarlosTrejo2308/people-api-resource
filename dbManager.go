package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func generatePath() string {
	dbUser := "root"
	dbPwd := "example"
	dbURL := os.Getenv("MONGO_URL")
	port := 27017

	str := fmt.Sprintf("mongodb://%s:%s@%s:%d/?connect=direct", dbUser, dbPwd, dbURL, port)

	return str
}

func connect(uri string) *mongo.Client {
	if uri == "" {
		uri = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	}

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPwd + "@" + dbURL + ":" + port + "/?connect=direct"))

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27020/"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	abortOnError(err)

	err = client.Ping(ctx, readpref.Primary())
	abortOnError(err)

	return client
}
