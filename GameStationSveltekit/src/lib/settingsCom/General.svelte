<script lang="ts">
    // 
	import CameraPlus from '$lib/svgs/cameraPlus.svelte';

    let GeneralData = {
		Name: '',
		Image: '',
		Email: '',
		Mobile: '',
		Userid: '',
		Address: '',
		Region: '',
		City: '',
		Postcode: '',
		Country: '',
		CompanyName: '',
		CompanyEmail: '',
		CompanyMobile: '',
		CompanyAddress: '',
		CompanyRegion: '',
		CompanyCity: '',
		CompanyPostcode: ''
	};
	let ErrorMsg = {
		Name: [false, 'invalid/wrong email'],
		Email: [false, 'Password is too short']
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

	const onFileSelected = (e: any) => {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		let name = e.target.files[0].name;
		reader.onload = (e) => {
			GeneralData.Image = String(e.target?.result);

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
		console.log('🚀 ~ file: inputGameProfile.svelte ~ line 56 ~ UploadImage ~ imgLink', imgLink);
		GeneralData.Image = imgLink.Link;
	}
</script>

<style>
    /* your styles go here */
</style>



<div class=" relative  flex w-full xs:w-full flex-row justify-center transition-all ease-linear duration-200  ">
			{#if GeneralData.Image != ''}
				<!-- banner image -->
				<img
					src={GeneralData.Image}
					alt="ProfileImage"
					class=" h-56 w-56 xs:h-24 xs:w-24 self-center rounded-full object-cover  "
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

		<p class="ml-5 mt-1 text-sm">
			Manage the account details you've shared with Epic Games including your name, contact info and
			more
		</p>

		<p class=" ml-5 mt-3 text-2xl font-semibold text-slate-200">Account Info</p>

		<p class="ml-5 mt-1 text-sm">ID: f4019c2a2d1f487c8360a65c816a28a8</p>
		<div class="mt-4 flex  xl:flex-row md:flex-col lg:flex-row  justify-center xs:flex-col sm:flex-col sm flex-wrap gap-4">
			<!-- Name -->
			<!-- Name -->
			<div class="flex w-[45%] xs:w-full flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Postcode :
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
					placeholder="Userid"
					class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<!-- Name -->
			<div class="flex w-[45%] flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Postcode :
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
					placeholder="Userid"
					class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<!-- Name -->
			<div class="flex w-[45%] flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Postcode :
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
					placeholder="Userid"
					class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<!-- Name -->
			<div class="flex w-[45%] flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Postcode :
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
					placeholder="Userid"
					class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
		</div>

		<!-- ///////////////////////////// -->
		<p class=" ml-5 mt-3 text-2xl font-semibold text-slate-200">Address Info :</p>

		<div class="mt-4 flex  flex-row justify-center gap-2">
			<!-- Name -->
			<div class="flex flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Address :
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
					placeholder="Name"
					class=" mx-4 my-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<!-- Email -->
			<div class="flex flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						City :
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
					placeholder="Email"
					class=" mx-4 my-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
		</div>
		<div class="mt-4 flex  flex-row flex-wrap justify-center gap-2">
			<!-- Mobile -->
			<div class="flex w-[40%] flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Region :
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
					placeholder="mobile"
					class=" mx-4 my-2 h-12 w-full  self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<!-- Name -->
			<div class="flex w-[40%] flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Postcode :
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
					placeholder="Userid"
					class=" mx-4 my-2 h-12 w-full self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
			<div class="flex flex-col ">
				<div class=" flex flex-row">
					<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
						Country :
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
					placeholder="Userid"
					class=" mx-4 my-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
				/>
			</div>
		</div>
		<!-- Submit Button -->
		<div class=" mb-10 mt-5 flex  flex-row  justify-center">
			<button
				class=" ml-4 h-[50px] w-[320px] rounded-lg bg-sky-500 font-Poppins text-lg font-semibold text-white hover:bg-sky-600"
			>
				Save Changes
			</button>
			<button
				class=" ml-4 h-[50px] w-[320px] rounded-lg border-2 border-slate-200 bg-slate-600  font-Poppins text-lg font-semibold text-white hover:bg-opacity-60"
			>
				Discard Changes
	</button>
</div>
