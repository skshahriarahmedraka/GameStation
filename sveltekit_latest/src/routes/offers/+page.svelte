<script lang="ts">
    // your script goes here
    import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';
	// import UpcommingmidGameCom from '$lib/cards/UpcommingmidGameCom.svelte';
	import OffermidGameCom from '$lib/cards/OffermidGameCom.svelte';

	export let data: PageData;
	
	let { Userdata ,OfferGameList } = data;
	console.log("🚀 ~ file: +page.svelte ~ line 10 ~ Userdata", Userdata)
	UserProData.update((d) => (d = Userdata));
</script>

<style>
    /* your styles go here */
</style>


<div
	class=" flex max-h-fit min-h-screen w-[1440px] items-center  flex-row justify-center md:flex-col-reverse md:justify-evenly sm:justify-evenly xs:justify-evenly md:w-full md:px-4 sm:flex-col-reverse sm:w-full sm:px-4 xs:flex-col-reverse xs:w-full xs:px-4 gap-6  text-slate-100"
>
	<div class="h-fit w-[1178px] self-center  md:w-full xs:w-full sm:w-full md:justify-evenly sm:justify-evenly xs:justify-evenly">
		<p class="">Filtered Games</p>
		<div class=" flex flex-row  flex-wrap gap-4 md:justify-evenly sm:justify-evenly xs:justify-evenly">

			{#if OfferGameList != null}

			{#each OfferGameList as i}
			<OffermidGameCom  Obj={i} />
			{/each}

			{:else}
			<p class="text-2xl">No Games Found</p>
			{/if}
		</div>
		<!--  -->
	</div>
</div>