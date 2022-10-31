package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewsModel struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	NewsID        string             `json:"NewsID" bson:"NewsID"`
	Title         string             `json:"Title" bson:"Title"`
	BannerImg     string             `json:"BannerImg" bson:"BannerImg"`
	Date          string             `json:"Date" bson:"Date"`
	WrittenBy     string             `json:"WrittenBy" bson:"WrittenBy"`
	Detail        string             `json:"Detail" bson:"Detail"`
	FooterComment string             `json:"FooterComment" bson:"FooterComment"`
}
