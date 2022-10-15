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
		// c.JSON(http.StatusOK, "Login successful !!! ")

		//ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
		//// ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
		//defer cancel()
		//
		//var user model.User
		//var DBUser model.User
		//
		//
		//err:=c.BindJSON(&user)
		//LogError.LogError("❌ ~ file: login.go ~ line 23 ~ returnfunc ~ err : ", err)
		//if err!=nil {
		//	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		//	return
		//}
		//// MONGO FIND ONE
		//err=H.Mongo.Collection("usercol").FindOne(ctx,bson.M{"email":user.Email}).Decode(&DBUser)
		//LogError.LogError("❌ ~ file: login.go ~ line 31 ~ returnfunc ~ err : ", err)
		//if err!=nil {
		//	c.JSON(http.StatusInternalServerError,gin.H{"error":"email or password incorrect"})
		//	return
		//}
		//
		//validitymsg,b:=BcryptVerifyPass(user.Password,DBUser.Password)
		//if !b{
		//	c.JSON(http.StatusInternalServerError,gin.H{"error":validitymsg})
		//	fmt.Println("🚀 ~ file: login.go ~ line 40 ~ returnfunc ~ validitymsg : ", validitymsg)
		//	return
		//}
		//token,refreshToken,err := token.TokenGenerate(DBUser.Email,DBUser.FirstName,DBUser.LastName,DBUser.UserID)
		//c.SetCookie("Auth1",token,60*60*24,"/","localhost",false , true)
		//c.SetCookie("Auth1Refresh",refreshToken,60*60*24,"/","localhost",false , true)
		//LogError.LogError("❌ ~ file: login.go ~ line 48 ~ returnfunc ~ err : ", err)
		//fmt.Println("🚀 ~ file: login.go ~ line 46 ~ returnfunc ~ refreshToken : ", refreshToken)
		//fmt.Println("🚀 ~ file: login.go ~ line 46 ~ returnfunc ~ token : ", token)
		//c.JSON(http.StatusOK,DBUser)

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		var user model.LoginModel
		var dbUser model.UserData
		// err := json.NewDecoder(r.Body).Decode(&user)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ l : ", user.Password)

		//err := json.NewDecoder(r.Body).Decode(&user)
		err := c.BindJSON(&user)
		if err != nil {
			LogError.LogError("❌🔥 error in c.bindjson() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			//http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("🚀", user)

		//var resError model.ErrorRes

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		// SEARCH EMAIL
		// count, err := H.Mongo.Collection("usercol").CountDocuments(ctx, bson.M{"Email": user.Email})
		err = H.Mongo.Collection("usercol").FindOne(ctx, bson.M{"Email": user.Email}).Decode(&dbUser)
		if err != nil {
			LogError.LogError("❌🔥 error in mongodb connection  ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
			Name:   dbUser.Name,
			Email:  dbUser.Email,
			UserID: dbUser.UserID,
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
