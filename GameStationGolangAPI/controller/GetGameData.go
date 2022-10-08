package controller

import (
	// "app/LogError"
	// "app/model"
	"context"
	"fmt"
	// "log"
	"time"

	// "fmt"
	"app/LogError"
	"app/model"
	"errors"
	"net/http"

	// "time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) GetGameData() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		// QueryId := c.Query("gameid")
		QueryId := c.Param("gameid")
        fmt.Println("🚀 ~ file: GetGameData.go ~ line 26 ~ returnfunc ~ QueryId : ", QueryId)

		if QueryId == "" {
			LogError.LogError("🚀 ~ file: GetGameData.go ~ line 20 ~ returnfunc ~ QueryId : ", errors.New("Query ID no found"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid  Game id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var result model.Gamedata
err := H.Mongo.Collection("GameData").FindOne(ctx, bson.D{{"GameID", QueryId}}).Decode(&result)
if err != nil {
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
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

		c.JSON(http.StatusOK, result)

	}
}
