package token

import (
	"app/errors"
	"app/model"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func TokenGenerate(email, firstName, lastName, UID string) (string, string, error) {
	claims := model.SignDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UID:       UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	RefreshClaims := model.SignDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	errors.ERROR("🚀 ~ file: tokenGenerate.go ~ line 29 ~ funcTokenGenerater ~ err : ", err)

	if err != nil {
		return "", "", err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims).SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	errors.ERROR("🚀 ~ file: tokenGenerate.go ~ line 37 ~ funcTokenGenerater ~ err : ", err)
	if err != nil {
		return "", "", err
	}
	return token, refreshtoken, err
}

func ValidateJWT(s string) (claims *model.SignDetails, msg string) {
	token, err := jwt.ParseWithClaims(s, &model.SignDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	errors.ERROR("🚀 ~ file: tokenGenerate.go ~ line 45 ~ token,err:=jwt.ParseWithClaims ~ err : ", err)

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*model.SignDetails)
	if !ok {
		msg = "token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg

}

func ValidateRefreshJWT(s string) (claims *model.SignDetails, msg string) {
	token, err := jwt.ParseWithClaims(s, &model.SignDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})
	errors.ERROR("🚀 ~ file: tokenGenerate.go ~ line 68 ~ token,err:=jwt.ParseWithClaims ~ err : ", err)

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*model.SignDetails)
	if !ok {
		msg = "token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg

}

func UpdateAllToken(signedtoken string, refreshtoken string, userid string){
	ctx,cancel:= context.WithTimeout(context.Background(),100*time.Second)
    fmt.Println("🚀 ~ file: tokenGenerate.go ~ line 95 ~ funcUpdateAllToken ~ ctx : ", ctx)
	defer cancel()
	var updateObj primitive.D
	updateObj=append(updateObj, bson.E{Key:"token",Value:signedtoken})
	updateObj=append(updateObj, bson.E{Key:"refreshToken",Value:refreshtoken})
	updatedAt,err:= time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
    fmt.Println("🚀 ~ file: tokenGenerate.go ~ line 99 ~ funcUpdateAllToken ~ updatedAt : ", updatedAt)
    errors.ERROR("🚀 ~ file: tokenGenerate.go ~ line 99 ~ funcUpdateAllToken ~ err : ", err)
	if err!=nil {
		return
	}
}