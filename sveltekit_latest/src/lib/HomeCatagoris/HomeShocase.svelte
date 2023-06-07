<script lang="ts">
	import ImageShow from '$lib/Carousels/imageShow.svelte';
	import Showcase from '$lib/Carousels/showcase.txt';
	import Blog from '$lib/cards/midCard.txt';
	import Nav from '$lib/navbar/nav.svelte';
	import Smallcard from '$lib/cards/smallcard.svelte';
	import Daisyui1 from '$lib/Carousels/daisyui1.svelte';
	import type { GameDataType } from '$lib/Store/types';
	let mirage = new URL('../smallImages/mirage.jpg', import.meta.url).href;
	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;

	export let CarouselList:GameDataType[] = [] as GameDataType[]
	// let Obj = {
	// 	ID: '1321432',
	// 	Name: "Assassin's Creed® Mirage",
	// 	Image: mirage,
	// 	Price: '43.65'
	// };
	// let listOfImage = [war,forbiddenwest,anthem,fifa]
	// let liObj = [
	// 	{
	// 		ID: '1321432',
	// 		Name: "Assassin's Creed® Mirage",
	// 		Image: mirage,
	// 		Price: '43.65',
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Moto:"In Marvel’s Spider-Man Remastered, the worlds of Peter Parker and Spider-Man collide in an original, action-packed story. Play as an experienced Peter Parker, fighting big crime and iconic villains in Marvel’s New York. Web-swing through vibrant neighborhoods and defeat villains with epic takedowns."
	// 	},
	// 	{
	// 		ID: '1321432',
	// 		Name: "Assassin's Creed® Mirage",
	// 		Image: anthem,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Price: '43.65',
	// 		Moto:"In Marvel’s Spider-Man Remastered, the worlds of Peter Parker and Spider-Man collide in an original, action-packed story. Play as an experienced Peter Parker, fighting big crime and iconic villains in Marvel’s New York. Web-swing through vibrant neighborhoods and defeat villains with epic takedowns.",

	// 	},
	// 	{
	// 		ID: '1321432',
	// 		Name: "Assassin's Creed® Mirage",
	// 		Image: fifa,
	// 		Rating: 9.7,
	// 		NumberOfPlayers: 34821571533,
	// 		Moto:"In Marvel’s Spider-Man Remastered, the worlds of Peter Parker and Spider-Man collide in an original, action-packed story. Play as an experienced Peter Parker, fighting big crime and iconic villains in Marvel’s New York. Web-swing through vibrant neighborhoods and defeat villains with epic takedowns.",

	// 		Price: '43.65'
	// 	},
	// 	{
	// 		ID: '1321432',
	// 		Name: "Assassin's Creed® Mirage",
	// 		Image: forbiddenwest,
	// 		Rating: 9.7,
	// 		Moto:"In Marvel’s Spider-Man Remastered, the worlds of Peter Parker and Spider-Man collide in an original, action-packed story. Play as an experienced Peter Parker, fighting big crime and iconic villains in Marvel’s New York. Web-swing through vibrant neighborhoods and defeat villains with epic takedowns.",
	// 		NumberOfPlayers: 34821571533,
	// 		Price: '43.65'
	// 	},
	// 	{
	// 		ID: '1321432',
	// 		Name: "Assassin's Creed® Mirage",
	// 		Image: anthem,
	// 		Rating: 9.7,
	// 		Moto:"In Marvel’s Spider-Man Remastered, the worlds of Peter Parker and Spider-Man collide in an original, action-packed story. Play as an experienced Peter Parker, fighting big crime and iconic villains in Marvel’s New York. Web-swing through vibrant neighborhoods and defeat villains with epic takedowns.",
	// 		NumberOfPlayers: 34821571533,
	// 		Price: '43.65'
	// 	},
	// 	{
	// 		ID: '1321432',
	// 		Name: "Assassin's Creed® Mirage",
	// 		Image: fifa,
	// 		Rating: 9.7,
	// 		Moto:"In Marvel’s Spider-Man Remastered, the worlds of Peter Parker and Spider-Man collide in an original, action-packed story. Play as an experienced Peter Parker, fighting big crime and iconic villains in Marvel’s New York. Web-swing through vibrant neighborhoods and defeat villains with epic takedowns.",
	// 		NumberOfPlayers: 34821571533,
	// 		Price: '43.65'
	// 	}
	// ];
</script>

<!-- markup (zero or more items) goes here -->

<!-- <Daisyui1 OtherImages={listOfImage}/> -->
<ImageShow {CarouselList} />
<div class="flex flex-col gap-2 overflow-hidden md:flex-row xl:flex-col lg:flex-row   h-fit">
	{#if CarouselList.length > 0}
		<!-- {#each CarouselList as Obj}
			<Smallcard {Obj} />
		{/each} -->
	{:else}
		<div class="flex flex-col gap-2 overflow-hidden md:flex-row xl:flex-col lg:flex-row   h-fit">
			<Smallcard  />
			<Smallcard  />
			<Smallcard  />
			<Smallcard  />
			<Smallcard />
			<Smallcard />
		</div>
	{/if}
	
</div>

<style>
	/* your styles go here */
</style>
