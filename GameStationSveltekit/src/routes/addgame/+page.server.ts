// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ locals }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }
	// else if (){

	// }

	//   ONLY FOR ADMIN	
	// const MyCookie = cookies.get('Auth1') || '';
	// const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	// let Userdata:any 
	// if (MyCookie!= ''){
		
	// 	const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// 	console.log("decoded: ",decoded);
	//   //   let resdata
	// 	if (decoded.Accounttype != "admin"){
	// 	  throw redirect(302,"/")
	// 	}
	// }
	
	
	return {}
	
}
