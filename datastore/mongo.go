package datastore

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (b *Books) Initialize() {

	clientOptions := options.Client().ApplyURI("mongodb://172.17.0.4:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to MongoDB!")
	//b.Store = loader.LoadData(file)
}
