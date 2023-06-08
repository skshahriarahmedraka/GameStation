

export type GameDataType =  {
    GameID: string;
    Name: string;
    Moto: string;
    LogoImage: string;
    BigPosterImage: string;
    SmallPosterImage: string;
    OtherImages: string[];
    Genres: string[];
    Feature: string[];
    Description: string;
    Rating: number;
    RatingGivenBy: {
        'PC Gamer': string;
        IGN: string;
        'Game Informer': string;
    };
    Minspec: {
        Name: string;
        Value: string;
    }[];
    Recomendedspec: {
        Name: string;
        Value: string;
    }[];
    Discount: number;
    Price: number;
    Developer: string;
    Publisher: string;
    Released: string;
    Platform: string[];
    Players: number;
    Comment: {
        Name: string;
        Rating: number;
        Description: string;
    }[];
    FollowUs: {
        Facebook: string;
        Discord: string;
        Youtube: string;
        Twitter: string;
        Site: string;
    };
}[]

export type UserDataType = {
	Name: string;
	UserID: string;
	Accounttype: string;
	ProfileImg: string;
	BannerImg: string;
	Email: string;
	Mobile: string;
	Address: string;
	Country: string;
	City: string;
	ZipCode: string;
	Coin: number;
	TransactionHistory: string[];
	WishList: string[];
	ContactAdminMsg: string[];
	ContactDevMsg: string[];
	UserCart: string[];
}

export type CookieInfo1Type=  {
    Email    :   string 
    Name      :  string 
    UserID   :  string 
    Accounttype :string 
}