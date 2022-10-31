

<script lang="ts">
    // 

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
	import GameProfile from "$lib/cards/gameProfile.svelte";

    import { page } from '$app/stores';
	import Error1 from '$lib/Error/error1.svelte';
	import Error2 from '$lib/Error/error2.svelte';
    import { afterUpdate } from 'svelte';
    import Loading from '$lib/loading/saturnLoading.svelte';

    let loadingState: boolean = true;
	afterUpdate(() => {
		loadingState = false;
	});
</script>

<style>
    /* your styles go here */
</style>


{#if loadingState}
	<Loading class=" scale-150" />
{:else if $page.error}
	<!-- <Error1/> -->
	<Error2 />
{:else}

<!-- <slot/> -->

<div class=" min-h-screen min-w-full max-w-full bg-[#121212] flex flex-col gap-5 justify-between overflow-hidden">
    <Nav />
   
	<div class=" w-full h-fit flex flex-row justify-center gap-2">
        <!-- showcase Images and small game cards -->
        
        
    </div>
    
    <div class="flex flex-col justify-center items-center ">
        <!-- <MostPopular /> -->
        <slot />
        
    </div>
    <Footer />
</div>
{/if}




