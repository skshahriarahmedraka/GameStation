import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
 
export const load = (({ params }) => {

 
//   throw error(404, 'Not found');
}) satisfies PageLoad;