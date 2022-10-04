<script lang="ts">
	export let value = '';
	export let minRows = 1;
	export let maxRows:any;
	
	$: minHeight = `${1 + minRows * 1.2}em`;
	$: maxHeight = maxRows ? `${1 + maxRows * 1.2}em` : `auto`;
</script>


<textarea bind:value  class=" {$$props.class} text-[#e7e9eb] border-gray-600  border-2 p-2 text-base font-Poppins scrol max-h-fit h-60 overflow-y-scroll scrol3 bg-inherit w-full overflow-hidden rounded-md"  placeholder="Write Here"></textarea>	

<style>

</style>