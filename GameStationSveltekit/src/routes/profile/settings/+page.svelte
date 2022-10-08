<script lang="ts">
	import Twitter from '$lib/foot/icons/twitter.svelte';
	import ContractAdimin from '$lib/settingsCom/ContractAdimin.svelte';
	import ContractGameDev from '$lib/settingsCom/ContractGameDev.svelte';
	import General from '$lib/settingsCom/General.svelte';
	import PassSec from '$lib/settingsCom/Pass&Sec.svelte';
	import PaymentManagement from '$lib/settingsCom/PaymentManagement.svelte';
	import RequestToken from '$lib/settingsCom/RequestToken.svelte';
	import Wishlist from '$lib/settingsCom/wishlist.svelte';
	import CameraPlus from '$lib/svgs/cameraPlus.svelte';
	import Coin from '$lib/svgs/coin2.svelte';
	import List from '$lib/svgs/list.svelte';
	import Message from '$lib/svgs/message.svelte';
	import Paycard from '$lib/svgs/paycard.svelte';
	import Security from '$lib/svgs/security.svelte';
	import Transaction from '$lib/svgs/transaction.svelte';
	import User from '$lib/svgs/user.svelte';

	let list = [
		{ name: 'General', link: '', icon: User },
		{ name: 'Payment Mangement', link: '', icon: Paycard },
		{ name: 'Request Token', link: '', icon: Coin },
		{ name: 'Password & Security', link: '', icon: Security },
		{ name: 'Transaction', link: '', icon: Transaction },
		{ name: 'Wishlist', link: '', icon: List },
		{ name: 'Contact Admin', link: '', icon: Message },
		{ name: 'Contact Game Developer', link: '', icon: Message }
	];
	let active = list[0].name;
</script>

<!-- markup (zero or more items) goes here -->

<div class="flex flex-row justify-center gap-1">

	<div class="flex flex-col gap-0.5 xs:hidden  ">
		{#each list as l}
			<div
				on:click={() => {
					active = l.name;
				}}
				class=" flex h-14 w-72 flex-row justify-center gap-2   rounded-md bg-gray-700 hover:cursor-pointer hover:bg-opacity-80  {active ===
				l.name
					? 'border-x-4 bg-opacity-60  border-blue-500 '
					: ''} items-center "
			>
				<svelte:component this={l.icon} class=" h-6 w-6 stroke-slate-200 " />
				<p class=" font-Poppins text-slate-200 hover:text-white">{l.name}</p>
			</div>
		{/each}
	</div>
	<div
		class=" flex h-fit min-h-[300px] w-[880px] flex-col rounded-lg  bg-slate-700 bg-opacity-30 font-Poppins"
	>
		{#if active === 'General'}
			<General />
			
		{:else if active === 'Payment Mangement'}
			<PaymentManagement />
		{:else if active === 'Request Token'}
			<RequestToken />
		{:else if active ==='Password & Security'}
			<PassSec />
		{:else if  active ==='Transaction'}
			<Transaction />
		{:else if active ==='Wishlist'}
			<Wishlist />
		{:else if  active ==='Contact Admin'}
			<ContractAdimin />
		{:else if active ==='Contact Game Developer'}
			<ContractGameDev />
		{:else}
			<!-- else content here -->
		{/if}
	</div>
</div>

<style>
	/* your styles go here */
</style>
