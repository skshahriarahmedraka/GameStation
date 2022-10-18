import { json } from '@sveltejs/kit';

export async function POST({ request }) {
	const data = await request.json();
	let resData:any;
	await fetch('http://localhost:8001/admin/addproduct', {
		// credentials: 'same-origin',
		method: 'POST',
		mode: 'no-cors',
		headers: {
			'Content-Type': 'application/json',
			Accept: 'application/json'
		},

		body: JSON.stringify(data)
	})
		.then((response) => {
			return response.json();
		})
		.then((data) => {
			resData = data;
			console.log('🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ REsponse', data);
		});

	console.log('🚀 ~ file: +server.ts ~ line 23 ~ POST ~ resData', resData);
	return json({ status: resData.status, GameID: resData.GameID });
}
