import { json } from "@sveltejs/kit";


import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ( {request} ) =>{
	const data = await request.json();
    let resData
    await fetch(`http://${process.env.GO_HOST}/game/addcomment`, {
			// credentials: 'same-origin',
			method: 'POST',
			mode: 'no-cors',
			headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(data)
		})
			.then((response) => response.json())
			.then((data) => {
                resData=data
                console.log("🚀 ~ file: +server.ts ~ line 20 ~ .then ~ resData", resData)
		});

    return json(resData)
}