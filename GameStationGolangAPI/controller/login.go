package controller

import (
	"app/errors"
	"app/model"
	"app/token"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusOK, "Login successful !!! ")

		ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
		// ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
		defer cancel()

		var user model.User
		var DBUser model.User
	
		
		err:=c.BindJSON(&user)
        errors.ERROR("🚀 ~ file: login.go ~ line 23 ~ returnfunc ~ err : ", err)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		} 
		err=UserCollection.FindOne(ctx,bson.M{"email":user.Email}).Decode(&DBUser)
        errors.ERROR("🚀 ~ file: login.go ~ line 31 ~ returnfunc ~ err : ", err)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"email or password incorrect"})
			return 
		}

		validitymsg,b:=BcryptVerifyPass(*user.Password,*DBUser.Password)
		if !b{
			c.JSON(http.StatusInternalServerError,gin.H{"error":validitymsg})
            fmt.Println("🚀 ~ file: login.go ~ line 40 ~ returnfunc ~ validitymsg : ", validitymsg)
			return 
		}
		token,refreshToken,err := token.TokenGenerate(*DBUser.Email,*DBUser.FirstName,*DBUser.LastName,*DBUser.UserID)
        fmt.Println("🚀 ~ file: login.go ~ line 46 ~ returnfunc ~ refreshToken : ", refreshToken)
        fmt.Println("🚀 ~ file: login.go ~ line 46 ~ returnfunc ~ token : ", token)
		c.JSON(http.StatusOK,DBUser)

		
	}
}
