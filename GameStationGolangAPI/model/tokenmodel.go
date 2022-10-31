package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Email       string `json:"Email"`
	Name        string `json:"Name"`
	UserID      string `json:"UserID"`
	Accounttype string `json:"Accounttype"`
	jwt.StandardClaims
}
