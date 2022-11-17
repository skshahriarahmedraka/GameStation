<script lang="ts">
	  import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';

	export let data: PageData;
	
	let { Userdata } = data;
	console.log("🚀 ~ file: +page.svelte ~ line 10 ~ Userdata", Userdata)
	UserProData.update((d) => (d = Userdata));

	import { goto } from '$app/navigation';

	import { UserSettingSelect } from '$lib/Store/store';
	import CameraPlus from '$lib/svgs/cameraPlus.svelte';
	import Profile from '$lib/svgs/profile.svelte';
	import WalpaperLogo from '$lib/svgs/walpaperLogo.svelte';
	// your script goes here

	// export let data
	// let {UserProData} = data
	// console.log("🚀 ~ file: +page.svelte ~ line 6 ~ UserProData", UserProData)

	const mirageLogo = new URL(
		'../../../lib/images/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.jpeg',
		import.meta.url
	).href;
</script>

<!-- markup (zero or more items) goes here -->

<div class=" flex min-h-fit w-full   flex-col items-center   gap-3  ">
	<!-- <p class=" text-white ">{printdata.profile}</p> -->
	{#if $UserProData.BannerImg.trim() != ''}
		<!-- content here -->
		<!-- src={mirageLogo} -->
		<img
			src={$UserProData.BannerImg}
			alt="banner logo"
			class=" aspect-[16/9] h-[500px] sm:w-full sm:h-fit sm:m-4 xs:w-full xs:h-fit xs:m-4  rounded-lg object-cover "
		/>
	{:else}
		<div
			class=" flex aspect-[16/9]  h-[500px] sm:w-full sm:h-fit sm:m-4 xs:w-full xs:h-fit xs:m-4 items-center justify-center rounded-lg border-2   border-slate-200 stroke-slate-200 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
		>
			<WalpaperLogo
				class="aspect-[16/9] h-[300px]  stroke-slate-200  stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
			/>
		</div>
		<!-- <CameraPlus
			class="  h-56 w-56 rounded-xl border-2 border-slate-200 stroke-slate-200 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
		/> -->
	{/if}

	<div class=" flex  w-[70%] min-h-fit flex-row sm:flex-col xs:flex-col justify-center gap-8 ">
		{#if $UserProData.ProfileImg.trim() != ''}
			<img
				src={$UserProData.ProfileImg}
				alt="banner logo"
				class=" h-56 w-56 rounded-full object-cover "
			/>
		{:else}
			<Profile
				class="  h-56 w-56 rounded-xl border-2 border-slate-200 stroke-slate-200 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
			/>
		{/if}
		<div class="flex flex-col ">
			<div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white">Name :</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
					{$UserProData.Name}
				</div>
			</div>

            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white">UserID :</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
                    @{$UserProData.UserID}
				</div>
			</div>
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> Email:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
		            {$UserProData.Email}
				</div>
			</div>
			
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> Password:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
		            ***************
				</div>
			</div>
			
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> Mobile:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
		            {$UserProData.Mobile ? $UserProData.Mobile : "Not Given" }
				</div>
			</div>
			
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white">Address :</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
                    {$UserProData.Address ? $UserProData.Address : "Not Given"}
				</div>
			</div>
			
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> City:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>{$UserProData.City ? $UserProData.City : "Not Given"}
				</div>
			</div>
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white">Country :</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
                {$UserProData.Country ? $UserProData.Country : "Not Given"}
				</div>
			</div>

            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> ZipCode:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
                {$UserProData.ZipCode ? $UserProData.ZipCode : "Not Given" }
				</div>
			</div>

            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> Coin:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
                {$UserProData.Coin ? $UserProData.Coin  : "0" }
				</div>
			</div>

            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> Account Type:</p>
				<div
					class=" mx-4 my-2 h-12 w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
                {$UserProData.Accounttype ? $UserProData.Accounttype : "Normal"}
				</div>
			</div>

            <div class=" flex flex-row gap-2 w-full items-center ">
				<p class=" text-xl text-white"> Transaction History:</p>
				<div
					class=" mx-4 my-2 h-fit w-fit self-center flex flex-col rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
				{#if $UserProData.TransactionHistory!=null}
					 <!-- content here -->
					 {#each $UserProData.TransactionHistory as i}
					 <p class="">
					 {i}
					 </p>
					 {/each}
				{:else}
					 <!-- else content here -->
					 <p class="">
						 No History
					 </p>
				{/if}
				</div>
			</div>

            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> Contract Admin message:</p>
				<div
					class=" mx-4 my-2 h-fit w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
				{#if $UserProData.ContactAdminMsg!=null}
					 <!-- content here -->
					 {#each $UserProData.ContactAdminMsg as i}
					 <p class="">
					 {i}
					 </p>
					 {/each}
				{:else}
					 <!-- else content here -->
					 <p class="">
						 No History
					 </p>
				{/if}
				</div>
			</div>
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> WishList:</p>
				<div
					class=" mx-4 my-2 h-fit w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
				{#if $UserProData.WishList!=null}
					 <!-- content here -->
					 {#each $UserProData.WishList as i}
					 <p class="">
					 {i}
					 </p>
					 {/each}
				{:else}
					 <!-- else content here -->
					 <p class="">
						 No wish List
					 </p>
				{/if}
                <!-- {$UserProData.WishList ? $UserProData.WishList : "" } -->
		            
				</div>
			</div>
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white">Contract Developer Message  :</p>
				<div
					class=" mx-4 my-2 h-fit w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
				{#if $UserProData.ContactDevMsg!=null}
					 <!-- content here -->
					 {#each $UserProData.ContactDevMsg as i}
					 <p class="">
					 {i}
					 </p>
					 {/each}
				{:else}
					 <!-- else content here -->
					 <p class="">
						 No History
					 </p>
				{/if}
                <!-- {$UserProData.ContactDevMsg ? $UserProData.ContactDevMsg : "No History"} -->
				</div>
			</div>
            <div class=" flex flex-row gap-2 w-full items-center">
				<p class=" text-xl text-white"> User Cart:</p>
				<div
					class=" mx-4 my-2 h-fit w-fit self-center rounded-lg border-0 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				>
				{#if $UserProData.UserCart!=null}
					 <!-- content here -->
					 {#each $UserProData.UserCart as i}
					 <p class="">
					 {i}
					 </p>
					 {/each}
				{:else}
					 <!-- else content here -->
					 <p class="">
						 No Product In Cart"
					 </p>
				{/if}
                <!-- {$UserProData.UserCart ? $UserProData.UserCart : " } -->
				</div>
			</div>
			
			
			
            <div class=" self-end" >
                <button
                    on:click={() => {
		UserSettingSelect.update((n)=>n="General");

                        goto('/profile/settings');
                    }}
                    class="h-12 w-fit rounded-xl bg-blue-600 px-6 font-Poppins font-semibold text-white hover:bg-sky-500 active:bg-blue-800  "
                >
                    Go to Settings
                </button>
            </div>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
