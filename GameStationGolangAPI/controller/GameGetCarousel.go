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
	// "app/LogError"
	"app/model"
	// "errors"
	"net/http"

	// "time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) GetCarousel() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")

		// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		// defer cancel()

		var results []model.Gamedata

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		// 1
		var result model.Gamedata
		err := H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", os.Getenv("CAROUSEL_GAMEID_1")}}).Decode(&result)
        fmt.Println("🚀 ~ file: GameGetCarousel.go ~ line 42 ~ returnfunc ~ err : ", err)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		results = append(results, result)
		//2
		err = H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", os.Getenv("CAROUSEL_GAMEID_2")}}).Decode(&result)
        fmt.Println("🚀 ~ file: GameGetCarousel.go ~ line 55 ~ returnfunc ~ err : ", err)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		results = append(results, result)

		//3
		err = H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", os.Getenv("CAROUSEL_GAMEID_3")}}).Decode(&result)
        fmt.Println("🚀 ~ file: GameGetCarousel.go ~ line 69 ~ returnfunc ~ err : ", err)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		results = append(results, result)

		// 4
		err = H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", os.Getenv("CAROUSEL_GAMEID_4")}}).Decode(&result)
        fmt.Println("🚀 ~ file: GameGetCarousel.go ~ line 83 ~ returnfunc ~ err : ", err)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		results = append(results, result)
		// 5
		err = H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", os.Getenv("CAROUSEL_GAMEID_5")}}).Decode(&result)
        fmt.Println("🚀 ~ file: GameGetCarousel.go ~ line 96 ~ returnfunc ~ err : ", err)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		results = append(results, result)
		// 6
		err = H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOne(ctx, bson.D{{"GameID", os.Getenv("CAROUSEL_GAMEID_6")}}).Decode(&result)
        fmt.Println("🚀 ~ file: GameGetCarousel.go ~ line 109 ~ returnfunc ~ err : ", err)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": " Game not Found"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game not found & other error"})
			return
		}

		results = append(results, result)

		// fmt.Println("🚀 ~ file: GameMostPopular.go ~ line 94 ~ returnfunc ~ result : ", results)
		c.JSON(http.StatusOK, results)
	}
}
