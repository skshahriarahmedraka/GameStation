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
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusCreated, "Successfully Signed Up!!")
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var u model.User
		err := c.BindJSON(&u)
		errors.ERROR("🚀 ~ file: signup.go ~ line 22 ~ returnfunc ~ err : ", err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = Validate.Struct(u)
		errors.ERROR("🚀 ~ file: signup.go ~ line 27 ~ returnfunc ~ err : ", err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// SEARCH EMAIL
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": u.Email})
		errors.ERROR("🚀 ~ file: signup.go ~ line 36 ~ returnfunc ~ err : ", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User Already register"})
			return
		}
		//MOBILE 
		count, err = UserCollection.CountDocuments(ctx, bson.M{"mobile": u.Mobile})
        errors.ERROR("🚀 ~ file: signup.go ~ line 50 ~ returnfunc ~ err : ", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mobile already in use"})
			return
		}

		hashpass,err:= BcryptHashPass(*u.Password)
        errors.ERROR("🚀 ~ file: signup.go ~ line 61 ~ returnfunc ~ err : ", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
			return
		}
		u.Password =&hashpass
		u.CreatedAt ,err =time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
        errors.ERROR("🚀 ~ file: signup.go ~ line 68 ~ returnfunc ~ err : ", err)
		u.UpdatedAt,err=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
        errors.ERROR("🚀 ~ file: signup.go ~ line 70 ~ returnfunc ~ err : ", err)
		u.ID=primitive.NewObjectID()
		*u.UserID = u.ID.Hex()

		jwtToken,refreshToken,err:= token.TokenGenerate(*u.Email,*u.FirstName,*u.LastName, *u.UserID)
        errors.ERROR("🚀 ~ file: signup.go ~ line 77 ~ returnfunc ~ err : ", err)

		u.Token=&jwtToken
		u.RefreshToken=&refreshToken
		u.UserCart=make([]model.UserProduct, 0)
		u.Address=make([]model.Address, 0)
		u.OrderStatus=make([]model.Order, 0)
		res,err:=UserCollection.InsertOne(ctx,u)
		errors.ERROR("🚀 ~ file: signup.go ~ line 83 ~ returnfunc ~ err : ", err)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"not created"})
			return 
		}else {

			fmt.Println("🚀 ~ file: signup.go ~ line 83 ~ returnfunc ~ res : ", res)
		}
		c.JSON(http.StatusOK,"successfully signed up")
	}
}


func BcryptHashPass(s string) (string,error){
	b,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.DefaultCost)
    errors.ERROR("🚀 ~ file: signup.go ~ line 55 ~ funcBcryptHashPass ~ err : ", err)
	if err!=nil {
		return "",err
	}else {
		return string(b),nil
	}
}

func BcryptVerifyPass(userpass string, givenpass string) (string,bool){
	err:=bcrypt.CompareHashAndPassword([]byte(userpass),[]byte(givenpass))
    errors.ERROR("🚀 ~ file: signup.go ~ line 65 ~ funcBcryptVerifyPass ~ err : ", err)
	if err!=nil{
		msg:="Email or Password is incorrect"
		return msg, false 
	} else {
		return "",true 
	}
}