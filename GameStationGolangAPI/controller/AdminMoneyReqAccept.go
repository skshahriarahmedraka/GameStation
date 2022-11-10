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
	"github.com/golang-jwt/jwt/v4"

	// "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) AdminMoneyReqAccept() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		type ReqDataStruct struct {
			UserID string 		`json:"UserID", bson:"UserID"`
			Amount float64	`json:"Amount", bson:"Amount"`
		}
		var ReqData ReqDataStruct
		err := c.BindJSON(&ReqData)
		if err != nil {
			LogError.LogError("bind json error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "binding error"})
			return
		}
		if ReqData.UserID == "" || ReqData.Amount == 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid or  amount is empty"})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		count, err := H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).CountDocuments(ctx, bson.M{"UserID": ReqData.UserID})
        fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 45 ~ returnfunc ~ count : ", count)
		if err != nil {
			LogError.LogError("🚀 ~ file: userTokenReq.go ~ line 36 ~ returnfunc ~ err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {

			// var SendJwt model.UserMoneyReq

			// SendJwt.JWT = "Your Request is Pending for Admin Acceptance"
			// SendJwt.Amount = ReqData.Amount
			// SendJwt.Status = false

			//update := bson.D{{"$push", ReqComment.Comment}}

			expirationTime := time.Now().Add(time.Hour * 100000)
			myClaim := &model.JWTModel{
				UserID: ReqData.UserID,
				Amount: ReqData.Amount,

				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_MONEY")))
            fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 73 ~ returnfunc ~ tokenString : ", tokenString)

			if err != nil {
				LogError.LogError("❌🔥 error in token.SignedString ", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			//inserting token in db
			// opts := options.FindOneAndUpdate().SetUpsert(true)
			filter := bson.D{{"UserID", ReqData.UserID}}
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

			var ReqListData []model.UserMoneyReq
			ReqListData = resData.ReqList
			for i, v := range ReqListData {
				if v.Amount == ReqData.Amount && v.Status==false {
					ReqListData[i].JWT = tokenString
					ReqListData[i].Status = true
					break
				}
			}
			fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 99 ~ returnfunc ~ ReqListData : ", ReqListData)
			res,err:=H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).ReplaceOne(ctx,filter,bson.D{{"ReqList",ReqListData}})
            fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 107 ~ returnfunc ~ res : ", res)
			if err != nil {
				LogError.LogError("mongodb find one and update error", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
				return
			}

			c.JSON(http.StatusOK,res )

		} else {
			LogError.LogError("❌🔥 User is not found ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
			return

		}

	}
}
