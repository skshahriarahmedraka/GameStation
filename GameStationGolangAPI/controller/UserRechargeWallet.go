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
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) ProfileRechargeWallet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")
		type ReqStruct struct {
			UserID string `json:"UserID" bson:"UserID"`
			JWT    string `json:"JWT" bson:"JWT"`
		}

		var ReqUserData ReqStruct

		err := c.BindJSON(&ReqUserData)
		if err != nil {
			LogError.LogError("bind json error", err)
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
			filter := bson.D{{"UserID", ReqUserData.UserID}}
			// update := bson.D{{"$set", bson.D{{"ReqList.$.JWT",  tokenString}}}}
			//update2 := bson.D{{
			//	"$addToSet", bson.D{{"Comment", bson.D{{
			//		"$each": []model.CommentModel{ReqComment.Comment}}}}}}}
			var resData model.UserMoney

			err = H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOne(ctx, filter).Decode(&resData)
			fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 90 ~ returnfunc ~ resData : ", resData)
			if err != nil {
				LogError.LogError("mongodb find one and update error", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
				return
			}

			//var ReqListData []model.UserMoneyReq
			//ReqListData = resData.ReqList
			for i, v := range resData.ReqList {
				if v.JWT == ReqUserData.JWT && v.Status == "accepted" {
					resData.Coin += v.Amount
					resData.ReqList[i].Status = "used"

					opts := options.Update().SetUpsert(true)
					filter := bson.D{{"UserID", ReqUserData.UserID}}
					update := bson.D{
						{"$inc", bson.D{{"Coin", v.Amount}}},
					}
					res, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).UpdateOne(ctx, filter, update, opts)
                    fmt.Println("🚀 ~ file: UserRechargeWallet.go ~ line 72 ~ returnfunc ~ res : ", res)
					if err != nil {
						LogError.LogError("mongodb update one error", err)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb update one error"})
						return

					} 
					break
				}
			}
			res, err := H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).ReplaceOne(ctx, filter, resData)
			fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 107 ~ returnfunc ~ res : ", res)
			if err != nil {
				LogError.LogError("mongodb find one and update error", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
				return
			}

			c.JSON(http.StatusOK, res)

		} else {
			LogError.LogError("❌🔥 User is not found ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
			return

		}

	}
}
