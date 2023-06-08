// import * as jwt from 'jsonwebtoken';

// import { UserProData } from '$lib/Store/store';
import type { PageServerLoad } from './$types';
// import { redirect } from '@sveltejs/kit'
// import * as jwt from 'jsonwebtoken';
// import { UserProData } from '$lib/Store/store';

export const  load: PageServerLoad = async () => {
	
	

	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies);
	//   console.log('🚀 ~ file: +layout.server.ts ~ line 4 ~ cookies', cookies.get('Auth1'));
	//   const MyCookie = cookies.get('Auth1') || '';
	//   const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	//   let Userdata:UserDataType = {} as UserDataType 
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
	return {
		// Userdata
	}
}


//  function GetUserData(decoded) {
	
// }


// import { error, json } from '@sveltejs/kit';




