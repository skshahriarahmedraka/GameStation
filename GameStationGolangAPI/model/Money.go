package model

import "github.com/golang-jwt/jwt/v4"

// type AdminMoneyManagement struct {
// 	ID      primitive.ObjectID `json:"_id" bson:"_id"`
// 	AdminID string             `json:"AdminID" bson:"AdminID"`
// 	UserList []UserMoney           `json:"UserList" bson:"UserList"`
// }

type UserMoney struct {
	UserID  string         `json:"UserID" bson:"UserID"`
	Coin    float64        `json:"Coin" bson:"Coin"`
	ReqList []UserMoneyReq `json:"ReqList" bson:"ReqList"`
}
type UserMoneyReq struct {
	Amount float64 `json:"Amount" bson:"Amount"`
	JWT    string  `json:"JWT" bson:"JWT"`
	Status string  `json:"Status" bson:"Status"` // true = already used token, false = Not used token
}

type JWTModel struct {
	UserID string  `json:"UserID" bson:"UserID"`
	Amount float64 `json:"Amount" bson:"Amount"`
	jwt.StandardClaims
}
