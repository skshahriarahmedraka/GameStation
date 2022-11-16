import { json } from '@sveltejs/kit';

import type { RequestHandler } from './$types';

export const POST: RequestHandler = async({ request })=> {
	const data = await request.json();
	let resData:any;
	await fetch(`http://${process.env.GO_HOST}/game/${data.GameID}`, {
		// credentials: 'same-origin',
		method: 'GET',
		mode: 'no-cors',
		headers: {
			'Content-Type': 'application/json',
			Accept: 'application/json'
		},

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
