package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) AddNews() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")

		var ReqNews model.NewsModel
		err := c.BindJSON(&ReqNews)
		LogError.LogError("error", err)
		if err != nil {
			LogError.LogError("ReqNews error :", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "news list bind json error"})
		}
		fmt.Println("ReqNewsList : ", ReqNews)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		ReqNews.ID = primitive.NewObjectID()
		id := uuid.New()
		ReqNews.NewsID = id.String()
		res, err := H.Mongo.Collection(os.Getenv("NEWSDATA_COL")).InsertOne(ctx, ReqNews)
		if err != nil {
			LogError.LogError("🚀 ~ file: NewsAdd.go ~ line 29 ~ returnfunc ~ err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		fmt.Println("🚀 ~ file: NewsAdd.go ~ line 29 ~ returnfunc ~ res : ", res)

		c.JSON(http.StatusOK, ReqNews)
	}
}
