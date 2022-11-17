import { json } from '@sveltejs/kit';

import type { RequestHandler } from './$types';

export const GET: RequestHandler = async()=> {
	// const data = await request.json();
	let resData:any;
	await fetch(`http://${process.env.GO_HOST}/game/upcomming`, {
		// credentials: 'same-origin',
		method: 'GET',
		mode: 'no-cors',
		headers: {
			'Content-Type': 'application/json',
			Accept: 'application/json'
		},
		// body: JSON.stringify(data)
		

	})
		.then((response) => {
			return response.json();
		})
		.then((data) => {
			resData = data;
			console.log('🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ REsponse', data);
		});

	console.log('🚀 ~ file: +server.ts ~ line 23 ~ POST ~ resData', resData);
	return json(resData);
}
