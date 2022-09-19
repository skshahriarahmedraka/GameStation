/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./src/**/*.{html,js,svelte,ts}',    "./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}",],
		theme: {
		extend: {
			fontFamily:{
				"Poppins":["Poppins","sans-serif"]
			}
		}
	},
	plugins: [
		require('tailwindcss-question-mark'),
		require('@tailwindcss/line-clamp'),
		require('@tailwindcss/typography'),
		require('@tailwindcss/aspect-ratio'),
		require('@tailwindcss/forms'),
		require('prettier-plugin-tailwindcss'),
		require('flowbite/plugin')

	],
	darkMode: 'class',

};
