<!-- <script lang="ts">
    import {slide} from 'svelte/transition'
//     const gallery_items = [{
//     url: '<https://picsum.photos/id/237/400/300>',
//     description: 'Dog'
//   }, {
//     url: '<https://picsum.photos/id/238/400/300>',
//     description: 'Building'
//   },{
//     url: '<https://picsum.photos/id/240/400/300>',
//     description: 'Staircase'
//   }]
  const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
    import { Carousel, CarouselTransition } from 'flowbite-svelte'



    const gallery_items = [
  {
    id: 0,
    name: "Cosmic timetraveler",
    imgurl: anthem,
    attribution: "cosmic-timetraveler-pYyOZ8q7AII-unsplash.com",
  },
  {
    id: 1,
    name: "Cristina Gottardi",
    imgurl: fifa,
    attribution: "cristina-gottardi-CSpjU6hYo_0-unsplash.com",
  },
  {
    id: 2,
    name: "Johannes Plenio",
    imgurl: forbiddenwest,
    attribution: "johannes-plenio-RwHv7LgeC7s-unsplash.com",
  },
  {
    id: 3,
    name: "Jonatan Pie",
    imgurl: war,
    attribution: "jonatan-pie-3l3RwQdHRHg-unsplash.com",
  },
  
 
];

  let currentSlideItem = 0;
</script>

<style>
    /* your styles go here */
</style>



{#each [gallery_items[currentSlideItem]] as item (currentSlideItem)}
    <img transition:slide src={item.imgurl} alt={item.name} class="w-[1152px] h-[648px] object-cover rounded-xl transition-all ease-linear duration-200"/>
{/each} -->