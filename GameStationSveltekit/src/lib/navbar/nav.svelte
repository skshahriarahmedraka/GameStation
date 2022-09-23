<script lang="ts">
	import Blog from '$lib/cards/midCard.svelte';
	import Game from '$lib/logo/game.svelte';
	import Search from '$lib/logo/Search.svelte';
	import CartIcon from '$lib/svgs/cartIcon.svelte';
	import Coin from '$lib/svgs/coin.svelte';
	import Profile from '$lib/svgs/profile.svelte';
	import World from '$lib/svgs/world.svelte';
	import {goto} from "$app/navigation"
	import List from '$lib/svgs/list.svelte';
	// import Gamestation from "$lib/logo/gamestation.svelte"
	// import GameStation from "$lib/logo/gamestation.svelte"
	// const game = new URL('../logo/game.svelte', import.meta.url).href;
	const gamestation = new URL('../logo/gamestation.png', import.meta.url).href;
</script>

<div class="flex h-[52px] w-full flex-row items-center  justify-center bg-gray-900">
	<button on:click="{()=>{goto("/")}}" class=" flex h-16 w-fit flex-row gap-2 items-center ">
		<Game class=" h-16 w-fit " />
		<!-- <GameStation class=" h-28 w-fit " /> -->
		<!-- <img src="{game}" alt="game" class=" h-16 w-fit"> -->
		<img src={gamestation} alt="gamestation" class=" h-12 w-fit " />
	</button>

	<button on:click="{()=>{goto("/browse")}}" class=" text-base text-slate-300 hover:text-white font-Poppins mx-3">
		Browse
	</button>
	<button on:click="{()=>{goto("/upcomming")}}" class=" text-base text-slate-300 hover:text-white font-Poppins mx-3">
		Upcomming
    </button>
	<button on:click="{()=>{goto("/offers")}}" class=" text-base text-slate-300 hover:text-white font-Poppins mx-3">
		Offers
	</button>
	<button on:click="{()=>{goto("/news")}}" class=" text-base text-slate-300 hover:text-white font-Poppins mx-3">
		News
	</button>
	<!-- search bar  -->
	<div class="mx-2 flex justify-center  text-gray-200">
		<div class=" xl:w-[600px]">
			<div class=" bg-transparent relative flex w-full flex-row items-stretch ">
				<!-- search input -->
				<input
					type="search"
					class="form-control bg-transparent  relative mx-1 block h-12 w-full min-w-0 flex-auto rounded-lg  border-2 border-solid  border-gray-700 bg-slate-600 bg-opacity-100  px-3 py-1 text-base font-normal transition ease-in-out  focus:border-blue-600 focus:bg-gray-700 focus:outline-none"
					placeholder="Search"
					aria-label="Search"
					aria-describedby="button-addon2"
				/>
				<!-- search button -->
				<button
					class="  bg-transparent bg-opacity-100 bg-black  items-center rounded-lg  border-2 border-sky-500 px-6 py-1 text-xs font-medium   text-white shadow-md transition duration-150  ease-in-out hover:bg-blue-600 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800  active:shadow-lg"
					type="button"
					id="button-addon2"
				>
					<Search class=" scale-110"/>
				</button>
			</div>
		</div>
	</div>
	<!-- world -->
	<List on:click={()=>{}} class="stroke-[1px] stroke-sky-400 hover:stroke-sky-300  active:stroke-sky-500 h-7 w-7  rounded-xl hover:bg-sky-200 hover:bg-opacity-25  mr-2 cursor-pointer"/>
	<!-- <World on:click={()=>{}} class="fill-sky-400 hover:fill-sky-300  active:fill-sky-500 h-8 w-8  rounded-xl hover:bg-sky-200 hover:bg-opacity-25  mr-2 cursor-pointer"/> -->
	<div class=" h-[60%] w-fit flex flex-row gap-2 rounded-2xl mx-1  border-blue-500 border-0 bg-sky-300 bg-opacity-30 items-center pr-2">
		<Coin class="h-8 w-8" />
		<p class="text-slate-200 ">{826753}</p>
	</div>
		<Profile  on:click={()=>{}} class=" fill-sky-400 h-6 w-6 scale-125   mr-2   hover:fill-sky-300 active:fill-sky-500  rounded-xl hover:bg-sky-200 hover:bg-opacity-25  cursor-pointer" />
	<CartIcon   on:click={()=>{}} class=" h-6 w-6 stroke-1 scale-150  hover:fill-sky-300 active:fill-sky-500  rounded-xl hover:bg-sky-200 hover:bg-opacity-25 cursor-pointer " />
</div>

<style>
	/* your styles go here */
</style>
