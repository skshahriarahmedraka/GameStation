<script lang="ts">
	//
	import CameraPlus from '$lib/svgs/cameraPlus.svelte';
	import { UserProData } from '$lib/Store/store';

	let GeneralData = {
		UserID: $UserProData.UserID,
		Name: $UserProData.Name,
		Accounttype: $UserProData.Accounttype,

		ProfileImg: $UserProData.ProfileImg,
		BannerImg: $UserProData.BannerImg,

		Email: $UserProData.Email,
		Mobile: $UserProData.Mobile,

		Address: $UserProData.Address,
		// Region: '',
		Country: $UserProData.Country,
		City: $UserProData.City,
		ZipCode: $UserProData.ZipCode,

		Coin: $UserProData.Coin,

		TransactionHistory: $UserProData.TransactionHistory || [],
		WishList: $UserProData.WishList || [],
		ContactAdminMsg: $UserProData.ContactAdminMsg || [],
		ContactDevMsg: $UserProData.ContactDevMsg || [],
		UserCart: $UserProData.UserCart || []
	};
	console.log('🚀 ~ file: General.svelte ~ line 31 ~ GeneralData', GeneralData);
	let ErrorMsg = {
		Userid: [false, ''],
		Accounttype: [false, ''],

		Name: [false, 'invalid/wrong email'],
		ProfileImg: [false, ''],
		BannerImg: [false, ''],

		Email: [false, 'Password is too short'],
		Mobile: [false, ''],

		Address: [false, ''],
		// Region: '',
		Country: [false, ''],
		City: [false, ''],
		ZipCode: [false, ''],

		Coin: [false, ''],

		TransactionHistory: [false, ''],
		WishList: [false, ''],
		ContactAdminMsg: [false, ''],
		ContactDevMsg: [false, ''],
		UserCart: [false, '']
	};
	// let SuccessMsg:string ="Success"
	const ValidateEmail = (E: string) => {
		return String(E)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<script>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	let FileInput: any;
	let FileInput2: any;
	let ThisProfileImg: any;
	let ThisBannerImg: any;
	// banner image
	const onBannerSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name = e.target.files[0].name;
		reader.onload = (e) => {
			GeneralData.BannerImg = String(e.target?.result);

			UploadBannerImage(String(e.target?.result), name);
		};
	};

	async function UploadBannerImage(img: string, name: string) {
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
		const imgLink = await res.json();
		console.log('🚀 ~ file: General.svelte ~ line 95 ~ UploadBannerImage ~ BannerImg', imgLink);
		GeneralData.BannerImg = imgLink.Link;
	}

	// Profile image
	const onFileSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name = e.target.files[0].name;
		reader.onload = (e) => {
			GeneralData.ProfileImg = String(e.target?.result);

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
		const imgLink = await res.json();
		console.log('🚀 ~ file: General.svelte ~ line 129 ~ UploadImage ~ ProfileImg', imgLink);
		GeneralData.ProfileImg = imgLink.Link;
	}

	//

	async function UpdateProfileData() {
		await fetch('/api/profile/general', {
			// credentials: 'same-origin',
			method: 'POST',
			// mode: 'cors',
			body: JSON.stringify(GeneralData)
		})
			.then((res) => {
				return res.json();
			})
			.then((resdata) => {
				console.log('🚀 ~ file: General.svelte ~ line 148 ~ .then ~ resdata', resdata);
				UserProData.update((n) => {
					return (n = resdata);
				});
				GeneralData = resdata;
				console.log('🚀 ~ file: General.svelte ~ line 152 ~ .then ~ UserProData', $UserProData);
			});
	}
	function DiscardChanges() {
		GeneralData = {
			UserID: $UserProData.UserID,
			Name: $UserProData.Name,
			Accounttype: $UserProData.Accounttype,

			ProfileImg: $UserProData.ProfileImg,
			BannerImg: $UserProData.BannerImg,

			Email: $UserProData.Email,
			Mobile: $UserProData.Mobile,

			Address: $UserProData.Address,
			// Region: '',
			Country: $UserProData.Country,
			City: $UserProData.City,
			ZipCode: $UserProData.ZipCode,

			Coin: $UserProData.Coin,

			TransactionHistory: $UserProData.TransactionHistory || [],
			WishList: $UserProData.WishList || [],
			ContactAdminMsg: $UserProData.ContactAdminMsg || [],
			ContactDevMsg: $UserProData.ContactDevMsg || [],
			UserCart: $UserProData.UserCart || []
		};
	}
</script>

<!-- Banner Img -->
<p class="m-4 text-slate-200">Banner Image :</p>
<div
	class=" relative  mb-5 mt-1 flex w-full flex-row justify-center transition-all duration-200 ease-linear xs:w-full "
>
	{#if GeneralData.BannerImg != ''}
		<!-- banner image -->
		<img
			src={GeneralData.BannerImg}
			bind:this={ThisBannerImg}
			on:error={() => {
				ThisBannerImg.src = GeneralData.BannerImg;
			}}
			alt="ProfileImage"
			class=" h-56 w-[80%] self-center rounded-full object-cover   "
		/>
	{:else}
		<CameraPlus
			class="  h-56 w-[80%] rounded-lg  border-2 border-slate-200 stroke-slate-200 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
		/>

		<!-- <CameraPlus class="  rounded-xl  h-56 w-56 stroke-slate-200 stroke-[1px] border-2 border-slate-200 transition-all ease-linear duration-150 hover:stroke-slate-500 hover:border-slate-500  hover:cursor-pointer"   /> -->
	{/if}
	<button class="  absolute  h-56 w-full self-end  " on:click={() => FileInput2.click()}>
		<input
			bind:this={FileInput2}
			on:change={(e) => onBannerSelected(e)}
			type="file"
			class=" hidden h-10 w-10 "
			accept=".jpg, .jpeg, .png, .svg"
		/>
	</button>
</div>
<!-- ProfileImg -->
<p class="m-4 text-slate-200">Profile Image :</p>

<div
	class=" relative  flex w-full flex-row justify-center transition-all duration-200 ease-linear xs:w-full  "
>
	{#if GeneralData.ProfileImg != ''}
		<!-- banner image -->
		<img
			src={GeneralData.ProfileImg}
			bind:this={ThisProfileImg}
			on:error={() => {
				ThisProfileImg.src = GeneralData.ProfileImg;
			}}
			alt="ProfileImage"
			class=" h-56 w-56 self-center rounded-full object-cover xs:h-24 xs:w-24  "
		/>
	{:else}
		<CameraPlus
			class="  h-56 w-56 rounded-full  border-2 border-slate-200 stroke-slate-200 stroke-[1px] transition-all duration-150 ease-linear hover:cursor-pointer hover:border-slate-500 hover:stroke-slate-500"
		/>

		<!-- <CameraPlus class="  rounded-xl  h-56 w-56 stroke-slate-200 stroke-[1px] border-2 border-slate-200 transition-all ease-linear duration-150 hover:stroke-slate-500 hover:border-slate-500  hover:cursor-pointer"   /> -->
	{/if}
	<button class=" absolute   h-56 w-56   self-end  " on:click={() => FileInput.click()}>
		<input
			bind:this={FileInput}
			on:change={(e) => onFileSelected(e)}
			type="file"
			class=" hidden h-10 w-10 "
			accept=".jpg, .jpeg, .png, .svg"
		/>
	</button>
</div>
<p class=" ml-5 mt-3 text-2xl font-semibold text-slate-200">General Settings</p>

<p class="ml-5 mt-1 text-sm text-slate-400">
	Manage the account details you've shared with Epic Games including your name, contact info and
	more
</p>

<p class=" ml-5 mt-3 text-2xl font-semibold text-slate-200">Account Info</p>

<p class="ml-5 mt-1 text-sm text-slate-400">ID: {GeneralData.UserID}</p>
<div class="mt-4 flex flex-row flex-wrap   justify-center  gap-4 ">
	<!-- Name -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Name :
			</div>
			{#if ErrorMsg.Name[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Name[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.Name}
			type="text"
			placeholder="User Name"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- Email -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Email :
			</div>
			{#if ErrorMsg.Email[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Email[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.Email}
			type="text"
			placeholder="Email Address"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- Mobile -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Mobile :
			</div>
			{#if ErrorMsg.Mobile[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Mobile[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.Mobile}
			type="text"
			placeholder="Mobile Number"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- Coin -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Coin :
			</div>
		</div>
		<div
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		>
			{GeneralData.Coin}
		</div>
	</div>
</div>

<!-- ///////////////////////////// -->
<p class=" ml-5 mt-3 text-2xl font-semibold text-slate-200">Address Info :</p>

<div class="mt-4 flex flex-row flex-wrap   justify-center  gap-4 ">
	<!-- Address -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Address :
			</div>
			{#if ErrorMsg.Address[0]}
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Address[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.Address}
			type="text"
			placeholder="Your Address"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>

	<!-- City -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				City :
			</div>
			{#if ErrorMsg.City[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.City[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.City}
			type="text"
			placeholder="Your City"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- ZipCode -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Zip Code :
			</div>
			{#if ErrorMsg.ZipCode[0]}
				<!-- content here -->
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.ZipCode[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.ZipCode}
			type="text"
			placeholder="Your Location Zip Code"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
	<!-- Country -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Country :
			</div>
			{#if ErrorMsg.Country[0]}
				<div class=" mt-2 ml-3 inline-flex ">
					<svg
						class="h-6 w-6  fill-red-500"
						fill="currentColor"
						viewBox="0 0 20 20"
						xmlns="http://www.w3.org/2000/svg"
						><path
							fill-rule="evenodd"
							d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
							clip-rule="evenodd"
						/></svg
					>
					<p class="text-red-300">{ErrorMsg.Country[1]}</p>
				</div>
			{/if}
		</div>
		<input
			bind:value={GeneralData.Country}
			type="text"
			placeholder="Your Country"
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		/>
	</div>
</div>
<p class=" ml-5 mt-3 text-2xl font-semibold text-slate-200">My Stats :</p>
<div class="mt-4 flex flex-row flex-wrap   justify-center  gap-4 ">
	<!-- Address -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Number Of Transactions :
			</div>
		</div>
		<div
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-transparent p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		>
			{GeneralData.TransactionHistory.length}
		</div>
	</div>
	<!-- Address -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Number Of WishList :
			</div>
		</div>
		<div
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-transparent  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		>
			{GeneralData.WishList.length}
		</div>
	</div>
	<!-- Address -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Number Of Developer Message :
			</div>
		</div>
		<div
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-transparent  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		>
			{GeneralData.ContactDevMsg.length}
		</div>
	</div>
	<!-- Address -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Number Of Admin Message :
			</div>
		</div>
		<div
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-transparent  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		>
			{GeneralData.ContactAdminMsg.length}
		</div>
	</div>
	<!-- Address -->
	<div class="flex w-[45%] flex-col xs:w-full ">
		<div class=" flex flex-row">
			<div class=" ml-3 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
				Number Of Carts :
			</div>
		</div>
		<div
			class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-transparent   p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
		>
			{GeneralData.UserCart.length}
		</div>
	</div>
</div>

<div class="mt-4 flex  flex-row flex-wrap justify-center gap-2" />
<!-- Submit Button -->
<div class=" mb-10 mt-5 flex  flex-row  justify-center">
	<button
		on:click={() => {
			UpdateProfileData();
		}}
		class=" ml-4 h-[50px] w-[320px] rounded-lg bg-sky-500 font-Poppins text-lg font-semibold text-white hover:bg-sky-600"
	>
		Save Changes
	</button>
	<button
		on:click={() => {
			DiscardChanges();
		}}
		class=" ml-4 h-[50px] w-[320px] rounded-lg border-2 border-slate-200 bg-slate-600  font-Poppins text-lg font-semibold text-white hover:bg-opacity-60"
	>
		Discard Changes
	</button>
</div>

<style>
	/* your styles go here */
</style>
