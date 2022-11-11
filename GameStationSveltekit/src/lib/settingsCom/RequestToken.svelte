<script lang="ts">
	// your script goes here

	import Coin2 from '$lib/svgs/coin2.svelte';
	import { UserProData } from '$lib/Store/store';
	import Nop from '$lib/logo/nop.svelte';
	let AllowedReq = [10, 20, 50, 100, 200, 500, 1000, 10000];

	// let MoneyReqList2 = [
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
	// 	},
	// 	{ Amount: '10', Details: 'Your Request is Pending for Admin Acceptance', Status: 'pending' },
	// 	{
	// 		Amount: '100',
	// 		Details: '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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',
	// 		Status: 'accepted'
	// 	},
	// 	{ Amount: '100', Details: 'Your Request is Pending for Admin Acceptance', Status: 'pending' },
	// 	{ Amount: '1000', Details: 'Your Request is Pending for Admin Acceptance', Status: 'pending' },
	// 	{
	// 		Amount: '10',
	// 		Details: 'ghp_Vax5DhHRGpqdZvUh3WKYvKz9Xxv73U5Bn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',
	// 		Status: 'accepted'
	// 	},
	// 	{ Amount: '10000', Details: 'Your Request is Pending for Admin Acceptance', Status: 'pending' }
	// ];

	let MoneyReqList: {
		Amount: string;
		JWT: string;
		Status: string;
	}[] = [];
	let NumOfReq = {
		accepted: 0,
		pending: 0,
		used: 0
	};
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
				console.log('🚀 ~ file: RequestToken.svelte ~ line 22 ~ requestMoney ~ data', data);
				MoneyReqList = data;
				return data;
			});
			console.log("🚀 ~ file: RequestToken.svelte ~ line 82 ~ .then ~ MoneyReqList", MoneyReqList)
		console.log('🚀 ~ file: RequestToken.svelte ~ line 26 ~ requestMoney ~ resdata', resdata);
		SeeThroughMoneyReqList();
	}
	GetReqList();
	SeeThroughMoneyReqList();
	async function requestMoney(amount: number) {
		// let resdata
		let resdata = await fetch('/api/profile/moneytokenreq', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				UserID: $UserProData.UserID,
				Amount: amount
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
		GetReqList();
		SeeThroughMoneyReqList();
	}

	
	let showUserid = {
		id: '',
		show: false
	};
</script>

<!-- markup (zero or more items) goes here -->
<div class="flex flex-col p-4 transition-all duration-200  ease-linear">
	<div class=" flex flex-col">
		<p class=" m-3 font-Poppins text-2xl text-slate-200 ">Accepted Request :</p>
		<div class="flex flex-row flex-wrap justify-around gap-2  ">
			{#if NumOfReq.accepted === 0}
				<Nop class=" h-20 w-20 stroke-slate-400  " />
			{:else}
				{#each MoneyReqList as req}
					{#if req.Status === 'accepted'}
						<div
							class=" flex h-44 w-72 flex-col justify-center gap-1 rounded-lg border-2 border-slate-400 p-3"
						>
							<div class=" flex flex-row items-center gap-2 ">
								<Coin2 class="h-8 w-8 stroke-yellow-400" />
								<p class=" text-xl font-semibold text-white">{req.Amount}</p>
							</div>
							{#if showUserid.id === req.JWT && showUserid.show}
								<!-- <p
								class=" cursor-pointer text-sm text-gray-500 line-clamp-1 hover:text-gray-400"
							>
								@{$UserProData.UserID}
							</p> -->
								<p
									on:click={() => {
										navigator.clipboard.writeText(req.JWT);
									}}
									on:keypress={() => {}}
									class="{setTimeout(() => {
										showUserid.show = false;
									}, 1000)} cursor-pointer text-2xl font-semibold text-blue-500 line-clamp-1 "
								>
									Copied !!
								</p>
							{:else}
								<p
									on:click={() => {
										navigator.clipboard.writeText(req.JWT)
										showUserid.show = true;
										showUserid.id = req.JWT;
									}}
									on:keypress={() => {}}
									class=" cursor-pointer break-words text-sm text-slate-200 line-clamp-4 hover:text-blue-400 active:text-blue-600"
								>
									{req.JWT}
								</p>
							{/if}

							<button
								class="h-10 w-fit self-center rounded-xl bg-green-500 p-2 px-4 font-semibold text-gray-700 hover:bg-green-600 active:bg-green-700 "
								>Accepted</button
							>
						</div>
					{/if}
				{:else}
					<!-- empty list -->
				{/each}
			{/if}
		</div>
	</div>
	<p class=" m-3 font-Poppins text-2xl text-slate-200">Pending Money Request :</p>
	<div class="flex flex-row flex-wrap justify-around gap-2">
		{#if NumOfReq.pending === 0}
			<Nop class=" h-20 w-20 stroke-slate-400  " />
		{:else}
			{#each MoneyReqList as req}
				<!-- content here -->
				{#if req.Status == 'pending'}
					<div
						class=" flex h-44 w-72 flex-col justify-center gap-3 rounded-lg border-2 border-slate-400 p-3 "
					>
						<div class=" flex flex-row items-center gap-2 ">
							<Coin2 class="h-8 w-8 stroke-yellow-400" />
							<p class=" text-xl font-semibold text-white">{req.Amount}</p>
						</div>
						<p class=" text-sm text-slate-300">{req.JWT}</p>
						<button
							class="h-10 w-fit self-center rounded-xl bg-yellow-400 p-2 px-4 font-semibold text-gray-700 hover:bg-yellow-500 "
							>Pending</button
						>
					</div>
				{/if}
			{:else}
				<!-- empty list -->
			{/each}
		{/if}
	</div>
	<p class=" m-3 font-Poppins text-2xl text-slate-200">Send Request to Admin for Money :</p>
	<div class="flex flex-row flex-wrap justify-around gap-2 space-y-3">
		{#each AllowedReq as MoneyReq}
			<div
				class=" flex h-44 w-72 flex-col justify-center gap-3 rounded-lg border-2 border-slate-400 p-3 "
			>
				<div class=" flex flex-row items-center gap-2 ">
					<Coin2 class="h-8 w-8 stroke-yellow-400" />
					<p class=" text-xl font-semibold text-white">{MoneyReq}</p>
				</div>
				<p class=" text-sm text-slate-300">Your Money request will send to the admin.</p>
				<button
					on:click={() => {
						requestMoney(MoneyReq);
					}}
					class="h-10 w-fit self-center rounded-xl bg-blue-500 p-2 px-4 font-semibold text-slate-200 "
					>Send Request</button
				>
			</div>
		{/each}
	</div>
</div>

<style>
	/* your styles go here */
</style>
