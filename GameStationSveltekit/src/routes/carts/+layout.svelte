<script lang="ts">
    // your script goes here
import Nav from "$lib/navbar/nav.svelte";
import Footer from '$lib/foot/footer.svelte';
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

<!-- markup (zero or more items) goes here -->
<div class=" min-h-screen min-w-full bg-[#1c1c1c] flex flex-col gap-5 justify-between overflow-hidden ">
    <Nav />
   
	
    <slot></slot>
    
    
    <Footer />
</div>
{/if}
