<script lang="ts">
    // your script goes here
    import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';

	export let data: PageData;
	
	let { Userdata } = data;
	console.log("🚀 ~ file: +page.svelte ~ line 10 ~ Userdata", Userdata)
	UserProData.update((d) => (d = Userdata));




    let NewsData =data.NewsData 
    let NewsDatax ={
        Title:"The Epic Games Store Halloween Sale Ending Soon for 2022",
        Banner:"https://cdn2.unrealengine.com/epic-games-halloween-sale-ending-soon-1920x1080-2267b349b6e1.jpg",
        Date:"10/28/2022",
        WrittenBy:"Craig Pearson",
        Detail:`You thought you could get away, didn’t you? Well, we’ve caught you, and now it’s time for you to pay … less for your games. Oh no, PLEASE STOP SCREAMING! We’re just trying to remind you that the Epic Games Store Halloween Sale is still allllliiivveee, with prices slashed. And there’s a horde of ongoing in-game events for you to enjoy.

We’ve already covered a number of games and events in the previous Halloween Sale blog, but the discount virus has spread all over the store. You’ll find a second list of highlights below, including the classic Fallout franchise’s debut on the Epic Games Store.

The sale is live now and ends on Nov. 1, 2022, 11 AM ET. But that’s not the end. With Halloween in-game events in full swing, now’s the perfect time to find out just what the ‘Pumpkintiles’ are capable of in Dying Light 2: Stay Human and to work out what all that screaming is about in Dauntless. `,
        FooterComment:"Ghostbusters: Spirits Unleashed is available now on the Epic Games Store. "
    }
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->

<div class=" w-full h-fit flex flex-col items-center gap-2  ">
    <p class=" text-4xl font-Poppins font-semibold text-slate-200 w-[60%] text-center">{NewsData.Title}</p>
    <img src="{NewsData.BannerImg}" alt="banner" class="w-[60%] h-fit object-cover">
    <div class="w-[70%] flex flex-col items-start">
        <p class=" font-mono text-sm text-slate-300">Published At : {NewsData.Date}</p>
        <p class="font-mono text-sm mb-4 text-slate-300">Written By : {NewsData.WrittenBy}</p>
        <p class="text-slate-100 font-Poppins mb-5">{NewsData.Detail}</p>
        <p class=" text-blue-600 font-Poppins ">{NewsData.FooterComment}</p>
        
    </div>
</div>