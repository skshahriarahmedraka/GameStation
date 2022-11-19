<script lang="ts">
	import Star from '$lib/svgs/star.svelte';
	import Users from '$lib/svgs/users.svelte';
	import Dot from '$lib/svgs/dot.svelte';
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

<style>
	/* your styles go here */
</style>
