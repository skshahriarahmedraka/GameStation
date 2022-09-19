

<script lang="ts">
	import '../app.css';
    // 
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

<slot/>

{/if}