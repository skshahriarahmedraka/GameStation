<script lang="ts">
	import { goto } from '$app/navigation';
     import type { PageData } from './$types';

export let data: PageData;

    let NewsDataList:any =data.NewsDataList 
    let liNews=[
        {
            ID:"8745",
            Image:"https://cdn2.unrealengine.com/quickly-grab-legendary-gear-in-destiny-2-and-show-off-your-fortnite-ornaments-1920x1080-5b5534ccd00d.jpg",
            Date:"23 jul, 22",
            Title:"Quickly grab Legendary gear in Destiny 2 and show off your Fortnite ornaments",
            Content:"With Fortnite and Destiny 2 hosting a historic crossover event to celebrate Destiny2's arrival on the Epic Games Store, both games received lovely new cosmetics to celebrate. In Fortnite, you can buy the new Outfits in the Item Store and wear them out in battle. The same is true in Destiny 2! The only difference is that in Destiny 2, the ornaments are worn on the armor you already have. So, let’s get you some of that armor!  "
        },
        {
            ID:"8745",
            Image:"https://cdn2.unrealengine.com/quickly-grab-legendary-gear-in-destiny-2-and-show-off-your-fortnite-ornaments-1920x1080-5b5534ccd00d.jpg",
            Date:"23 jul, 22",
            Title:"Quickly grab Legendary gear in Destiny 2 and show off your Fortnite ornaments",
            Content:"With Fortnite and Destiny 2 hosting a historic crossover event to celebrate Destiny2's arrival on the Epic Games Store, both games received lovely new cosmetics to celebrate. In Fortnite, you can buy the new Outfits in the Item Store and wear them out in battle. The same is true in Destiny 2! The only difference is that in Destiny 2, the ornaments are worn on the armor you already have. So, let’s get you some of that armor!  "
        },
        {
            ID:"8745",
            Image:"https://cdn2.unrealengine.com/quickly-grab-legendary-gear-in-destiny-2-and-show-off-your-fortnite-ornaments-1920x1080-5b5534ccd00d.jpg",
            Date:"23 jul, 22",
            Title:"Quickly grab Legendary gear in Destiny 2 and show off your Fortnite ornaments",
            Content:"With Fortnite and Destiny 2 hosting a historic crossover event to celebrate Destiny2's arrival on the Epic Games Store, both games received lovely new cosmetics to celebrate. In Fortnite, you can buy the new Outfits in the Item Store and wear them out in battle. The same is true in Destiny 2! The only difference is that in Destiny 2, the ornaments are worn on the armor you already have. So, let’s get you some of that armor!  "
        },
        {
            ID:"8745",
            Image:"https://cdn2.unrealengine.com/quickly-grab-legendary-gear-in-destiny-2-and-show-off-your-fortnite-ornaments-1920x1080-5b5534ccd00d.jpg",
            Date:"23 jul, 22",
            Title:"Quickly grab Legendary gear in Destiny 2 and show off your Fortnite ornaments",
            Content:"With Fortnite and Destiny 2 hosting a historic crossover event to celebrate Destiny2's arrival on the Epic Games Store, both games received lovely new cosmetics to celebrate. In Fortnite, you can buy the new Outfits in the Item Store and wear them out in battle. The same is true in Destiny 2! The only difference is that in Destiny 2, the ornaments are worn on the armor you already have. So, let’s get you some of that armor!  "
        },
        {
            ID:"8745",
            Image:"https://cdn2.unrealengine.com/quickly-grab-legendary-gear-in-destiny-2-and-show-off-your-fortnite-ornaments-1920x1080-5b5534ccd00d.jpg",
            Date:"23 jul, 22",
            Title:"Quickly grab Legendary gear in Destiny 2 and show off your Fortnite ornaments",
            Content:"With Fortnite and Destiny 2 hosting a historic crossover event to celebrate Destiny2's arrival on the Epic Games Store, both games received lovely new cosmetics to celebrate. In Fortnite, you can buy the new Outfits in the Item Store and wear them out in battle. The same is true in Destiny 2! The only difference is that in Destiny 2, the ornaments are worn on the armor you already have. So, let’s get you some of that armor!  "
        },
    ]
</script>

<style>
    /* your styles go here */
</style>

<div class=" flex flex-col gap-0 items-center">
    {#each NewsDataList  as i}
        <div on:click={()=> goto(`/news/${i.NewsID}`)} on:keypress={()=>{}}  class="flex flex-row justify-center hover:cursor-pointer items-center max-w-[1420px] h-[133.5px] hover:bg-slate-900  border-y-[1px] border-gray-700">
            <img src="{i.BannerImg}" alt="" class=" ml-4 h-[112.5px] w-[200px]">
            <div class="flex flex-col gap-1 ml-2 font-Poppins text-slate-100 ">
                <p class=" text-sm ">{i.Date}</p>
                <p class="text-base font-semibold">{i.Title}</p>
                <p class=" text-sm line-clamp-2">{i.Detail}</p>
            </div>
        </div>
    {/each}
</div>