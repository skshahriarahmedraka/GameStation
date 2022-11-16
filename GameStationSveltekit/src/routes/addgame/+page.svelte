<script lang="ts">
	  import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';

	export let data: PageData;
	
	let { Userdata } = data;
	console.log("🚀 ~ file: +page.svelte ~ line 10 ~ Userdata", Userdata)
	UserProData.update((d) => (d = Userdata));
	
	// import GameProfile from '$lib/cards/gameProfile.svelte';
	// import Showcase from '$lib/Carousels/showcase.txt';
	// // import Showcase from "$lib/Carousels/showcase.svelte";
	// import Daisyui1 from '$lib/Carousels/daisyui1.svelte';
	// import RatingsStar from '$lib/logo/ratingsStar.txt';
	import Rating3 from '$lib/logo/rating4.svelte';
	// import Textarea from '$lib';

	// const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	// const fifa = new URL('../images/fifa.png', import.meta.url).href;
	// const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	// const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
	import Youtube from '$lib/svgs/youtube.svelte';
	// import Game from '$lib/logo/game.svelte';
	import Peoples from '$lib/svgs/Peoples.svelte';
	import Fb from '$lib/foot/icons/fb.svelte';
	// const syn1 = new URL('$lib/images/assassinscreedsyndicate/1.jpg', import.meta.url).href;
	// const syn2 = new URL('$lib/images/assassinscreedsyndicate/2.jpg', import.meta.url).href;
	// const syn3 = new URL('$lib/images/assassinscreedsyndicate/3.jpg', import.meta.url).href;
	// const syn4 = new URL('$lib/images/assassinscreedsyndicate/4.jpg', import.meta.url).href;
	// const syn5 = new URL('$lib/images/assassinscreedsyndicate/5.jpg', import.meta.url).href;
	// const syn6 = new URL('../../lib/images/assassinscreedsyndicate/6.jpg', import.meta.url).href;
	const syn7 = new URL('$lib/images/assassinscreedsyndicate/7.jpg', import.meta.url).href;
	import Discord from '$lib/foot/icons/discord.svelte';
	// import { goto } from '$app/navigation';

	// your script goes here

	import Twitter from '$lib/foot/icons/twitter.svelte';
	import Textarea from '$lib/otherComponents/textarea.svelte';
	import InputGameProfile from '$lib/cards/inputGameProfile.svelte';
	import InputCarousels from '$lib/Carousels/inputCarousels.svelte';
	// import Game from '$lib/logo/game.svelte';

	// const images = [
	// 	{
	// 		id: 0,
	// 		name: 'Cosmic timetraveler',
	// 		imgurl: anthem,
	// 		attribution: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com'
	// 	},
	// 	{
	// 		id: 1,
	// 		name: 'Cristina Gottardi',
	// 		imgurl: fifa,
	// 		attribution: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com'
	// 	},
	// 	{
	// 		id: 2,
	// 		name: 'Johannes Plenio',
	// 		imgurl: forbiddenwest,
	// 		attribution: 'johannes-plenio-RwHv7LgeC7s-unsplash.com'
	// 	},
	// 	{
	// 		id: 3,
	// 		name: 'Jonatan Pie',
	// 		imgurl: war,
	// 		attribution: 'jonatan-pie-3l3RwQdHRHg-unsplash.com'
	// 	}
	// ];

	// 	let Gamedata = {
	// 		GameID: '1234',
	// 		Name: "Assassin's Creed Syndicate",
	// 		Moto: 'London, 1868. In the heart of the Industrial Revolution, play as Jacob Frye - a brash and charismatic Assassin.',

	// 		LogoImage:
	// 			'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACV_UCS12483_EGST_Logo_Color_1000x407-1000x407-48fad9206ee895c599365e2b7c5bebcab099bd26.png?h=270&resize=1&w=480',
	// 		BigPosterImage:
	// 			'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACS-STD-2560x1440-635b7b6c86f18730071426375e7c4fe0bd831ddd.jpg',
	// 		SmallPosterImage:
	// 			'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640',

	// 		OtherImages: [syn1, syn2, syn3, syn4, syn5, syn7],
	// 		Genres: ['Action', 'RPG', 'Adventure'],
	// 		Feature: ['Single Player', 'Achivement', 'Control', 'Support'],
	// 		Description: `As Jacob Frye, a young and reckless Assassin, use your skills to help those trampled by the march of progress.
	// Travel the city at the height of the Industrial Revolution and meet iconic historical figures. From Westminster to Whitechapel, you will come across Darwin, Dickens, Queen Victoria… and many more.
	// As a gang leader, strengthen your stronghold and rally rival gang members to your cause, in order to take back the capital from the Templars’ hold.`,
	// 		FollowUs: {
	// 			facebook: 'https://www.facebook.com/epicgames',
	// 			discord: 'https://discord.com/app',
	// 			youtube: 'https://www.youtube.com/channel/UC5Qk8mWBwtMyEj7iQQYRk1A',
	// 			twitter: 'https://twitter.com/epicgames',
	// 			site: 'https://store.epicgames.com/en-US/p/assassins-creed-syndicate'
	// 		},
	// 		Rating: 9,
	// 		RatingGivenBy: {
	// 			'PC Gamer':
	// 				'The Assassin’s Creed series has been running for eight years, and that kind of longevity doesn’t happen without taking a few risks. To combat fatigue, each entry attempts to punctuate the familiar with new elements. Not all of these experiments pay off, ',
	// 			IGN: 'Syndicate smartly negotiates this internal conflict by dramatising it in the form of its twin playable characters, Evie and Jacob Frye. The former is a devout Assassin, intent on stopping the Templars by tracking down the remaining pieces of Eden. Her brother Jacob, however, is a pragmatist – a social reformer who rails against the Assassin’s burdensome legacy',
	// 			'Game Informer':
	// 				"Syndicate moves Assassin's Creed forward with a solid new adventure, a beautiful London playground and a renewed sense of fun"
	// 		},
	// 		Minspec: {
	// 			OS: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)',
	// 			CPU: 'Intel Core i5 2400s @ 2.5 GHz / AMD FX 6350 @ 3.9 GHz',
	// 			GPU: 'NVIDIA GeForce GTX 660 / AMD Radeon R9 270 (2GB VRAM with Shader Model 5.0)',
	// 			Memory: '6 GB RAM',
	// 			Storage: '100 GB',
	// 			'Sound Card': 'DirectX Compatible Sound Card with latest drivers'
	// 		},
	// 		Recomendedspec: {
	// 			OS: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)',
	// 			CPU: 'Intel Core i7 3770 @ 3.5 GHz / AMD FX 8350 @ 4.0 GHz',
	// 			GPU: 'NVIDIA GeForce GTX 760 (4GB) or the newer GTX 970 (4GB) / AMD Radeon R9 280X (3GB of VRAM) or better',
	// 			Memory: '8 GB RAM',
	// 			Storage: '200 GB',
	// 			'Sound Card': 'DirectX Compatible Sound Card with latest drivers'
	// 		},
	// 		Price: 34.55,
	// 		Discount: 25,
	// 		Developer: 'Ubisoft',
	// 		Publisher: 'Ubisoft',
	// 		Released: '02/20/20',
	// 		Platform: ['windows', 'linux', 'mac'],
	// 		Players: 23478238
	// 	};

	function RoundNumOfPeople(x: number) {
		if (x === null) {
			return 0;
		} else if (x >= 1000000000) {
			x /= 1000000000;
			return String(x.toFixed(2)) + 'B';
		} else if (x >= 1000000) {
			x /= 1000000;
			return String(x.toFixed(2)) + 'M';
		} else if (x >= 1000) {
			x /= 1000;
			return String(x.toFixed(2)) + 'K';
		} else {
			return String(x);
		}
	}
	let FormData3 = {
		GameID: '',
		Name: "Assassin's Creed® Mirage",
		Moto: 'Experience the story of Basim, a cunning street thief seeking answers and justice as he navigates the bustling streets of ninth–century Baghdad.',
		LogoImage: '',
		BigPosterImage: '',
		SmallPosterImage: '',
		OtherImages: [],
		Genres: ['Action', 'RPG', 'Open World', 'Adventure'],
		Feature: ['Achievements', 'Controller Support', 'Single Player'],
		Description:
			'Cyberpunk 2077 is an open-world, action-adventure RPG set in the megalopolis of Night City, where you play as a cyberpunk mercenary wrapped up in a do-or-die fight for survival. Improved and featuring all-new free additional content, customize your character and playstyle as you take on jobs, build a reputation, and unlock upgrades. The relationships you forge and the choices you make will shape the story and the world around you. Legends are made here. What will yours be?CREATE YOUR OWN CYBERPUNKBecome a cyberpunk, an urban mercenary equipped with cybernetic enhancements and build your legend on the streets of Night City.EXPLORE THE CITY OF THE FUTURENight City is packed to the brim with things to do, places to see, and people to meet. And it’s up to you where to go, when to go, and how to get there.BUILD YOUR LEGENDGo on daring adventures and build relationships with unforgettable characters whose fates are shaped by the choices you make.EQUIPPED WITH IMPROVEMENTSExperience Cyberpunk 2077 with a host of changes and improvements to gameplay and economy, the city, map usage, and more.INCLUDES FREE ADDITIONAL CONTENT Get your hands on a haul of free items including new guns and melee weapons, as well as extra customization options and more.',
		FollowUs: {
			Facebook: 'http://facebook.com',
			Discord: 'http://discord.com',
			Youtube: 'http://youtube.com',
			Twitter: 'http://twitter.com',
			Site: 'http://epicgames.com'
		},
		Rating: 7,
		RatingGivenBy: {
			'PC Gamer':
				'Experience Cyberpunk 2077 with a host of changes and improvements to gameplay and economy, the city, map usage, and more.INCLUDES FREE ADDITIONAL CONTENTGet your hands on a haul of free items including new guns and melee weapons, as well as extra customization options and more.',
			IGN: 'Cyberpunk 2077 with a host of changes and improvements to gameplay and economy, the city, map usage, and more.INCLUDES FREE ADDITIONAL CONTENTGet your hands on a haul of free items including new guns and melee weapons, as well as extra customization options and more.',
			'Game Informer':
				' with a host of changes and improvements to gameplay and economy, the city, map usage, and more.INCLUDES FREE ADDITIONAL CONTENTGet your hands on a haul of free items including new guns and melee weapons, as well as extra customization options and more.'
		},
		Minspec: [
			{
				Name: 'OS',
				Value: 'Windows 10 /11'
			},
			{
				Name: 'Storage',
				Value: '100GB'
			},
			{
				Name: 'Memory',
				Value: '6GB'
			},
			{
				Name: 'CPU',
				Value: 'Intel xeon 11Gen'
			},
			{
				Name: 'GPU',
				Value: 'nvidia 3090ti'
			}
		],
		Recomendedspec: [
			{
				Name: 'OS',
				Value: 'Ubuntu / Red Hat'
			},
			{
				Name: 'Storage',
				Value: '200GB'
			},
			{
				Name: 'Memory',
				Value: '10GB'
			},
			{
				Name: 'GPU',
				Value: 'Nvidia 4090'
			},
			{
				Name: 'CPU',
				Value: 'ryzen 7 7gen'
			}
		],
		Price: 58.0,
		Discount: 30,
		Developer: 'Epic Games',
		Publisher: 'Sony Entertainment',
		Released: '23/8/22',
		Platform: ['windows', 'linux', 'mac'],
		Players: 11869785
	};
	let FormData: {
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
	} = {
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
	};

	let FormDataErr: {
		GameID: boolean;
		Name: boolean;
		Moto: boolean;
		LogoImage: boolean;
		BigPosterImage: boolean;
		SmallPosterImage: boolean;
		OtherImages: boolean;
		Genres: boolean;
		Feature: boolean;
		Description: boolean;
		Rating: boolean;
		RatingGivenBy: {
			'PC Gamer': boolean;
			IGN: boolean;
			'Game Informer': boolean;
		};
		Minspec: {
			Name: string;
			Value: boolean;
		}[];
		Recomendedspec: {
			Name: string;
			Value: boolean;
		}[];
		Discount: boolean;
		Price: boolean;
		Developer: boolean;
		Publisher: boolean;
		Released: boolean;
		Platform: boolean;
		Players: boolean;

		FollowUs: {
			Facebook: boolean;
			Discord: boolean;
			Youtube: boolean;
			Twitter: boolean;
			Site: boolean;
		};
	} = {
		GameID: true,
		Name: true,
		Moto: true,
		LogoImage: true,
		BigPosterImage: true,
		SmallPosterImage: true,
		OtherImages: true,
		Genres: true,
		Feature: true,
		Description: true,
		FollowUs: {
			Facebook: true,
			Discord: true,
			Youtube: true,
			Twitter: true,
			Site: true
		},
		Rating: true,
		RatingGivenBy: {
			'PC Gamer': true,
			IGN: true,
			'Game Informer': true
		},
		Minspec: [
			{
				Name: 'OS',
				Value: true
			},
			{
				Name: 'Storage',
				Value: true
			},
			{
				Name: 'Memory',
				Value: true
			},
			{
				Name: 'CPU',
				Value: true
			},
			{
				Name: 'GPU',
				Value: true
			}
		],
		Recomendedspec: [
			{
				Name: 'OS',
				Value: true
			},
			{
				Name: 'Storage',
				Value: true
			},
			{
				Name: 'Memory',
				Value: true
			},
			{
				Name: 'GPU',
				Value: true
			},
			{
				Name: 'CPU',
				Value: true
			}
		],
		Price: true,
		Discount: true,
		Developer: true,
		Publisher: true,
		Released: true,
		Platform: true,
		Players: true
	};

	let Gamedata = {
		GameID: '',
		Name: "Assassin's Creed Syndicate",
		Moto: 'London, 1868. In the heart of the Industrial Revolution, play as Jacob Frye - a brash and charismatic Assassin.',

		LogoImage:
			'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACV_UCS12483_EGST_Logo_Color_1000x407-1000x407-48fad9206ee895c599365e2b7c5bebcab099bd26.png?h=270&resize=1&w=480',
		BigPosterImage:
			'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACS-STD-2560x1440-635b7b6c86f18730071426375e7c4fe0bd831ddd.jpg',
		SmallPosterImage:
			'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640',

		OtherImages: ['', ''],
		Genres: ['Action', 'RPG', 'Adventure'],
		Feature: ['Single Player', 'Achivement', 'Control', 'Support'],
		Description: `As Jacob Frye, a young and reckless Assassin, use your skills to help those trampled by the march of progress.
		Travel the city at the height of the Industrial Revolution and meet iconic historical figures. From Westminster to Whitechapel, you will come across Darwin, Dickens, Queen Victoria… and many more.	
		As a gang leader, strengthen your stronghold and rally rival gang members to your cause, in order to take back the capital from the Templars’ hold.`,
		FollowUs: {
			Facebook: 'https://www.facebook.com/epicgames',
			Discord: 'https://discord.com/app',
			Youtube: 'https://www.youtube.com/channel/UC5Qk8mWBwtMyEj7iQQYRk1A',
			Twitter: 'https://twitter.com/epicgames',
			Site: 'https://store.epicgames.com/en-US/p/assassins-creed-syndicate'
		},
		Rating: 9,
		RatingGivenBy: {
			'PC Gamer':
				'The Assassin’s Creed series has been running for eight years, and that kind of longevity doesn’t happen without taking a few risks. To combat fatigue, each entry attempts to punctuate the familiar with new elements. Not all of these experiments pay off, ',
			IGN: 'Syndicate smartly negotiates this internal conflict by dramatising it in the form of its twin playable characters, Evie and Jacob Frye. The former is a devout Assassin, intent on stopping the Templars by tracking down the remaining pieces of Eden. Her brother Jacob, however, is a pragmatist – a social reformer who rails against the Assassin’s burdensome legacy',
			'Game Informer':
				"Syndicate moves Assassin's Creed forward with a solid new adventure, a beautiful London playground and a renewed sense of fun"
		},
		Minspec: [
			{ Name: 'OS', Value: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)' },
			{
				Name: 'Storage',
				Value: '100 GB'
			},
			{
				Name: 'Memory',
				Value: '6 GB RAM'
			},
			{
				Name: 'GPU',
				Value: 'NVIDIA GeForce GTX 660 / AMD Radeon R9 270 (2GB VRAM with Shader Model 5.0)'
			},
			{
				Name: 'CPU',
				Value: 'Intel Core i5 2400s @ 2.5 GHz / AMD FX 6350 @ 3.9 GHz'
			},
			{
				Name: 'Sound Card',
				Value: 'DirectX Compatible Sound Card with latest drivers'
			}
		],
		Recomendedspec: [
			{ Name: 'OS', Value: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)' },
			{
				Name: 'Storage',
				Value: '100 GB'
			},
			{
				Name: 'Memory',
				Value: '6 GB RAM'
			},
			{
				Name: 'GPU',
				Value: 'NVIDIA GeForce GTX 660 / AMD Radeon R9 270 (2GB VRAM with Shader Model 5.0)'
			},
			{
				Name: 'CPU',
				Value: 'Intel Core i5 2400s @ 2.5 GHz / AMD FX 6350 @ 3.9 GHz'
			},
			{
				Name: 'Sound Card',
				Value: 'DirectX Compatible Sound Card with latest drivers'
			}
		],
		Price: 34.55,
		Discount: 25,
		Developer: 'Ubisoft',
		Publisher: 'Ubisoft',
		Released: '02/20/20',
		Platform: ['windows', 'linux', 'mac'],
		Players: 23478238
	};
	let thisPlayer: any;
	// $: console.log(data.Minspec[0].name);
	// $: console.log(data.Minspec[0].value);
	let GenreObj: { Input: string; Focus: boolean; List: string[]; Suggested: string[] } = {
		Input: '' as string,
		Focus: false as boolean,
		List: [] as string[],
		Suggested: [] as string[]
	};
	function onKeyPressGenre(e: { keyCode: any }) {
		switch (e.keyCode) {
			case 32:
				onKeyGenre();
				break;
			case 13:
				onKeyGenre();
				break;
		}
	}
	function onKeyGenre() {
		if (GenreObj.Input.trim() != '') {
			GenreObj.List = [...GenreObj.List, GenreObj.Input.trim()];
			FormData.Genres = GenreObj.List;
			GenreObj.Input = '';
		}
	}

	let FeaturesObj: { Input: string; Focus: boolean; List: string[]; Suggested: string[] } = {
		Input: '' as string,
		Focus: false as boolean,
		List: [] as string[],
		Suggested: [] as string[]
	};
	function onKeyPressFeatures(e: { keyCode: any }) {
		switch (e.keyCode) {
			case 32:
				onKeyFeatures();
				break;
			case 13:
				onKeyFeatures();
				break;
		}
	}
	function onKeyFeatures() {
		if (FeaturesObj.Input.trim() != '') {
			FeaturesObj.List = [...FeaturesObj.List, FeaturesObj.Input.trim()];
			FormData.Feature = FeaturesObj.List;
			FeaturesObj.Input = '';
		}
	}
	// $: console.log("🚀 ~ file: +page.svelte ~ line 137 ~ FormData", FormData)

	// SUBMIT A GAME
	let responseClone;
	import { error, json } from '@sveltejs/kit';
	// import axios from 'axios';
	import Globe from '$lib/svgs/globe.svelte';
	import World from '$lib/svgs/world.svelte';
	import { goto } from '$app/navigation';

	function Checkerror() {
		if (FormData.Name.trim() === '') {
			FormDataErr.Name = true;
		}
	}
	async function PostGameData() {
		// Checkerror()
		// console.log('🚀 ~ file: +page.svelte ~ line 306 ~ PostGameData', data);
		fetch('/api/addgame', {
			// credentials: 'same-origin',
			method: 'POST',
			// mode: 'no-cors',
			// headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(FormData)
		})
			.then((response) => response.json())
			.then((data) => {
				console.log('🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ REsponse', data);
				console.log(
					'🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ response data.GameID',
					data.GameID
				);
				if (data.GameID) {
					console.log('🚀 ~ file: +page.svelte ~ line 627 ~ .then ~ data.GameID', data.GameID);
					// goto(`/${data.GameID}`);
				} else {
					throw error(404, 'Not Found');
				}
			});

		// ////////////////////////////////
		// fetch('https://jsonplaceholder.typicode.com/todos/1')
		// 	.then((response) => response.json())
		// 	.then((json) => console.log(json));
	}

	///////////////////////////////////////////////////////////////////////////////////
	// 	async function PostGameData2() {
	// 		axios.defaults.baseURL = '127.0.0.1:8001';
	//   axios.defaults.headers.post['Content-Type'] ='application/json;charset=utf-8';
	//   axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';
	// 		axios
	// 			.post(
	// 				'http://localhost:8001/admin/addproduct',
	// 				{
	// 					firstName: 'Fred',
	// 					lastName: 'Flintstone'
	// 				},
	// 				{
	// 					method: 'POST',
	// 					// mode: 'no-cors',
	// 					withCredentials: true,
	// 					headers: {
	// 						'Access-Control-Allow-Origin': "*",
	// 						'Content-Type': 'application/json',

	// 					},
	// 					// proxy: {
	// 					// 	host: '127.0.0.1',
	// 					// 	port: 8001
	// 					// }
	// 					// credentials: 'same-origin'
	// 				}
	// 			)
	// 			.then(function (response) {
	// 				console.log(response);
	// 			})
	// 			.catch(function (error) {
	// 				console.log(error);
	// 			});
	// 	}
	let ThisGenreInput: any;
	function RemoveGenreFromList(s: number) {
		GenreObj.List.splice(s, 1);
		ThisGenreInput.focus();
	}
	let ThisFeatureInput: any;
	function RemoveFeatureFromList(s: number) {
		FeaturesObj.List.splice(s, 1);
		ThisFeatureInput.focus();
	}
</script>

<!-- markup (zero or more items) goes here -->

<div class=" flex h-fit w-full flex-row justify-center  gap-6  text-slate-100">
	<div class="flex flex-col gap-3 ">
		<!-- Game Name -->
		<input
			bind:value={FormData.Name}
			type="text"
			placeholder="Game Name"
			class=" my-2 h-16 w-[70%] self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2  text-5xl font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800  "
		/>
		<!-- <p class=" text-5xl text-slate-100">{GameFormData.Name}</p> -->
		<div class=" flex flex-row gap-4 ">
			<Rating3 bind:value={FormData.Rating} />
			<input
				bind:value={FormData.Rating}
				type="number"
				placeholder="Ratings"
				class=" h-8 w-24 self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<Peoples class="h-6 w-6 stroke-blue-400" />
			<p class=" -ml-2 text-lg">{RoundNumOfPeople(FormData.Players)}</p>
			<input
				type="number"
				placeholder="Ratings"
				bind:value={FormData.Players}
				class=" h-8 w-24 self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
		</div>
		<div
			class="flex justify-center gap-3  sm:flex-col-reverse md:flex-col-reverse lg:flex-row  xl:flex-row xs:flex-col-reverse"
		>
			<div class="flex max-w-[1080px] flex-col gap-4 ">
				<!-- image show case -->
				<!-- <Daisyui1 OtherImages={GameFormData.OtherImages} /> -->

				<InputCarousels
					bind:OtherImages={FormData.OtherImages}
					bind:BigPosterImage={FormData.BigPosterImage}
				/>

				<p class=" text-2xl  ">Game Moto :</p>
				<textarea
					type="text"
					placeholder="Game Moto"
					bind:value={FormData.Moto}
					class=" scrol3 h-20 w-full self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 font-Poppins text-lg  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
				<!-- Genres and features -->
				<div class="mt-12 flex w-full flex-row  font-Poppins">
					<!-- GENRE -->
					<div class="flex w-1/2 flex-col border-l-[1px] border-slate-400 border-opacity-30 p-3">
						<p class=" font-normal text-gray-500">Genres</p>
						<div class=" flex  flex-col gap-2 text-base">
							<!--  -->

							<div
								class="mx-4 mt-3 h-fit w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium text-[#98999e]  outline-none   {GenreObj.Focus
									? 'border-sky-500'
									: ''} active:border-gray-800 "
							>
								<ul class="flex flex-row flex-wrap gap-2">
									{#each GenreObj.List as t, i}
										<li
											class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
										>
											{t}
											<svg
												on:click={() => {
													RemoveGenreFromList(i);
												}}
												on:keypress={() => {}}
												class="inline-flex h-6 w-6 hover:stroke-red-500"
												fill="none"
												stroke="currentColor"
												viewBox="0 0 24 24"
												xmlns="http://www.w3.org/2000/svg"
												><path
													stroke-linecap="round"
													stroke-linejoin="round"
													stroke-width="2"
													d="M6 18L18 6M6 6l12 12"
												/></svg
											>
										</li>
									{/each}
									<input
										bind:this={ThisGenreInput}
										on:keypress={onKeyPressGenre}
										maxlength="20"
										bind:value={GenreObj.Input}
										type="text"
										on:focus={() => {
											GenreObj.Focus = true;
										}}
										on:blur={() => {
											GenreObj.Focus = false;
										}}
										class=" inline w-fit grow rounded-lg border-0  bg-inherit shadow-none outline-0 focus:border-0 focus:outline-none"
									/>
								</ul>
							</div>
							<!--  -->
							<!-- {#each GameFormData.Genres as G}
								<p class="">{G} ,</p>
							{/each} -->
						</div>
					</div>
					<!-- features -->
					<div class="flex w-1/2 flex-col border-l-[1px] border-slate-400 border-opacity-30 p-3">
						<p class=" font-normal text-gray-500">features</p>
						<div class=" flex  flex-row gap-2 text-base">
							<div
								class="mx-4 mt-3 h-fit w-[95%] rounded-lg border-2 border-[#24262b] bg-[#303338] p-2  text-lg font-medium text-[#98999e]  outline-none   {FeaturesObj.Focus
									? 'border-sky-500'
									: ''} active:border-gray-800 "
							>
								<ul class="flex flex-row flex-wrap gap-2 ">
									{#each FeaturesObj.List as t, i}
										<li
											class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
										>
											{t}
											<svg
												on:click={() => {
													RemoveFeatureFromList(i);
												}}
												on:keypress={() => {}}
												class="inline-flex h-6 w-6 hover:stroke-red-500"
												fill="none"
												stroke="currentColor"
												viewBox="0 0 24 24"
												xmlns="http://www.w3.org/2000/svg"
												><path
													stroke-linecap="round"
													stroke-linejoin="round"
													stroke-width="2"
													d="M6 18L18 6M6 6l12 12"
												/></svg
											>
										</li>
									{/each}
									<input
										bind:this={ThisFeatureInput}
										on:keypress={onKeyPressFeatures}
										maxlength="20"
										bind:value={FeaturesObj.Input}
										type="text"
										on:focus={() => {
											FeaturesObj.Focus = true;
										}}
										on:blur={() => {
											FeaturesObj.Focus = false;
										}}
										class=" inline w-fit grow rounded-lg border-0  bg-inherit shadow-none outline-0 focus:border-0 focus:outline-none"
									/>
								</ul>
							</div>

							<!-- {#each GameFormData.Feature as G}
								<p class="">{G} ,</p>
							{/each} -->
						</div>
					</div>
				</div>
				<!-- description -->
				<div class=" flex flex-col gap-2 ">
					<p class=" text-xl font-semibold">Description :</p>
					<!-- <p class="">
						{GameFormData.Description}
					</p> -->
					<Textarea bind:value={FormData.Description} minRows={4} maxRows={40} />
				</div>
				<div class="flex flex-col gap-2 ">
					<p class="text-xl font-semibold">Follow Us :</p>
					<div class=" flex flex-col justify-center gap-5 rounded-lg bg-[#262626] py-3 ">
						<div class="flex flex-row justify-center gap-2">
							<button class="">
								<!-- <Fb  /> -->
								<World class="h-9 w-9 fill-slate-200" />
							</button>
							<input
								type="text"
								placeholder="WebSite"
								bind:value={FormData.FollowUs.Site}
								class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
							/>
						</div>
						<div class="flex flex-row justify-center gap-2">
							<button class="">
								<Fb class="h-7 w-7 fill-slate-200" />
							</button>
							<input
								type="text"
								placeholder="Facebook"
								bind:value={FormData.FollowUs.Facebook}
								class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
							/>
						</div>
						<div class="flex flex-row justify-center gap-2">
							<button class="">
								<Youtube class="h-8 w-8 fill-slate-200" />
							</button>
							<input
								type="text"
								placeholder="youtube"
								bind:value={FormData.FollowUs.Youtube}
								class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
							/>
						</div>
						<div class="flex flex-row justify-center gap-2">
							<button class="">
								<Discord class="h-8 w-8 fill-slate-200" />
							</button>
							<input
								type="text"
								placeholder="Discord"
								bind:value={FormData.FollowUs.Discord}
								class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
							/>
						</div>
						<div class="flex flex-row justify-center gap-2">
							<button class="">
								<Twitter class="h-8 w-8 fill-slate-200" />
							</button>
							<input
								type="text"
								placeholder="twitter"
								bind:value={FormData.FollowUs.Twitter}
								class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
							/>
						</div>
					</div>
				</div>
				<!-- Org Review -->
				<p class="text-xl font-semibold">Organizations Review :</p>
				<div class=" flex flex-row justify-around gap-2">
					<!-- content here -->
					<div class="h-[300px] w-[250px] rounded-xl bg-[#262626] p-4  ">
						<p class="border-b-[1px] border-gray-500  border-opacity-10 pb-2 text-lg ">
							Game Informer
						</p>
						<!-- <p class=" pt-2 text-sm text-slate-400">{Value}</p> -->
						<Textarea
							class="pt-2 text-sm text-slate-400 "
							bind:value={FormData.RatingGivenBy['Game Informer']}
							minRows={4}
							maxRows={40}
						/>
					</div>

					<div class="h-[300px] w-[250px] rounded-xl bg-[#262626] p-4  ">
						<p class="border-b-[1px] border-gray-500  border-opacity-10 pb-2 text-lg ">IGN</p>
						<!-- <p class=" pt-2 text-sm text-slate-400">{Value}</p> -->
						<Textarea
							class="pt-2 text-sm text-slate-400"
							bind:value={FormData.RatingGivenBy.IGN}
							minRows={4}
							maxRows={40}
						/>
					</div>
					<div class="h-[300px] w-[250px] rounded-xl bg-[#262626] p-4  ">
						<p class="border-b-[1px] border-gray-500  border-opacity-10 pb-2 text-lg ">PC Gamer</p>
						<!-- <p class=" pt-2 text-sm text-slate-400">{Value}</p> -->
						<Textarea
							class="pt-2 text-sm text-slate-400"
							bind:value={FormData.RatingGivenBy['PC Gamer']}
							minRows={4}
							maxRows={40}
						/>
					</div>
				</div>
				<!-- spec -->
				<div class="flex flex-col rounded-md bg-[#262626] p-4">
					<div class="flex h-16 w-full flex-row gap-4 ">
						<p class=" text-lg underline ">Windows</p>
						<p class="text-lg underline">Linux</p>
						<p class="text-lg underline">Mac</p>
					</div>
					<div class="flex h-fit w-full flex-row justify-start rounded-md bg-opacity-50 ">
						<div class=" flex  w-1/2 flex-col gap-4 ">
							<p class=" text-lg font-semibold">Minimum Requirement :</p>
							<div class=" flex flex-col gap-3 ">
								{#each FormData.Minspec as option}
									<div class="flex flex-col gap-1">
										<p class=" text-gray-200">{option.Name} :</p>
										<!-- <p class="">{Value}</p> -->
										<input
											type="text"
											placeholder={option.Name}
											bind:value={option.Value}
											class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
										/>
									</div>
								{/each}
							</div>
						</div>
						<div class=" flex  w-1/2 flex-col gap-4 ">
							<p class=" text-lg font-semibold">Recomended Requirement :</p>
							<div class=" flex flex-col gap-3 ">
								{#each FormData.Recomendedspec as option}
									<div class="flex flex-col gap-1">
										<p class=" text-gray-200">{option.Name} :</p>
										<!-- <p class="">{Value}</p> -->
										<input
											type="text"
											placeholder={option.Name}
											bind:value={option.Value}
											class=" h-12 w-[80%] self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
										/>
									</div>
								{/each}
								<!-- {#each Object.entries(GameFormData.Recomendedspec) as [title, Value]}
									<div class="flex flex-col gap-1">
										<p class=" text-gray-500">{title}</p>
										<p class="">{Value}</p>
									</div>
								{/each} -->
							</div>
						</div>
					</div>
				</div>
			</div>
			<InputGameProfile
				bind:LogoImage={FormData.LogoImage}
				bind:SmallPosterImage={FormData.SmallPosterImage}
				bind:Released={FormData.Released}
				bind:Price={FormData.Price}
				bind:Platform={FormData.Platform}
				bind:Publisher={FormData.Publisher}
				bind:Develper={FormData.Developer}
				bind:Discount={FormData.Discount}
			/>
		</div>
		<div class=" flex w-full justify-center">
			<button
				on:click={() => {
					PostGameData();
				}}
				class=" ml-4 h-[50px] w-[320px] rounded-lg bg-sky-500 font-Poppins text-lg font-semibold text-white hover:bg-sky-600"
			>
				Add Game
			</button>
		</div>
	</div>
	<!-- <Showcase /> -->
</div>

<style>
	/* your styles go here */
	/* Chrome, Safari, Edge, Opera */
	input::-webkit-outer-spin-button,
	input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}

	/* Firefox */
	input[type='number'] {
		-moz-appearance: textfield;
	}
</style>
