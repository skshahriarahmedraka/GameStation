// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import * as jwt from 'jsonwebtoken';



/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ locals,cookies }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }
	// let UserProData
	// console.log("🚀 ~ file: +page.server.ts ~ line 13 ~ GET ~ params.productid", params.profileid)
	// await fetch(`http://localhost:8001/user/${params.profileid}`,{
	// 	mode:"no-cors"
	// }).then((res)=>{
	//  return res.json()
	// }).then((d)=>{
	// 	UserProData=d
	// 	console.log("🚀 ~ file: +page.server.ts ~ line 16 ~ constload:PageServerLoad= ~ UserProData", UserProData)
	// })

	const MyCookie = cookies.get('Auth1') || '';
	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	let Userdata
	: {
		Name: string;
		UserID: string;
		Accounttype: string;
		ProfileImg: string;
		BannerImg: string;
		Email: string;
		Mobile: string;
		Address: string;
		Country: string;
		City: string;
		ZipCode: string;
		Coin: number;
		TransactionHistory: string[];
		WishList: string[];
		ContactAdminMsg: string[];
		ContactDevMsg: string[];
		UserCart: string[];
	} = {
		Name: '',
		UserID: '',
		Accounttype: '',

		ProfileImg: '',
		BannerImg: '',

		Email: '',
		Mobile: '',

		Address: '',
		// Region: '',
		Country: '',
		City: '',
		ZipCode: '',

		Coin: 0,

		TransactionHistory: [] as string[],
		WishList: [] as string[],
		ContactAdminMsg: [] as string[],
		ContactDevMsg: [] as string[],
		UserCart: [] as string[]
	};

	if (MyCookie != '') {
		

		interface tokeninterface{
			Email: string 
			Name: string,
			UserID: string,
			Accounttype: string,
			exp: number
		  }

		const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
		

		console.log(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UserID}`);
		await fetch(`http://${process.env.GO_HOST}/user/${(decoded as tokeninterface).UserID}`, {
			mode: 'no-cors'
		})
			.then((res) => {
				return res.json();
			})
			.then((da) => {
				// console.log('🚀 ~ file: +page.server.ts ~ line 106 ~ myfetch ~ da', da);

				Userdata = da;
			});
	}
	return {
		Userdata
	};
	
	// return {UserProData}
	// if (UserProData){

		
	// }
	// else {
	// 	throw error(404, 'Not found');
	// }
}
