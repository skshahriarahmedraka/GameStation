package model

type RegModel struct {
	UserName string `json:"UserName" bson:"UserName"`
	Email    string `json:"Email" bson:"Email"`
	Password string `json:"Password" bson:"Password"`
}
