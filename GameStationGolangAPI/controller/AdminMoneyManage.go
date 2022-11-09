package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) AdminMoneyManagement() gin.HandlerFunc {
	return func(c *gin.Context) {
		type ReqStruct struct {
			UserID string  `json:"UserID" bson:"UserID"`
			Amount float64 `json:"Amount" bson:"Amount"`
		}
		var ReqData ReqStruct

		err := c.BindJSON(&ReqData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}
		if ReqData.UserID == "" || ReqData.Amount == 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid or  amount is empty"})
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		count, err := H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).CountDocuments(ctx, bson.M{"UserID": ReqData.UserID})
		if err != nil {
			LogError.LogError("🚀 ~ file: userTokenReq.go ~ line 36 ~ returnfunc ~ err : ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {

			//building token
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

			if err != nil {
				LogError.LogError("❌🔥 error in token.SignedString ", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			var SendJwt model.UserMoneyReq
			// struct {
			// 	// UserID string `json:"UserID" bson:"UserID"`
			// 	Amount float64 `json:"Amount" bson:"Amount"`
			// 	JWT    string
			// 	Status string `json:"Status" bson:"Status"`
			// }
			SendJwt.JWT = tokenString
			SendJwt.Amount = ReqData.Amount
			SendJwt.Status = false

			//inserting token in db
			opts := options.FindOneAndUpdate().SetUpsert(true)
			filter := bson.D{{"UserID", ReqData.UserID}}
			//update := bson.D{{"$push", ReqComment.Comment}}
			update := bson.M{"$push": bson.M{
				"ReqList": SendJwt,
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

			c.JSON(http.StatusOK, SendJwt)

		} else {
			LogError.LogError("❌🔥 error in User already registered ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
			return

		}

		// authToken, err := c.Cookie("Auth1")
		// if authToken == "" {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthenticated User, cookie Auth1 not found"})
		// 	c.Abort()
		// 	return
		// }
		// claims, str := token.ValidateMoneyJWT(ReqData.JWT)
		// if str != "" {
		// 	LogError.LogError("🚀 ~ file: auth.go ~ line 22 ~ returnfunc ~ str : ", errors.New("jwt error"))
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Claims error"})
		// 	c.Abort()
		// 	return

		// }
		// c.Set("Email", claims.)
		// c.Set("UserID", claims.UserID)
		// c.Set("FirstName", claims.Name)

		// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		// defer cancel()
		// H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOneAndUpdate(ctx, bson.M{"UserID": ReqData.UserID}, bson.M{"$inc": bson.M{"Amount": ""}})
	}
}
