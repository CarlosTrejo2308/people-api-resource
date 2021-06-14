package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/CarlosTrejo2308/peopleApiResource/abort"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// GeneratePath returns a formated string of the mongo url
// given global enviroments of the credentials, url and port
func GeneratePath() string {
	dbUser := os.Getenv("MONGO_USR")
	dbPwd := os.Getenv("MONGO_PWD")
	dbURL := os.Getenv("MONGO_URL")
	port := os.Getenv("MONGO_PORT")

	str := fmt.Sprintf("mongodb://%s:%s@%s:%s/?connect=direct", dbUser, dbPwd, dbURL, port)

	return str
}

// Connect recives a mongo url string and returns a
// mongo client of that connection.
func Connect(uri string) *mongo.Client {
	if uri == "" {
		uri = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
	}

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPwd + "@" + dbURL + ":" + port + "/?connect=direct"))

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27020/"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	abort.AbortOnError(err)

	err = client.Ping(ctx, readpref.Primary())
	abort.AbortOnError(err)

	return client
}
