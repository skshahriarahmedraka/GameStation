<script lang="ts">
	import Star from "$lib/svgs/star.svelte";
	import Users from "$lib/svgs/users.svelte";
	import Dot from "$lib/svgs/dot.svelte";
	import { goto } from "$app/navigation";
	import { object_without_properties } from "svelte/internal";

    export let Obj:any

	function RoundNumOfPeople(x:number){
		if (x>=1000000000){
			x/=1000000000
			return String(x.toFixed(2))+"B"
		}
		else if (x>=1000000){
			x/=1000000
			return String(x.toFixed(2))+"M"
		}else if (x>=1000){
			x/=1000
			return String(x.toFixed(2))+"K"
		}else {
			return String(x)
		}
	}

</script>


<button
on:click={()=>{goto(`/${Obj.GameID}`)}}
class=" w-[215px] min-h-fit hover:scale-125 max-h-fit mt-4 mb-4 pb-2 gap-1 hover:bg-[#202022] shadow-lg hover:shadow-gray-600    hover:cursor-pointer    transition-all ease-linear duration-100 text-slate-200 rounded-md flex flex-col">
	<img src={Obj.SmallPosterImage} alt="assassign" class="w-full object-cover rounded-t-lg top-0 h-[285px]">
	<p class=" text-sm text-start ml-2 inline-flex font-Poppins  ">{Obj.Rating}/10 <Star class="inline-flex h-4 w-4 stroke-yellow-200" /> <Dot class=" self-center mx-2"/> {RoundNumOfPeople(Obj.Players)} <Users class="ml-1 inline-flex h-4 w-4 " /> </p>
	<p class=" text-start text-base ml-1 font-Poppins line-clamp-2">{Obj.Name}</p>
	<p class=" font-thin ml-2 text-sm text-start font-Poppins ">${Obj.Price}</p>
</button>

<style>
	/* your styles go here */
</style>
