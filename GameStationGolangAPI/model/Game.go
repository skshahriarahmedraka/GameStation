package model

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Gamedata struct {
	GameID           string   `json:"GameID" bson:"GameID"`
	Name             string   `json:"Name" bson:"Name"`
	Moto             string   `json:"Moto" bson:""`
	LogoImage        string   `json:"LogoImage" bson:"LogoImage"`
	BigPosterImage   string   `json:"BigPosterImage" bson:"BigPosterImage"`
	SmallPosterImage string   `json:"SmallPosterImage" bson:"SmallPosterImage"`
	OtherImages      []string `json:"OtherImages" bson:"OtherImages"`
	Genres           []string `json:"Genres" bson:"Genres"`
	Feature          []string `json:"Feature" bson:"Feature"`
	Description      string   `json:"Description" bson:"Description"`
	FollowUs         struct {
		Facebook string `json:"Facebook" bson:"Facebook"`
		Youtube  string `json:"Youtube" bson:"Youtube"`
		Discord  string `json:"Discord" bson:"Discord"`
		Twitter  string `json:"Twitter" bson:"Twitter"`
		Site     string `json:"Site" bson:"Site"`
	} `json:"FollowUs" bson:"FollowUs"`
	Rating        int `json:"Rating" bson:"Rating"`
	RatingGivenBy struct {
		PC_Gamer      string `json:"PC Gamer" bson:"PC Gamer"`
		IGN           string `json:"IGN" bson:"IGN"`
		Game_Informer string `json:"Game Informer" bson:"Game Informer"`
	} `json:"RatingGivenBy" bson:"RatingGivenBy"`
	Minspec []struct {
		Name  string `json:"Name" bson:"Name"`
		Value string `json:"Value" bson:"Value"`
	} `json:"Minspec" bson:"Minspec"`
	Recomendedspec []struct {
		Name  string `json:"Name" bson:"Name"`
		Value string `json:"Value" bson:"Value"`
	} `json:"Recomendedspec" bson:"Recomendedspec"`
	Price     float64        `json:"Price" bson:"Price"`
	Discount  int            `json:"Discount" bson:"Discount"`
	Developer string         `json:"Developer" bson:"Developer"`
	Publisher string         `json:"Publisher" bson:"Publisher"`
	Released  string         `json:"Released" bson:"Released"`
	Platform  []string       `json:"Platform" bson:"Platform"`
	Players   int            `json:"Players" bson:"Players"`
	Comment   []CommentModel `json:"Comment" bson:"Comment"`
}

type CommentModel struct {
	Name        string `json:"Name"`
	Rating      int    `json:"Rating"`
	Description string `json:"Description"`
}
