package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserData struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserID    string             `json:"Userid"`
	Image     string             `json:""`
	FirstName string             `json:"firstName" validate:"required,min=3,max=30"`
	LastName  string             `json:"LastName" validate:"required,min=2,max=30"`
	Email     string             `json:"Email" validate:"required,min=2,max=40"`
	Password  string             `json:"Password" validate:"required,min=3,max=50"`
	Mobile    string             `json:"Mobile"`

	Address string `json:"Address" bson:"Address"`
	City    string `json:"City"`
	Country string `json:"Country"`
	ZipCode string `json:"ZipCode"`

	Coin string `json:"Coin"`

	TransactionHistory []string `json:"TransactionHistory"`
	WishList           []string `json:"WishList"`

	ContactAdminMsg []string `json:"ContactAdminMsg"`
	ContactDevMsg   []string `json:"ContactDevMsg"`

	// Token 	*string 	`json:"token"`
	// RefreshToken *string `json:"refreshtoken"`
	// CreatedAt   time.Time     `json:"created_at"`
	// UpdatedAt   time.Time     `json:"updated_at"`
	UserCart    []string `json:"UserCart" bson:"UserCart"`
	// OrderList []string       `json:"orderstatus" bson:"orderstatus"`
}
