package model

type RegModel struct {
	Name string `json:"Name" bson:"Name"`
	Email    string `json:"Email" bson:"Email"`
	Password string `json:"Password" bson:"Password"`
}
