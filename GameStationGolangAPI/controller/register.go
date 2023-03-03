package controller

import (
	"app/LogError"
	"app/model"
	// "app/token"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

func (H *DatabaseCollections)Register() gin.HandlerFunc {
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


		} else {
			user.Accounttype = "normal"
			myid := uuid.New()
			user.UserID = myid.String()

		}
		user.BannerImg = ""
		user.Coin = 0.0
		user.ProfileImg = ""

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
		fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 115 ~ returnfunc ~ res : ", res)
		if err != nil {
			LogError.LogError("❌🔥 error in InsertOne() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//  AUTH1 GENEGRATION
		expirationTime := time.Now().Add(time.Hour * 100)
		claims := &model.TokenClaims{
			Name:        user.Name,
			Email:       user.Email,
			UserID:      user.UserID,
			Accounttype: user.Accounttype,

			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(os.Getenv("COOKIE_SECRET_JWT_AUTH1")))

		if err != nil {
			LogError.LogError("❌🔥 error in StatusInternalServerError token generation  ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

			return
		}
	
		//  noAUTH1 GENEGRATION
		expirationTime2 := time.Now().Add(time.Hour * 100)
		claims2 := &model.TokenClaims{
			Name:        user.Name,
			Email:       user.Email,
			UserID:      user.UserID,
			Accounttype: user.Accounttype,

			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime2.Unix(),
			},
		}

		token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)

		tokenString2, err := token2.SignedString([]byte(os.Getenv("COOKIE_SECRET_NOAUTH1")))

		if err != nil {
			LogError.LogError("❌🔥 error in StatusInternalServerError token generation  ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

			return
		}

		c.SetCookie("Auth1",tokenString,60*60*24,"/",os.Getenv("DOMAIN_ADDR"),false , true)
		c.SetCookie("noAuth1",tokenString2,60*60*24,"/",os.Getenv("DOMAIN_ADDR"),false , false)
		fmt.Println("😍 Register Successfull")

		c.JSON(http.StatusOK,gin.H{"message":"successfully signed up"})
	}
}


