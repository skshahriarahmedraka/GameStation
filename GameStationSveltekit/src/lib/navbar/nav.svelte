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
	import ProfileList from '$lib/svgs/profileList.svelte';
	import { UserProData ,UserSettingSelect } from '$lib/Store/store';
	import Notification from '$lib/svgs/notification.svelte';
	// import Gamestation from "$lib/logo/gamestation.svelte"
	// import GameStation from "$lib/logo/gamestation.svelte"
	// const game = new URL('../logo/game.svelte', import.meta.url).href;
	// let myval: any;
	// UserProData.subscribe((val) => {
	// 	myval = val;
	// });
	// console.log('🚀 ~ file: nav.svelte ~ line 19 ~ myval', myval);
	const gamestation = new URL('../logo/gamestation.png', import.meta.url).href;

	let ShowSearchBar = false;
	let ShowProfile = false;
	// console.log('myval.UserID', myval.UserID);
</script>

<div
	class="flex h-[52px] w-full flex-row items-center justify-center bg-gray-900 sm:justify-around sm:gap-2 md:gap-4 lg:gap-4 xl:gap-4 xs:justify-around  xs:gap-1  "
>
	{#if !ShowSearchBar}
		<button
			on:click={() => {
				goto('/');
			}}
			class=" flex h-16 w-fit flex-row items-center gap-2 "
		>
			<!-- content here -->
			<Game class=" h-16 w-fit sm:hidden xs:hidden  xs:h-10  " />

			<!-- <GameStation class=" h-28 w-fit " /> -->
			<!-- <img src="{game}" alt="game" class=" h-16 w-fit"> -->
			<img src={gamestation} alt="gamestation" class=" h-12 w-fit xs:h-8 " />
		</button>

		<button
			on:click={() => {
				goto('/browse');
			}}
			class="  font-Poppins text-base text-slate-300  hover:text-white sm:text-sm xs:text-xs"
		>
			Browse
		</button>
		<button
			on:click={() => {
				goto('/upcomming');
			}}
			class="  font-Poppins text-base text-slate-300 hover:text-white sm:text-sm xs:text-xs"
		>
			Upcomming
		</button>
		<button
			on:click={() => {
				goto('/offers');
			}}
			class="  font-Poppins text-base text-slate-300 hover:text-white sm:text-sm xs:text-xs"
		>
			Offers
		</button>
		<button
			on:click={() => {
				goto('/news');
			}}
			class="  font-Poppins text-base text-slate-300 hover:text-white sm:text-sm xs:text-xs"
		>
			News
		</button>
	{/if}

	<!-- search bar  -->
	<div class=" flex justify-center{ ShowSearchBar ? "w-full" : "" }  font-Poppins text-gray-200">
		<div class=" xl:w-[600px]">
			<div class="  relative flex w-full flex-row items-center">
				<!-- search input -->

				<!-- content here -->
				<input
					type="search"
					class=" form-control relative mx-1 block h-12 w-full min-w-0 flex-auto rounded-lg border-2  border-solid border-gray-700  bg-slate-600 px-3   py-1 text-base font-normal transition ease-in-out focus:border-blue-600  focus:bg-gray-700 focus:outline-none sm:hidden xs:hidden"
					placeholder="Search"
					aria-label="Search"
					style=" background-color: rgb(71 85 105 / var(--tw-bg-opacity));"
					aria-describedby="button-addon2"
				/>
				{#if ShowSearchBar}
					<input
						type="search"
						class=" form-control  relative mx-1 block h-12 w-full min-w-0 flex-auto rounded-lg  border-2 border-solid  border-gray-700 bg-slate-600   px-3 py-1 text-base font-normal transition ease-in-out  focus:border-blue-600 focus:bg-gray-700 focus:outline-none"
						placeholder="Search"
						aria-label="Search"
						style=" background-color: rgb(71 85 105 / var(--tw-bg-opacity));"
						aria-describedby="button-addon2"
					/>
					<svg
						on:click={() => {
							ShowSearchBar = false;
						}}
						on:keydown={() => {}}
						class="h-8 w-8  cursor-pointer"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
						xmlns="http://www.w3.org/2000/svg"
						><path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M6 18L18 6M6 6l12 12"
						/></svg
					>
				{:else}
					<!-- else content here -->
				{/if}

				<button
					on:click={() => {
						// if ( window.screen.width < 897){
						// 	ShowSearchBar = true;

						// }
						if (window.matchMedia('screen and (max-width: 897px)').matches) {
							ShowSearchBar = true;
							
						}
						// console.log("🚀 ~ file: nav.svelte ~ line 132 ~ window.matchMedia('screen and (max-width: 897px)').matches", window.matchMedia('screen and (max-width: 897px)').matches)
						// console.log("🚀 ~ file: nav.svelte ~ line 132 ~ screen.width", window.outerWidth)
					}}
					class="  items-center rounded-lg border-2   border-sky-500 bg-opacity-30  px-6 py-3  text-xs font-medium text-white shadow-md transition duration-150   ease-in-out hover:bg-blue-600 hover:bg-opacity-50 hover:shadow-lg  focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:bg-opacity-30 active:shadow-lg xs:px-1  xs:py-2"
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
	<div 
		on:click={() => { 
		UserSettingSelect.update((n)=>n="Wishlist");
		goto("/profile/settings")


	 }}
	 on:keypress={()=>{}}
	class="sm:hidden md:hidden lg:hidden xl:hidden xs:hidden xxl:block">
		<List
			class=" h-7 w-7 cursor-pointer rounded-xl stroke-sky-400  stroke-[1px] hover:bg-sky-200 hover:bg-opacity-25  hover:stroke-sky-300 active:stroke-sky-500"
		/>
	</div>
	<!-- <World on:click={()=>{}} class="fill-sky-400 hover:fill-sky-300  active:fill-sky-500 h-8 w-8  rounded-xl hover:bg-sky-200 hover:bg-opacity-25  mr-2 cursor-pointer"/> -->
	<div
	on:click={() => { 
		UserSettingSelect.update((n)=>n="Payment Mangement");

		goto("/profile/settings")
		

	 }}
	 on:keypress={()=>{}}
		class="  flex flex-row gap-2 justify-center items-center bg-blue-500 bg-opacity-30 hover:bg-opacity-40 hover:cursor-pointer p-2 rounded-2xl  sm:hidden md:hidden lg:hidden xl:flex xs:hidden xxl:flex "
	>
		<Coin class="storke-[1px] h-8 w-8 stroke-yellow-300" />
		<p class="text-slate-200 ">{$UserProData.Coin}</p>
	</div>
	<div
		on:click={() => {
			goto(`/profile/${$UserProData.UserID}`);
		}}
		on:keydown={() => {}}
		class="sm:hidden md:hidden lg:hidden xl:hidden xs:hidden xxl:contents"
	>
		{#if $UserProData.ProfileImg!= ""}
			 <!-- content here -->
			<img src={$UserProData.ProfileImg} alt="U" 
			class="  h-6 w-6 scale-125   cursor-pointer object-cover rounded-xl"
			>
		{:else}
			 <Profile
				 class="  h-6 w-6 scale-125   cursor-pointer   rounded-xl fill-sky-400  hover:bg-sky-200 hover:bg-opacity-25 hover:fill-sky-300  active:fill-sky-500"
			 />
		{/if}
	</div>
	<div 
		on:click={() => { 
		UserSettingSelect.update((n)=>n="Notifications");

		goto("/profile/settings")
		

	 }}
	 on:keypress={()=>{}}
	class="sm:hidden md:hidden lg:hidden xl:hidden xs:hidden xxl:block">
	 	<Notification 
		 class=" h-7 w-7 cursor-pointer rounded-xl stroke-sky-400  stroke-[1px] hover:bg-sky-200 hover:bg-opacity-25  hover:stroke-sky-300 active:stroke-sky-500"
		/>
		
	</div>
	<div
		on:click={() => {
			goto('/carts');
		}}
		on:keydown={() => {}}
		class="sm:hidden md:hidden lg:hidden xl:hidden xs:hidden xxl:contents  "
	>
		<CartIcon
			class=" h-6 w-6 scale-150 cursor-pointer  rounded-xl stroke-1  hover:bg-sky-200 hover:bg-opacity-25 hover:fill-sky-300 active:fill-sky-500 "
		/>
		{#if $UserProData.UserCart.length > 0}
					<div
						class=" -ml-4 rounded-full border-2 border-[#202225] bg-red-500  px-1 text-xs font-semibold text-white "
					>
						{$UserProData.UserCart.length}
					</div>
				{/if}
	</div>

	<div class=" xl:contents xs:h-6  xs:w-6 xxl:hidden ">
		<ProfileList class=" h-9 w-9 xs:h-6 xs:w-6 " />
	</div>
</div>

<style>
	/* your styles go here */
</style>
