package database

import (
	"fmt"
	"app/controller"

	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseInitialization(db1 *mongo.Database) controller.DatabaseCollections {
	fmt.Println("🚀 ~ file: database.go ~ line 15 ~ funcDatabaseInitialization ~ mongodb Database : ", db1)
	return controller.DatabaseCollections{Mongo: db1}
}
