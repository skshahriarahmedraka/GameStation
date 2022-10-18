<script lang="ts">
    // your script goes here

    export let data
    let {UserProData} = data
    console.log("🚀 ~ file: +page.svelte ~ line 6 ~ UserProData", UserProData)

    
	const mirageLogo = new URL('../../../lib/images/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.jpeg', import.meta.url).href;

</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->


<div class=" w-full h-fit flex flex-col gap-3   items-center  ">
    <!-- <p class=" text-white ">{printdata.profile}</p> -->
    <img src="{mirageLogo}" alt="banner logo" class=" object-cover h-[400px] rounded-lg w-[800px]  ">
     <!-- <div class=" h-[400px] w-[800px] ">
     </div> -->
     <div class=" h-[500px] aspect-[4/2] flex flex-row gap-2 justify-center ">
        <img src="{mirageLogo}" alt="banner logo" class=" h-24 w-44 rounded-md object-cover ">
        <div class="flex flex-col ">

            <p class=" text-white text-xl">Sk Shahriar Ahmed Raka</p>
            <p class=" text-white text-lg">@skshahriarraka</p>
        </div>
     </div>

</div>