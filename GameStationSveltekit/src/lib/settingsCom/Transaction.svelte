<script lang="ts">
	// your script goes here\
	import { UserProData } from '$lib/Store/store';
	// import Transaction from '$lib/svgs/transaction.svelte';

let TransactionGameList: {
		GameID: string;
		Name: string;
		Moto: string;
		LogoImage: string;
		BigPosterImage: string;
		SmallPosterImage: string;
		OtherImages: string[];
		Genres: string[];
		Feature: string[];
		Description: string;
		Rating: number;
		RatingGivenBy: {
			'PC Gamer': string;
			IGN: string;
			'Game Informer': string;
		};
		Minspec: {
			Name: string;
			Value: string;
		}[];
		Recomendedspec: {
			Name: string;
			Value: string;
		}[];
		Discount: number;
		Price: number;
		Developer: string;
		Publisher: string;
		Released: string;
		Platform: string[];
		Players: number;
		Comment: {
			Name: string;
			Rating: number;
			Description: string;
		}[];
		FollowUs: {
			Facebook: string;
			Discord: string;
			Youtube: string;
			Twitter: string;
			Site: string;
		};
	}[] = [] as {
		GameID: '';
		Name: '';
		Moto: '';
		LogoImage: '';
		BigPosterImage: '';
		SmallPosterImage: '';
		OtherImages: [];
		Genres: [''];
		Feature: [''];
		Description: '';
		FollowUs: {
			Facebook: '';
			Discord: '';
			Youtube: '';
			Twitter: '';
			Site: '';
		};
		Rating: 0;
		RatingGivenBy: {
			'PC Gamer': '';
			IGN: '';
			'Game Informer': '';
		};
		Minspec: [
			{
				Name: 'OS';
				Value: '';
			},
			{
				Name: 'Storage';
				Value: '';
			},
			{
				Name: 'Memory';
				Value: '';
			},
			{
				Name: 'CPU';
				Value: '';
			},
			{
				Name: 'GPU';
				Value: '';
			}
		];
		Recomendedspec: [
			{
				Name: 'OS';
				Value: '';
			},
			{
				Name: 'Storage';
				Value: '';
			},
			{
				Name: 'Memory';
				Value: '';
			},
			{
				Name: 'GPU';
				Value: '';
			},
			{
				Name: 'CPU';
				Value: '';
			}
		];
		Price: 0;
		Discount: 0;
		Developer: '';
		Publisher: '';
		Released: '';
		Platform: ['', '', ''];
		Players: 0;
		Comment: [];
	}[];
    async function GetTransactionGamedata(){
		let reqStruct =$UserProData.TransactionHistory
		


		await fetch(`/api/transactiondata`, {
			method: 'POST',
			mode: 'no-cors',
			body: JSON.stringify(reqStruct)
		})
			.then((res) => {
				return res.json();
			})
			.then((da) => {
				// console.log('🚀 ~ file: +page.server.ts ~ line 106 ~ myfetch ~ da', da);

				TransactionGameList = da;
			});
	}
    GetTransactionGamedata()
</script>

<!-- markup (zero or more items) goes here -->
<div class="transition-all duration-200 ease-linear flex flex-col ">
	<p class=" text-2xl m-5 self-center">Purchased Games </p>

	{#if TransactionGameList.length !=0  }
		<!-- content here -->
		{#each TransactionGameList as i}
			<!-- content here -->
			<div
				class=" flex h-56  w-full flex-row items-center rounded-md  bg-[#202020] pl-3 hover:bg-slate-800 sm:w-full md:w-full xs:w-full "
			>
				<img src={i.SmallPosterImage} alt={i.Name} class=" h-[200px] w-[150px] " />
    <div class=" ml-4 mr-2 flex h-full w-full flex-row">
        <div class="flex h-full w-fit flex-col">
            <p class=" mt-4 text-2xl">{i.Name}</p>
            <div class="flex flex-row gap-2">
                {#each i.Feature as j}
                    <p class=" text-sm font-thin text-blue-500">{j}</p>
                {/each}
            </div>
            <p class=" font-semibold">Rating {i.Rating}/10</p>
            <div class="grow" />
            <div class=" mb-4 flex flex-row gap-2">
                <!-- {#each i.Platform as j}
                    {#if j == 'mac'}
                        <Apple class="h-6 w-6 fill-slate-400" />
                    {:else if j == 'linux'}
                        <Linux class="h-6 w-6 fill-slate-400" />
                    {:else}
                        <Windows class="h-6 w-6 fill-slate-400" />
                    {/if}
                {/each} -->
            </div>
        </div>
        <div class="grow" />
        <div class=" mt-4 mb-3 flex h-full w-fit flex-col">
            <!-- <del class=" text-gray-400">${i.Price} </del> -->
            <p class="text-xl font-bold">${i.Price - (i.Price * i.Discount) / 100}</p>
            <p class="">Owned By You</p>
            <div class="grow" />
            <!-- <button
                on:click={() => {
                    RemoveFromCart(i.GameID);
                }}
                class=" mb-8 text-gray-500 underline hover:text-gray-200 ">Remove</button
            > -->
        </div>
    </div>
			</div>
		{/each}
	{:else}
		<!-- else content here -->
		<p class=" text-3xl text-slate-200">NO Transaction history</p>
	{/if}
</div>

<style>
	/* your styles go here */
</style>
