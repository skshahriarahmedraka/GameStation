package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserData struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	UserID string             `json:"UserID" bson:"UserID"`
	//Image  string             `json:"Image" bson:"Image"`
	Name string `json:"Name" validate:"required,min=3,max=30" bson:"Name"`
	//LastName  string             `json:"LastName" validate:"required,min=2,max=30"`
	Email    string `json:"Email" validate:"required,min=2,max=40" bson:"Email"`
	Password string `json:"Password" validate:"required,min=3,max=50" bson:"Password"`
	Mobile   string `json:"Mobile" bson:"Mobile"`

	Address string `json:"Address" bson:"Address"`
	City    string `json:"City" bson:"City"`
	Country string `json:"Country"  bson:"Country"`
	ZipCode string `json:"ZipCode" bson:"ZipCode"`

	Coin float64 `json:"Coin" bson:"Coin"`

	ProfileImg         string   `gorm:"type:varchar(200);not null" json:"ProfileImg" bson:"ProfileImg"`
	BannerImg          string   `gorm:"type:varchar(200);not null" json:"BannerImg" bson:"BannerImg"`
	Accounttype        string   `gorm:"type:varchar(20);not null" json:"Accounttype" bson:"Accounttype"`
	TransactionHistory []string `json:"TransactionHistory" bson:"TransactionHistory"`
	WishList           []string `json:"WishList" bson:"WishList"`

	ContactAdminMsg []string `json:"ContactAdminMsg" bson:"ContactAdminMsg"`
	ContactDevMsg   []string `json:"ContactDevMsg" bson:"ContactDevMsg"`

	// Token 	*string 	`json:"token"`
	// RefreshToken *string `json:"refreshtoken"`
	// CreatedAt   time.Time     `json:"created_at"`
	// UpdatedAt   time.Time     `json:"updated_at"`
	UserCart []string `json:"UserCart" bson:"UserCart"`
	// OrderList []string       `json:"orderstatus" bson:"orderstatus"`
}
