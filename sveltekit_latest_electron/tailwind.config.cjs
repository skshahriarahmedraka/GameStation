/** @type {import('tailwindcss').Config}*/
/** @type {import('tailwindcss').Config}*/
const config = {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		extend: {
		  fontFamily: {
					'poppins': ['Poppins', 'sans-serif'],
					"sf-pro": ['SF Pro Display', 'sans-serif'],
					"raleway": ['Raleway', 'sans-serif']
				  },
				  screens:{
					xs: { min:"1px", max: '575px' }, // Mobile (iPhone 3 - iPhone XS Max).
			  sm: { min: '576px', max: '897px' }, // Mobile (matches max: iPhone 11 Pro Max landscape @ 896px).
			  md: { min: '898px', max: '1199px' }, // Tablet (matches max: iPad Pro @ 1112px).
			  lg: { min: '1200px' }, // Desktop smallest.
			  xl: { min: '1159px' }, // Desktop wide.
			  xxl: { min: '1359px' }, // Desktop widescreen.
				}
		}
	  },
	
	  plugins: [
		require('tailwindcss-question-mark'),
			require('@tailwindcss/typography'),
			require('@tailwindcss/aspect-ratio'),
			require('@tailwindcss/forms'),
			require('prettier-plugin-tailwindcss')
	  ]
};

module.exports = config;
