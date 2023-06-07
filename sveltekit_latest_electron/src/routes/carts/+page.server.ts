// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
// import {UserProData} from '../../lib/Store/store'
import * as jwt from 'jsonwebtoken';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ cookies,locals }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }

	  const MyCookie = cookies.get('Auth1') || '';
	  const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
  
	  let Userdata
	  : {
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

	  let CartGameList: {
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
	}[]=[] as  {
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
	interface tokeninterface{
		Email: string 
		Name: string,
		UserID: string,
		Accounttype: string,
		exp: number
	  }

	const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	
	const reqStruct = {
		UserID: (decoded as tokeninterface).UserID
	}

	// console.log(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UserID}`);
	  await fetch(`http://${process.env.GO_HOST}/user/cart`, {
			method: 'POST',
			  mode: 'no-cors',
			body : JSON.stringify(reqStruct)
		  })
			  .then((res) => {
				  return res.json();
			  })
			  .then((da) => {
				  // console.log('🚀 ~ file: +page.server.ts ~ line 106 ~ myfetch ~ da', da);
  
				  CartGameList = da;
			  });
	  return {
		  Userdata,
		  CartGameList
	  };

	
	
	
}
