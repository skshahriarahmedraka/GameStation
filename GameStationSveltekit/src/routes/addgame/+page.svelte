<script lang="ts">
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

	let data: {
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
		FollowUs: {
			Facebook: string;
			Youtube: string;
			Discord: string;
			Twitter: string;
			Site: string;
		};
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
		Price: number;
		Discount: number;
		Developer: string;
		Publisher: string;
		Released: string;
		Platform: string[];
		Players: number;
	} = {
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
			data.Genres = GenreObj.List;
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
			data.Feature = FeaturesObj.List;
			FeaturesObj.Input = '';
		}
	}
	// $: console.log("🚀 ~ file: +page.svelte ~ line 137 ~ data", data)

	// SUBMIT A GAME
	let responseClone;
	import { json } from '@sveltejs/kit';
	import axios from 'axios';
	import Globe from '$lib/svgs/globe.svelte';
	import World from '$lib/svgs/world.svelte';
	async function PostGameData() {
		// console.log('🚀 ~ file: +page.svelte ~ line 306 ~ PostGameData', data);
		fetch('/api/addgame', {
			// credentials: 'same-origin',
			method: 'POST',
			// mode: 'no-cors',
			headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(data)
		})
			.then((response) => response.json())
			.then((data) => {
				console.log('🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ REsponse', data);
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
</script>

<!-- markup (zero or more items) goes here -->

<div class=" flex h-fit w-full flex-row justify-center  gap-6  text-slate-100">
	<div class="flex flex-col gap-3 ">
		<!-- Game Name -->
		<input
			bind:value={data.Name}
			type="text"
			placeholder="Game Name"
			class=" my-2 h-16 w-[70%] self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2  text-5xl font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
		<!-- <p class=" text-5xl text-slate-100">{Gamedata.Name}</p> -->
		<div class=" flex flex-row gap-4 ">
			<Rating3 value={data.Rating} />
			<input
				bind:value={data.Rating}
				type="text"
				placeholder="Ratings"
				class=" h-8 w-24 self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<Peoples class="h-6 w-6 stroke-blue-400" />
			<p class=" -ml-2 text-lg">{RoundNumOfPeople(data.Players)}</p>
			<input
				type="text"
				placeholder="Ratings"
				bind:value={data.Players}
				class=" h-8 w-24 self-start rounded-xl border-2 border-[#24262b]  bg-[#303338]  p-2 text-base  font-medium text-slate-100   outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
		</div>
		<div class="flex xl:flex-row gap-3  xs:flex-col-reverse sm:flex-col-reverse md:flex-col-reverse  lg:flex-row justify-center">
			<div class="flex max-w-[1080px] flex-col gap-4 ">
				<!-- image show case -->
				<!-- <Daisyui1 OtherImages={Gamedata.OtherImages} /> -->

				<InputCarousels
					bind:OtherImages={data.OtherImages}
					bind:BigPosterImage={data.BigPosterImage}
				/>

				<p class=" text-2xl  ">Game Moto :</p>
				<textarea
					type="text"
					placeholder="Game Moto"
					bind:value={data.Moto}
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
									{#each GenreObj.List as t}
										<li
											class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
										>
											{t}
										</li>
									{/each}
									<input
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
							<!-- {#each Gamedata.Genres as G}
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
								<ul class="flex flex-row flex-wrap gap-2">
									{#each FeaturesObj.List as t}
										<li
											class="m-1 rounded-md bg-[#3d4951] px-2 py-1 text-[#9bc0da] hover:cursor-pointer hover:bg-slate-600 hover:text-teal-200"
										>
											{t}
										</li>
									{/each}
									<input
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

							<!-- {#each Gamedata.Feature as G}
								<p class="">{G} ,</p>
							{/each} -->
						</div>
					</div>
				</div>
				<!-- description -->
				<div class=" flex flex-col gap-2 ">
					<p class=" text-xl font-semibold">Description :</p>
					<!-- <p class="">
						{Gamedata.Description}
					</p> -->
					<Textarea bind:value={data.Description} minRows={4} maxRows={40} />
				</div>
				<div class="flex flex-col gap-2 ">
					<p class="text-xl font-semibold">Follow Us :</p>
					<div class=" flex flex-col justify-center gap-5 rounded-lg bg-[#262626] py-3 ">
						<div class="flex flex-row justify-center gap-2">
							<button class="">
								<!-- <Fb  /> -->
								<World class="h-9 w-9 fill-slate-200"/>
							</button>
							<input
								type="text"
								placeholder="WebSite"
								bind:value={data.FollowUs.Site}
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
								bind:value={data.FollowUs.Facebook}
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
								bind:value={data.FollowUs.Youtube}
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
								bind:value={data.FollowUs.Discord}
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
								bind:value={data.FollowUs.Twitter}
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
							bind:value={data.RatingGivenBy['Game Informer']}
							minRows={4}
							maxRows={40}
						/>
					</div>

					<div class="h-[300px] w-[250px] rounded-xl bg-[#262626] p-4  ">
						<p class="border-b-[1px] border-gray-500  border-opacity-10 pb-2 text-lg ">IGN</p>
						<!-- <p class=" pt-2 text-sm text-slate-400">{Value}</p> -->
						<Textarea
							class="pt-2 text-sm text-slate-400"
							bind:value={data.RatingGivenBy.IGN}
							minRows={4}
							maxRows={40}
						/>
					</div>
					<div class="h-[300px] w-[250px] rounded-xl bg-[#262626] p-4  ">
						<p class="border-b-[1px] border-gray-500  border-opacity-10 pb-2 text-lg ">PC Gamer</p>
						<!-- <p class=" pt-2 text-sm text-slate-400">{Value}</p> -->
						<Textarea
							class="pt-2 text-sm text-slate-400"
							bind:value={data.RatingGivenBy['PC Gamer']}
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
								{#each data.Minspec as option}
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
								{#each data.Recomendedspec as option}
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
			<InputGameProfile
				bind:LogoImage={data.LogoImage}
				bind:SmallPosterImage={data.SmallPosterImage}
				bind:Released={data.Released}
				bind:Price={data.Price}
				bind:Platform={data.Platform}
				bind:Publisher={data.Publisher}
				bind:Develper={data.Developer}
				bind:Discount={data.Discount}
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
</style>
