<script lang="ts">
    // your script goes here
import Nav from "$lib/navbar/nav.svelte";
import Footer from '$lib/foot/footer.svelte';

import navigationState from '$lib/Store/store';



</script>

<style>
    /* your styles go here */
</style>



<!-- markup (zero or more items) goes here -->
<div class=" min-h-screen w-full bg-[#121212] flex flex-col gap-5 justify-between overflow-hidden">
    <Nav />
   
    <slot></slot>
    
    <Footer />
</div>

