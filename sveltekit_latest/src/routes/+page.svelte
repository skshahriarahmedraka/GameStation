<script lang="ts">
	//
	import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';

	export let data: PageData;
	
	let { MostPopularList,NewReleaseList,TopRatedList, TopSoledList,TrendingList,CarouselList } = data;
	console.log("🚀 ~ file: +page.svelte:9 ~ MostPopularList:", MostPopularList)
	// UserProData.update((d) => (d = Userdata));

	import ImageShow from '$lib/Carousels/imageShow.svelte';
	import Showcase from '$lib/Carousels/showcase.txt';
	import Blog from '$lib/cards/midCard.txt';
	import Nav from '$lib/navbar/nav.svelte';
	import Smallcard from '$lib/cards/smallcard.svelte';
	import MostPopular from '$lib/HomeCatagoris/MostPopular.svelte';
	import HomeShocase from '$lib/HomeCatagoris/HomeShocase.svelte';
	import Footer from '$lib/foot/footer.svelte';
	import Trending from '$lib/HomeCatagoris/Trending.svelte';
	import NewRelease from '$lib/HomeCatagoris/NewRelease.svelte';
	import TopRated from '$lib/HomeCatagoris/TopRated.svelte';
	import TopSeller from '$lib/HomeCatagoris/TopSeller.svelte';
	// import { json } from '@sveltejs/kit';

	//     import type { PageData } from './$types';
	// 	import { UserProData } from "$lib/Store/store";

	// export let data: PageData;

	// let {Userdata}=data
	// 	console.log("🚀 ~ file: +page.svelte ~ line 25 ~ Userdata", Userdata)
	// 	UserProData.update((d)=>d=Userdata)

	// async function sub() {
	// 	const res = await fetch('http://localhost:8002',{
	// 		method: 'GET',
	// 		mode: 'no-cors',
	// 		headers: {
	// 			'Content-Type': 'application/json'
	// 		}
	// 	}).then((res) => res.json()).then((data) => {
	// 		console.log('🚀 ~ file: +page.svelte ~ line 44 ~ data', data);
	// 		return data;
	// 	});
		
		
	// }

	// sub()
	
	
	
	
</script>

<div class=" flex max-h-fit min-h-screen w-full flex-col gap-5 overflow-hidden bg-[#121212]">
	<Nav />

	<div
		class=" flex h-fit w-full flex-row justify-center gap-2 md:flex-col md:flex-wrap xl:flex-row  "
	>
		<!-- showcase Images and small game cards -->
		<HomeShocase {CarouselList} />
	</div>

	<div class="flex flex-col items-center justify-center ">
		<MostPopular {MostPopularList} />
		<Trending {TrendingList} />
		<NewRelease {NewReleaseList} />
		<TopRated {TopRatedList} />
		<TopSeller {TopSoledList} />
	</div>
	<Footer />
</div>

<style>
	/* your styles go here */
</style>
