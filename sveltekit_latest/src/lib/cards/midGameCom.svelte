<script lang="ts">
	import Star from '$lib/svgs/star.svelte';
	import Users from '$lib/svgs/users.svelte';
	import Dot from '$lib/svgs/dot.svelte';
	import { goto } from '$app/navigation';
	import type { GameDataType } from '$lib/Store/types';

	export let Obj: GameDataType={} as GameDataType
	function RoundNumOfPeople(x: number) {
		if (x >= 1000000000) {
			x /= 1000000000;
			return String(x.toFixed(2)) + 'B';
		} else if (x >= 1000000) {
			x /= 1000000;
			return String(x.toFixed(2)) + 'M';
		} else if (x >= 1000) {
			x /= 1000;
			return String(x.toFixed(2)) + 'K';
		} else {
			return String(x);
		}
	}
</script>

{#if JSON.stringify(Obj) != '{}'  }
<button
on:click={()=>{
					goto(`/${Obj.GameID}`);
					
	}}
	class=" mt-4 mb-4 flex max-h-fit min-h-fit w-[215px] flex-col gap-1 rounded-md pb-2 text-slate-200    shadow-lg    transition-all duration-100 ease-linear hover:scale-125 hover:cursor-pointer hover:bg-[#202022] hover:shadow-gray-600"
	>
	<img
	src={Obj.SmallPosterImage}
	alt="Game Poster"
		class="top-0 h-[285px] w-full rounded-t-lg object-cover"
	/>
	<p class=" ml-2 inline-flex text-start font-Poppins text-sm  ">
		{Obj.Rating}/10 <Star class="inline-flex h-4 w-4 stroke-yellow-200" />
		<Dot class=" mx-2 self-center" />
		{RoundNumOfPeople(Obj.Players)}
		<Users class="ml-1 inline-flex h-4 w-4 " />
	</p>
	<p class=" ml-1 text-start font-Poppins text-base line-clamp-2">{Obj.Name}</p>
	<p class=" ml-2 text-start font-Poppins text-sm font-thin ">${Obj.Price}</p>
</button>
{:else} 
<button

	class=" mt-4 mb-4 flex max-h-fit min-h-fit w-[215px] flex-col gap-1 rounded-md pb-2 text-slate-200    shadow-lg  bg-[#1c1c1c]  transition-all duration-100 ease-linear hover:scale-125 hover:cursor-pointer hover:bg-[#202022] hover:shadow-gray-600"
	>
	<div
			class="top-0 h-[285px] w-full rounded-t-lg object-cover bg-[#252525] "
	/>
	<div class=" m-2 inline-flex  text-start font-Poppins text-sm bg-[#252525] h-6 ">
		<!-- {0}/10 <Star class="inline-flex h-4 w-4 stroke-yellow-200" />
		<Dot class=" mx-2 self-center" />
		{RoundNumOfPeople(0)}
		<Users class="ml-1 inline-flex h-4 w-4 " /> -->
	</div>
	<div class=" m-1 h-4 text-start font-Poppins text-base line-clamp-2 bg-[#252525]"></div>
	<!-- <p class=" ml-2 text-start font-Poppins text-sm font-thin ">${0}</p> -->
</button>
{/if}

<style>
	/* your styles go here */
</style>
