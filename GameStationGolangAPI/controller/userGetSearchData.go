package controller

import (
	// "app/LogError"
	// "app/model"
	"context"
	"fmt"
	// "fmt"
	"os"

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

func (H *DatabaseCollections) GetSearchData() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		// QueryId := c.Query("gameid")
		QueryId := c.Param("search")
        fmt.Println("🚀 ~ file: userGetSearchData.go ~ line 32 ~ returnfunc ~ QueryId : ", QueryId)
		// fmt.Println("🚀 ~ file: GetGameData.go ~ line 26 ~ returnfunc ~ QueryId : ", QueryId)

		if QueryId == "" {
			LogError.LogError("🚀 ~ file: GetGameData.go ~ line 20 ~ returnfunc ~ QueryId : ", errors.New("Query ID no found"))
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid  Game id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		// var result model.Gamedata
		// query := bson.M{
		// 	"$text": bson.M{
		// 	"$search": QueryId,
		//    },
		// }
// 		matchStage := bson.D{{"$match", bson.D{{"$text", bson.D{{"$search", QueryId}}}}}}
// cursor, err := H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).Aggregate(context.TODO(), mongo.Pipeline{matchStage})

		// filter := bson.D{{"$text", bson.D{{"$search", QueryId}}}}
		
		// filter := bson.D{{"Name" , bson.D{{"$regex" , QueryId}}}}
		filter :=bson.M{"Name": bson.M{"$regex": QueryId, "$options": "i"}}
        // fmt.Println("🚀 ~ file: userGetSearchData.go ~ line 49 ~ returnfunc ~ QueryId : ", QueryId)
		cursor,err := H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).Find(ctx, filter)
		fmt.Println("🚀 ~ file: userGetSearchData.go ~ line 52 ~ returnfunc ~ cursor : ", cursor)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		var results []model.Gamedata
		if err = cursor.All(ctx, &results); err != nil {
			LogError.LogError("🚀 ~ file: GetGameData.go ~ line 51 ~ returnfunc ~ res : ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Game not found & other error"})
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

		c.JSON(http.StatusOK, results)

	}
}
