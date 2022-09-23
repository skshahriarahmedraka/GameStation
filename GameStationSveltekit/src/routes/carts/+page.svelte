<script lang="ts">
	import Apple from "$lib/svgs/apple.svelte";
	import Linux from "$lib/svgs/linux.svelte";
	import Windows from "$lib/svgs/windows.svelte";

	let li: {
    name: string;
    midCoverImage: string;
    feature: string[];
    price: number;
    discount: number;
    discountEnd: string;
    supportedPlatform: string[];
    rating: number;
}[] = [
		
    {
			name: 'Cyberpunk 2077',
			midCoverImage:
				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
			feature: ['Action', 'RPG', 'Explore'],
            price:29.45,
            discount:50,
            discountEnd: "9/25/2022",
            supportedPlatform : ["mac", "windows", "linux"],
            rating:9.4,
		},
        {
			name: 'Cyberpunk 2077',
			midCoverImage:
				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
			feature: ['Action', 'RPG', 'Explore'],
            price:29.45,
            discount:50,rating:9.4,
            discountEnd: "9/25/2022",
            supportedPlatform : ["mac", "windows", "linux"]
		},
        {
			name: 'Cyberpunk 2077',
			midCoverImage:
				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
			feature: ['Action', 'RPG', 'Explore'],
            price:29.45,rating:9.4,
            discount:50,
            discountEnd: "9/25/2022",
            supportedPlatform : ["mac", "windows", "linux"]
		},
        {
			name: 'Cyberpunk 2077',
			midCoverImage:
				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
			feature: ['Action', 'RPG', 'Explore'],
            price:29.45,rating:9.4,
            discount:50,
            discountEnd: "9/25/2022",
            supportedPlatform : ["mac", "windows", "linux"]
		},
	];

    let total ={
        price :0,
        discount:0,
        
    }
    function CountTotal(){
        total.price=0
        total.discount=0
        for (let i=0;i<li.length;i++ ){
            total.price+=li[i].price 
            total.discount+=(li[i].price *li[i].discount /100)
        }
    }
    CountTotal()
</script>

<div class="flex max-h-fit min-h-[90%] w-full flex-col justify-start lg:px-48  ">
	<p class="  font-Poppins text-4xl text-slate-200 ">My Cart</p>
	<div class="flex w-full flex-row font-Poppins text-slate-200 ">
		<!-- CART GAMES -->
		<div class="flex w-full flex-col gap-2">
            {#if li.length!=0}
                 <!-- content here -->
                 {#each li as  i}
                     <!-- content here -->
                     <div class=" h-56 w-[1110px]  flex flex-row items-center hover:bg-slate-800 bg-[#202020] rounded-md pl-3 " >
                         <img src="{i.midCoverImage}" alt="{i.name}" class=" w-[150px] h-[200px] ">
                         <div class=" w-full h-full ml-4 mr-2 flex flex-row">
                             <div class="w-fit h-full flex flex-col">
                                 <p class=" text-2xl mt-4">{i.name}</p>
                                 <div class="flex flex-row gap-2">
     
                                     {#each i.feature as j}
                                     <p class=" text-lg font-thin text-blue-500">{j}</p>
                                     {/each}
                                 </div>
                                 <p class=" font-semibold">Rating {i.rating}/10</p>
                                 <div class="grow"></div>
                                 <div class=" flex flex-row mb-4 gap-2">
                                     {#each i.supportedPlatform as j}
                                          {#if j=="mac"}
                                             <Apple class="h-6 w-6 fill-slate-400" />
                                          {:else if  j=="linux"}
                                             <Linux class="h-6 w-6 fill-slate-400" />
                                          {:else}
                                             <Windows class="h-6 w-6 fill-slate-400" />
                                          {/if}
                                     {/each}
                                 </div>
                             </div>
                             <div class="grow"></div>
                             <div class=" w-fit h-full flex flex-col mt-4 mb-3">
                                  <del class=" text-gray-400">${i.price} </del> <p class="text-xl font-bold">${i.price-(i.price*.5)}</p>
                                 <p class="">sales Ends in {i.discountEnd}</p>
                                 <div class="grow"></div>
                                 <button class=" mb-8 text-gray-500 underline hover:text-gray-200 ">Remove</button>
                             </div>
                         </div>
                     </div>
     
                 {/each}
            {:else}
                 <!-- else content here -->
                 <div class=" h-full w-full flex  justify-center items-center">
                    <p class=" text-2xl font-thin">No Game in Cart</p>
                 </div>
            {/if}
		</div>
		<!-- SUMMARY -->
		<div class=" flex h-fit w-80 flex-col gap-4">
			<p class=" mb-6 text-4xl">Games and Apps Summary</p>
			<div class="flex flex-row">
				<p class="">Price</p>
				<div class=" grow " />
				<p class="">${total.price}</p>
			</div>
			<div class="flex flex-row">
				<p class="">Sale Discount</p>
				<div class=" grow " />
				<p class="">-${total.discount}</p>
			</div>
			<div class="flex flex-row">
				<p class="">Taxes</p>
				<div class=" grow " />
				<p class=" text-gray-500">Calculated at Checkout</p>
			</div>

			<div class="mt-2 flex flex-row border-t-[1px] border-gray-700 pt-2 ">
				<p class=" text-lg font-semibold">SubTotal</p>
				<div class=" grow " />
				<p class=" text-lg font-semibold">${total.price-total.discount}</p>
			</div>

			<button class=" h-[50px] w-[296px] bg-blue-500 transition-all ease-linear duration-150 hover:bg-sky-500 hover:font-semibold"> Buy Now </button>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
