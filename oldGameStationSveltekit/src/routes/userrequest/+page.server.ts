// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit'
import type { PageServerLoad } from './$types'
import * as jwt from 'jsonwebtoken';
// import { UserProData } from '$lib/Store/store';


/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ cookies,locals }) => {

	if (!locals.user.Authenticated) {
		console.log("🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user", locals.user)
		throw redirect(302, '/login')
	  }

	  const MyCookie = cookies.get('Auth1') || '';
	  const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
  
	  let Userdata: {
		  Name: string
		  UserID: string
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
		  
  
		  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
		  console.log('🚀 ~ file: +page.server.ts ~ line 136 ~ myfetch ~ decoded', decoded);
		  if ((decoded as tokeninterface).Accounttype != "admin"){
					throw redirect(302,"/")
				  }
  
				  interface tokeninterface{
					Email: string 
					Name: string,
					UserID: string,
					Accounttype: string,
					exp: number
				  }
		
				
		
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

	let UserReqList: {
		UserID: string;
		Coin: number;
		ReqList: {
			Amount: number;
			JWT: string;
			Status: string;
		}[]
	}[]= []
	  await fetch(`http://${process.env.GO_HOST}/admin/adminmoneymanagement`, {
	  		method: 'GET',
	  		mode: 'no-cors',

	  }).then((res)=>{return res.json()}).then((data)=>{
	  console.log("🚀 ~ file: +page.server.ts ~ line 112 ~ constload:PageServerLoad= ~ data", data)
	  UserReqList=data
	  })

	  return {
		  Userdata,
		  UserReqList
	  };

	//   const MyCookie = cookies.get('Auth1') || '';
	//   const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	 
	  
	//   if (MyCookie!= ''){
	// 	let temp1:  { Name: string;
	// 		UserID: string;
	// 		Accounttype: string;
	// 		ProfileImg: string;
	// 		BannerImg: string;
	// 		Email: string;
	// 		Mobile: string;
	// 		Address: string;
	// 		Country: string;
	// 		City: string;
	// 		ZipCode: string;
	// 		Coin: number;
	// 		TransactionHistory: string[],
	// 		WishList: string[],
	// 		ContactAdminMsg: string[],
	// 		ContactDevMsg: string[],
	// 		UserCart: string[];
	// 		}={
	// 		  Name: "",
	// 		  UserID: "",
	// 		  Accounttype: '',
		  
	// 		  ProfileImg: '',
	// 		  BannerImg: '',
		  
	// 		  Email: '',
	// 		  Mobile: '',
		  
	// 		  Address: '',
	// 		  // Region: '',
	// 		  Country: '',
	// 		  City: '',
	// 		  ZipCode: '',
		  
	// 		  Coin: 0,
		  
	// 		  TransactionHistory: [] as string[],
	// 		  WishList: [] as string[],
	// 		  ContactAdminMsg: [] as string[],
	// 		  ContactDevMsg: [] as string[],
	// 		  UserCart: [] as  string[]
	// 	  }
	// 		UserProData.subscribe((d)=>{
	// 		  temp1=d
	// 		})
	// 		if (temp1.UserID===""){
	// 		  const myfetch =async ()=>{
	// 			  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// 			console.log("decoded: ",decoded);
			  
	// 			console.log(`http://${process.env.GO_HOST}/user/${decoded.UserID}`)
	// 			await fetch(`http://${process.env.GO_HOST}/user/${decoded.UserID}`,{
	// 				mode:"no-cors"
	// 			}).then((res)=>{
	// 				return res.json()
	// 			}).then((da)=>{
					
	// 			  UserProData.update((n)=>{
	// 				  da
	// 			  })
			  
	// 			})
	// 		  }
	// 		  myfetch()
	// 		}
	// 	  // ONLY FOR ADMIN 
	// 	  const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// 	  console.log("decoded: ",decoded);
	// 	//   let resdata
	// 	  if (decoded.Accounttype != "admin"){
	// 		throw redirect(302,"/")
	// 	  }
	//   }

	
	
	// return {}
	
}
