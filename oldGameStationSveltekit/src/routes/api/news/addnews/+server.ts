import { json } from "@sveltejs/kit";


import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ( {request} ) =>{
	const data = await request.json();
    console.log("🚀 ~ file: +server.ts ~ line 8 ~ constPOST:RequestHandler= ~ data", data)
    let resData
    await fetch(`http://${process.env.GO_HOST}/news/addnews`, {
			// credentials: 'same-origin',
			method: 'POST',
			mode: 'no-cors',
			headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(data)
		})
			.then((response) => response.json())
			.then((data) => {
                resData=data
                console.log("🚀 ~ file: +server.ts ~ line 20 ~ .then ~ resData add News", resData)
		});

    return json(resData)
}