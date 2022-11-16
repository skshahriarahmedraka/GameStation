package controller

import (
	// "app/LogError"
	"app/LogError"
	"app/model"
	"context"

	// "errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) BrowseGames() gin.HandlerFunc {
	return func(c *gin.Context) {

		// NOT COMPLETE
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")

		// QueryId := c.Param("newsid")
		// fmt.Println("QueryId : ", QueryId)

		// if QueryId == "" {
		// 	LogError.LogError("🚀 ~ file:  QueryId : ", errors.New("Query ID no found"))
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid  newsid id"})
		// 	return
		// }

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var results []model.Gamedata
		cursor, err := H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).Find(ctx, bson.D{})
		if err != nil {
			if err == mongo.ErrNoDocuments {
				LogError.LogError("🚀 ~ file: BrowseList.go ~ line 43 ~ returnfunc ~ err : ", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": " News not Found"})
				return
			}
			LogError.LogError("🚀 ~ file: BrowseList.go ~ line 43 ~ returnfunc ~ err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "News not found & other error"})
			return
		}
		// var results []model.UserMoney
		if err = cursor.All(context.TODO(), &results); err != nil {
			LogError.LogError("🚀 ~ file: BrowseList.go ~ line 48 ~ iferr=cursor.All ~ err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": " News not Found"})
			return
		}
		fmt.Println("result : ", results)
		c.JSON(http.StatusOK, results)

		// filter := bson.D{{"GameID", bson.D{{"$et", QueryId}}}}
		// opts := options.FindOne().SetSort(bson.D{})

		// var res bson.M
		// err := H.Mongo.Collection("GameData").FindOne(ctx, filter, opts).Decode(&res)

		// fmt.Println("🚀 ~ file: GetGameData.go ~ line 51 ~ returnfunc ~ res : ", res)

		// if err != nil {
		// 	log.Fatalln("❌ findOne : ", err)
		// }
		// fmt.Println("🚀✨ FindOne successful & result: ", res)
		//c.JSON(http.StatusOK, "Get News By id ok !!! ")
	}
}
