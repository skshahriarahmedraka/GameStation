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


var intervalId = setInterval(function() {
  ChangeImage()
}, 3000);

import {fade} from 'svelte/transition'


</script>

<style>
    /* your styles go here */
</style>




<div class="">
    <img  transition:fade src={liObj[activeImage].Image} alt="" class="w-[1152px] h-[648px] object-cover rounded-xl transition-all ease-linear duration-200 ">
</div>