<script lang="ts">
	import { goto } from '$app/navigation';
	export let Obj: {
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
	} = {} as {
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
	};

	// your script goes here
	// const syndicate = new URL('../smallImages/syndicate.jpeg', import.meta.url).href;
	// const mirage = new URL('../smallImages/mirage.jpg', import.meta.url).href;

	//     export let Obj: {
	//     ID: string;
	//     Name: string;
	//     Image: string;
	//     Price: string;
	//     Rating: number;
	//     NumberOfPlayers: number;
	// }
// console.log("🚀 ~ file: smallcard.svelte:142 ~ Obj", Obj)
</script>

<!-- markup (zero or more items) goes here -->

{#if JSON.stringify(Obj) != '{}' }
	<div
		on:click={() => {
			goto(`/${Obj.GameID}`);
		}}
		on:keypress={() => {}}
		class="flex h-[106px] w-[256px] flex-row items-center rounded-xl bg-opacity-50  ?  hover:cursor-pointer hover:bg-[#252525]"
	>
		<img
			src={Obj.SmallPosterImage}
			alt="syndicate"
			class="m-3 h-[80px] w-[60px] rounded-lg object-cover "
		/>
		<div class="flex h-full w-full flex-col items-start justify-center">
			<p class=" font-Poppins text-lg text-slate-100">{Obj.Name}</p>
			<p class=" font-Poppins text-sm text-slate-300">${Obj.Price}</p>
		</div>
	</div>
{:else}
	<div
		class="flex h-[106px] w-[256px] flex-row items-center rounded-xl bg-[#252525] bg-opacity-50 hover:cursor-pointer  hover:bg-[#252525]"
	>
		<div class="m-3 h-[80px] w-[60px] rounded-lg bg-[#252525] object-cover ">
			<!--  -->
		</div>
		<div class="flex h-full w-full flex-col items-start justify-center">
			<!-- <p class=" text-lg text-slate-100 font-Poppins">{""}</p> -->
			<!-- <p class=" text-sm text-slate-300 font-Poppins">${""}</p> -->
		</div>
	</div>
{/if}

<style>
	/* your styles go here */
</style>
