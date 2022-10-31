// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

/** @type {import('./$types').PageServerLoad} */
export const load: PageServerLoad = async ({ params, locals }) => {
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 7 ~ constload:PageServerLoad= ~ locals.user',
			locals.user
		);
		throw redirect(302, '/login');
	}

	let NewsData: {
		NewsID: string;
		Title: string;
		BannerImg: string;
		Date: string;
		WrittenBy: string;
		Detail: string;
		FooterComment: string;
	}

	console.log(
		'🚀 ~ file: +page.server.ts ~ line 23 ~ constload:PageServerLoad= ~ params.newsid',
		params.newsid
	);
	await fetch(`http://${process.env.GO_HOST}/news/${params.newsid}`, {
		mode: 'no-cors'
	})
		.then((res) => {
			return res.json();
		})
		.then((d) => {
			NewsData = d;
			console.log('🚀 ~ file: +page.server.ts ~ line 14 ~ load ~ NewsData', NewsData);
		});
		return { NewsData };
		

};
