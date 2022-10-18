// package controller

// import (
// 	"app/database"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var UserCollection *mongo.Collection=database.UserDatabase(database.Client,"usercoll")
// var ProductCollection *mongo.Collection=database.ProductDatabase(database.Client,"productcoll")

package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
	// "gorm.io/gorm"
)

type DatabaseCollections struct {
	Mongo *mongo.Database
	// Postgres *gorm.DB 
}

//func (H *DatabaseCollections) SveltekitLogin() gin.HandlerFunc {
//
//}

