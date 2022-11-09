package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func (H *DatabaseCollections) SveltekitLogin() gin.HandlerFunc {
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

		// if err!=nil{
		// 	response.WriteHeader(http.StatusInternalServerError)
		// 	response.Write([]byte(`{"message":"`+err.Error()+`"}`))
		// 	return
		//  }
		userPass := []byte(user.Password)
		dbPass := []byte(dbUser.Password)
		passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
		if passErr != nil {
			LogError.LogError("❌🔥 error in bcrypt error ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//  jwtToken, err := GenerateJWT()
		//  if err != nil{
		//  response.WriteHeader(http.StatusInternalServerError)
		//  response.Write([]byte(`{"message":"`+err.Error()+`"}`))
		//  return
		//  }
		//  response.Write([]byte(`{"token":"`+jwtToken+`"}`))
		// if user.ID != 0  {
		//CREATE COOKIE
		expirationTime := time.Now().Add(time.Hour * 1000)
		myClaim := &model.TokenClaims{
			Name:        dbUser.Name,
			Email:       dbUser.Email,
			UserID:      dbUser.UserID,
			Accounttype: dbUser.Accounttype,

			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		// }
		// LOGIN SUCCESSFUL

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			LogError.LogError("❌🔥 error in token.SignedString ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
