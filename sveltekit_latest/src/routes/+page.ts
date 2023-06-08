// import { error } from '@sveltejs/kit';
import type { GameDataType } from '$lib/Store/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	// let MostPopularList: GameDataType[] = [] as GameDataType[];

	// console.log(
	// 	'🚀 +page.ts MostPopularList: ',
	// 	'http://${import.meta.env.VITE_GO_HOST}/api/game/mostpopular',
	// 	MostPopularList
	// );
	// await fetch(`http://${import.meta.env.VITE_GO_HOST}/game/mostpopular`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		MostPopularList = data;
	// 		console.log('🚀 ~ MostPopularList', MostPopularList);
	// 	});

	// console.log('🚀 +page.ts MostPopularList:', MostPopularList);

	

	// let NewReleaseList: GameDataType[] = [] as GameDataType[];

	// await fetch(`http://${process.env.GO_HOST}/game/newrelease`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		NewReleaseList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			NewReleaseList
	// 		);
	// 	});


	// let TopRatedList: GameDataType[] = [] as GameDataType[];
	// await fetch(`http://${process.env.GO_HOST}/game/toprated`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		TopRatedList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ TopRatedList',
	// 			TopRatedList
	// 		);
	// 	});
	// let TopSoledList: GameDataType[] = [] as GameDataType[];
	// await fetch(`http://${process.env.GO_HOST}/game/topsold`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		TopSoledList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ TopSoledList',
	// 			TopSoledList
	// 		);
	// 	});
	// let TrendingList: GameDataType[] = [] as GameDataType[];
	// await fetch(`http://${process.env.GO_HOST}/game/trending`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		TrendingList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ TrendingList',
	// 			TrendingList
	// 		);
	// 	});
	// let CarouselList: GameDataType[] = [] as GameDataType[];
	// await fetch(`http://${process.env.GO_HOST}/game/carousel`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		CarouselList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ CarouselList',
	// 			CarouselList
	// 		);
	// 	});

	
	// return {
	// 	// Userdata,
	// 	MostPopularList,
	// 	NewReleaseList,
	// 	TopRatedList,
	// 	TopSoledList,
	// 	TrendingList,
	// 	CarouselList
	// };
	return{
		
	}
	//   throw error(404, 'Not found');
};
