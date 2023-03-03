// import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

type TypeGameData = {
	GameID: '';
	Name: '';
	Moto: '';
	LogoImage: '';
	BigPosterImage: '';
	SmallPosterImage: '';
	OtherImages: [];
	Genres: [''];
	Feature: [''];
	Description: '';
	FollowUs: {
		Facebook: '';
		Discord: '';
		Youtube: '';
		Twitter: '';
		Site: '';
	};
	Rating: 0;
	RatingGivenBy: {
		'PC Gamer': '';
		IGN: '';
		'Game Informer': '';
	};
	Minspec: [
		{
			Name: 'OS';
			Value: '';
		},
		{
			Name: 'Storage';
			Value: '';
		},
		{
			Name: 'Memory';
			Value: '';
		},
		{
			Name: 'CPU';
			Value: '';
		},
		{
			Name: 'GPU';
			Value: '';
		}
	];
	Recomendedspec: [
		{
			Name: 'OS';
			Value: '';
		},
		{
			Name: 'Storage';
			Value: '';
		},
		{
			Name: 'Memory';
			Value: '';
		},
		{
			Name: 'GPU';
			Value: '';
		},
		{
			Name: 'CPU';
			Value: '';
		}
	];
	Price: 0;
	Discount: 0;
	Developer: '';
	Publisher: '';
	Released: '';
	Platform: ['', '', ''];
	Players: 0;
	Comment: [];
} 
type TypeUserData = {
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
export const load:PageLoad = async ({ fetch }) => {

    let MostPopularList: TypeGameData[] = [] as TypeGameData[];
    
    console.log("🚀 +page.ts MostPopularList: ","http://${import.meta.env.VITE_GO_HOST}/api/game/mostpopular", MostPopularList)
	await fetch(`http://${import.meta.env.VITE_GO_HOST}/api/game/mostpopular`, {
        method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
            MostPopularList = data;
			console.log(
                '🚀 ~ MostPopularList',
				MostPopularList
                );
            });
            
            console.log("🚀 +page.ts MostPopularList:", MostPopularList)
    
			const NewReleaseList: TypeGameData[] = [] as TypeGameData[];
			const TopRatedList: TypeGameData[] = [] as TypeGameData[];
			const TopSoledList: TypeGameData[] = [] as TypeGameData[];
			const TrendingList: TypeGameData[] = [] as TypeGameData[];
			const CarouselList: TypeGameData[] = [] as TypeGameData[];

			const Userdata: TypeUserData = {} as TypeUserData;

    return {
		Userdata,
        MostPopularList,
		NewReleaseList,
		TopRatedList,
		TopSoledList,
		TrendingList,
		CarouselList
    }
 
//   throw error(404, 'Not found');
}
