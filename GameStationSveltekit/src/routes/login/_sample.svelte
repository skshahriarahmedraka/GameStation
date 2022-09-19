<script lang='ts'>
import { goto } from "$app/navigation";


    async function name() {
        
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
    }
    
</script>

<style>
    /* your styles go here */
</style>

<!-- markup (zero or more items) goes here -->