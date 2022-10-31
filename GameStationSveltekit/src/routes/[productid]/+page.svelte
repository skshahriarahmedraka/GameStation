<script lang="ts">
	export let data;
	console.log('🚀 ~ file: +page.svelte ~ line 3 ~ data', data.mydata);

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
	} = data.mydata;
	import GameProfile from '$lib/cards/gameProfile.svelte';
	import {UserProData} from "$lib/Store/store"
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

	let Gamedata3 = {
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
		Minspec: {
			OS: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)',
			CPU: 'Intel Core i5 2400s @ 2.5 GHz / AMD FX 6350 @ 3.9 GHz',
			GPU: 'NVIDIA GeForce GTX 660 / AMD Radeon R9 270 (2GB VRAM with Shader Model 5.0)',
			Memory: '6 GB RAM',
			Storage: '100 GB',
			'Sound Card': 'DirectX Compatible Sound Card with latest drivers'
		},
		Recomendedspec: {
			OS: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)',
			CPU: 'Intel Core i7 3770 @ 3.5 GHz / AMD FX 8350 @ 4.0 GHz',
			GPU: 'NVIDIA GeForce GTX 760 (4GB) or the newer GTX 970 (4GB) / AMD Radeon R9 280X (3GB of VRAM) or better',
			Memory: '8 GB RAM',
			Storage: '200 GB',
			'Sound Card': 'DirectX Compatible Sound Card with latest drivers'
		},
		Price: 34.55,
		Discount: 25,
		Developer: 'Ubisoft',
		Publisher: 'Ubisoft',
		Released: '02/20/20',
		Platform: ['windows', 'linux', 'mac'],
		Players: 23478238
	};

	let Gamedatax = {
		GameID: '',
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
		Players: 23478238,
		Comment: [
			{
				Name: 'raka',
				Rating: 7,
				Description:
					'when you assign the cursor returned from the find() method to a variable using the var keyword, the cursor does not automatically iterate. You can use the cursor method forEach() to iterate the cursor and access the documents, as in the following example: var myCursor = db.'
			},
			{
				Name: 'raka',
				Rating: 7,
				Description:
					'when you assign the cursor returned from the find() method to a variable using the var keyword, the cursor does not automatically iterate. You can use the cursor method forEach() to iterate the cursor and access the documents, as in the following example: var myCursor = db.'
			}
		]
	};
	let TempComment = { Name: $UserProData.Name, Rating: 0 as number, Description: '' };
	function AddComment() {
		Gamedata.Comment.push(TempComment);
	}
	async function SubmitComment(){
		let mydata={
					GameID : Gamedata.GameID as string, 
					Comment : TempComment 
				}
		
		await fetch('/api/productid/postcomment', {
				// credentials: 'same-origin',
				method: 'POST',
				// mode: 'cors',
				body: JSON.stringify(mydata)
			})
				.then((res) => {
					return res.json();
				})
				.then((resdata: any) => {
					console.log("🚀 ~ file: +page.svelte ~ line 240 ~ .then ~ resdata comment", resdata)
					Gamedata.Comment.push(TempComment)
					console.log("🚀 ~ file: +page.svelte ~ line 242 ~ .then ~ Gamedata", Gamedata)
					console.log("🚀 ~ file: +page.svelte ~ line 242 ~ .then ~ Gamedata commnet", Gamedata.Comment)
				});
	}

	let mydata = {
		GameID: '',
		Name: '',
		Moto: '',
		LogoImage: '',
		BigPosterImage: '',
		SmallPosterImage: '',

		OtherImages: ['', '', '', '', '', ''],
		Genres: ['', '', ''],
		Feature: ['', '', '', ''],
		Description: '',
		FollowUs: {
			Facebook: '',
			Discord: '',
			Youtube: '',
			Twitter: '',
			Site: ''
		},
		Rating: Number(),
		RatingGivenBy: {
			'PC Gamer': '',
			IGN: '',
			'Game Informer': ''
		},
		Minspec: [
			{ Name: 'OS', Value: '' },
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
		Recomendedspec: [
			{ Name: 'OS', Value: '' },
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
		Price: Number(),
		Discount: Number(),
		Developer: '',
		Publisher: '',
		Released: '',
		Platform: ['', '', ''],
		Players: Number()
	};

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

	import Happysad from '$lib/svgs/happysad.svelte';
	import Globe from '$lib/svgs/globe.svelte';
	import WriteOnPaper from '$lib/svgs/writeonpaper.svelte';
	import LoveEmoji from '$lib/svgs/loveImoji.svelte';
	import Mouse from '$lib/svgs/mouse.svelte';
	import Cup from '$lib/svgs/cup.svelte';
	import { xlink_attr } from 'svelte/internal';
	import { json } from '@sveltejs/kit';
	import Control from '$lib/svgs/control.svelte';
	import { navigating } from '$app/stores';
	import SaturnLoading from '$lib/loading/saturnLoading.svelte';
	import World from '$lib/svgs/world.svelte';
	import Textarea from '$lib/otherComponents/textarea.svelte';
	import { UserSettingSelect } from '$lib/Store/store';
	import { comment } from 'postcss';
	let TheGameHas = [
		{ title: 'Amazing Characters', icon: Happysad },
		{ title: 'Diverse Characters', icon: Globe },
		{ title: 'Amazing Storytelling', icon: WriteOnPaper },
		{ title: 'Highly Recommended', icon: LoveEmoji },
		{ title: 'Quickly Understood Controls', icon: Control },
		{ title: 'Amazing Characters', icon: Cup }
	];
</script>

<div class=" flex h-fit w-full flex-row justify-center  gap-6  text-slate-100">
	<div class="flex flex-col gap-3 ">
		<p class=" text-5xl text-slate-100">{Gamedata.Name}</p>
		<div class=" flex flex-row gap-4">
			<Rating3 value={Gamedata.Rating} />
			<Peoples class="h-6 w-6 stroke-blue-400" />
			<p class=" -ml-2 text-lg">{RoundNumOfPeople(Gamedata.Players)}</p>
		</div>
		<div class="flex flex-row gap-3 ">
			<div class="flex w-[1080px] flex-col gap-4">
				<!-- image show case -->
				<Daisyui1 OtherImages={Gamedata.OtherImages} />
				<p class=" text-2xl  ">{Gamedata.Moto}</p>
				<!-- Genres and features -->
				<div class="mt-12 flex w-full flex-row  font-Poppins">
					<!-- GENRE -->
					<div class="flex w-1/2 flex-col border-l-[1px] border-slate-400 border-opacity-30 p-3">
						<p class=" font-normal text-gray-500">Genres</p>
						<div class=" flex  flex-row gap-2 text-base">
							{#each Gamedata.Genres as G}
								<p class="">{G} ,</p>
							{/each}
						</div>
					</div>
					<!-- features -->
					<div class="flex w-1/2 flex-col border-l-[1px] border-slate-400 border-opacity-30 p-3">
						<p class=" font-normal text-gray-500">features</p>
						<div class=" flex  flex-row gap-2 text-base">
							{#each Gamedata.Feature as G}
								<p class="">{G} ,</p>
							{/each}
						</div>
					</div>
				</div>
				<!-- description -->
				<div class=" flex flex-col ">
					<p class=" text-xl font-semibold">Description :</p>
					<p class="">
						{Gamedata.Description}
					</p>
				</div>
				<div class="flex flex-col gap-2 ">
					<p class="text-xl font-semibold">Follow Us :</p>
					<div class=" flex flex-row justify-center gap-5 rounded-lg bg-[#262626] py-3 ">
						{#if Gamedata.FollowUs.Site.trim() != ''}
							<button class="">
								<!-- <Fb class="h-7 w-7 fill-slate-200" /> -->
								<World class="h-9 w-9 fill-slate-200" />
							</button>
						{/if}
						{#if Gamedata.FollowUs.Facebook.trim() != ''}
							<button class="">
								<Fb class="h-7 w-7 fill-slate-200" />
							</button>
						{/if}
						{#if Gamedata.FollowUs.Youtube.trim() != ''}
							<button class="">
								<Youtube class="h-8 w-8 fill-slate-200" />
							</button>
						{/if}
						{#if Gamedata.FollowUs.Discord.trim() != ''}
							<button class="">
								<Discord class="h-8 w-8 fill-slate-200" />
							</button>
						{/if}
						{#if Gamedata.FollowUs.Twitter.trim() != ''}
							<button class="">
								<Twitter class="h-8 w-8 fill-slate-200" />
							</button>
						{/if}
					</div>
				</div>
				<!-- Org Review -->
				<p class="text-xl font-semibold">Organizations Review :</p>
				<div class=" flex flex-row justify-around gap-2">
					{#each Object.entries(Gamedata.RatingGivenBy) as [title, Value]}
						<!-- content here -->
						<div class="h-[300px] w-[250px] rounded-xl bg-[#262626] p-4  ">
							<p class="border-b-[1px] border-gray-500  border-opacity-10 pb-2 text-lg ">
								{title}
							</p>
							<p class=" pt-2 text-sm text-slate-400">{Value}</p>
						</div>
					{/each}
				</div>
				<div class="flex flex-row flex-wrap justify-center gap-3	 ">
					{#each TheGameHas as q}
						<div
							class=" flex h-72 w-80 flex-col items-center justify-center rounded-md bg-[#262626] p-4 align-middle "
						>
							<svelte:component this={q.icon} />
							<!-- <p class="">{x.}</p> -->
							<div class=" w-fit text-sm">The game has</div>
							<div class="w-fit text-center text-3xl font-semibold ">{q.title}</div>
						</div>
					{/each}
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
								{#each Gamedata.Minspec as option}
									<div class="flex flex-col gap-1">
										<p class=" font-semibold text-gray-400">{option.Name} :</p>
										<!-- <p class="">{Value}</p> -->
										<p class=" ml-2">{option.Value}</p>
									</div>
								{/each}
								<!-- {#each Object.entries(Gamedata.Minspec) as [title, Value]}
									<div class="flex flex-col gap-1">
										<p class=" text-gray-500">{title}</p>
										<p class="">{Value}</p>
									</div>
								{/each} -->
							</div>
						</div>
						<div class=" flex  w-1/2 flex-col gap-4 ">
							<p class=" text-lg font-semibold">Recomended Requirement :</p>
							<div class=" flex flex-col gap-3 ">
								{#each Gamedata.Recomendedspec as option}
									<div class="flex flex-col gap-1">
										<p class=" font-semibold text-gray-400">{option.Name} :</p>
										<!-- <p class="">{Value}</p> -->
										<p class=" ml-2">{option.Value}</p>
									</div>
								{/each}
								<!-- {#each Object.entries(Gamedata.Recomendedspec) as [title, Value]}
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
			<GameProfile
				LogoImage={Gamedata.LogoImage}
				Released={Gamedata.Released}
				Price={Gamedata.Price}
				Platform={Gamedata.Platform}
				Publisher={Gamedata.Publisher}
				Develper={Gamedata.Developer}
				Discount={Gamedata.Discount}
			/>
		</div>
		<p class=" text-2xl font-Poppins">Reviews : </p>
		<!-- show COMMENT -->
		<div class=" flex flex-col gap-3">
			{#if Gamedata.Comment != null}
				{#each Gamedata.Comment as comment}
				<div class=" flex flex-col gap-1">
					<p class=" font-semibold text-blue-600">{comment.Name}</p>
					<Rating3 value={comment.Rating} />
					<p class="">{comment.Description}</p>
					
				</div>
				{/each}
				 <!-- content here -->
			{/if}
		</div>
		<!-- COMMENT -->
		<div class=" flex flex-col gap-2">
			<p class=" text-2xl  ">Leave Your Review :</p>
			<div class=" gap-2 flex flex-row">

				<Rating3 value={TempComment.Rating} />
				<input
				bind:value={TempComment.Rating}
				type="number"
				placeholder="Ratings"
				class=" h-8 w-24 self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<textarea
				type="text"
				placeholder="Game Moto"
				bind:value={TempComment.Description}
				class=" scrol3 h-20 w-full self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 font-Poppins text-lg  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<div class=" self-end">
				<button
					on:click={() => {
						SubmitComment()
					}}
					class="h-10 w-fit rounded-xl bg-blue-600 px-3 font-Poppins font-semibold text-white hover:bg-sky-500 active:bg-blue-800  "
				>
					Post Comment 
				</button>
			</div>
		</div>
	</div>
	<!-- <Showcase /> -->
</div>

<style>
	/* your styles go here */
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
