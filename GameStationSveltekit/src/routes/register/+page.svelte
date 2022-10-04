
<script lang="ts">
	import { goto } from '$app/navigation';
	import Amazon from "$lib/svgs/Amazon.svelte";
	import Github from "$lib/svgs/Github.svelte";
	import Google from "$lib/svgs/google.svelte";
	import Twitter from "$lib/svgs/twitter.svelte";
	import Bg1 from "$lib/images/mix3.jpeg"
	const RegBg = new URL(Bg1, import.meta.url).href;

	let RegisterData = {
		UserName: '',
		Email: '',
		Password: ''
	};
	let ErrorMsg = {
		UserName: [false, 'UserName is not Valid'],
		Email: [false, 'invalid/wrong email'],
		Password: [false, 'Password is too short']
	};
	// let b:boolean=false
	const ValidateEmail = (E: string) => {
		return String(E)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	function ValidateUserName(e: string) {
		return e.match(/^[a-zA-Z\s]*$/);
	}

	async function OnSubmit() {
		if (RegisterData.UserName.trim() === '') {
			ErrorMsg.UserName[1] = 'User Name is empty';
			ErrorMsg.UserName[0] = true;
		} else if (!ValidateUserName(RegisterData.UserName)) {
			ErrorMsg.UserName[1] = 'invalid/wrong User Name';
			ErrorMsg.UserName[0] = true;
		} else {
			ErrorMsg.UserName[0] = false;
		}
		if (RegisterData.Email.trim() === '') {
			ErrorMsg.Email[1] = 'Email is Empty';
			ErrorMsg.Email[0] = true;
		} else if (!ValidateEmail(RegisterData.Email)) {
			ErrorMsg.Email[1] = 'invalid/wrong email ';
			ErrorMsg.Email[0] = true;
		} else {
			ErrorMsg.Email[0] = false;
		}
		if (RegisterData.Password.trim() === '') {
			ErrorMsg.Password[1] = 'Password is Empty';
			ErrorMsg.Password[0] = true;
		} else if (RegisterData.Password.length < 8 || RegisterData.Password.length > 40) {
			ErrorMsg.Password[1] = 'Too big or too small';
			ErrorMsg.Password[0] = true;
		} else {
			ErrorMsg.Password[0] = false;
		}
		if (!ErrorMsg.UserName[0] && !ErrorMsg.Email[0] && !ErrorMsg.Password[0]) {
			await fetch('auth/register', {
				credentials: 'same-origin',
				method: 'POST',
				mode: 'cors',
				body: JSON.stringify(RegisterData)
			})
				.then((response) => response.json())
				.then((value) => {
					// console.log("🚀 ~ file: index.svelte ~ line 71 ~ .then ~ value : ", value)
					// console.log("🚀 ~ file: index.svelte ~ line 71 ~ .then ~ value.ID : ", value.ID)
					ErrorMsg.UserName[1] = 'Success';
					ErrorMsg.UserName[0] = true;
					RegisterData = {
						UserName: '',
						Email: '',
						Password: ''
					};
					let s:string="/"+String(value.ID)
                    // console.log("🚀 ~ file: index.svelte ~ line 73 ~ .then ~ s : ", s)
                    // console.log("🚀 ~ file: index.svelte ~ line 73 ~ .then ~ s typeof : ", typeof(s))
					goto(s);
				});
		}
	}
</script>
<!-- bg-[#181818] -->
<div class="grid h-screen w-full  place-items-center  " style="background-image: url('{RegBg}');">
	<a
		href="/login"
		on:click={() => {
			RegisterData['UserName'] = '';
			RegisterData['Email'] = '';
			RegisterData['Password'] = '';
		}}
		class="absolute  top-0 right-0  m-10"
	>
		<button
			class="  h-12 w-60 rounded-3xl bg-sky-600 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
		>
			Login
		</button>
	</a>
	<div class="">
		<div class=" flex h-fit w-[500px] flex-col space-y-4 rounded-xl bg-[#2d2d2d] bg-opacity-90 pb-10">
			<div
				class="mt-5 flex w-full items-center justify-center text-3xl font-semibold text-gray-200"
			>
				Sign Up
			</div>
			<div class="flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					User Name :
				</div>
				{#if ErrorMsg.UserName[0]}
					<div class=" mt-2 ml-3 inline-flex ">
						<svg
							class="h-5 w-5  fill-red-500"
							fill="currentColor"
							viewBox="0 0 20 20"
							xmlns="http://www.w3.org/2000/svg"
							><path
								fill-rule="evenodd"
								d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
								clip-rule="evenodd"
							/></svg
						>
						<p class="text-sm text-red-300 ">{ErrorMsg.UserName[1]}</p>
					</div>
				{:else}
					<!-- else content here -->
				{/if}
			</div>
			<input
				bind:value={RegisterData['UserName']}
				type="text"
				class=" mx-4 mt-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<div class="flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Email :
				</div>
				{#if ErrorMsg.Email[0]}
					<div class=" mt-2 ml-3 inline-flex ">
						<svg
							class="h-5 w-5  fill-red-500"
							fill="currentColor"
							viewBox="0 0 20 20"
							xmlns="http://www.w3.org/2000/svg"
							><path
								fill-rule="evenodd"
								d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
								clip-rule="evenodd"
							/></svg
						>
						<p class="text-sm text-red-300">{ErrorMsg.Email[1]}</p>
					</div>
				{:else}
					<!-- else content here -->
				{/if}
			</div>
			<input
				bind:value={RegisterData['Email']}
				type="text"
				class=" mx-4 mt-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<div class="flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Password :
				</div>
				{#if ErrorMsg.Password[0]}
					<div class=" mt-2 ml-3 inline-flex ">
						<svg
							class="h-5 w-5  fill-red-500"
							fill="currentColor"
							viewBox="0 0 20 20"
							xmlns="http://www.w3.org/2000/svg"
							><path
								fill-rule="evenodd"
								d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
								clip-rule="evenodd"
							/></svg
						>
						<p class="text-sm text-red-300">{ErrorMsg.Password[1]}</p>
					</div>
				{:else}
					<!-- else content here -->
				{/if}
			</div>
			<input
				bind:value={RegisterData['Password']}
				type="password"
				class=" mx-4 mt-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<div class="mt-8 ml-4 text-center text-xs font-semibold text-gray-500">
				By registering, you agree to Accord's <p
					class="inline cursor-pointer text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500"
				>
					Terms of Service
				</p>
				and
				<p
					class="inline cursor-pointer text-sm font-bold text-sky-600 hover:text-sky-300 active:text-sky-500"
				>
					privacy Policy
				</p>
			</div>

			<button
				on:click={OnSubmit}
				class=" m-4 h-12 w-60 self-center rounded-xl bg-sky-700 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
			>
				Register
			</button>
			<div class=" flex flex-row gap-1 pt-2 ">
					<!-- google -->
					<Google />
					<!-- Github -->
					<Github />
					<!-- Amazon -->
					<Amazon />
					<!-- Twitter -->
					<Twitter />
			</div>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
	input:-webkit-autofill,
input:-webkit-autofill:hover, 
input:-webkit-autofill:focus, 
input:-webkit-autofill:active{
    -webkit-box-shadow: 0 0 0 30px #303338 inset !important;
}

</style>
