import { json } from "@sveltejs/kit";
// import * as cookie from 'cookie';

import type { RequestHandler } from './$types';
 
export const POST: RequestHandler = async ({ cookies , request }) =>{
	const data = await request.json();
    console.log("🚀 ~ file: +server.ts ~ line 8 ~ POST ~ data", data)
    let resData:any
    await fetch(`http://${process.env.GO_HOST}/sveltekit/register`, {
			// credentials: 'same-origin',
				// credentials: 'include',

			method: 'POST',
			mode: 'no-cors',
			headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(data)
		})
			.then((response) => response.json())
			.then((data) => {
                resData=data
                console.log("🚀 ~ file: +server.ts ~ line 22 ~ .then ~ resData", resData)
		});
        if (resData.JWT){

            
                cookies.set('Auth1', resData.JWT, {
                    // send cookie for every page
                    path: '/',
                    // server side only cookie so you can't use `document.cookie`
                    httpOnly: true,
                    // Domain:"127.0.0.1",
                    // only requests from same site can send cookies
                    // https://developer.mozilla.org/en-US/docs/Glossary/CSRF
                    sameSite: 'strict',
                    // only sent over HTTPS in production
                    secure: process.env.NODE_ENV === 'production',
                    // set cookie to expire after a month
                    maxAge: 60 * 60 * 24 * 30 *30,
                  })
                  cookies.set('noAuth1', resData.Base64, {
                    // send cookie for every page
                    path: '/',
                    // server side only cookie so you can't use `document.cookie`
                    httpOnly: false,
                    // Domain:"127.0.0.1",
                    // only requests from same site can send cookies
                    // https://developer.mozilla.org/en-US/docs/Glossary/CSRF
                    sameSite: 'strict',
                    // only sent over HTTPS in production
                    secure: process.env.NODE_ENV === 'production',
                    // set cookie to expire after a month
                    maxAge: 60 * 60 * 24 * 30 *30,
                  })
                }
               
                else {
                    return json({"status":resData.error})
                }
                return json({"status":"OK"})
        

}