<script lang="ts">
	import { goto } from '$app/navigation';

    //   import { flip } from "svelte/animate";
    // 	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	// const fifa = new URL('../images/fifa.png', import.meta.url).href;
	// const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	// const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
    // import { Carousel, CarouselTransition } from 'flowbite-svelte'


//     export let liObj: {
//     ID: string;
//     Name: string;
//     Image: string;
//     Price: string;
//     Rating: number;
//     NumberOfPlayers: number;
//     Moto:string 
// }[]
export let CarouselList: {
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
    // const images = [anthem,forbiddenwest,fifa,war]
    let activeImage=0
    let ChangeImage=()=>{
        if (activeImage+1>CarouselList.length -1){
            activeImage=0
        }else{
            activeImage+=1
        }
    }

    let ActiveIndex:number=0

    function ChangeActiveImagePositive(){
        if (ActiveIndex+1>CarouselList.length-1) {
            ActiveIndex=0
        }
        else {
            ActiveIndex+=1
        }
    }

var intervalId = setInterval(function() {
  ChangeImage()
}, 3000);



import {elasticIn, elasticOut } from 'svelte/easing';

	import { fade, slide } from 'svelte/transition';

    var intervalId = setInterval(function() {
  ChangeActiveImagePositive()
}, 4000);
</script>

<style>
    /* your styles go here */
</style>

{#if CarouselList.length>0}


<div class="carousel aspect-[16/9] h-[600px]">
    <!-- content here -->
	
    <div on:click={()=>{
		goto(`/${CarouselList[ActiveIndex].GameID}`)
        
    }} 
    on:keypress={()=>{}}
    transition:slide="{{ duration: 1000, easing:elasticOut}}"   id="slide1" class="hover:cursor-pointer  hover:opacity-90 carousel-item  w-full relative">
        <img    src="{CarouselList[ActiveIndex].BigPosterImage}" alt="" class="w-full object-cover rounded-xl transition-all ease-linear duration-200" />
        <!-- <div class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between">
            <button on:click={()=>{ChangeActiveImageNagative()}} class="btn btn-circle">❮</button>
            <button on:click={()=>{ChangeActiveImagePositive()}} class="btn btn-circle">❯</button>
        </div> -->
        <div class=" absolute left-10 bottom-10 flex flex-col gap-3">
			<p class="text-xl w-[40%] line-clamp-3 text-white v">{CarouselList[ActiveIndex].Moto}</p>
            <div class="flex flex-row gap-3">
				<button class=" w-fit px-4 py-2 object-center hover:font-semibold bg-slate-200 text-gray-900 font-Poppins rounded-lg text-xl">Buy Now</button>
                <button class="w-fit px-4 py-2 object-center bg-transparent hover:font-semibold backdrop-blur-lg text-white border-2 border-transparent hover:border-2 hover:border-white font-Poppins rounded-lg text-xl">Add to WishList</button>
            </div>
        </div>
		
    </div>
	
</div>
{:else}


<div class="carousel aspect-[16/9] h-[600px]">
    <!-- content here -->
	
    <div 

    transition:slide="{{ duration: 1000, easing:elasticOut}}"   id="slide1" class=" rounded-xl  bg-[#252525]  hover:opacity-90 carousel-item  w-full relative">
        <!-- <img    src="" alt="" class="w-full object-cover rounded-xl transition-all ease-linear duration-200" /> -->
        <!-- <div class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between">
            <button on:click={()=>{ChangeActiveImageNagative()}} class="btn btn-circle">❮</button>
            <button on:click={()=>{ChangeActiveImagePositive()}} class="btn btn-circle">❯</button>
        </div> -->
        <div class=" absolute left-10 bottom-10 flex flex-col gap-3">
			<!-- <p class="text-xl w-[40%] line-clamp-3 text-white v">{""}</p> -->
            <div class="flex flex-row gap-3">
				<button class=" w-fit px-4 py-2 object-center hover:font-semibold bg-slate-200 text-gray-900 font-Poppins rounded-lg text-xl">Buy Now</button>
                <button class="w-fit px-4 py-2 object-center bg-transparent hover:font-semibold backdrop-blur-lg text-white border-2 border-transparent hover:border-2 hover:border-white font-Poppins rounded-lg text-xl">Add to WishList</button>
            </div>
        </div>
		
    </div>
	
</div>

{/if}

<!-- <div class="">
    <img  transition:fade src={liObj[activeImage].Image} alt="" class="max-w-[1152px] max-h-[648px] min-h-fit min-w-fit  object-cover rounded-xl transition-all ease-linear duration-200 ">
</div> -->