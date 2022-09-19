

<script lang="ts">
    // 

	import ImageShow from "$lib/Carousels/imageShow.svelte";
import Showcase from "$lib/Carousels/showcase.svelte";
import Blog from "$lib/cards/midCard.svelte";
import Nav from "$lib/navbar/nav.svelte";
	import Smallcard from "$lib/cards/smallcard.svelte";
	import MostPopular from "$lib/HomeCatagoris/MostPopular.svelte";
	import HomeShocase from "$lib/HomeCatagoris/HomeShocase.svelte";
	import Footer from '$lib/foot/footer.svelte';
	import Trending from "$lib/HomeCatagoris/Trending.svelte";
	import NewRelease from "$lib/HomeCatagoris/NewRelease.svelte";
	import TopRated from "$lib/HomeCatagoris/TopRated.svelte";
	import TopSeller from "$lib/HomeCatagoris/TopSeller.svelte";


</script>

<style>
    /* your styles go here */
</style>

<div class=" min-h-screen min-w-full bg-[#121212] flex flex-col gap-5">
    <Nav />
   
	<div class=" w-full h-fit flex flex-row justify-center gap-2">
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



