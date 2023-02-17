package controller

import (
	"app/LogError"
	"app/model"
	"app/token"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections)Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusOK, "Login successful !!! ")

		ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
		// ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		var user model.User
		var DBUser model.User
	
		
		err:=c.BindJSON(&user)
        LogError.LogError("❌ ~ file: login.go ~ line 23 ~ returnfunc ~ err : ", err)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		} 
		// MONGO FIND ONE
		err=H.Mongo.Collection("usercol").FindOne(ctx,bson.M{"email":user.Email}).Decode(&DBUser)
        LogError.LogError("❌ ~ file: login.go ~ line 31 ~ returnfunc ~ err : ", err)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"email or password incorrect"})
			return 
		}

		validitymsg,b:=BcryptVerifyPass(user.Password,DBUser.Password)
		if !b{
			c.JSON(http.StatusInternalServerError,gin.H{"error":validitymsg})
            fmt.Println("🚀 ~ file: login.go ~ line 40 ~ returnfunc ~ validitymsg : ", validitymsg)
			return 
		}
		token,refreshToken,err := token.TokenGenerate(DBUser.Email,DBUser.FirstName,DBUser.LastName,DBUser.UserID)
		c.SetCookie("Auth1",token,60*60*24,"/",os.Getenv("DOMAIN_ADDR"),false , true)
		c.SetCookie("Auth1Refresh",refreshToken,60*60*24,"/",os.Getenv("DOMAIN_ADDR"),false , true)
        LogError.LogError("❌ ~ file: login.go ~ line 48 ~ returnfunc ~ err : ", err)
        fmt.Println("🚀 ~ file: login.go ~ line 46 ~ returnfunc ~ refreshToken : ", refreshToken)
        fmt.Println("🚀 ~ file: login.go ~ line 46 ~ returnfunc ~ token : ", token)
		c.JSON(http.StatusOK,DBUser)

		
	}
}
