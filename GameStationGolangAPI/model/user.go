package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserID      string        `json:"userid"`
	FirstName string             `json:"firstname" validate:"required,min=3,max=30"`
	LastName  string             `json:"lastlame" validate:"required,min=2,max=30"`
	Email     string             `json:"email" validate:"required,min=2,max=40"`
	Password  string             `json:"password" validate:"required.min=3,max=50"`
	Mobile    string             `json:"mobile"`
	// Token 	*string 	`json:"token"`
	// RefreshToken *string `json:"refreshtoken"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	UserCart    []UserProduct `json:"usercart" bson:"usercart"`
	Address     []Address     `json:"address" bson:"address"`
	OrderStatus []Order       `json:"orderstatus" bson:"orderstatus"`
}

type Product struct {
	ProductID   primitive.ObjectID `json:"_id" bson:"_id"`
	ProductName *string            `json:"productname" bson:"productname"`
	Price       *uint64            `json:"price" bson:"price"`
	Rating      *uint8             `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}

type UserProduct struct {
	ProductID   primitive.ObjectID `json:"_id" bson:"_id"`
	ProductName *string            `json:"productname" bson:"productname"`
	Price       *uint64            `json:"price" bson:"price"`
	Rating      *uint8             `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}

type Address struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	House   *string            `json:"house" bson:"house"`
	Street  *string            `json:"street" bson:"street"`
	City    *string            `json:"city" bson:"city"`
	Country *string            `json:"country" bson:"country"`
	Pincode *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	OrderID       primitive.ObjectID `json:"_id" bson:"_id"`
	OrderCart     []UserProduct      `json:"ordercart" bson:"ordercart"`
	OrderedAt     time.Time          `json:"orderedat" bson:"orderedat"`
	Price         uint64             `json:"price" bson:"price"`
	Discount      uint64             `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"paymentmethod" bson:"paymentmethod"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"`
}
