<script lang="ts">
	import Discord from '$lib/foot/icons/discord.svelte';
	import Apple from '$lib/svgs/apple.svelte';
	import Linux from '$lib/svgs/linux.svelte';
	import Windows from '$lib/svgs/windows.svelte';
	import Plus from '$lib/svgs/plus.svelte';
	import Camera from '$lib/svgs/camera.svelte';
	import CameraPlus from '$lib/svgs/cameraPlus.svelte';
	import Checked from '$lib/svgs/checked.svelte';
	import axios from 'axios';

	const mirageLogo = new URL('../images/mirageLogo.png', import.meta.url).href;
	export let LogoImage: string;
	export let Price: number = 0;
	export let Discount: number = 25;
	export let Develper: string = '';
	export let Publisher: string = '';
	export let Platform: string[] = ['', '', ''];
	export let Released: string = '';
	let avatar: any = '';
	let FileInput: any;

	const onFileSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		reader.onload = (e) => {
			avatar = e.target?.result;
			UploadImage(String(e.target?.result));
		};
	};

	async function UploadImage(img: string) {
		const data = { img: '' };
		const imgData = img.split(',');
		data['img'] = imgData[1];

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
		const imgLink = await res.json();
		console.log('🚀 ~ file: inputGameProfile.svelte ~ line 56 ~ UploadImage ~ imgLink', imgLink);
		avatar=imgLink.Link 
	}

	
	



	
</script>

<!-- markup (zero or more items) goes here -->

<div class="flex  h-fit w-80 flex-col gap-3  ">
	<div class=" relative  flex w-full flex-row justify-center ">
		{#if avatar != ''}
			<!-- banner image -->
			<img
				src={avatar}
				alt="ProfileImage"
				class=" aspect-[18/10]  w-full self-center rounded-lg object-contain "
			/>
		{:else}
			<!-- banner image -->
			<!-- <Accord_Default class="mt-6 aspect-[16/9]  w-[80%] object-contain   rounded-lg  bg-slate-600 " /> -->
			<!-- <Bd class="mt-6 aspect-[16/9]  w-[80%] object-cover   rounded-lg  bg-slate-600 " />  -->
			<!-- <Accord class="mt-6 w-[80%] aspect-[16/9] h-80 self-center rounded-lg object-fill bg-slate-600 " /> -->
			<CameraPlus class="  h-56 w-56  stroke-1 " />
			<!-- <img
				src="{imgUrl}"
				alt="ProfileImage"
				class=" mt-6 aspect-[16/9] w-[80%]  self-center rounded-lg object-cover "
			/> -->
		{/if}
		<button class=" absolute   h-full   w-full self-end  " on:click={() => FileInput.click()}>
			<!-- <img src="https://static.thenounproject.com/png/625182-200.png" alt="uploadImage" class=" h-28 w-28"> -->
			<!-- banner camera logo -->
			<!-- <Camera class="h-10 w-10 stroke-slate-100 stroke-1 opacity-50  hover:opacity-100" /> -->
			<input
				bind:this={FileInput}
				on:change={(e) => onFileSelected(e)}
				type="file"
				class=" hidden h-10 w-10 "
				accept=".jpg, .jpeg, .png, .svg"
			/>
		</button>
	</div>

	<!-- <img src={LogoImage} alt="" class=" h-fit w-full object-cover " /> -->

	<!-- <div class=" ml-4 flex flex-row gap-2">
		<p class=" h-[25px] w-fit rounded-lg  bg-blue-500 px-4 py-1 text-center text-sm text-white">
			-{Discount}%
		</p>
		<del class=" text-slate-400">${Price}</del>
		<p class=" text-slate-100">${(Price - (Price * Discount) / 100).toFixed(2)}</p>
	</div> -->
	<!-- <button
		class=" ml-4 h-[50px] w-[320px] rounded-lg bg-sky-500 hover:bg-sky-600 font-Poppins text-lg font-semibold text-white"
		>Buy Now</button
	>
	<button
		class=" ml-4 h-[50px] w-[320px] rounded-lg border-2 border-white font-Poppins text-lg font-semibold hover:border-slate-300  text-white hover:text-slate-300"
		>Add To Cart</button
	>
	<button
		class=" ml-4 flex h-[30px] w-[320px] flex-row  items-center justify-center rounded-lg border-2  border-white hover:border-slate-300 font-Poppins font-semibold hover:text-slate-200 text-white"
	>
		<Plus class="h-4 w-4" />
		<p class="text-sm">Add To Wishlist</p></button
	> -->

	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Price :</p>
		<input
			bind:value={Price}
			type="text"
			placeholder="Price"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Discount(%) :</p>
		<input
			bind:value={Discount}
			type="text"
			placeholder="Discount"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- Developer -->
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Developer :</p>
		<input
			bind:value={Develper}
			type="text"
			placeholder="Developer"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- Publisher -->
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Publisher :</p>
		<input
			bind:value={Publisher}
			type="text"
			placeholder="Publisher"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Release Date :</p>
		<input
			bind:value={Released}
			type="text"
			placeholder="Released Date"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- <div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Developer :</p>
		<input
			bind:value={Develper}
			type="text"
			placeholder="Developer"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Developer :</p>
		<input
			bind:value={Develper}
			type="text"
			placeholder="Developer"
			class=" my-2 h-8 w-full self-start rounded-2xl border-2  border-[#24262b]  bg-[#303338] p-2   font-medium  text-slate-100 outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div> -->
	<!-- <div
		class=" ml-6   mt-2 flex w-[320px] flex-row justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Publisher</p>
		<p class=" ">{Publisher}</p>
	</div>
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-row justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Release Date</p>
		<p class=" ">{Released}</p>
	</div> -->
	<div
		class=" ml-6   mt-2 flex w-[320px] flex-col justify-between border-b-[1px] border-slate-500 border-opacity-40 pr-4 pb-2 font-Poppins text-slate-100 "
	>
		<p class=" text-gray-400 ">Platform</p>
		<div class=" mb-4 flex flex-row gap-2">
			<button
				class="flex flex-row stroke-2"
				on:click={() => {
					Platform[0] === 'linux' ? (Platform[0] = '') : (Platform[0] = 'linux');
				}}
			>
				<Checked
					class=" h-5 w-5 {Platform[0] === 'linux' ? ' stroke-green-400 ' : ' stroke-gray-600 '}"
				/>
				<Linux class="h-5 w-5 fill-slate-100" />
			</button>

			<button
				class="flex flex-row stroke-2"
				on:click={() => {
					Platform[1] === 'mac' ? (Platform[1] = '') : (Platform[1] = 'mac');
				}}
			>
				<Checked
					class=" h-5 w-5 {Platform[1] === 'mac' ? ' stroke-green-400 ' : ' stroke-gray-600 '}"
				/>
				<Apple class="h-5 w-5 fill-slate-100" />
			</button>
			<button
				class="flex flex-row stroke-2"
				on:click={() => {
					Platform[2] === 'windows' ? (Platform[2] = '') : (Platform[2] = 'windows');
				}}
			>
				<Checked
					class=" h-5 w-5 {Platform[2] === 'windows' ? ' stroke-green-400 ' : ' stroke-gray-600 '}"
				/>
				<Windows class="h-5 w-5 fill-slate-100" />
			</button>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
