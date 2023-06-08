// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import * as jwt from 'jsonwebtoken';

export const load: PageServerLoad = async ({ cookies, locals }) => {
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 10 ~ constload:PageServerLoad= ~ user',
			locals.user.Authenticated
		);
		throw redirect(302, '/login');
	}

	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies);
	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies.get('Auth1'));
	//   const MyCookie = cookies.get('Auth1') || '';
	//   const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	//   let Userdata:any
	//   if (MyCookie!= ''){

	// 	  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// 	  console.log("decoded: ",decoded);
	// 	//   let resdata
	// 	  console.log(`http://${process.env.GO_HOST}/user/${decoded.UserID}`)
	// 	  await fetch(`http://${process.env.GO_HOST}/user/${decoded.UserID}`,{
	// 		  mode:"no-cors"
	// 	  }).then((res)=>{
	// 		  return res.json()
	// 	  }).then((d)=>{

	// 		Userdata=d

	// 	  })

	//   }
	// return {
	// 	Userdata
	// }

	const MyCookie = cookies.get('Auth1') || '';
	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	let Userdata: UserDataType = {
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
		interface tokeninterface {
			Email: string;
			Name: string;
			UserID: string;
			Accounttype: string;
			exp: number;
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

	// MOST POPULAR
	let MostPopularList: GameDataType[] = [] as GameDataType[]
	await fetch(`http://${process.env.GO_HOST}/game/mostpopular`, {
		method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			MostPopularList = data;
			console.log(
				'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
				MostPopularList
			);
		});

	// NewReleaseList
	let NewReleaseList: GameDataType[]  = [] as GameDataType[] 
	await fetch(`http://${process.env.GO_HOST}/game/newrelease`, {
		method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			NewReleaseList = data;
			console.log(
				'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
				NewReleaseList
			);
		});

	// TopRatedList
	let TopRatedList: GameDataType[]  = [] as GameDataType[] 
	await fetch(`http://${process.env.GO_HOST}/game/toprated`, {
		method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			TopRatedList = data;
			console.log(
				'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
				TopRatedList
			);
		});
	// TopSoledList
	let TopSoledList: GameDataType[]  = [] as GameDataType[] 
	await fetch(`http://${process.env.GO_HOST}/game/topsold`, {
		method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			TopSoledList = data;
			console.log(
				'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
				TopSoledList
			);
		});

	// TrendingList
	let TrendingList: GameDataType[]  = [] as GameDataType[] 
	await fetch(`http://${process.env.GO_HOST}/game/trending`, {
		method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			TrendingList = data;
			console.log(
				'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
				TrendingList
			);
		});
	// CarouselList
	let CarouselList: GameDataType[]  = [] as GameDataType[] 
	await fetch(`http://${process.env.GO_HOST}/game/carousel`, {
		method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
			CarouselList = data;
			console.log(
				'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
				CarouselList
			);
		});
	return {
		Userdata,
		MostPopularList,
		NewReleaseList,
		TopRatedList,
		TopSoledList,
		TrendingList,
		CarouselList
	};
};