package controller

import (
	// "app/LogError"
	// "app/model"
	"context"
	"fmt"
	"os"

	// "log"
	"time"

	// "fmt"
	"app/LogError"
	"app/model"
	// "errors"
	"net/http"

	// "time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) TopRated() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")


		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		var result []model.Gamedata
		filter := bson.D{}
		opts := options.Find().SetSort(bson.D{{"Rating", -1}})
		cursor, err := H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).Find(ctx, filter, opts)
		if err != nil {
        LogError.LogError("🚀 ~ file: GameMostPopular.go ~ line 48 ~ returnfunc ~ err : ", err)
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}
		

		var x int =0
		for cursor.Next(ctx) {
			var temp model.Gamedata
			err := cursor.Decode(&temp)

			if err != nil {
            LogError.LogError("🚀 ~ file: GameMostPopular.go ~ line 70 ~ forcursor.Next ~ err : ", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			result = append(result, temp)
			if x >13 {
				break
			}else { x+=1}
		}

	

        fmt.Println("🚀 ~ file: GameMostPopular.go ~ line 94 ~ returnfunc ~ result : ", result)
		c.JSON(http.StatusOK, result)

	}
}
