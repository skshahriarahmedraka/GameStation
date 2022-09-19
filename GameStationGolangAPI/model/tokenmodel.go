package model


import "github.com/golang-jwt/jwt/v4"



type SignDetails struct {
	Email string 
	FirstName string 
	LastName string 
	UID string 
	jwt.StandardClaims
}
