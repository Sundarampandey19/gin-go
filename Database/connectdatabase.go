package database

import (
	"context"
	"log"
	"time"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Client *mongo.Client

func ConnectMongoDB(){
	clientOptions  := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = client 
}

