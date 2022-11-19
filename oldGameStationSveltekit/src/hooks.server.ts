// import * as cookie from "cookie"
import * as jwt from 'jsonwebtoken';
import dotenv from 'dotenv';

dotenv.config();

import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// get cookies from browser
	const MyCookie: string = event.cookies.get('Auth1') || '';
	console.log('🚀 ~ file: hooks.server.ts ~ line 12 ~ consthandle:Handle= ~ MyCookie', MyCookie);
	// event.locals.user.Authenticated=false

	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;
	console.log(
		'🚀 ~ file: hooks.server.ts ~ line 14 ~ consthandle:Handle= ~ JWT_Auth_KEY',
		JWT_Auth_KEY
	);
	// const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);
	// console.log(decoded);
	jwt.verify(MyCookie, JWT_Auth_KEY, (err) => {
		// console.log("🚀 ~ file: hooks.server.ts ~ line 23 ~ jwt.verify ~ decode", decode)
		// console.log("🚀 ~ file: hooks.ts ~ line 21 ~ jwt.verify ~ err : ", err)
		if (err) {
			event.locals.user = {
				Authenticated: false
			};
			// return {user:{authenticated:false}}
		} else {
			event.locals.user = {
				Authenticated: true
			};
		}
	});

	// if `user` exists set `events.local`

	console.log(
		'🚀 ~ file: hooks.server.ts ~ line 29 ~ consthandle:Handle= ~ event.locals.user',
		event.locals.user.Authenticated
	);
	// event.locals.user = {
	//     Authenticated:false
	// }

	// load page as normal
	return await resolve(event);
};
