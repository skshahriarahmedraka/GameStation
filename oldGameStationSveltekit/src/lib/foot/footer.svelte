<script lang="ts">
	import Gamestation from '$lib/logo/gamestation.svelte';
	import Discord from './icons/discord.svelte';
	import {goto} from "$app/navigation"
	import {UserProData} from "$lib/Store/store"

	// your script goes here

	import Fb from './icons/fb.svelte';
	import Twitter from './icons/twitter.svelte';
	import Game from '$lib/logo/game.svelte';
	const gamestation = new URL('../logo/gamestation.png', import.meta.url).href;

	import Youtube from './icons/youtube.svelte';
</script>

<div class=" h-96 md:px-20 lg:px-20 xl:px-20 w-full sm:px-5 xs:px-2 bg-[#202020] ">
	<!-- <h1>developed by Sk Shahriar ahmed raka</h1> -->
	<!-- social medias -->
	<div class="flex h-16 w-full flex-row gap-3 md:ml-16 flex-wrap  ">
		<!--  -->
		<button  on:click={()=>{window.open("https://www.facebook.com/dhrubo.nil.180 ",'_blank')?.focus()}} class="">
			<Fb class="h-7 w-7 fill-slate-200" />
		</button>
		<button on:click={()=>{window.open("https://www.youtube.com/c/PlayStation",'_blank')?.focus()}} class="">
			<Youtube class="h-8 w-8 fill-slate-200" />
		</button>
		<button on:click={()=>{window.open("https://discord.com/channels/@me/1010975553039380630",'_blank')?.focus()}} class="">
			<Discord class="h-8 w-8 fill-slate-200" />
		</button>
		<button on:click={()=>{window.open("https://twitter.com/sveltejs",'_blank')?.focus()}} class="">
			<Twitter class="h-8 w-8 fill-slate-200" />
		</button>
	</div>
	<div class="flex flex-row gap-4 text-sm md:ml-16 text-gray-300 flex-wrap ">
		<!--  -->
		<!-- resources -->
		<div class="flex flex-col gap-2 flex-wrap">
			<p class=" font-semibold text-gray-600">Resources</p>
			<div class="flex flex-row gap-5 flex-wrap">
				<div class=" flex flex-col">
					<p class="" />
					<p>Support-A-Creator</p>
					<p>Distribute on Game Station</p>
					<p>Careers</p>
					{#if $UserProData.Accounttype==="admin"}
					<!--  -->
					
					<button class="hover:font-bold underline hover:cursor-pointer " on:click="{()=>{goto('/addgame')}}">Add Game</button>
					{/if}
				</div>
				<div class=" flex flex-col">
					<p>Fan Art Policy</p>
					<p>UX Research</p>
					<p>Store EULA</p>
					{#if $UserProData.Accounttype==="admin"}
					<!--  -->
					
					<button class="hover:font-bold underline hover:cursor-pointer " on:click="{()=>{goto('/news/addnews')}}">Add News</button>
					{/if}
				</div>
				<div class="flex flex-col">
					<p>Online Services</p>
					<p>Community Rules</p>
					<p>Epic Newsroom</p>
					{#if $UserProData.Accounttype==="admin"}
					<!--  -->
					<button class="hover:font-bold underline hover:cursor-pointer " on:click="{()=>{goto('/userrequest')}}">User Request</button>
					{/if}
				</div>
			</div>
		</div>
		<div class=" flex flex-col gap-2">
			<p class=" font-semibold text-gray-600">Made By Game Station</p>
			<div class="flex flex-row gap-5">
				<div class="flex flex-col">
					<p>Battle Breakers</p>
					<p>Fortnite</p>
					<p>Infinity Blade</p>
				</div>

				<div class=" flex flex-col">
					<p>Robo Recall</p>
					<p>Shadow Complex</p>
					<p>Unreal Tournament</p>
				</div>
			</div>
		</div>
	</div>

	<!-- @ -->
	<div class="mt-5 md:ml-16">
		<div class="h-[1px] w-full bg-slate-400 opacity-25" />
		<!--  -->
		<p class=" mt-5 text-xs md:mr-24  text-gray-300">
			© 2022, Game Station Games, Inc. All rights reserved. Epic, Epic Games, the Epic Games logo, Fortnite,
			the Fortnite logo, Unreal, Unreal Engine, the Unreal Engine logo, Unreal Tournament, and the
			Unreal Tournament logo are trademarks or registered trademarks of Epic Games, Inc. in the
			United States of America and elsewhere. Other brands or product names are the trademarks of
			their respective owners. Non-US transactions through Epic Games International, S.à r.l.
		</p>
	</div>
	<!-- terms and condition -->
	<div class="mt-5  flex w-[95%] flex-row gap-5 text-sm md:ml-16 text-gray-300">
		<!--  -->
		<button class="w-fit">Terms of Service</button>
		 <button class="">Privacy Policy</button>
		<button class="">Store Refund Policy</button>
		<div class=" grow" />

		<button on:click="{()=>{goto("/")}}" class=" flex h-16 w-fit flex-row gap-2 items-center ">
			<Game class=" h-16 w-fit " />
			<!-- <GameStation class=" h-28 w-fit " /> -->
			<!-- <img src="{game}" alt="game" class=" h-16 w-fit"> -->
			<img src={gamestation} alt="gamestation" class=" h-12 w-fit " />
		</button>
	
	</div>
</div>
