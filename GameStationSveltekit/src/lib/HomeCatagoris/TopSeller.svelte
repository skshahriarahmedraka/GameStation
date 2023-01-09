<script lang="ts">
	import MidGameCom from '$lib/cards/midGameCom.svelte';
	import Plus from '$lib/svgs/plus.svelte';

	export let TopSoledList: {
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
	}[] = [] as {
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
	}[];
	// let Obj = {
	// 	Name: "Assassin's Creed Syndicate Standard Edition",
	// 	Price: 87.43,
	// 	Rating: 9.7,
	// 	NumberOfPlayers: 34821571533,
	// 	Image:
	// 		'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640'
	// };
	// let LiObj = [
	// 	{
	// 		ID: '1267842',
	// 		Name: "Assassin's Creed Syndicate Standard Edition",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:
	// 			'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640'
	// 	},
	// 	{
	// 		ID: '1267842',
	// 		Name: "Marvel's Spider-Man Remastered",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:"https://cdn2.unrealengine.com/egs-spider-man-remaster-carousel-thumb-1200x1600-c51cad22ccfc.jpg?h=1280&resize=1&w=960"
	// 	},
	// 	{
	// 		ID: '1267842',
	// 		Name: "Genshin Impact",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:"https://cdn2.unrealengine.com/egs-genshin-impact-3-0-carousel-thumb-en-1200x1600-c0b892d410a8.jpg?h=1280&resize=1&w=960"
	// 	},
	// 	{
	// 		ID: '1267842',
	// 		Name: "Rocket League",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:"https://cdn2.unrealengine.com/egs-rl-season-8-carousel-thumb-315x399-a58151517871.jpg?h=1280&resize=1&w=960"
	// 	},
	// 	{
	// 		ID: '1267842',
	// 		Name: "Assassin's Creed Syndicate Standard Edition",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:
	// 			'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640'
	// 	},
	// 	{
	// 		ID: '1267842',
	// 		Name: "Grand Theft Auto V: Premium Edition",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:"https://cdn1.epicgames.com/0584d2013f0149a791e7b9bad0eec102/offer/GTAV_EGS_Artwork_1200x1600_Portrait Store Banner-1200x1600-382243057711adf80322ed2aeea42191.jpg?h=854&resize=1&w=640"
	// 	},
	// 	{
	// 		ID: '1267842',
	// 		Name: "Shadow of the Tomb Raider: Definitive Edition",
	// 		Price: 87.43,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Image:"https://cdn1.epicgames.com/offer/4b5461ca8d1c488787b5200b420de066/egs-shadowofthetombraiderdefinitiveedition-eidosmontralcrystaldynamicsnixxessoftware-s4-1200x1600-7ee40d6fa744_1200x1600-950cdb624cc75d04fe3c8c0b62ce98de?h=854&resize=1&w=640"
	// 	}
	// ];
</script>

<div class="flex h-fit w-full flex-col justify-center ">
	<div class="mx-24 flex h-16 flex-row">
		<button class=" text-xl text-slate-200 hover:text-white">Most Sold</button>
		<div class="grow" />
		<button class=" mr-28 text-lg text-slate-400  hover:text-white"
			><Plus class="inline-flex h-7 w-7" /> See more</button
		>
	</div>
	<div class=" mx-24 flex flex-row flex-wrap space-x-4 space-y-4   ">
		{#if TopSoledList != null}
			{#each TopSoledList as Obj}
				<MidGameCom {Obj} />
			{/each}
		{:else}
			{#each Array(10) as _, i}
				<MidGameCom />
			{/each}
		{/if}
		<!-- {#each TopSoledList as Obj}
			<MidGameCom {Obj} />
		{/each} -->
	</div>
</div>

<style>
	/* your styles go here */
</style>
