package controller

import (
	"app/LogError"
	"app/model"
	// "app/token"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (H *DatabaseCollections) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		var user model.LoginModel
		var dbUser model.UserData

		err := c.BindJSON(&user)
		if err != nil {
			LogError.LogError("❌🔥 error in c.bindjson() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			return
		}
		fmt.Println("🚀", user)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		// SEARCH EMAIL
		err = H.Mongo.Collection(os.Getenv("USERDATA_COL")).FindOne(ctx, bson.M{"Email": user.Email}).Decode(&dbUser)
		if err != nil {
			LogError.LogError("❌🔥 error in mongodb connection  ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
			return
		}

		userPass := []byte(user.Password)
		dbPass := []byte(dbUser.Password)
		passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
		if passErr != nil {
			LogError.LogError("❌🔥 error in bcrypt error ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//  AUTH1 GENEGRATION
		expirationTime := time.Now().Add(time.Hour * 100)
		claims := &model.TokenClaims{
			Name:        dbUser.Name,
			Email:       dbUser.Email,
			UserID:      dbUser.UserID,
			Accounttype: dbUser.Accounttype,

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
			Name:        dbUser.Name,
			Email:       dbUser.Email,
			UserID:      dbUser.UserID,
			Accounttype: dbUser.Accounttype,

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

		c.SetCookie("Auth1", tokenString, 60*60*24, "/", os.Getenv("DOMAIN_ADDR"), false, true)
		c.SetCookie("noAuth1", tokenString2, 60*60*24, "/", os.Getenv("DOMAIN_ADDR"), false, false)
		fmt.Println("😍 Login Successfull")
		c.JSON(http.StatusOK, dbUser)

	}
}
