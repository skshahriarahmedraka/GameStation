<script lang="ts">
	import Coin2 from '$lib/svgs/coin2.svelte';
	import { UserProData } from '$lib/Store/store';
	import Nop from '$lib/logo/nop.svelte';

</script>

<!-- markup (zero or more items) goes here -->
<div class="flex flex-col p-4 transition-all duration-200  ease-linear">
	<div class=" flex flex-col items-center">
		<p class=" m-3 font-Poppins text-2xl text-slate-200">Reset Your Password :</p>
		<div class=" flex h-56 w-[80%] flex-col rounded-lg border-2 border-slate-200">
			<div class="flex flex-row">
				<p class="m-6 font-Poppins text-base text-slate-200">change your password for increase security </p>
				<!-- <div class=" ml-12 mt-2 self-start text-left text-base font-semibold text-gray-300 ">
					Password :
				</div> -->
				
			</div>
			<input
				type="text"
				class=" mx-4 my-2 h-12 w-[95%] self-center rounded-2xl border-2  border-[#24262b] bg-[#303338]  p-2 text-lg font-medium text-[#98999e] outline-none  focus:border-sky-500 active:border-gray-800 "
			/>
			<button
				
				class="m-3 h-10 w-fit self-center rounded-xl bg-blue-500 p-2 px-4 font-semibold text-white hover:bg-green-600 active:bg-green-700 "
				>Reset</button
			>
		</div>
	</div>
	
</div>

<style>
	/* your styles go here */
</style>
