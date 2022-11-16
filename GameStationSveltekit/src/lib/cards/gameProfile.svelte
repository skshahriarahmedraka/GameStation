<script lang="ts">
	import Discord from '$lib/foot/icons/discord.svelte';
	import Apple from '$lib/svgs/apple.svelte';
	import Linux from '$lib/svgs/linux.svelte';
    import Windows from '$lib/svgs/windows.svelte';
	import Plus from '$lib/svgs/plus.svelte';
	import  {UserProData} from '$lib/Store/store';
	import { goto } from '$app/navigation';

	let GeneralData = {
		UserID: $UserProData.UserID,
		Name: $UserProData.Name,
		Accounttype: $UserProData.Accounttype,

		ProfileImg: $UserProData.ProfileImg,
		BannerImg: $UserProData.BannerImg,

		Email: $UserProData.Email,
		Mobile: $UserProData.Mobile,

		Address: $UserProData.Address,
		// Region: '',
		Country: $UserProData.Country,
		City: $UserProData.City,
		ZipCode: $UserProData.ZipCode,

		Coin: $UserProData.Coin,

		TransactionHistory: $UserProData.TransactionHistory || [],
		WishList: $UserProData.WishList || [],
		ContactAdminMsg: $UserProData.ContactAdminMsg || [],
		ContactDevMsg: $UserProData.ContactDevMsg || [],
		UserCart: $UserProData.UserCart || []
	};
	// const mirageLogo = new URL('../images/mirageLogo.png', import.meta.url).href;
	export let GameID: string;
	export let LogoImage: string = '';
	export let Price: number = 0;
	export let Discount: number = 25;
	export let Develper: string = '';
	export let Publisher: string = '';
	export let Platform: string[] = ['', '', ''];
	export let Released: string = '';
	async function UpdateProfileData() {
		console.log("🚀 ~ file: gameProfile.svelte ~ line 49 ~ UpdateProfileData ~ GeneralData", GeneralData)
		await fetch('/api/profile/general', {
			// credentials: 'same-origin',
			method: 'POST',
			// mode: 'cors',
			body: JSON.stringify(GeneralData)
		})
			.then((res) => {
				return res.json();
			})
			.then((resdata) => {
				console.log('🚀 ~ file: General.svelte ~ line 148 ~ .then ~ resdata', resdata);
				UserProData.update((n) => {
					return (n = resdata);
				});
				// GeneralData = resdata;
				console.log('🚀 ~ file: General.svelte ~ line 152 ~ .then ~ UserProData', $UserProData);
			});
	}
	function AddToCart(){
		if ($UserProData.UserCart ===null){
			UserProData.update((data)=>{
			data.UserCart=[GameID];
			return data;
		})
		}else {
			for (let i = 0; i < $UserProData.UserCart.length; i++) {
				if ($UserProData.UserCart[i] === GameID) {
					console.log(" Already addled game GameID", GameID)
					
					return;
				}
			}
			UserProData.update((data)=>{
				data.UserCart.push(GameID);
				return data;
			})
		}
		GeneralData.UserCart=$UserProData.UserCart;
		UpdateProfileData()
		// $UserProData.UserCart.push(GameID);
		console.log("🚀 ~ file: gameProfile.svelte ~ line 20 ~ AddToCart ~ UserProData", $UserProData)
	}
	function BuyNow(){
		if ($UserProData.UserCart ===null){
			UserProData.update((data)=>{
			data.UserCart=[GameID];
			return data;
		})
		}else {
			for (let i = 0; i < $UserProData.UserCart.length; i++) {
				if ($UserProData.UserCart[i] === GameID) {
					console.log(" Already addled game GameID", GameID)
					
					return;
				}
			}
			UserProData.update((data)=>{
				data.UserCart.push(GameID);
				return data;
			})
		}
		GeneralData.UserCart=$UserProData.UserCart;
		UpdateProfileData()
		// Simulate a mouse click:
// window.location.href = "/cart";
		goto("/carts")
		// $UserProData.UserCart.push(GameID);
		console.log("🚀 ~ file: gameProfile.svelte ~ line 20 ~ AddToCart ~ UserProData", $UserProData)
	}

</script>

<!-- markup (zero or more items) goes here -->

<div class="flex  h-fit w-80 flex-col gap-3  ">
	<img src={LogoImage} alt="" class=" h-fit w-full object-cover " />

	<div class=" ml-4 flex flex-row gap-2">
		<p class=" h-[25px] w-fit rounded-lg  bg-blue-500 px-4 py-1 text-center text-sm text-white">
			-{Discount}%
		</p>
		<del class=" text-slate-400">${Price}</del>
		<p class=" text-slate-100">${(Price - (Price * Discount) / 100).toFixed(2)}</p>
	</div>
	{#if $UserProData.TransactionHistory.includes(GameID)}
		 <!-- content here -->
		 <button
		 on:click={()=>{goto("/purchased")}}
			 class=" ml-4 h-[50px] w-[320px] rounded-lg bg-green-500 hover:bg-green-600 active:bg-green-700 font-Poppins text-lg font-semibold text-slate-800"
			 >Purchased</button
		 >
	{:else}
		 <!-- else content here -->
		 <button
		 on:click={()=>{BuyNow()}}
			 class=" ml-4 h-[50px] w-[320px] rounded-lg bg-sky-500 hover:bg-sky-600 font-Poppins text-lg font-semibold text-white"
			 >Buy Now</button
		 >
		 <button
			 on:click={()=>{AddToCart()}}
	 
			 class=" ml-4 h-[50px] w-[320px] rounded-lg border-2 border-white hover:bg-gray-600 hover:bg-opacity-70 hover:text-white font-Poppins text-lg font-semibold hover:border-slate-300  text-white "
			 >Add To Cart</button
		 >
	{/if}
	<button
		class=" ml-4 flex h-[30px] w-[320px] flex-row  items-center justify-center rounded-lg border-2 hover:bg-gray-600 hover:bg-opacity-70 hover:text-white border-white hover:border-slate-300 font-Poppins font-semibold  text-white"
	>
		<Plus class="h-4 w-4" />
		<p class="text-sm">Add To Wishlist</p></button
	>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-row justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Developer</p>
		<p class=" ">{Develper}</p>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-row justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Publisher</p>
		<p class=" ">{Publisher}</p>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-row justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Release Date</p>
		<p class=" ">{Released}</p>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-row justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Platform</p>
		<div class=" flex flex-row mb-4 gap-2">
            {#each Platform as j}
                 {#if j=="mac"}
                    <Apple class="h-5 w-5 fill-slate-100" />
                 {:else if  j=="linux"}
                    <Linux class="h-5 w-5 fill-slate-100" />
                 {:else if j=="windows"}
                    <Windows class="h-5 w-5 fill-slate-100" />
                 {/if}
            {/each}
        </div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
