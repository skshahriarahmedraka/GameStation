<script lang="ts">
    // your script goes here
    import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';

	export let data: PageData;
	
	let { Userdata } = data;
	console.log("🚀 ~ file: +page.svelte ~ line 10 ~ Userdata", Userdata)
	UserProData.update((d) => (d = Userdata));
</script>

<style>
    /* your styles go here */
</style>


<div class=" text-white">
    Offers
</div>