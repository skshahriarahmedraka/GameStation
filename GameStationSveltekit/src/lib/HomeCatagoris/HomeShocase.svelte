<script lang="ts">
	import ImageShow from '$lib/Carousels/imageShow.svelte';
	import Showcase from '$lib/Carousels/showcase.svelte';
	import Blog from '$lib/cards/midCard.svelte';
	import Nav from '$lib/navbar/nav.svelte';
	import Smallcard from '$lib/cards/smallcard.svelte';
	let mirage = new URL('../smallImages/mirage.jpg', import.meta.url).href;
	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;

	// let Obj = {
	// 	ID: '1321432',
	// 	Name: "Assassin's Creed® Mirage",
	// 	Image: mirage,
	// 	Price: '43.65'
	// };

	let liObj = [
		{
			ID: '1321432',
			Name: "Assassin's Creed® Mirage",
			Image: mirage,
			Price: '43.65',
			Rating: 9.7,
			NumberOfPlayers: 34821571533
		},
		{
			ID: '1321432',
			Name: "Assassin's Creed® Mirage",
			Image: anthem,
			Rating: 9.7,
			NumberOfPlayers: 34821571533,
			Price: '43.65'
		},
		{
			ID: '1321432',
			Name: "Assassin's Creed® Mirage",
			Image: fifa,
			Rating: 9.7,
			NumberOfPlayers: 34821571533,
			Price: '43.65'
		},
		{
			ID: '1321432',
			Name: "Assassin's Creed® Mirage",
			Image: forbiddenwest,
			Rating: 9.7,
			NumberOfPlayers: 34821571533,
			Price: '43.65'
		},
		{
			ID: '1321432',
			Name: "Assassin's Creed® Mirage",
			Image: anthem,
			Rating: 9.7,
			NumberOfPlayers: 34821571533,
			Price: '43.65'
		},
		{
			ID: '1321432',
			Name: "Assassin's Creed® Mirage",
			Image: fifa,
			Rating: 9.7,
			NumberOfPlayers: 34821571533,
			Price: '43.65'
		}
	];
</script>

<!-- markup (zero or more items) goes here -->

<ImageShow {liObj} />
<div class="flex flex-col gap-2 overflow-hidden">
	{#each liObj as Obj}
		<Smallcard {Obj} />
	{/each}
</div>

<style>
	/* your styles go here */
</style>
