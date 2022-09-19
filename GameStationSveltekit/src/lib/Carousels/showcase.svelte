<script lang="ts">
	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
    import { Carousel, CarouselTransition } from 'flowbite-svelte'



    const images = [
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
let showThumbs=false
let showCaptions=false
</script>

<style>
    /* your styles go here */
</style>


  
  <div class="max-w-4xl">
    <CarouselTransition {images} loop transitionType="fade" transitionParams="{{ duration: 1000 }}" {showCaptions} {showThumbs} duration="3000" />
  </div>