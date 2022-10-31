<script lang="ts">
	import GameProfile from '$lib/cards/gameProfile.svelte';
	import Showcase from '$lib/Carousels/showcase.txt';
	// import Showcase from "$lib/Carousels/showcase.svelte";
	import Daisyui1 from '$lib/Carousels/daisyui1.svelte';
	import RatingsStar from '$lib/logo/ratingsStar.txt';
	import Rating3 from '$lib/logo/rating4.svelte';

	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
	import Youtube from '$lib/svgs/youtube.svelte';
	import Game from '$lib/logo/game.svelte';
	import Peoples from '$lib/svgs/Peoples.svelte';
	import Fb from '$lib/foot/icons/fb.svelte';
	const syn1 = new URL('../../lib/images/assassinscreedsyndicate/1.jpg', import.meta.url).href;
	const syn2 = new URL('../../lib/images/assassinscreedsyndicate/2.jpg', import.meta.url).href;
	const syn3 = new URL('../../lib/images/assassinscreedsyndicate/3.jpg', import.meta.url).href;
	const syn4 = new URL('../../lib/images/assassinscreedsyndicate/4.jpg', import.meta.url).href;
	const syn5 = new URL('../../lib/images/assassinscreedsyndicate/5.jpg', import.meta.url).href;
	// const syn6 = new URL('../../lib/images/assassinscreedsyndicate/6.jpg', import.meta.url).href;
	const syn7 = new URL('../../lib/images/assassinscreedsyndicate/7.jpg', import.meta.url).href;
	import Discord from '../../lib/foot/icons/discord.svelte';
	import { goto } from '$app/navigation';

	// your script goes here

	import Twitter from '$lib/foot/icons/twitter.svelte';
	import { each } from 'svelte/internal';
	import EmidGameCom from '$lib/cards/EmidGameCom.svelte';
	// import Game from '$lib/logo/game.svelte';

	const images = [
		{
			id: 0,
			name: 'Cosmic timetraveler',
			imgurl: anthem,
			attribution: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com'
		},
		{
			id: 1,
			name: 'Cristina Gottardi',
			imgurl: fifa,
			attribution: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com'
		},
		{
			id: 2,
			name: 'Johannes Plenio',
			imgurl: forbiddenwest,
			attribution: 'johannes-plenio-RwHv7LgeC7s-unsplash.com'
		},
		{
			id: 3,
			name: 'Jonatan Pie',
			imgurl: war,
			attribution: 'jonatan-pie-3l3RwQdHRHg-unsplash.com'
		}
	];
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

	let Gamedata : {
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
	}  = {
		GameID: '1234',
		Name: "Assassin's Creed Syndicate",
		Moto: 'London, 1868. In the heart of the Industrial Revolution, play as Jacob Frye - a brash and charismatic Assassin.',

		LogoImage:
			'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACV_UCS12483_EGST_Logo_Color_1000x407-1000x407-48fad9206ee895c599365e2b7c5bebcab099bd26.png?h=270&resize=1&w=480',
		BigPosterImage:
			'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACS-STD-2560x1440-635b7b6c86f18730071426375e7c4fe0bd831ddd.jpg',
		SmallPosterImage:
			'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640',

		OtherImages: [syn1, syn2, syn3, syn4, syn5, syn7],
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
		Minspec:
		[
			{
				Name: 'OS',
				Value: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)'
			},
			{
				Name: 'Storage',
				Value: '100 GB'
			},
			{
				Name: 'Memory',
				Value: '6 GB RAM'
			},
			{
				Name: 'CPU',
				Value: 'Intel Core i5 2400s @ 2.5 GHz / AMD FX 6350 @ 3.9 GHz'
			},
			{
				Name: 'GPU',
				Value: 'NVIDIA GeForce GTX 660 / AMD Radeon R9 270 (2GB VRAM with Shader Model 5.0)'
			}
		],
		Recomendedspec: [
			{
				Name: 'OS',
				Value: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)'
			},
			{
				Name: 'Storage',
				Value: '200 GB'
			},
			{
				Name: 'Memory',
				Value: '8 GB RAM'
			},
			{
				Name: 'GPU',
				Value: 'NVIDIA GeForce GTX 760 (4GB) or the newer GTX 970 (4GB) / AMD Radeon R9 280X (3GB of VRAM) or better'
			},
			{
				Name: 'CPU',
				Value: 'Intel Core i7 3770 @ 3.5 GHz / AMD FX 8350 @ 4.0 GHz'
			}
		],
		
		
		Price: 34.55,
		Discount: 25,
		Developer: 'Ubisoft',
		Publisher: 'Ubisoft',
		Released: '02/20/20',
		Platform: ['windows', 'linux', 'mac'],
		Players: 23478238,
		Comment: []
	};



	let ListOfGameData=[Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata,Gamedata]

	function RoundNumOfPeople(x: number) {
		if (x >= 1000000000) {
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
	let filterCatagories = [
		{ Name: 'Event', Show: false, List: ['Halloween Sale', 'Holyday sale', 'Eid Sale'] },
		{
			Name: 'Price',
			Show: false,
			List: [
				'Free',
				'Under $5.00',
				'Under $10.00',
				'Under $20.00',
				'Under $30.00',
				'$14.99 and above',
				'Discounted'
			]
		},
		{
			Name: 'Genre',
			Show: false,
			List: [
				'Action',
				'Adventure',
				'Application',
				'Card Game',
				'Casual',
				'City Builder',
				'Comedy',
				'Dungeon Crawler',
				'Exploration',
				'Fighting',
				'First Person',
				'Horror',
				'Indie',
				'Music',
				'Narration',
				'Open World',
				'Party',
				'Platformer',
				'Puzzle',
				'Racing',
				'Retro',
				'Rogue',
				'RPG',
				'Shooter',
				'Simulation',
				'Space',
				'Sports',
				'Stealth',
				'Strategy',
				'Survival',
				'Tower Defense',
				'Trivia',
				'Turn-Based'
			]
		},
		{
			Name: 'Feature',
			Show: false,
			List: [
				'Achievements',
				'Challenging Combat',
				'Cloud Saves',
				'Co-op',
				'Competitive',
				'Controller Support',
				'Cross Platform',
				'MMO',
				'Multiplayer',
				'Single Player',
				'VR'
			]
		},
		{
			Name: 'Types',
			Show: false,
			List: ['Apps', 'Editor', 'Game', 'Game Add-On', 'Game Bundle', 'Game Demo', 'Game Edition']
		},
		{ Name: 'Platform', Show: false, List: ['mac', 'windows', 'linux'] }
	];
	let SelectedField: any = {
		Event: {},
		Price: {},
		Genre: {},
		Feature: {},
		Types: {},
		Platform: {}
	};

	function OnclickRestFiler(){
		SelectedField= {
		Event: {},
		Price: {},
		Genre: {},
		Feature: {},
		Types: {},
		Platform: {}
	}
	}
</script>

<!-- markup (zero or more items) goes here -->

<div
	class=" flex max-h-fit min-h-screen w-[1440px]  flex-row justify-center  gap-6  text-slate-100"
>
	<div class="h-fit w-[1178px] ">
		<p class="">Filtered Games</p>
		<div class=" flex flex-row  flex-wrap gap-4">

			{#each ListOfGameData as i}
			<EmidGameCom  Obj={i} />
			<!-- content here -->
			{/each}
		</div>
		<!--  -->
	</div>

	<div class=" flex h-screen w-[245px] flex-col gap-1">
		<button on:click={()=>{OnclickRestFiler()}} class="h-6  w-fit border-[1px] flex flex-row bg-gray-800 text-sm rounded-md self-end text-center pb-1 px-2">Reset Filter</button>
		<div class="flex h-fit w-full flex-row flex-wrap">
			{#each Object.entries(SelectedField) as [Name, obj]}
				{#each Object.entries(obj) as [fieldname, value]}
					{#if value}
						<!-- content here -->
						<div class="h-6  w-fit border-[1px] flex flex-row bg-gray-600 text-sm rounded-lg text-center pb-1 px-2">
							<p class="">{fieldname}</p>
							<!-- <svg class="w-4 h-4 ml-1 my-1 hover:cursor-pointer" on:click={()=>{ value=false}} on:keypress={()=>{}} fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg> -->
						</div>
					{/if}
				{/each}
			{/each}
		</div>
		<!--  -->
		{#each filterCatagories as filter}
			<div
				on:click={() => {
					filter.Show = !filter.Show;
				}}
				on:keypress={() => {}}
				class="relative flex h-10 w-full flex-row justify-around rounded-lg border-2 border-gray-300 border-opacity-50 bg-gray-800 py-1 text-center font-Poppins hover:cursor-pointer hover:bg-opacity-80"
			>
				<p class=" ">{filter.Name}</p>
				{#if filter.Show}
					<svg
						class="absolute right-2 h-3  w-3 self-center "
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M5 15l7-7 7 7"
						/></svg
					>
				{:else}
					<svg
						class="  absolute right-2 h-3 w-3 self-center"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M19 9l-7 7-7-7"
						/></svg
					>
				{/if}
			</div>
			{#if filter.Show}
				{#each filter.List as i}
					<div
						on:click={() => {
							SelectedField[filter.Name][i] = !SelectedField[filter.Name][i];
						}}
						on:keypress={() => {}}
						class=" relative h-8 w-full rounded-lg p-2 hover:cursor-pointer hover:bg-slate-700 hover:bg-opacity-25"
					>
						<p class=" text-sm">{i}</p>
						{#if SelectedField[filter.Name][i]}
							<svg
								class=" absolute right-2 top-2 h-4 w-4 self-center stroke-green-400 "
								fill="none"
								stroke="currentColor"
								viewBox="0 0 24 24"
								xmlns="http://www.w3.org/2000/svg"
								><path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
								/></svg
							>
						{/if}
					</div>
				{/each}
			{/if}
			<!-- content here -->
		{/each}
	</div>
	<!-- <Showcase /> -->
</div>

<style>
	/* your styles go here */
</style>
