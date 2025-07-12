package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClientMain *mongo.Client
var MongoClientSrc *mongo.Client

func SetMainClient(client *mongo.Client) {
	MongoClientMain = client
}
func SetSrcClient(client *mongo.Client) {
	MongoClientSrc = client
}

func GetCollection(collectionName string) *mongo.Collection {
	database := os.Getenv("MONGODB")
	if database == "" {
		database = "shop_dashboard"
	}
	return MongoClientMain.Database(database).Collection(collectionName)
}

func GetMainDB() *mongo.Database {
	database := os.Getenv("MONGODB")
	if database == "" {
		database = "shop_dashboard"
	}
	return MongoClientMain.Database(database)
}

func GetSrcDB() *mongo.Database {
	database := os.Getenv("MONGODB")
	if database == "" {
		database = "shop_dashboard"
	}
	return MongoClientSrc.Database(database)
}
