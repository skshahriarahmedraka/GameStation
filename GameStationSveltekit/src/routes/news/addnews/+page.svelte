


<script lang="ts">
	import { goto } from '$app/navigation';
	import Textarea from '$lib/otherComponents/textarea.svelte';
    // your script goes here
    let NewsData ={
        NewsID:"",
        Title:"",
        BannerImg:"",
        Date:"",
        WrittenBy:"",
        Detail:"",
        FooterComment:""
    }
    let NewsData2 ={
        NewsID:"",
        Title:"Blankos Block Party - Exclusive Epic Halloween Bundle",
        BannerImg:"",
        Date:"10/28/2022",
        WrittenBy:"Craig Pearson",
        Detail:`You thought you could get away, didn’t you? Well, we’ve caught you, and now it’s time for you to pay … less for your games. Oh no, PLEASE STOP SCREAMING! We’re just trying to remind you that the Epic Games Store Halloween Sale is still allllliiivveee, with prices slashed. And there’s a horde of ongoing in-game events for you to enjoy.

We’ve already covered a number of games and events in the previous Halloween Sale blog, but the discount virus has spread all over the store. You’ll find a second list of highlights below, including the classic Fallout franchise’s debut on the Epic Games Store.

The sale is live now and ends on Nov. 1, 2022, 11 AM ET. But that’s not the end. With Halloween in-game events in full swing, now’s the perfect time to find out just what the ‘Pumpkintiles’ are capable of in Dying Light 2: Stay Human and to work out what all that screaming is about in Dauntless. `,
        FooterComment:"Ghostbusters: Spirits Unleashed is available now on the Epic Games Store. "
    }
    import CameraPlus from '$lib/svgs/cameraPlus.svelte';
	import { error } from '@sveltejs/kit';
    let FileInput: any;
    let ThisBannerImg:HTMLImageElement

	const onFileSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name =e.target.files[0].name
		reader.onload = (e) => {
			NewsData.BannerImg = String(e.target?.result)
			
			UploadImage(String(e.target?.result),name);
		};
	};

	async function UploadImage(img: string,name:string) {
        console.log("🚀 ~ file: +page.svelte ~ line 72 ~ UploadImage ~ UploadImage")
		const data = { img: '',name:"" };
		const imgData = img.split(',');
		data['img'] = imgData[1];
		data["name"]=name 

		// console.log(data);
        let ResData:any
		await fetch('/api/imgupload', {
			method: 'POST',
			// mode: 'no-cors',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify(data)
		}).then((res)=>{
            return res.json()
        }).then((d)=>{
            ResData=d
        })
        console.log("🚀 ~ file: +page.svelte ~ line 67 ~ UploadImage ~ ResData", ResData)
		
		NewsData.BannerImg=ResData.Link 
	}

    async function SubmitNews(){
        console.log("🚀 ~ file: +page.svelte ~ line 69 ~ SubmitNews ~ SubmitNews")
        
        await fetch('/api/news/addnews', {
			// credentials: 'same-origin',
			method: 'POST',
			// mode: 'no-cors',
			// headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(NewsData)
		})
			.then((response) => response.json())
			.then((data) => {
				console.log("🚀 ~ file: +page.svelte ~ line 78 ~ .then ~ data news added", data)
				console.log("🚀 ~ file: +page.svelte ~ line 78 ~ .then ~ data news added id", data.NewsID)
				
				if (data.NewsID) {
					console.log('🚀 ~ file: +page.svelte ~ line 627 ~ .then ~ data.GameID', data.NewsID);
					goto(`/news/${data.NewsID}`);
				} else {
					throw error(404, 'Not Found');
				}
			});
    }
</script>




<div class=" w-full h-fit flex flex-col items-center gap-2  ">
    <input bind:value={NewsData.Title} type="text"
    placeholder="News Title"
			class=" text-4xl font-Poppins font-semibold  my-2 h-16 w-[60%]  rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2    text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800  "
    
    >
    <!-- <p class=" text-4xl font-Poppins font-semibold text-slate-200">{NewsData.Title}</p> -->
    <div class=" relative  flex w-[60%] max-h-[600px] aspect-[18/10]    flex-row justify-center ">
	{#if NewsData.BannerImg != ''}
			<!-- banner image -->
			<div class="w-full h-full">
				<svg
					on:click={() => {
						NewsData.BannerImg=""
					}}
					on:keydown={()=>{}}
					class="absolute top-5 left-5 h-8  w-8 stroke-white hover:cursor-pointer hover:fill-gray-500"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
					/></svg
				>

				<img
				src={NewsData.BannerImg}
				bind:this={ThisBannerImg}
				on:error={ ()=>{ThisBannerImg.src=NewsData.BannerImg}}
				alt="BannerImg"
				class=" aspect-[18/10]  w-full self-center rounded-lg object-contain "
				/>
			</div>
		{:else}
		<CameraPlus class="  h-full w-full rounded-xl stroke-slate-200 border-2 border-slate-200 transition-all ease-linear duration-150 hover:stroke-slate-500 hover:border-slate-500 stroke-[1px] hover:cursor-pointer" />
			
			<!-- <CameraPlus class="  rounded-xl  h-56 w-56 stroke-slate-200 stroke-[1px] border-2 border-slate-200 transition-all ease-linear duration-150 hover:stroke-slate-500 hover:border-slate-500  hover:cursor-pointer"   /> -->
			
			<button class=" absolute   h-full   w-full self-end  " on:click={() => FileInput.click()}>
			
				<input
					bind:this={FileInput}
					on:change={(e) => onFileSelected(e)}
					type="file"
					class=" hidden h-10 w-10 "
					accept=".jpg, .jpeg, .png, .svg  .webp"
				/>
			</button>
		{/if}
	</div>
    <!-- <img src="{NewsData.Banner}" alt="banner" class="w-[60%] h-fit object-cover"> -->
    <div class="w-[70%] flex flex-col items-start">
        <!-- <p class=" font-mono text-sm text-slate-300"> : {NewsData.Date}</p> -->
        <div class=" flex flex-col gap-2 w-[90%] font-mono text-sm text-slate-300 ">
            <p class=" "> Published At (e.g. 10/28/2022) :</p>
            <input bind:value={NewsData.Date} type="text"
			class="  font-Poppins  my-2 h-10 w-[40%]  rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2    text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800  "   >
        </div>
        <!-- <p class=""> : {NewsData.WrittenBy}</p> -->
        <div class=" flex flex-col gap-2 w-[90%] font-mono text-sm mb-4 text-slate-300">
            <p class=" "> Written By :</p>
            <input bind:value={NewsData.WrittenBy} type="text"
			class="  font-Poppins  my-2 h-10 w-[40%]  rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2    text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800  "   >
        </div>
        <div class=" flex flex-col gap-2 w-[90%] ">
            <p class=" text-xl font-semibold text-slate-200">Description :</p>
          
            <Textarea class=" w-full h-[500px]  " bind:value={NewsData.Detail} minRows={4} maxRows={40} />
        </div>
        <div class=" flex flex-col gap-2 w-[90%] ">
            <p class=" text-xl font-semibold text-slate-200"> conclusion :</p>
            <input bind:value={NewsData.FooterComment} type="text"
			class=" text-base font-Poppins  my-2 h-10 w-[90%]  rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2    text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800  "   >
        </div>
        <!-- <p class="text-slate-100 font-Poppins mb-5">{NewsData.Detail}</p> -->
        <!-- <p class=" text-blue-600 font-Poppins ">{NewsData.FooterComment}</p>  -->
        
    </div>
    <div class=" flex w-full justify-center">
        <button
            on:click={() => {
                SubmitNews();
            }}
            class=" ml-4 h-[50px] w-[320px] rounded-lg bg-sky-500 font-Poppins text-lg font-semibold text-white hover:bg-sky-600"
        >
            Add News
        </button>
    </div>
</div>