package controller

import (
	"app/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection=database.UserDatabase(database.Client,"usercoll")
var ProductCollection *mongo.Collection=database.ProductDatabase(database.Client,"productcoll")