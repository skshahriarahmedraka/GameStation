// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async () => {
	throw redirect(302, '/login');
};
