<script context="module" lang="ts">
	export async function load({session}:any){
		// if (!session.user.authenticated){
        //     return {
        //         status:302,
        //         redirect:"/login"
        //     }
        // }
		return {}
	}
</script>
<script lang="ts">
	import { goto } from "$app/navigation";


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

            let res = await fetch('http://localhost:50000/login', {
            // credentials: 'same-origin',
			
			method: 'POST',
			mode: 'no-cors',
			credentials: 'include',
			body: JSON.stringify({
				Username: 'raka',
				Password: 'password1'
			})
		});
		console.log('🚀 ~ file: index@blank.svelte ~ line 45 ~ OnSubmit ~ res', res);
		if (res.ok) {
			// let value= await res.json()
			console.log('🚀 ~ file: index.svelte ~ line 49 ~ OnSubmit ~ value : ', 200);
			// LoginData = {
			// 	Email: '',
			// 	Password: ''
			// };
			// let s:string="/"+String(value.ID)
            
			goto("/skraka");
		}

		
			// let res =await fetch('/auth/login', {
			// 	credentials: 'same-origin',
			// 	method: 'POST',
			// 	mode: 'cors',
			// 	body: JSON.stringify(LoginData)
			// })
			
			// if (res.ok){
			// 	let value= await res.json()
            //     console.log("🚀 ~ file: index.svelte ~ line 49 ~ OnSubmit ~ value : ", value)
			// 	LoginData = {
			// 		Email: '',
			// 		Password: ''
			// 	};
			// 	let s:string="/"+String(value.ID)
                 
			// 	goto(s);
			// }else{
			// 	console.log("🚀 ~ file: index.svelte ~ line 57 ~ OnSubmit ~ err : ")
			// 	ErrorMsg.Email[1] = 'invalid/wrong email or password';
			// 	ErrorMsg.Email[0] = true;
			// }
			
			
		}
	}
</script>

<div class="grid h-screen w-full  place-items-center bg-[#181818] ">
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
		<div class=" flex h-[420px] w-[500px] flex-col space-y-4 rounded-xl bg-[#2d2d2d]">
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
				bind:value={LoginData['Email']}
				type="text"
				class=" mx-4 mt-2 h-12 w-[415px] self-center rounded-2xl border-2 border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
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
				class=" m-4 h-12 w-60  self-center rounded-xl bg-sky-700 text-xl font-bold text-gray-300 hover:bg-sky-600 active:bg-sky-500"
			>
				Login
			</button>
		</div>
		<div class=" flex flex-row gap-1 pt-2 ">
			<svg
				class="h-16 w-full rounded-2xl px-10 py-3 hover:border-2   hover:border-sky-700"
				enable-background="new 0 0 512 512"
				id="Layer_1"
				version="1.1"
				viewBox="0 0 512 512"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><g
					><path
						d="M42.4,145.9c15.5-32.3,37.4-59.6,65-82.3c37.4-30.9,80.3-49.5,128.4-55.2c56.5-6.7,109.6,4,158.7,33.4   c12.2,7.3,23.6,15.6,34.5,24.6c2.7,2.2,2.4,3.5,0.1,5.7c-22.3,22.2-44.6,44.4-66.7,66.8c-2.6,2.6-4,2.4-6.8,0.3   c-64.8-49.9-159.3-36.4-207.6,29.6c-8.5,11.6-15.4,24.1-20.2,37.7c-0.4,1.2-1.2,2.3-1.8,3.5c-12.9-9.8-25.9-19.6-38.7-29.5   C72.3,169,57.3,157.5,42.4,145.9z"
						fill="#E94335"
					/><path
						d="M126,303.8c4.3,9.5,7.9,19.4,13.3,28.3c22.7,37.2,55.1,61.1,97.8,69.6c38.5,7.7,75.5,2.5,110-16.8   c1.2-0.6,2.4-1.2,3.5-1.8c0.6,0.6,1.1,1.3,1.7,1.8c25.8,20,51.7,40,77.5,60c-12.4,12.3-26.5,22.2-41.5,30.8   c-43.5,24.8-90.6,34.8-140.2,31C186.3,501.9,133,477.5,89,433.5c-19.3-19.3-35.2-41.1-46.7-66c10.7-8.2,21.4-16.3,32.1-24.5   C91.6,329.9,108.8,316.9,126,303.8z"
						fill="#34A853"
					/><path
						d="M429.9,444.9c-25.8-20-51.7-40-77.5-60c-0.6-0.5-1.2-1.2-1.7-1.8c8.9-6.9,18-13.6,25.3-22.4   c12.2-14.6,20.3-31.1,24.5-49.6c0.5-2.3,0.1-3.1-2.2-3c-1.2,0.1-2.3,0-3.5,0c-40.8,0-81.7-0.1-122.5,0.1c-4.5,0-5.5-1.2-5.4-5.5   c0.2-29,0.2-58,0-87c0-3.7,1-4.7,4.7-4.7c74.8,0.1,149.6,0.1,224.5,0c3.2,0,4.5,0.8,5.3,4.2c6.1,27.5,5.7,55.1,2,82.9   c-3,22.2-8.4,43.7-16.7,64.5c-12.3,30.7-30.4,57.5-54.2,80.5C431.6,443.8,430.7,444.3,429.9,444.9z"
						fill="#4285F3"
					/><path
						d="M126,303.8c-17.2,13.1-34.4,26.1-51.6,39.2c-10.7,8.1-21.4,16.3-32.1,24.5C34,352.1,28.6,335.8,24.2,319   c-8.4-32.5-9.7-65.5-5.1-98.6c3.6-26,11.1-51,23.2-74.4c15,11.5,29.9,23.1,44.9,34.6c12.9,9.9,25.8,19.7,38.7,29.5   c-2.2,10.7-5.3,21.2-6.3,32.2c-1.8,20,0.1,39.5,5.8,58.7C125.8,301.8,125.9,302.8,126,303.8z"
						fill="#FABB06"
					/></g
				></svg
			>
			<!-- <svg class="h-12 w-12" xmlns="http://www.w3.org/2000/svg" width="12" height="12" focusable="false" viewBox="0 0 12 12"> <path fill="currentColor" d="M6 0a6 6 0 110 12A6 6 0 016 0zm0 .98C3.243.98 1 3.223 1 6a5.02 5.02 0 003.437 4.77.594.594 0 00.045.005c.203.01.279-.129.279-.25l-.007-.854c-1.39.303-1.684-.674-1.684-.674-.227-.58-.555-.734-.555-.734-.454-.312.034-.306.034-.306.365.026.604.288.708.43l.058.088c.446.767 1.17.546 1.455.418.046-.325.174-.546.317-.672-1.11-.127-2.277-.558-2.277-2.482 0-.548.195-.996.515-1.348l-.03-.085c-.064-.203-.152-.658.079-1.244l.04-.007c.124-.016.548-.013 1.335.522A4.77 4.77 0 016 3.408c.425.002.853.058 1.252.17.955-.65 1.374-.516 1.374-.516.272.692.1 1.202.05 1.33.32.35.513.799.513 1.347 0 1.93-1.169 2.354-2.283 2.478.18.155.34.462.34.93l-.006 1.378c0 .13.085.282.323.245A5.02 5.02 0 0011 6C11 3.223 8.757.98 6 .98z"/> </svg> -->
			<svg
				class="h-16 w-full rounded-2xl px-10 py-3 hover:border-2  hover:border-sky-700"
				enable-background="new 0 0 128 128"
				id="Social_Icons"
				version="1.1"
				viewBox="0 0 128 128"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><g id="_x31__stroke"
					><g id="Github_1_"
						><rect
							clip-rule="evenodd"
							fill="none"
							fill-rule="evenodd"
							height="128"
							width="128"
						/><path
							clip-rule="evenodd"
							d="M63.996,1.333C28.656,1.333,0,30.099,0,65.591    c0,28.384,18.336,52.467,43.772,60.965c3.2,0.59,4.368-1.394,4.368-3.096c0-1.526-0.056-5.566-0.088-10.927    c-17.804,3.883-21.56-8.614-21.56-8.614c-2.908-7.421-7.104-9.397-7.104-9.397c-5.812-3.988,0.44-3.907,0.44-3.907    c6.42,0.454,9.8,6.622,9.8,6.622c5.712,9.819,14.98,6.984,18.628,5.337c0.58-4.152,2.236-6.984,4.064-8.59    c-14.212-1.622-29.152-7.132-29.152-31.753c0-7.016,2.492-12.75,6.588-17.244c-0.66-1.626-2.856-8.156,0.624-17.003    c0,0,5.376-1.727,17.6,6.586c5.108-1.426,10.58-2.136,16.024-2.165c5.436,0.028,10.912,0.739,16.024,2.165    c12.216-8.313,17.58-6.586,17.58-6.586c3.492,8.847,1.296,15.377,0.636,17.003c4.104,4.494,6.58,10.228,6.58,17.244    c0,24.681-14.964,30.115-29.22,31.705c2.296,1.984,4.344,5.903,4.344,11.899c0,8.59-0.08,15.517-0.08,17.626    c0,1.719,1.152,3.719,4.4,3.088C109.68,118.034,128,93.967,128,65.591C128,30.099,99.344,1.333,63.996,1.333"
							fill="#3E75C3"
							fill-rule="evenodd"
							id="Github"
						/></g
					></g
				></svg
			>
			<svg
				class="h-16 w-full rounded-2xl px-10   py-3 hover:border-2 hover:border-sky-700"
				enable-background="new 0 0 128 128"
				id="Social_Icons"
				version="1.1"
				viewBox="0 0 128 128"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><g id="_x31__stroke"
					><g id="Amazon_1_"
						><rect clip-rule="evenodd" fill="none" fill-rule="evenodd" height="128" width="128" /><g
							id="Amazon"
							><path
								clip-rule="evenodd"
								d="M70.849,35.607c-3.673,0.279-7.917,0.557-12.154,1.114     c-6.495,0.858-12.99,1.988-18.348,4.533c-10.453,4.239-17.52,13.282-17.52,26.556c0,16.685,10.732,25.156,24.301,25.156     c4.516,0,8.188-0.572,11.567-1.408c5.381-1.702,9.889-4.811,15.255-10.466c3.108,4.239,3.966,6.227,9.325,10.744     c1.415,0.557,2.83,0.557,3.951-0.279c3.394-2.831,9.332-7.921,12.433-10.744c1.415-1.129,1.136-2.831,0.278-4.232     c-3.101-3.96-6.216-7.356-6.216-14.984V36.164c0-10.737,0.858-20.631-7.052-27.972C80.173,2.266,69.998,0,62.088,0h-3.394     C44.297,0.836,29.05,7.055,25.648,24.862c-0.564,2.266,1.136,3.11,2.265,3.388l15.819,1.98c1.693-0.286,2.544-1.702,2.822-3.102     c1.407-6.219,6.495-9.329,12.139-9.901h1.136c3.394,0,7.067,1.416,9.039,4.247c2.258,3.388,1.979,7.913,1.979,11.874V35.607z      M67.74,69.225c-1.979,3.953-5.373,6.498-9.046,7.356c-0.557,0-1.407,0.279-2.258,0.279c-6.209,0-9.881-4.804-9.881-11.866     c0-9.05,5.366-13.282,12.139-15.27c3.673-0.843,7.917-1.129,12.154-1.129v3.396C70.849,58.488,71.127,63.571,67.74,69.225z"
								fill="#343B45"
								fill-rule="evenodd"
							/><path
								clip-rule="evenodd"
								d="M109.464,103.642c-0.64-0.008-1.294,0.143-1.904,0.429     c-0.685,0.271-1.385,0.587-2.047,0.866l-0.971,0.407l-1.264,0.504v0.015c-13.735,5.572-28.162,8.84-41.513,9.126     c-0.489,0.015-0.986,0.015-1.46,0.015c-20.997,0.015-38.126-9.728-55.405-19.328c-0.602-0.316-1.227-0.482-1.829-0.482     c-0.775,0-1.573,0.294-2.152,0.836C0.339,96.58-0.007,97.37,0,98.176c-0.008,1.047,0.557,2.01,1.347,2.635     C17.565,114.899,35.342,127.985,59.251,128c0.467,0,0.941-0.015,1.415-0.023c15.21-0.339,32.406-5.481,45.757-13.869l0.083-0.053     c1.746-1.047,3.492-2.236,5.14-3.554c1.024-0.76,1.731-1.95,1.731-3.185C113.332,105.126,111.473,103.642,109.464,103.642z      M127.985,95.857v-0.008c-0.06-1.333-0.339-2.349-0.896-3.192l-0.06-0.083l-0.068-0.083c-0.564-0.617-1.106-0.851-1.693-1.107     c-1.754-0.678-4.305-1.039-7.375-1.047c-2.205,0-4.636,0.211-7.082,0.745l-0.008-0.166l-2.461,0.821l-0.045,0.023l-1.392,0.452     v0.06c-1.633,0.678-3.116,1.521-4.493,2.522c-0.858,0.64-1.565,1.491-1.603,2.793c-0.023,0.708,0.339,1.521,0.933,2.003     c0.595,0.482,1.287,0.648,1.896,0.648c0.143,0,0.278-0.008,0.399-0.03l0.12-0.008l0.09-0.015     c1.204-0.256,2.958-0.429,5.012-0.715c1.761-0.196,3.627-0.339,5.246-0.339c1.144-0.008,2.175,0.075,2.882,0.226     c0.354,0.075,0.617,0.166,0.76,0.241c0.053,0.015,0.09,0.038,0.113,0.053c0.03,0.098,0.075,0.354,0.068,0.708     c0.015,1.355-0.557,3.87-1.347,6.325c-0.768,2.455-1.701,4.917-2.318,6.551c-0.151,0.376-0.248,0.791-0.248,1.242     c-0.015,0.655,0.256,1.453,0.828,1.98c0.557,0.527,1.279,0.738,1.881,0.738h0.03c0.903-0.008,1.671-0.369,2.333-0.888     c6.246-5.617,8.421-14.592,8.512-19.644L127.985,95.857z"
								fill="#FF9A00"
								fill-rule="evenodd"
							/></g
						></g
					></g
				></svg
			>
			<svg
				class="h-16 w-full rounded-2xl px-10   py-3 hover:border-2 hover:border-sky-700"
				enable-background="new 0 0 128 128"
				id="Social_Icons"
				version="1.1"
				viewBox="0 0 128 128"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><g id="_x37__stroke"
					><g id="Twitter"
						><rect
							clip-rule="evenodd"
							fill="none"
							fill-rule="evenodd"
							height="128"
							width="128"
						/><path
							clip-rule="evenodd"
							d="M128,23.294    c-4.703,2.142-9.767,3.59-15.079,4.237c5.424-3.328,9.587-8.606,11.548-14.892c-5.079,3.082-10.691,5.324-16.687,6.526    c-4.778-5.231-11.608-8.498-19.166-8.498c-14.493,0-26.251,12.057-26.251,26.927c0,2.111,0.225,4.16,0.676,6.133    C41.217,42.601,21.871,31.892,8.91,15.582c-2.261,3.991-3.554,8.621-3.554,13.552c0,9.338,4.636,17.581,11.683,22.412    c-4.297-0.131-8.355-1.356-11.901-3.359v0.331c0,13.051,9.053,23.937,21.074,26.403c-2.201,0.632-4.523,0.948-6.92,0.948    c-1.69,0-3.343-0.162-4.944-0.478c3.343,10.694,13.035,18.483,24.53,18.691c-8.986,7.227-20.315,11.533-32.614,11.533    c-2.119,0-4.215-0.123-6.266-0.37c11.623,7.627,25.432,12.088,40.255,12.088c48.309,0,74.717-41.026,74.717-76.612    c0-1.171-0.023-2.342-0.068-3.49C120.036,33.433,124.491,28.695,128,23.294"
							fill="#00AAEC"
							fill-rule="evenodd"
							id="Twitter_1_"
						/></g
					></g
				></svg
			>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
