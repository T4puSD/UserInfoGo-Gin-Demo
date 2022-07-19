package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"user-info-service/config"
)

var DB *mongo.Client

func ConnectDB() {
	if DB != nil {
		return
	}

	// establishing connection
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetEnv().MongodbUri))
	if err != nil {
		log.Fatalln("Unable to connect to MongoDB", err)
	}

	// testing the connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to MongoDB")
	DB = client
}

func GetCollection(collectionName string) *mongo.Collection {
	ConnectDB()
	if DB == nil {
		log.Fatalln("Database client is not initialized!")
	}
	return DB.Database(config.GetEnv().MongodbDatabaseName).Collection(collectionName)
}
