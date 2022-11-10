package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) BrowseGames() gin.HandlerFunc {
	return func(c *gin.Context) {

		// NOT COMPLETE
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")

		QueryId := c.Param("newsid")
		fmt.Println("QueryId : ", QueryId)

		if QueryId == "" {
			LogError.LogError("🚀 ~ file:  QueryId : ", errors.New("Query ID no found"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid  newsid id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var result model.NewsModel
		err := H.Mongo.Collection(os.Getenv("NEWSDATA_COL")).FindOne(ctx, bson.D{{"NewsID", QueryId}}).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " News not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "News not found & other error"})
			return
		}

		// filter := bson.D{{"GameID", bson.D{{"$et", QueryId}}}}
		// opts := options.FindOne().SetSort(bson.D{})

		// var res bson.M
		// err := H.Mongo.Collection("GameData").FindOne(ctx, filter, opts).Decode(&res)

		// fmt.Println("🚀 ~ file: GetGameData.go ~ line 51 ~ returnfunc ~ res : ", res)

		// if err != nil {
		// 	log.Fatalln("❌ findOne : ", err)
		// }
		// fmt.Println("🚀✨ FindOne successful & result: ", res)
		fmt.Println("result : ", result)
		c.JSON(http.StatusOK, result)
		//c.JSON(http.StatusOK, "Get News By id ok !!! ")
	}
}
