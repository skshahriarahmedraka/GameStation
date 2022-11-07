// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import * as jwt from 'jsonwebtoken';
// import { UserProData } from '$lib/Store/store';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ cookies, locals }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }

	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies);
	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies.get('Auth1'));

	// ONLY FOR ADMIN 
	  const MyCookie = cookies.get('Auth1') || '';
	  const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	  let Userdata:any 
	  if (MyCookie!= ''){
		  
		  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
		  console.log("decoded: ",decoded);
		//   let resdata
		  if (decoded.Accounttype != "admin"){
			throw redirect(302,"/")
		  }
	  }
	
	
	return {}
	
}
