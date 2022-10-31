package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections) NewsList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ReqNewsList []model.NewsModel

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		filter := bson.D{{"NewsID", bson.D{{"$ne", ""}}}}
		cursor, err := H.Mongo.Collection(os.Getenv("NEWSDATA_COL")).Find(ctx, filter)
		fmt.Println("🚀 ~ file: NewsList.go ~ line 24 ~ returnfunc ~ cursor : ", cursor)
		if err != nil {
			LogError.LogError("🚀 ~ file: NewsList.go ~ line 25 ~ returnfunc ~ err : ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for cursor.Next(ctx) {
			var temp model.NewsModel
			err := cursor.Decode(&temp)

			if err != nil {
				LogError.LogError("🚀 ~ file: NewsList.go ~ line 36 ~ forcursor.Next ~ err : ", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ReqNewsList = append(ReqNewsList, temp)
		}

		c.JSON(http.StatusOK, ReqNewsList)
	}
}
