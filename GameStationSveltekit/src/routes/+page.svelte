

<script lang="ts">
    // 
    import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';


	export let data: PageData;

	let { Userdata } = data;
	console.log('🚀 ~ file: +layout.svelte ~ line 26 ~ Userdata', Userdata);
	UserProData.update((d) => (d = Userdata));

	import ImageShow from "$lib/Carousels/imageShow.svelte";
import Showcase from "$lib/Carousels/showcase.txt";
import Blog from "$lib/cards/midCard.txt";
import Nav from "$lib/navbar/nav.svelte";
	import Smallcard from "$lib/cards/smallcard.svelte";
	import MostPopular from "$lib/HomeCatagoris/MostPopular.svelte";
	import HomeShocase from "$lib/HomeCatagoris/HomeShocase.svelte";
	import Footer from '$lib/foot/footer.svelte';
	import Trending from "$lib/HomeCatagoris/Trending.svelte";
	import NewRelease from "$lib/HomeCatagoris/NewRelease.svelte";
	import TopRated from "$lib/HomeCatagoris/TopRated.svelte";
	import TopSeller from "$lib/HomeCatagoris/TopSeller.svelte";

//     import type { PageData } from './$types';
// 	import { UserProData } from "$lib/Store/store";

// export let data: PageData;

// let {Userdata}=data
// 	console.log("🚀 ~ file: +page.svelte ~ line 25 ~ Userdata", Userdata)
// 	UserProData.update((d)=>d=Userdata)

</script>

<style>
    /* your styles go here */
</style>

<div class=" min-h-screen max-h-fit w-full bg-[#121212] flex flex-col gap-5 overflow-hidden">
    <Nav />
   
	<div class=" w-full h-fit flex flex-row xl:flex-row justify-center gap-2 md:flex-col md:flex-wrap  ">
        <!-- showcase Images and small game cards -->
        <HomeShocase />
        
    </div>
    
    <div class="flex flex-col justify-center items-center ">
        <MostPopular />
        <Trending />
        <NewRelease/> 
        <TopRated/>
        <TopSeller />
    </div>
    <Footer />
</div>



