package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

//var Validate = validator.New()

func (H *DatabaseCollections) SveltekitRegister() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		var user model.UserData
		err := c.BindJSON(&user)
		if err != nil {
			LogError.LogError("❌🔥 error in c.bindjson() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("🚀", user)
		if user.Email == os.Getenv("ADMIN_EMAIL") {

			user.Accounttype = os.Getenv("ADMIN_ACCOUNT_TYPE")
			user.UserID = os.Getenv("ADMIN_ACCOUNT_ID")
			// ctx,cancel :=context.WithTimeout(context.Background(), 10*time.Second)
			// defer cancel()
			// res,err:=H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).InsertOne(ctx, bson.D{{"ID",primitive.NewObjectID()},{"AdminID", user.UserID}})
			// if err!=nil{
			// 	LogError.LogError("❌🔥 error in InsertOne() ", err)
			// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// 	return
			// } else {

			// 	fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 44 ~ ifuser.Email==os.Getenv ~ res : ", res)
			// }

		} else {
			user.Accounttype = "normal"
			myid := uuid.New()
			user.UserID = myid.String()

		}
		user.BannerImg = ""
		user.Coin = 0.0
		user.ProfileImg = ""
		//objID, err := primitive.ObjectIDFromHex(myid.String())
		//if err != nil {
		//	LogError.LogError("❌🔥 primitive.ObjectIDFromHex error:  ", err)
		//
		//}
		user.ID = primitive.NewObjectID()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		// SEARCH EMAIL
		count, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).CountDocuments(ctx, bson.M{"Email": user.Email})
		if err != nil {
			LogError.LogError("❌🔥 error in mongodb connection  ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			LogError.LogError("❌🔥 error in User already registered ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
			return
		}
		////MOBILE
		//count, err = H.Mongo.Collection(os.Getenv("USERDATA_COL")).CountDocuments(ctx, bson.M{"Mobile": user.Mobile})
		//if err != nil {
		//	LogError.LogError("❌🔥 error in mongodb countDocument mobile connection error ", err)
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb countDocument mobile connection error"})
		//
		//	return
		//
		//}
		//if count > 0 {
		//	LogError.LogError("❌🔥 error in mobile number already registered ", err)
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "mobile number already registered"})
		//	return
		//}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
		if err != nil {
			log.Println(err)
		}

		res, err := H.Mongo.Collection(os.Getenv("USERDATA_COL")).InsertOne(ctx, user)
		LogError.LogError("🚀 ~ file: register.go ~ line 102 ~ func ~ err : ", err)
		if err == nil {

			fmt.Println("🚀 ~ file: register.go ~ line 116 ~ func ~ res : ", res)
		}
		var uMoney model.UserMoney
		uMoney.UserID = user.UserID
		uMoney.Coin = 0.0
		uMoney.ReqList = []model.UserMoneyReq{}
		res, err = H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).InsertOne(ctx, uMoney)
		LogError.LogError("🚀 ~ file: SveltekitRegister.go ~ line 115 ~ returnfunc ~ err : ", err)
		fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 115 ~ returnfunc ~ res : ", res)
		if err != nil {
			LogError.LogError("❌🔥 error in InsertOne() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//mongoRes, err := H.Mongo.Collection("userdb").InsertOne(ctx, user)
		//fmt.Println("🚀 ~ file: register.go ~ line 80 ~ func ~ mongoRes : ", mongoRes)

		//if err != nil {
		//	LogError.LogError("❌🔥 error in mongodb Connection error ", err)
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb Connection error"})
		//
		//	return
		//}

		expirationTime := time.Now().Add(time.Hour * 100)

		claims := &model.TokenClaims{
			Name:        user.Name,
			Email:       user.Email,
			UserID:      user.UserID,
			Accounttype: user.Accounttype,
			// Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			LogError.LogError("❌🔥 error in StatusInternalServerError token generation  ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

			return
		}
		////////////////////////////////
		var SendJwt struct {
			JWT string
		}
		SendJwt.JWT = tokenString

		c.JSON(http.StatusOK, SendJwt)

	}
}

//func BcryptHashPass(s string) (string, error){
//	b,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.DefaultCost)
//	LogError.LogError("❌ ~ file: signup.go ~ line 55 ~ funcBcryptHashPass ~ err : ", err)
//	if err!=nil {
//		return "",err
//	}else {
//		return string(b),nil
//	}
//}
//
//func BcryptVerifyPass(userpass string, Databasepass string) (string,bool){
//	fmt.Println("🚀 ~ file: register.go ~ line 116 ~ funcBcryptVerifyPass ~ Databasepass : ", Databasepass)
//	fmt.Println("🚀 ~ file: register.go ~ line 116 ~ funcBcryptVerifyPass ~ userpass : ", userpass)
//	err:=bcrypt.CompareHashAndPassword([]byte(Databasepass),[]byte(userpass))
//	LogError.LogError("❌ ~ file: register.go ~ line 119 ~ funcBcryptVerifyPass ~ err : ", err)
//	if err!=nil{
//		msg:="Email or Password is incorrect"
//		return msg, false
//	} else {
//		return "",true
//	}
//}
