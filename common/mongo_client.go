package common

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"compress/zlib"

	user_schema "github.com/mnshah219/go_gin/auth/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client = nil

func GetClient() *mongo.Client {
	if client != nil {
		return client
	}
	fmt.Print("Initiating mongo with zlib compression ", zlib.BestCompression)
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	mongoURI := os.Getenv("MONGODB_URI")
	clientInstance, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error establising mongo connection", err)
	}
	err = clientInstance.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("Error establising mongo connection", err)
	}
	client = clientInstance
	ensureIndex(client)
	return client
}

func ensureIndex(client *mongo.Client) {
	database := client.Database(DB_NAME)
	// user table
	userIdx := database.Collection(user_schema.USER_TABLE).Indexes()

	emailIdx := mongo.IndexModel{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)}
	userIdx.CreateOne(context.TODO(), emailIdx)
}
