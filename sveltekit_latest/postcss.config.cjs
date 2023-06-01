// eslint-disable-next-line @typescript-eslint/no-var-requires
const tailwindcss = require('tailwindcss');
// eslint-disable-next-line @typescript-eslint/no-var-requires
const autoprefixer = require('autoprefixer');

const config = {
	plugins: [//Some plugins, like tailwindcss/nesting, need to run before Tailwind,
    //But others, like autoprefixer, need to run after,
    tailwindcss(), //Some plugins, like tailwindcss/nesting, need to run before Tailwind, tailwindcss(), //But others, like autoprefixer, need to run after, autoprefixer, autoprefixer
    autoprefixer,
]
};

module.exports = config;
