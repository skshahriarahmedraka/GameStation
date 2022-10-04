<script lang="ts">
	import Blog from '$lib/cards/midCard.txt';
	import Game from '$lib/logo/game.svelte';
	import Search from '$lib/logo/Search.svelte';
	import CartIcon from '$lib/svgs/cartIcon.svelte';
	import Coin from '$lib/svgs/coin2.svelte';
	import Profile from '$lib/svgs/profile.svelte';
	import World from '$lib/svgs/world.svelte';
	import { goto } from '$app/navigation';
	import List from '$lib/svgs/list.svelte';
	// import Gamestation from "$lib/logo/gamestation.svelte"
	// import GameStation from "$lib/logo/gamestation.svelte"
	// const game = new URL('../logo/game.svelte', import.meta.url).href;
	const gamestation = new URL('../logo/gamestation.png', import.meta.url).href;
</script>

<div class="flex h-[52px] w-full flex-row items-center  justify-center bg-gray-900">
	<button
		on:click={() => {
			goto('/');
		}}
		class=" flex h-16 w-fit flex-row items-center gap-2 "
	>
		<Game class=" h-16 w-fit " />
		<!-- <GameStation class=" h-28 w-fit " /> -->
		<!-- <img src="{game}" alt="game" class=" h-16 w-fit"> -->
		<img src={gamestation} alt="gamestation" class=" h-12 w-fit " />
	</button>

	<button
		on:click={() => {
			goto('/browse');
		}}
		class=" mx-3 font-Poppins text-base text-slate-300 hover:text-white"
	>
		Browse
	</button>
	<button
		on:click={() => {
			goto('/upcomming');
		}}
		class=" mx-3 font-Poppins text-base text-slate-300 hover:text-white"
	>
		Upcomming
	</button>
	<button
		on:click={() => {
			goto('/offers');
		}}
		class=" mx-3 font-Poppins text-base text-slate-300 hover:text-white"
	>
		Offers
	</button>
	<button
		on:click={() => {
			goto('/news');
		}}
		class=" mx-3 font-Poppins text-base text-slate-300 hover:text-white"
	>
		News
	</button>
	<!-- search bar  -->
	<div class="mx-2 flex justify-center  text-gray-200 font-Poppins">
		<div class=" xl:w-[600px]">
			<div class="  relative flex w-full flex-row items-stretch ">
				<!-- search input -->
				<input
					type="search"
					class=" form-control relative mx-1 block h-12 w-full min-w-0 flex-auto rounded-lg  border-2 border-solid  border-gray-700 bg-slate-600   px-3 py-1 text-base font-normal transition ease-in-out  focus:border-blue-600 focus:bg-gray-700 focus:outline-none"
					placeholder="Search"
					aria-label="Search"
					style=" background-color: rgb(71 85 105 / var(--tw-bg-opacity));"
					aria-describedby="button-addon2"
				/>
				<!-- search button -->
				<button
					class="  items-center rounded-lg border-2    border-sky-500 bg-opacity-30  px-6 py-1 text-xs font-medium text-white shadow-md   transition duration-150 ease-in-out hover:bg-blue-600  hover:bg-opacity-50 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:bg-opacity-30  active:shadow-lg"
					type="button"
					style=" 
					background-color: rgb(56 189 248 / var(--tw-bg-opacity));"
					id="button-addon2"
				>
					<Search class=" scale-110" />
				</button>
			</div>
		</div>
	</div>
	<!-- world -->
	<List
		on:click={() => {}}
		class="mr-2 h-7 w-7  cursor-pointer rounded-xl stroke-sky-400  stroke-[1px] hover:bg-sky-200 hover:bg-opacity-25  hover:stroke-sky-300 active:stroke-sky-500"
	/>
	<!-- <World on:click={()=>{}} class="fill-sky-400 hover:fill-sky-300  active:fill-sky-500 h-8 w-8  rounded-xl hover:bg-sky-200 hover:bg-opacity-25  mr-2 cursor-pointer"/> -->
	<div
		class=" mx-1 flex h-[60%] w-fit flex-row items-center gap-2  rounded-2xl border-0 border-blue-500 bg-sky-300 bg-opacity-30 pr-2"
	>
		<Coin class="h-8 w-8 storke-[1px] stroke-yellow-300" />
		<p class="text-slate-200 ">{826753}</p>
	</div>
	<div 	on:click={() => {
		goto('/carts');
	}}>
		<Profile
			class=" mr-2 h-6 w-6 scale-125   cursor-pointer   rounded-xl fill-sky-400  hover:bg-sky-200 hover:bg-opacity-25 hover:fill-sky-300  active:fill-sky-500"
		/>
	</div>
	<div 	on:click={() => {
		goto('/carts');
	}} class="">

		<CartIcon
			
			class=" h-6 w-6 scale-150 cursor-pointer  rounded-xl stroke-1  hover:bg-sky-200 hover:bg-opacity-25 hover:fill-sky-300 active:fill-sky-500 "
		/>
	</div>
</div>

<style>
	/* your styles go here */
</style>
