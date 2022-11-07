<script lang="ts">
	// your script goes here
	import Coin2 from '$lib/svgs/coin2.svelte';

	let UserReqList = [
		{
			UserID: 'skraka',
			Coin: 28.23,
			ReqList: [
				{ Amount: 245, JWT: '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', Status: 'pending' },
				{ Amount: 235, JWT: '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', Status: 'accepted' },
				{ Amount: 523, JWT: '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', Status: 'pending' }
			]
		},
		{
			UserID: 'linux',
			Coin: 78.23,
			ReqList: [
				{ Amount: 245, JWT: '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', Status: 'accepted' },
				{ Amount: 36, JWT: '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', Status: 'pending' },
				{ Amount: 36, JWT: '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', Status: 'pending' }
			]
		},
		{
			UserID: 'whatsup',
			Coin: 72.23,
			ReqList: [
				{ Amount: 36, JWT: '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', Status: 'pending' },
				{ Amount: 36, JWT: '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', Status: 'pending' },
				{ Amount: 36, JWT: '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', Status: 'pending' }
			]
		}
	];
</script>

<!-- markup (zero or more items) goes here -->

<div class="flex flex-col transition-all duration-200 ease-linear  w-[80%]  self-center">
	<p class="">All User Money request</p>
	<div class="flex flex-col">
		{#each UserReqList as req}
			<!-- content here -->
			<p class="">UserID : {req.UserID}</p>
			<div class="flex flex-row flex-wrap gap-2 ">
				{#each req.ReqList as i}
					<!-- content here -->
					<div
						class=" flex h-44 w-72 flex-col justify-center gap-1 rounded-lg border-2 border-slate-400 p-3"
					>
						<div class=" flex flex-row items-center gap-2 ">
							<Coin2 class="h-8 w-8 stroke-yellow-400" />
							<p class=" text-xl">{i.Amount}</p>
						</div>
						<p
							class=" cursor-pointer break-words text-sm line-clamp-4 hover:text-blue-400 active:text-blue-600"
						>
							{i.JWT}
						</p>
                        
						<button
							class="h-10 w-fit self-center rounded-xl { i.Status==="accepted" ? "bg-green-500 hover:bg-green-600 active:bg-green-700" : " bg-yellow-400 hover:bg-yellow-600 active:bg-yellow-700"  }  p-2 px-4 font-semibold text-gray-700 "
							>{i.Status}</button
						>
					</div>
				{/each}
			</div>
		{/each}
	</div>
</div>

<style>
	/* your styles go here */
</style>
