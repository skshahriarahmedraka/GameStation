package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"
)

func (H *DatabaseCollections) ProfileUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		var ReqUserData model.UserData
		err := c.BindJSON(&ReqUserData)
		LogError.LogError("bind json error", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "binding error"})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		count, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).CountDocuments(ctx, bson.M{"UserID": ReqUserData.UserID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb Count Document error"})
			return
		}
		if count > 0 {
			opts := options.Update().SetUpsert(true)
			filter := bson.D{{"UserID", ReqUserData.UserID}}
			update := bson.D{
				{"$set", bson.D{{"Name", ReqUserData.Name}}},
				{"$set", bson.D{{"Email", ReqUserData.Email}}},
				{"$set", bson.D{{"ProfileImg", ReqUserData.ProfileImg}}},
				{"$set", bson.D{{"BannerImg", ReqUserData.BannerImg}}},
				{"$set", bson.D{{"Mobile", ReqUserData.Mobile}}},
				{"$set", bson.D{{"Address", ReqUserData.Address}}},
				{"$set", bson.D{{"City", ReqUserData.City}}},
				{"$set", bson.D{{"Country", ReqUserData.Country}}},
				{"$set", bson.D{{"ZipCode", ReqUserData.ZipCode}}},
			}
			res, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).UpdateOne(ctx, filter, update, opts)
			if err != nil {
				LogError.LogError("mongodb update one error", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb update one error"})
				return

			} else {
				fmt.Println("response ", res)
				c.JSON(http.StatusOK, ReqUserData)
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request mongodb count 0"})

		}

	}
}
