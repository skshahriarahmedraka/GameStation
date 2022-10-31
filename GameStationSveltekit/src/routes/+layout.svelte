<script lang="ts">
	import '../app.css';
	// import "flowbite/dist/flowbite.css";
	//
	import { page } from '$app/stores';
	import Error1 from '$lib/Error/error1.svelte';
	import Error2 from '$lib/Error/error4.svelte';
	import { afterUpdate } from 'svelte';
	import Loading from '$lib/loading/saturnLoading.svelte';

	// console.log("UserProData : ",$UserProData)
	// let myval
	// UserProData.subscribe(val=>{
	// 	myval=val
	// })
	// console.log("🚀 ~ file: +layout.svelte ~ line 18 ~ UserProData", myval)

	
	let loadingState: boolean = true;
	afterUpdate(() => {
		loadingState = false;
	});
</script>

<svelte:head>
	<title>GameStation</title>
</svelte:head>

{#if loadingState}
	<Loading class=" " />
{:else if $page.error}
	<!-- <Error1/> -->
	<Error2 />
{:else}
	<slot />
{/if}

<style>
	/* your styles go here */
</style>
