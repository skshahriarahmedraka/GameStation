// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import * as jwt from 'jsonwebtoken';

export const load: PageServerLoad = async ({ cookies, locals }) => {
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 10 ~ constload:PageServerLoad= ~ user',
			locals.user.Authenticated
		);
		throw redirect(302, '/login');
	}

	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies);
	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies.get('Auth1'));
	//   const MyCookie = cookies.get('Auth1') || '';
	//   const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	//   let Userdata:any
	//   if (MyCookie!= ''){

	// 	  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// 	  console.log("decoded: ",decoded);
	// 	//   let resdata
	// 	  console.log(`http://${process.env.GO_HOST}/user/${decoded.UserID}`)
	// 	  await fetch(`http://${process.env.GO_HOST}/user/${decoded.UserID}`,{
	// 		  mode:"no-cors"
	// 	  }).then((res)=>{
	// 		  return res.json()
	// 	  }).then((d)=>{

	// 		Userdata=d

	// 	  })

	//   }
	// return {
	// 	Userdata
	// }

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
};
