import { json } from "@sveltejs/kit";



export async function POST({ request }) {
	const data = await request.json();
    let resData
    await fetch('http://localhost:8001/admin/addproduct', {
			// credentials: 'same-origin',
			method: 'POST',
			mode: 'no-cors',
			headers: new Headers({ 'Content-Type': 'application/json', Accept: 'application/json' }),

			body: JSON.stringify(data)
		})
			.then((response) => response.json())
			.then((data) => {
                resData=data
				console.log('🚀 ~ file: +page.svelte ~ line 312 ~ .then ~ REsponse', data);
		});

    return json(resData)
}