package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) AddGameComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ReqComment GameCommentModel
		err := c.BindJSON(&ReqComment)
		if err != nil {
			LogError.LogError("bind json error ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		
		fmt.Println("ReqComment : ", ReqComment)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()
		opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.D{{"GameID", ReqComment.GameID}}
		//update := bson.D{{"$push", ReqComment.Comment}}
		update := bson.M{"$push": bson.M{
			"Comment": ReqComment.Comment,
		},
		}
		//update2 := bson.D{{
		//	"$addToSet", bson.D{{"Comment", bson.D{{
		//		"$each": []model.CommentModel{ReqComment.Comment}}}}}}}
		var UpdatedDocument model.Gamedata
		err = H.Mongo.Collection(os.Getenv("GAMEDATA_COL")).FindOneAndUpdate(ctx, filter, update, opts).Decode(&UpdatedDocument)
		if err != nil {
			LogError.LogError("mongodb find one and update error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
			return
		}
		fmt.Println("UpdatedDocument comment", UpdatedDocument.Comment)
		var myComment model.CommentModel
		myComment.Name = "raka"
		myComment.Rating = 7
		myComment.Description = "when you assign the cursor returned from the find() method to a variable using the var keyword, the cursor does not automatically iterate. You can use the cursor method forEach() to iterate the cursor and access the documents, as in the following example: var myCursor = db."

		c.JSON(http.StatusOK, myComment)
	}
}

type GameCommentModel struct {
	GameID  string             `json:"GameID"`
	Comment model.CommentModel `json:"Comment"`
}
