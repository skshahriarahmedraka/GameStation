// import { error, json } from '@sveltejs/kit';

// import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ params }) => {
	let UserProData
	console.log("🚀 ~ file: +page.server.ts ~ line 13 ~ GET ~ params.productid", params.profileid)
	await fetch(`http://localhost:8001/user/${params.profileid}`,{
		mode:"no-cors"
	}).then((res)=>{
	 return res.json()
	}).then((d)=>{
		UserProData=d
		console.log("🚀 ~ file: +page.server.ts ~ line 16 ~ constload:PageServerLoad= ~ UserProData", UserProData)
	})
	
	return {UserProData}
	// if (UserProData){

		
	// }
	// else {
	// 	throw error(404, 'Not found');
	// }
}
