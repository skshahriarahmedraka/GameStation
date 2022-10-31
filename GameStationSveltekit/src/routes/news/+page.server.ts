// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({  locals }) => {
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user',
			locals.user
		);
		throw redirect(302, '/login');
	}

	let NewsDataList: {
		NewsID: string;
		Title: string;
		BannerImg: string;
		Date: string;
		WrittenBy: string;
		Detail: string;
		FooterComment: string;
	}[] =[]

	
	await fetch(`http://${process.env.GO_HOST}/news`, {
		mode: 'no-cors'
	})
		.then((res) => {
			return res.json();
		})
		.then((d) => {
			NewsDataList = d;
			console.log('🚀 ~ file: +page.server.ts ~ line 14 ~ load ~ NewsDataList list', NewsDataList);
		});
		return { NewsDataList };
		

};
