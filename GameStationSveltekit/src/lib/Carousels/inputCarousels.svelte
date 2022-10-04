<script lang="ts">
	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
    import CameraPlus from '$lib/svgs/cameraPlus.svelte';

	export let OtherImages:string[] =["",""]
	const images = [
		{
			id: 0,
			name: 'Cosmic timetraveler',
			imgurl: anthem,
			attribution: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com'
		},
		{
			id: 1,
			name: 'Cristina Gottardi',
			imgurl: fifa,
			attribution: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com'
		},
		{
			id: 2,
			name: 'Johannes Plenio',
			imgurl: forbiddenwest,
			attribution: 'johannes-plenio-RwHv7LgeC7s-unsplash.com'
		},
		{
			id: 3,
			name: 'Jonatan Pie',
			imgurl: war,
			attribution: 'jonatan-pie-3l3RwQdHRHg-unsplash.com'
		}
	];
    let activeImage =images[0]
    let ActiveIndex:number=0
    function ChangeActiveImagePositive(){
        if (ActiveIndex===OtherImages.length) {
            ActiveIndex=0
        }
        else {
            ActiveIndex+=1
        }
    }
    function ChangeActiveImageNagative(){
        if (ActiveIndex===0) {
            ActiveIndex=OtherImages.length
        }
        else {
            ActiveIndex-=1
        }
    }
    // import { flip } from "svelte/animate";
import {elasticIn, elasticOut } from 'svelte/easing';

	import { fade, slide } from 'svelte/transition';

//     var intervalId = setInterval(function() {
//   ChangeActiveImagePositive()
// }, 4000);


let showImages:string[]

</script>

<div class="carousel aspect-[16/9] h-[600px]">
         <!-- content here -->
            
         <div  transition:slide="{{ duration: 2000, easing:elasticOut}}"   id="slide1" class="carousel-item relative w-full">
			{#if ActiveIndex === OtherImages.length}
				<div class="">

					<img    src="{OtherImages[ActiveIndex]}" alt="" class="w-full object-cover rounded-xl transition-all ease-linear duration-200" />
				</div>
			{:else}
				 <!-- <button class=" absolute   self-end   h-full w-full  " on:click={() => FileInput.click()}>
				
					<input
						bind:this={FileInput}
						on:change={(e) => onFileSelected(e)}
						type="file"
						class=" hidden h-10 w-10 "
						accept=".jpg, .jpeg, .png, .svg"
					/>
				</button> -->
			{/if}


             <div class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between">
                 <button on:click={()=>{ChangeActiveImageNagative()}} class="btn btn-circle">❮</button>
                 <button on:click={()=>{ChangeActiveImagePositive()}} class="btn btn-circle">❯</button>
             </div>
         </div>
	
</div>

<style>
	/* your styles go here */
</style>
