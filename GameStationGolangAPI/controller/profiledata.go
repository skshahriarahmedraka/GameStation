package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"time"
)

func (H *DatabaseCollections) UserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusOK, "Login successful !!! ")
		//c.Param("profileid")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		var UserDataDB model.UserData
		//options.FindOne()
		err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).FindOne(ctx, bson.M{"UserID": c.Param("profileid")}).Decode(&UserDataDB)
		if err != nil {
			LogError.LogError("profiledata mongo find one error", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			objID, _ := primitive.ObjectIDFromHex("")
			UserDataDB.ID = objID
			UserDataDB.Password = ""
			c.JSON(http.StatusOK, UserDataDB)
		}

	}
}
