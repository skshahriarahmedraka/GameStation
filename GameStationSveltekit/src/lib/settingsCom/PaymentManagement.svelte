<script lang="ts">
	import Coin2 from '$lib/svgs/coin2.svelte';
	import { UserProData } from '$lib/Store/store';
	import Nop from '$lib/logo/nop.svelte';
	// your script goes here
	let MoneyReqList : {
		Amount: string;
		JWT: string;
		Status: string;
	}[] = [];
	// = [
	// 	{ Amount: '10', Details: 'Your Request is Pending for Admin Acceptance', Status: 'pending' },
	// 	{
	// 		Amount: '1000',
	// 		Details: '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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',
	// 		Status: 'used'
	// 	},
	// 	{ Amount: '10', Details: 'Your Request is Pending for Admin Acceptance', Status: 'used' },
	// 	{ Amount: '1000', Details: 'Your Request is Pending for Admin Acceptance', Status: 'pending' },
	// 	{
	// 		Amount: '10',
	// 		Details: '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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',
	// 		Status: 'used'
	// 	}
	// ];

	let inputJWT = '';

	let ErrorMsg = {
		Err: false,
		Msg: ''
	};
	let NumOfReq = {
		accepted: 0,
		pending: 0,
		used: 0
	};
	$: console.log("🚀 ~ file: PaymentManagement.svelte ~ line 38 ~ NumOfReq", NumOfReq)
	function SeeThroughMoneyReqList() {
		NumOfReq = {
			accepted: 0,
			pending: 0,
			used: 0
		};
		for (let i of MoneyReqList) {
			if (i.Status === 'accepted') {
				NumOfReq.accepted++;
			} else if (i.Status === 'pending') {
				NumOfReq.pending++;
			} else if (i.Status === 'used') {
				NumOfReq.used++;
			}
		}
	}
	async function RechargeWallet() {
		console.log('🚀 ~ file: PaymentManagement.svelte ~ line 30 ~ SubmitReq ~ inputJWT', inputJWT);
		if (inputJWT === '') {
			ErrorMsg.Err = true;
			ErrorMsg.Msg = 'Please Enter Your JWT';
			return;
		}

		let resdata = await fetch('/api/profile/rechargewallet', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				UserID: $UserProData.UserID,
				JWT: inputJWT
			})
		})
			.then((res) => {
				return res.json();
			})
			.then((data) => {
				console.log('🚀 ~ file: RequestToken.svelte ~ line 22 ~ requestMoney ~ data', data);

				return data;
			});
		console.log('🚀 ~ file: RequestToken.svelte ~ line 26 ~ requestMoney ~ resdata', resdata);
		// let data = await response.json()
		SeeThroughMoneyReqList()
	}
	async function GetReqList() {
		let resdata = await fetch('/api/profile/moneytokenreqlist', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				UserID: $UserProData.UserID
			})
		})
			.then((res) => {
				return res.json();
			})
			.then((data) => {
				MoneyReqList = data;
				return data;
			});
	
		SeeThroughMoneyReqList()
	
	}
	GetReqList()
	SeeThroughMoneyReqList()
</script>

<!-- markup (zero or more items) goes here -->
<div class="flex flex-col p-4 transition-all duration-200  ease-linear">
	<div class=" flex flex-col items-center">
		<p class=" m-3 font-Poppins text-2xl text-slate-200">Add Money to Wallet :</p>
		<div class=" flex h-56 w-[80%] flex-col rounded-lg border-2 border-slate-200">
			<div class="flex flex-row">
				<p class="m-6 font-Poppins text-base text-slate-200">Give your Token to Recharge Money</p>
				<!-- <div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Password :
				</div> -->
				{#if ErrorMsg.Err}
					<div class=" m-6 inline-flex ">
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
						<p class="text-sm  text-red-300">{ErrorMsg.Msg}</p>
					</div>
				{:else}
					<!-- else content here -->
				{/if}
			</div>
			<input
				bind:value={inputJWT}
				type="text"
				class=" mx-4 my-2 h-12 w-[95%] self-center rounded-2xl border-2  border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<button
				on:click={() => {
					RechargeWallet();
				}}
				class="m-3 h-10 w-fit self-center rounded-xl bg-blue-500 p-2 px-4 font-semibold text-white hover:bg-green-600 active:bg-green-700 "
				>Recharge</button
			>
		</div>
	</div>
	<div class=" flex flex-col">
		<p class=" m-3 font-Poppins text-2xl text-slate-200">Used Request :</p>
		<div class="flex flex-row flex-wrap justify-around gap-2 ">
			{#if NumOfReq.used === 0}
				<Nop class=" h-20 w-20 stroke-slate-400  " />
			{:else}

			{#each MoneyReqList as req}
				<!-- content here -->
				{#if req.Status === 'used'}
					<!-- content here -->
					<div
						class=" flex h-44 w-72 flex-col justify-center gap-1 rounded-lg border-2 border-slate-400 p-3"
					>
						<div class=" flex flex-row items-center gap-2 ">
							<Coin2 class="h-8 w-8 stroke-yellow-400" />
							<p class=" text-xl font-semibold text-white">{req.Amount}</p>
						</div>
						<p
							class=" cursor-pointer break-words text-sm text-slate-300 line-clamp-4 hover:text-blue-400 active:text-blue-600"
						>
							{req.JWT}
						</p>
						<button
							class="h-10 w-fit self-center rounded-xl bg-transparent text-slate-400 border-2 border-slate-400  p-2 px-4 font-semibold  "
							>Used</button
						>
					</div>
				{/if}
				{:else}
				<!-- empty list -->
				{/each}
				{/if}
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
