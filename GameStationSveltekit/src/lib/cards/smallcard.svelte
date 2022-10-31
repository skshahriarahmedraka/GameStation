<script lang="ts">
	import { goto } from "$app/navigation";


    // your script goes here
	// const syndicate = new URL('../smallImages/syndicate.jpeg', import.meta.url).href;
	// const mirage = new URL('../smallImages/mirage.jpg', import.meta.url).href;

    export let Obj: {
    ID: string;
    Name: string;
    Image: string;
    Price: string;
    Rating: number;
    NumberOfPlayers: number;
}
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div on:click={()=>{goto("/")}}
    on:K
    class="w-[256px] h-[106px] hover:cursor-pointer flex flex-row items-center bg-opacity-50  hover:bg-[#252525]  rounded-xl">
    <img src={Obj.Image} alt="syndicate" class="w-[60px] h-[80px] object-cover m-3 rounded-lg ">
    <div class="w-full h-full flex flex-col justify-center items-start">
        <p class=" text-lg text-slate-100 font-Poppins">{Obj.Name}</p>
        <p class=" text-sm text-slate-300 font-Poppins">${Obj.Price}</p>
    </div>
</div>