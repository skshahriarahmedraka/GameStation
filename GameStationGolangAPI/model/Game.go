package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gamedata struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	GameID           string             `json:"GameID" bson:"GameID"`
	Name             string             `json:"Name" bson:"Name"`
	Moto             string             `json:"Moto" bson:""`
	LogoImage        string             `json:"LogoImage" bson:"LogoImage"`
	BigPosterImage   string             `json:"BigPosterImage" bson:"BigPosterImage"`
	SmallPosterImage string             `json:"SmallPosterImage" bson:"SmallPosterImage"`
	OtherImages      []string           `json:"OtherImages" bson:"OtherImages"`
	Genres           []string           `json:"Genres" bson:"Genres"`
	Feature          []string           `json:"Feature" bson:"Feature"`
	Description      string             `json:"Description" bson:"Description"`
	FollowUs         struct {
		Facebook string `json:"" bson:""`
		Youtube  string `json:"" bson:""`
		Discord  string `json:"" bson:""`
		Twitter  string `json:"" bson:""`
		Site     string `json:"" bson:""`
	} `json:"FollowUs" bson:"FollowUs"`
	Rating        int `json:"" bson:""`
	RatingGivenBy struct {
		PC_Gamer      string `json:"" bson:""`
		IGN           string `json:"" bson:""`
		Game_Informer string `json:"" bson:""`
	} `json:"" bson:""`
	Minspec []struct {
		Name  string `json:"" bson:""`
		Value string `json:"" bson:""`
	} `json:"" bson:""`
	Recomendedspec []struct {
		Name  string `json:"" bson:""`
		Value string `json:"" bson:""`
	} `json:"" bson:""`
	Price     int      `json:"" bson:""`
	Discount  int      `json:"" bson:""`
	Developer string   `json:"" bson:""`
	Publisher string   `json:"" bson:""`
	Released  string   `json:"" bson:""`
	Platform  []string `json:"" bson:""`
	Players   int      `json:"" bson:""`
}