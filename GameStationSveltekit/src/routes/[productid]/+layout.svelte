<script lang="ts">
	// your script goes here
	import Nav from '$lib/navbar/nav.svelte';
	import Footer from '$lib/foot/footer.svelte';

	import navigationState from '$lib/Store/store';

	import { navigating } from '$app/stores';
	import SaturnLoading from '$lib/loading/saturnLoading.svelte';
</script>

<!-- markup (zero or more items) goes here -->
<div class=" flex min-h-screen w-full flex-col justify-between gap-5 overflow-hidden bg-[#121212]">
	<Nav />
	{#if $navigating}
		<!-- content here -->
		<SaturnLoading class=" scale-105 " />
	{:else}
		<slot />
	{/if}

	<Footer />
</div>

<style>
	/* your styles go here */
</style>
