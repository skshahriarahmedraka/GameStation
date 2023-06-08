/* eslint-disable @typescript-eslint/no-explicit-any */
// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import * as jwt from 'jsonwebtoken';
import type { GameDataType,UserDataType } from '$lib/Store/types';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ cookies,params,locals }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }


	let mydata: GameDataType = {} as GameDataType
	console.log("🚀 ~ file: +page.server.ts ~ line 13 ~ GET ~ params.productid", params.productid)
	await fetch(`http://${process.env.GO_HOST}/game/${params.productid}`,{
		mode:"no-cors"
	}).then((res)=>{
	 return res.json()
	}).then((d)=>{
		mydata=d
		console.log("🚀 ~ file: +page.server.ts ~ line 14 ~ load ~ mydata", mydata)
	})
	


	const MyCookie = cookies.get('Auth1') || '';
	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	let Userdata: UserDataType = {} as UserDataType;

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
	

	return {mydata,Userdata}
	// if (mydata){

		
	// }
	// else {
	// 	throw error(404, 'Not found');
	// }result
}
