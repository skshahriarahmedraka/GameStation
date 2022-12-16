/* eslint-disable @typescript-eslint/no-explicit-any */
// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import * as jwt from 'jsonwebtoken';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ cookies,params,locals }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }


	
	  let DefaultGameList: {
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
	}[] = [] as  {
		GameID: '',
		Name: '',
		Moto: '',
		LogoImage: '',
		BigPosterImage: '',
		SmallPosterImage: '',
		OtherImages: [],
		Genres: [''],
		Feature: [''],
		Description: '',
		FollowUs: {
			Facebook: '',
			Discord: '',
			Youtube: '',
			Twitter: '',
			Site: ''
		},
		Rating: 0,
		RatingGivenBy: {
			'PC Gamer': '',
			IGN: '',
			'Game Informer': ''
		},
		Minspec: [
			{
				Name: 'OS',
				Value: ''
			},
			{
				Name: 'Storage',
				Value: ''
			},
			{
				Name: 'Memory',
				Value: ''
			},
			{
				Name: 'CPU',
				Value: ''
			},
			{
				Name: 'GPU',
				Value: ''
			}
		],
		Recomendedspec: [
			{
				Name: 'OS',
				Value: ''
			},
			{
				Name: 'Storage',
				Value: ''
			},
			{
				Name: 'Memory',
				Value: ''
			},
			{
				Name: 'GPU',
				Value: ''
			},
			{
				Name: 'CPU',
				Value: ''
			}
		],
		Price: 0,
		Discount: 0,
		Developer: '',
		Publisher: '',
		Released: '',
		Platform: ['', '', ''],
		Players: 0,
		Comment: []
	}[]
	  await fetch(`http://${process.env.GO_HOST}/user/browse`,{
		  method: "GET",
		  mode: "no-cors",

	  }).then((res)=>{
		  return res.json()
	  }).then((data)=>{
		DefaultGameList=data
	  console.log("🚀 ~ file: +page.server.ts ~ line 103 ~ constload:PageServerLoad= ~ DefaultGameList", DefaultGameList)
	  })
	
	  let SearchGameList: {
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
	}[] = [] as  {
		GameID: '',
		Name: '',
		Moto: '',
		LogoImage: '',
		BigPosterImage: '',
		SmallPosterImage: '',
		OtherImages: [],
		Genres: [''],
		Feature: [''],
		Description: '',
		FollowUs: {
			Facebook: '',
			Discord: '',
			Youtube: '',
			Twitter: '',
			Site: ''
		},
		Rating: 0,
		RatingGivenBy: {
			'PC Gamer': '',
			IGN: '',
			'Game Informer': ''
		},
		Minspec: [
			{
				Name: 'OS',
				Value: ''
			},
			{
				Name: 'Storage',
				Value: ''
			},
			{
				Name: 'Memory',
				Value: ''
			},
			{
				Name: 'CPU',
				Value: ''
			},
			{
				Name: 'GPU',
				Value: ''
			}
		],
		Recomendedspec: [
			{
				Name: 'OS',
				Value: ''
			},
			{
				Name: 'Storage',
				Value: ''
			},
			{
				Name: 'Memory',
				Value: ''
			},
			{
				Name: 'GPU',
				Value: ''
			},
			{
				Name: 'CPU',
				Value: ''
			}
		],
		Price: 0,
		Discount: 0,
		Developer: '',
		Publisher: '',
		Released: '',
		Platform: ['', '', ''],
		Players: 0,
		Comment: []
	}[]


	  await fetch(`http://${process.env.GO_HOST}/search/${params.search}`,{
		mode:"no-cors"
	}).then((res)=>{
	 return res.json()
	}).then((d)=>{
		SearchGameList=d
		console.log("🚀 ~ file: +page.server.ts ~ line 279 ~ constload:PageServerLoad= ~ SearchGameList", SearchGameList)
	})


	const MyCookie = cookies.get('Auth1') || '';
	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	let Userdata: {
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
	} = {
		Name: '',
		UserID: '',
		Accounttype: '',

		ProfileImg: '',
		BannerImg: '',

		Email: '',
		Mobile: '',

		Address: '',
		// Region: '',
		Country: '',
		City: '',
		ZipCode: '',

		Coin: 0,

		TransactionHistory: [] as string[],
		WishList: [] as string[],
		ContactAdminMsg: [] as string[],
		ContactDevMsg: [] as string[],
		UserCart: [] as string[]
	};

	if (MyCookie != '') {
		interface tokeninterface{
			Email: string 
			Name: string,
			UserID: string,
			Accounttype: string,
			exp: number
		  }

		const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
		

		console.log(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UserID}`);
		await fetch(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UserID}`, {
			mode: 'no-cors'
		})
			.then((res) => {
				return res.json();
			})
			.then((da) => {
				// console.log('🚀 ~ file: +page.server.ts ~ line 106 ~ myfetch ~ da', da);

				Userdata = da;
			});
	}
	

	return {Userdata,DefaultGameList,SearchGameList};
	// if (mydata){

		
	// }
	// else {
	// 	throw error(404, 'Not found');
	// }result
}
