<script lang="ts">
    //   import { flip } from "svelte/animate";
    // 	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	// const fifa = new URL('../images/fifa.png', import.meta.url).href;
	// const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	// const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
    // import { Carousel, CarouselTransition } from 'flowbite-svelte'


    export let liObj: {
    ID: string;
    Name: string;
    Image: string;
    Price: string;
    Rating: number;
    NumberOfPlayers: number;
}[]
    // const images = [anthem,forbiddenwest,fifa,war]
    let activeImage=0
    let ChangeImage=()=>{
        if (activeImage+1>liObj.length -1){
            activeImage=0
        }else{
            activeImage+=1
        }
    }

    let ActiveIndex:number=0

    function ChangeActiveImagePositive(){
        if (ActiveIndex+1>liObj.length-1) {
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


<div class="carousel aspect-[16/9] h-[600px]">
    <!-- content here -->

    <div  transition:slide="{{ duration: 1000, easing:elasticOut}}"   id="slide1" class="carousel-item relative w-full">
        <img    src="{liObj[ActiveIndex].Image}" alt="" class="w-full object-cover rounded-xl transition-all ease-linear duration-200" />
        <!-- <div class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between">
            <button on:click={()=>{ChangeActiveImageNagative()}} class="btn btn-circle">❮</button>
            <button on:click={()=>{ChangeActiveImagePositive()}} class="btn btn-circle">❯</button>
        </div> -->
    </div>

</div>

<!-- <div class="">
    <img  transition:fade src={liObj[activeImage].Image} alt="" class="max-w-[1152px] max-h-[648px] min-h-fit min-w-fit  object-cover rounded-xl transition-all ease-linear duration-200 ">
</div> -->