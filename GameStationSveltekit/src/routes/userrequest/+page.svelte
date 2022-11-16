<script lang="ts">
	// your script goes here
	import Coin2 from '$lib/svgs/coin2.svelte';
	import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';
	import Nop from '$lib/logo/nop.svelte';

	export let data: PageData;

	let { Userdata, UserReqList } = data;
	console.log('🚀 ~ file: +page.svelte ~ line 10 ~ Userdata', Userdata);
	
	UserProData.update((d) => (d = Userdata));

	// let UserReqList = [
	// 	{
	// 		UserID: 'skraka',
	// 		Coin: 28.23,
	// 		ReqList: [
	// 			{ Amount: 245, JWT: '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', Status: 'pending' },
	// 			{ Amount: 235, JWT: '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', Status: 'accepted' },
	// 			{ Amount: 523, JWT: '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', Status: 'pending' }
	// 		]
	// 	},
	// 	{
	// 		UserID: 'linux',
	// 		Coin: 78.23,
	// 		ReqList: [
	// 			{ Amount: 245, JWT: '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', Status: 'accepted' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' }
	// 		]
	// 	},
	// 	{
	// 		UserID: 'whatsup',
	// 		Coin: 72.23,
	// 		ReqList: [
	// 			{ Amount: 36, JWT: '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', Status: 'pending' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' },
	// 			{ Amount: 36, JWT: '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', Status: 'pending' }
	// 		]
	// 	}
	// ];
	async function AcceptReq(Amount:number,UserID:string) {
		console.log("🚀 ~ file: +page.svelte ~ line 45 ~ AcceptReq ~ UserID", UserID)
		console.log("🚀 ~ file: +page.svelte ~ line 44 ~ AcceptReq ~ Amount", Amount)
		let res = await fetch('/api/profile/moneyreqaccept', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ 
				UserID: UserID,
				Amount: Amount,
			 })
		});
		let data = await res.json();
		console.log('🚀 ~ file: +page.svelte ~ line 45 ~ AcceptReq ~ data', data);
	}
</script>

<!-- markup (zero or more items) goes here -->

<div class="flex w-[80%] flex-col self-center transition-all  duration-200  ease-linear">
	<p class="self-center text-3xl text-white ">All User Money request</p>
	<div class="flex flex-col gap-2">
		{#each UserReqList as req}
			
			<p class=" text-xl text-white ">UserID : {req.UserID}</p>
			<div class="flex flex-row flex-wrap gap-2 ">
				{#if req.ReqList.length === 0}
					
					<Nop class=" ml-10 h-20 w-20" />
				{:else}
					
					{#each req.ReqList as i}
					
						<div
							class=" flex h-44 w-72 flex-col justify-center gap-1 rounded-lg border-2 border-slate-400 p-3"
						>
							<div class=" flex flex-row items-center gap-2 ">
								<Coin2 class="h-8 w-8 stroke-yellow-400" />
								<p class=" text-xl text-white">{i.Amount}</p>
							</div>
							<p
								class=" cursor-pointer break-words text-sm text-slate-300 line-clamp-4 hover:text-blue-400 active:text-blue-600"
							>
								{i.JWT}
							</p>

							<button
							on:click={()=>{AcceptReq(i.Amount ,req.UserID)}}
								class="h-10 w-fit self-center rounded-xl {i.Status === 'accepted'
									? 'bg-green-500 hover:bg-green-600 active:bg-green-700 text-gray-700'
									: i.Status === 'used' ? ' bg-transparent text-slate-400 border-2 border-slate-400 ' : ' bg-yellow-400 hover:bg-yellow-600 active:bg-yellow-700 text-gray-700'}  p-2 px-4 font-semibold  "
								>{i.Status}</button
							>
						</div>
					{/each}
				{/if}
			</div>
		{/each}
	</div>
</div>

<style>
	/* your styles go here */
</style>
