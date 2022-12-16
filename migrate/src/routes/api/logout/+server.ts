import { json } from '@sveltejs/kit';
// import * as cookie from 'cookie';

import type { RequestHandler } from './$types';
export const GET: RequestHandler = async ({ cookies }) => {
    const MyCookie = cookies.get('Auth1') || '';


	if (MyCookie!= '') {
		cookies.set('Auth1', "", {
			// send cookie for every page
			path: '/',
			expires: new Date(0)
		});
		return json({ status: 'logout' });
	} else {
		return json({ status: 'error' });
	}
}
