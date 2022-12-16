<script lang="ts">
	import { goto } from "$app/navigation";
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
	}

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
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div on:click={()=>{
    goto(`/${Obj.GameID}`)
    
}} 
on:keypress={()=>{}}
    class="w-[256px] h-[106px] hover:cursor-pointer flex flex-row items-center bg-opacity-50  hover:bg-[#252525]  rounded-xl">
    <img src={Obj.SmallPosterImage} alt="syndicate" class="w-[60px] h-[80px] object-cover m-3 rounded-lg ">
    <div class="w-full h-full flex flex-col justify-center items-start">
        <p class=" text-lg text-slate-100 font-Poppins">{Obj.Name}</p>
        <p class=" text-sm text-slate-300 font-Poppins">${Obj.Price}</p>
    </div>
</div>