package config

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() *mongo.Database {clientOptions := options.Client().ApplyURI("mongodb+srv://user_test:Ismacs2003@firstproyectwebengineer.b6xlw.mongodb.net/?retryWrites=true&w=majority&appName=FirstProyectWebEngineering")
	
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	DB = client.Database("MedicalRecord")
	return DB
}
