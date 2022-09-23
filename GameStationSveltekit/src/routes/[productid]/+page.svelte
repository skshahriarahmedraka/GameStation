<script lang="ts">
	import GameProfile from '$lib/cards/gameProfile.svelte';
	import Showcase from '$lib/Carousels/showcase.svelte';
	// import Showcase from "$lib/Carousels/showcase.svelte";
	import Daisyui1 from '$lib/Carousels/daisyui1.svelte';
	import RatingsStar from '$lib/logo/ratingsStar.svelte';
	import Rating3 from '$lib/logo/rating4.svelte';

	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
	import { Carousel, CarouselTransition } from 'flowbite-svelte';
	import Game from '$lib/logo/game.svelte';
	const syn1 = new URL('../../lib/images/assassinscreedsyndicate/1.jpg', import.meta.url).href;
	const syn2 = new URL('../../lib/images/assassinscreedsyndicate/2.jpg', import.meta.url).href;
	const syn3  = new URL('../../lib/images/assassinscreedsyndicate/3.jpg', import.meta.url).href;
	const syn4 = new URL('../../lib/images/assassinscreedsyndicate/4.jpg', import.meta.url).href;
	const syn5 = new URL('../../lib/images/assassinscreedsyndicate/5.jpg', import.meta.url).href;
	// const syn6 = new URL('../../lib/images/assassinscreedsyndicate/6.jpg', import.meta.url).href;
	const syn7 = new URL('../../lib/images/assassinscreedsyndicate/7.jpg', import.meta.url).href;


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

	let Gamedata = {
        GameID:"1234",
		Name: "Assassin's Creed Syndicate",
		Moto: 'London, 1868. In the heart of the Industrial Revolution, play as Jacob Frye - a brash and charismatic Assassin.',

		LogoImage: "https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACV_UCS12483_EGST_Logo_Color_1000x407-1000x407-48fad9206ee895c599365e2b7c5bebcab099bd26.png?h=270&resize=1&w=480",
		BigPosterImage: 'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACS-STD-2560x1440-635b7b6c86f18730071426375e7c4fe0bd831ddd.jpg',
		SmallPosterImage: 'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640',
		
        OtherImages: [syn1,syn2,syn3 ,syn4,syn5,syn7],
        Genres: ['Action', 'RPG', 'Adventure'],
		Feature: ['achivement', 'control', 'support'],
		Description: `As Jacob Frye, a young and reckless Assassin, use your skills to help those trampled by the march of progress.
Travel the city at the height of the Industrial Revolution and meet iconic historical figures. From Westminster to Whitechapel, you will come across Darwin, Dickens, Queen Victoria… and many more.
As a gang leader, strengthen your stronghold and rally rival gang members to your cause, in order to take back the capital from the Templars’ hold.`,
		FollowUs: { facebook: 'https://www.facebook.com/epicgames', discord: 'https://discord.com/app',youtube:"https://www.youtube.com/channel/UC5Qk8mWBwtMyEj7iQQYRk1A", twitter: 'https://twitter.com/epicgames', site: 'https://store.epicgames.com/en-US/p/assassins-creed-syndicate' },
		Rating: 9,
		RatingGivenBy: {
			'PC Gamer': "The Assassin’s Creed series has been running for eight years, and that kind of longevity doesn’t happen without taking a few risks. To combat fatigue, each entry attempts to punctuate the familiar with new elements. Not all of these experiments pay off, ",
			IGN: "Syndicate smartly negotiates this internal conflict by dramatising it in the form of its twin playable characters, Evie and Jacob Frye. The former is a devout Assassin, intent on stopping the Templars by tracking down the remaining pieces of Eden. Her brother Jacob, however, is a pragmatist – a social reformer who rails against the Assassin’s burdensome legacy",
			'Game Informer': "Syndicate moves Assassin's Creed forward with a solid new adventure, a beautiful London playground and a renewed sense of fun"
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
        Developer: "Ubisoft",
        Publisher:"Ubisoft",
        Released: "02/20/20",
        Platform:["windows","linux", "mac"]
	};
</script>

<!-- markup (zero or more items) goes here -->

<div class=" flex h-fit w-full flex-row justify-center  gap-6  text-slate-100">
	<div class="flex flex-col gap-3 ">
		<p class=" text-5xl text-slate-100">{Gamedata.Name}</p>
		<div class="">
			<Rating3 value={Gamedata.Rating} />
		</div>
		<div class="flex flex-row gap-3 ">
			<div class="flex flex-col w-[1080px]">
				<Daisyui1 OtherImages={Gamedata.OtherImages} />
                <p class=" text-2xl mt-4 ">{Gamedata.Moto}</p>
                <div class="w-full mt-12">
                    <!-- GENRE -->
                    <div class="w-1/2 flex flex-col">
                        <p class="">Genres</p>
                        <div class="">

                        </div>
                    </div>
                    <!-- features -->
                    <div class="w-1/2"></div>
                </div>

			</div>
			<GameProfile LogoImage={Gamedata.LogoImage} Released={Gamedata.Released} Price={Gamedata.Price} Platform={Gamedata.Platform} Publisher={Gamedata.Publisher} Develper={Gamedata.Publisher} Discount={Gamedata.Discount} />
		</div>
        <div class=""></div>
	</div>
	<!-- <Showcase /> -->
</div>

<style>
	/* your styles go here */
</style>
