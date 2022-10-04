
<script lang="ts">
import { goto } from "$app/navigation";
	import Amazon from "$lib/svgs/Amazon.svelte";
	import Github from "$lib/svgs/Github.svelte";
	import Google from "$lib/svgs/google.svelte";
	import Twitter from "$lib/svgs/twitter.svelte";
	import Bg1 from "$lib/images/mix3.jpeg"
	const LoginBg = new URL(Bg1, import.meta.url).href;


	let LoginData = {
		Email: '',
		Password: ''
	};
	let b: boolean = false;
	let ErrorMsg = {
		
		Email: [false, 'invalid/wrong email'],
		Password: [false, 'Password is too short']
	};
	// let SuccessMsg:string ="Success"
	const ValidateEmail = (E: string) => {
		return String(E)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	async function OnSubmit() {
		if (LoginData.Email.trim() === '') {
			ErrorMsg.Email[1] = 'Email is Empty';
			ErrorMsg.Email[0] = true;
		} else if (!ValidateEmail(LoginData.Email)) {
			ErrorMsg.Email[1] = 'invalid/wrong email or password';
			ErrorMsg.Email[0] = true;
		} 
		if (LoginData.Password.trim() === '') {
			ErrorMsg.Password[1] = 'Password is Empty';
			ErrorMsg.Password[0]=true
		} else if (LoginData.Password.length<8 || LoginData.Password.length>30){
			ErrorMsg.Password[1] = 'Password is Too short';
			ErrorMsg.Password[0]=true
		}
		 else {
			// await fetch("http://localhost:8080/login", {
			let res =await fetch('/auth/login', {
				credentials: 'same-origin',
				method: 'POST',
				mode: 'cors',
				body: JSON.stringify(LoginData)
			})
			
			if (res.ok){
				let value= await res.json()
                console.log("🚀 ~ file: index.svelte ~ line 49 ~ OnSubmit ~ value : ", value)
				LoginData = {
					Email: '',
					Password: ''
				};
				let s:string="/"+String(value.ID)
                 
				goto(s);
			}else{
				console.log("🚀 ~ file: index.svelte ~ line 57 ~ OnSubmit ~ err : ")
				ErrorMsg.Email[1] = 'invalid/wrong email or password';
				ErrorMsg.Email[0] = true;
			}
			
			// .then(()=>).then(response=>response.json()).then((value) => {
            //     // const UserURL = response.json()
            //     // console.log("🚀🚀🚀🚀 ~ file: index.svelte ~ line 37 ~ OnSubmit ~ UserURL : ", value.UserURL)
            //     // ErrorMsg = 'Success';
            //     // goto(`${value.UserURL}`)
			// 	// goto("/")
				
			// 	// b = true;
			// }).catch((err)=>{
                
			// });
		}
	}
</script>
<!-- bg-[#181818] -->

<div class="grid h-screen w-full  place-items-center  " style="background-image: url('{LoginBg}');" >
	<a
		href="/register"
		on:click={() => {
			LoginData['Email'] = '';
			LoginData['Password'] = '';
		}}
		class="absolute  top-0 right-0  m-10"
	>
		<button
			class="  h-12 w-60 rounded-3xl bg-sky-600 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
		>
			Register
		</button>
	</a>
	<div class="">
		<div class=" flex h-fit pb-6 w-[500px] flex-col space-y-4 rounded-xl bg-[#2d2d2d] bg-opacity-90">
			<div
				class="mt-5 flex w-full items-center justify-center text-3xl font-semibold text-gray-200"
			>
				Sign In
			</div>
			<div class=" flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Email :
				</div>
				{#if ErrorMsg.Email[0] }
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
			bind:value={LoginData.Email}
				type="text"
				class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "
			/>
			<div class="flex flex-row">
				<div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Password :
				</div>
				{#if ErrorMsg.Password[0] }
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
						<p class="text-red-300">{ErrorMsg.Password[1]}</p>
					</div>
				{/if}
			</div>
			<input
				bind:value={LoginData['Password']}
				type="password"
				class=" self-center h-12 w-[415px] p-2 text-lg font-medium text-[#98999e] outline-none focus:border-sky-500  bg-[#303338] border-2 mx-4 my-2 border-[#24262b]  active:border-gray-800 rounded-2xl "

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
				class=" m-4 h-12 w-60  self-center rounded-xl bg-sky-700 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
			>
				Login
			</button>
			<div class=" flex flex-row gap-1 pt-2  ">
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
