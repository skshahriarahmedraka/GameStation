<script>
    // your script goes here
import Nav from "$lib/navbar/nav.svelte";
import Footer from '$lib/foot/footer.svelte';


</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->
<div class=" min-h-screen min-w-full bg-[#121212] flex flex-col gap-5 justify-between">
    <Nav />
   
	
    <slot></slot>
    
    
    <Footer />
</div>