<script lang="ts">
	const anthem = new URL('../images/anthem.jpg', import.meta.url).href;
	const fifa = new URL('../images/fifa.png', import.meta.url).href;
	const forbiddenwest = new URL('../images/forbiddenwest.jpg', import.meta.url).href;
	const war = new URL('../images/GodOfWar.jpg', import.meta.url).href;
	import CameraPlus from '$lib/svgs/cameraPlus.svelte';

	let ThisCarouselImage:any
	export let OtherImages: string[];
	export let BigPosterImage: string = '';
	// const images = [
	// 	{
	// 		id: 0,
	// 		name: 'Cosmic timetraveler',
	// 		imgurl: anthem,
	// 		attribution: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com'
	// 	},
	// 	{
	// 		id: 1,
	// 		name: 'Cristina Gottardi',
	// 		imgurl: fifa,
	// 		attribution: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com'
	// 	},
	// 	{
	// 		id: 2,
	// 		name: 'Johannes Plenio',
	// 		imgurl: forbiddenwest,
	// 		attribution: 'johannes-plenio-RwHv7LgeC7s-unsplash.com'
	// 	},
	// 	{
	// 		id: 3,
	// 		name: 'Jonatan Pie',
	// 		imgurl: war,
	// 		attribution: 'jonatan-pie-3l3RwQdHRHg-unsplash.com'
	// 	}
	// ];
	// let activeImage = images[0];
	let showImages: string[] = [];
	let ActiveIndex: number = 0;
	function ChangeActiveImagePositive() {
		if (ActiveIndex >= showImages.length) {
			ActiveIndex = 0;
		} else {
			ActiveIndex += 1;
		}
	}
	function ChangeActiveImageNagative() {
		if (ActiveIndex === 0) {
			ActiveIndex = showImages.length;
		} else {
			ActiveIndex -= 1;
		}
	}
	// import { flip } from "svelte/animate";
	import { elasticIn, elasticOut } from 'svelte/easing';

	import { fade, slide } from 'svelte/transition';

	let FileInput: any;

	const onFileSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name = e.target.files[0].name;
		reader.onload = (e) => {
			// LogoImage = String(e.target?.result)

			UploadImage(String(e.target?.result), name);
		};
	};

	async function UploadImage(img: string, name: string) {
		const data = { img: '', name: '' };
		const imgData = img.split(',');
		data['img'] = imgData[1];
		data['name'] = name;

		// console.log(data);
		const res = await fetch('/api/imgupload', {
			method: 'POST',
			// mode: 'no-cors',
			headers: {
				'Content-Type': 'application/json',
				Accept: 'application/json'
			},
			body: JSON.stringify(data)
		});
		if (res.status === 200) {
			const imgLink = await res.json();
			// console.log('🚀 ~ file: inputCarousels.svelte ~ line 93 ~ UploadImage ~ imgLink', imgLink);

			showImages.push(imgLink.Link);
			// console.log(
			// 	'🚀 ~ file: inputCarousels.svelte ~ line 96 ~ UploadImage ~ showImages',
			// 	showImages
			// );
			ActiveIndex = showImages.length - 1;
			BigPosterImage = showImages[0];
			if (showImages.length > 1) {
				// OtherImages = showImages.slice(1);
				OtherImages = showImages;
			}
		}else {
			
			console.log("❌🔥 ~ file: inputCarousels.svelte ~ line 105 ~ UploadImage ~ Error")
		}
	}
</script>

<div class="carousel aspect-[16/9] h-[600px]">
	<!-- content here -->

	<div
		transition:slide={{ duration: 1000, easing: elasticOut }}
		id="slide1"
		class="carousel-item relative w-full"
	>
		{#if ActiveIndex === showImages.length}
			<!-- content here -->
			<button class=" absolute   h-full   w-full self-end  " on:click={() => FileInput.click()}>
				<CameraPlus
					class=" h-full w-full rounded-xl  border-2 border-slate-200 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
				/>
				<input
					bind:this={FileInput}
					on:change={(e) => onFileSelected(e)}
					type="file"
					class=" hidden h-0 w-0  hover:cursor-pointer "
					accept=".jpg, .jpeg, .png, .svg .webp"
				/>
			</button>
		{:else}
			<div class=" h-full w-full">
				<img
					src={showImages[ActiveIndex]}
					bind:this={ThisCarouselImage}
				on:error={ ()=>{ThisCarouselImage.src=showImages[ActiveIndex]}}
					alt=""
					class="h-full w-full rounded-xl object-cover transition-all duration-200 ease-linear "
				/>
			</div>
		{/if}
		<!-- {:else}
			<button class=" absolute   h-full   w-full self-end  " on:click={() => FileInput.click()}>
				<CameraPlus class=" h-full w-full rounded-xl  border-2 border-slate-400  stroke-[1px]" />
				<input
					bind:this={FileInput}
					on:change={(e) => onFileSelected(e)}
					type="file"
					class=" hidden h-10 w-10 "
					accept=".jpg, .jpeg, .png, .svg"
				/>
			</button>
		{/if} -->

		<!-- <div class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between"> -->
		<button
			on:click={() => {
				ChangeActiveImageNagative();
			}}
			class="btn btn-circle absolute top-1/2 left-2">❮</button
		>
		<button
			on:click={() => {
				ChangeActiveImagePositive();
			}}
			class="btn btn-circle absolute right-2 top-1/2">❯</button
		>
		<!-- </div> -->
	</div>
</div>

<style>
	/* your styles go here */
</style>
