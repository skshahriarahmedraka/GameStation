import { error, json } from '@sveltejs/kit';


/** @type {import('./$types').PageServerLoad} */
export async function load({ params, }) {
	let mydata
	console.log("🚀 ~ file: +page.server.ts ~ line 13 ~ GET ~ params.productid", params.productid)
	await fetch(`http://localhost:8001/game/${params.productid}`,{
		mode:"no-cors"
	}).then((res)=>{
	 return res.json()
	}).then((d)=>{
		mydata=d
		console.log("🚀 ~ file: +page.server.ts ~ line 14 ~ load ~ mydata", mydata)
	})
	
	return {mydata}
	// if (mydata){

		
	// }
	// else {
	// 	throw error(404, 'Not found');
	// }
}
