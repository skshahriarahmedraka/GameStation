// import { error, json } from '@sveltejs/kit';

import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import * as jwt from 'jsonwebtoken';
import { UserProData } from '$lib/Store/store';


export const load: PageServerLoad = async ({ cookies, locals }) => {
	if (!locals.user.Authenticated) {
		console.log(
			'🚀 ~ file: +page.server.ts ~ line 10 ~ constload:PageServerLoad= ~ user',
			locals.user.Authenticated
		);
		throw redirect(302, '/login');
	}



	const MyCookie = cookies.get('Auth1') || '';
	const JWT_Auth_KEY: string = process.env.JWT_SECRET as string;

	let Userdata: UserDataType

	if (MyCookie != '') {
		interface tokeninterface {
			Email: string;
			Name: string;
			UserID: string;
			Accounttype: string;
			exp: number;
		}

		const decoded = jwt.verify(MyCookie, JWT_Auth_KEY);

		console.log(`/user/${(decoded as tokeninterface).UserID}`);
		await fetch(`https://localhost/api/user/${(decoded as tokeninterface).UserID}`, {
			mode: 'no-cors'
		})
			.then((res) => {
				return res.json();
			})
			.then((da) => {
				// console.log('🚀 ~ file: +page.server.ts ~ line 106 ~ myfetch ~ da', da);

				Userdata = da;
			});
	}
	UserProData.update((d) => (d = Userdata));

	// MOST POPULAR
	// let MostPopularList: {
	// 	GameID: string;
	// 	Name: string;
	// 	Moto: string;
	// 	LogoImage: string;
	// 	BigPosterImage: string;
	// 	SmallPosterImage: string;
	// 	OtherImages: string[];
	// 	Genres: string[];
	// 	Feature: string[];
	// 	Description: string;
	// 	Rating: number;
	// 	RatingGivenBy: {
	// 		'PC Gamer': string;
	// 		IGN: string;
	// 		'Game Informer': string;
	// 	};
	// 	Minspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Recomendedspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Discount: number;
	// 	Price: number;
	// 	Developer: string;
	// 	Publisher: string;
	// 	Released: string;
	// 	Platform: string[];
	// 	Players: number;
	// 	Comment: {
	// 		Name: string;
	// 		Rating: number;
	// 		Description: string;
	// 	}[];
	// 	FollowUs: {
	// 		Facebook: string;
	// 		Discord: string;
	// 		Youtube: string;
	// 		Twitter: string;
	// 		Site: string;
	// 	};
	// }[] = [] as {
	// 	GameID: '';
	// 	Name: '';
	// 	Moto: '';
	// 	LogoImage: '';
	// 	BigPosterImage: '';
	// 	SmallPosterImage: '';
	// 	OtherImages: [];
	// 	Genres: [''];
	// 	Feature: [''];
	// 	Description: '';
	// 	FollowUs: {
	// 		Facebook: '';
	// 		Discord: '';
	// 		Youtube: '';
	// 		Twitter: '';
	// 		Site: '';
	// 	};
	// 	Rating: 0;
	// 	RatingGivenBy: {
	// 		'PC Gamer': '';
	// 		IGN: '';
	// 		'Game Informer': '';
	// 	};
	// 	Minspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Recomendedspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Price: 0;
	// 	Discount: 0;
	// 	Developer: '';
	// 	Publisher: '';
	// 	Released: '';
	// 	Platform: ['', '', ''];
	// 	Players: 0;
	// 	Comment: [];
	// }[];
	// await fetch(`http://${process.env.GO_HOST}/game/mostpopular`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		MostPopularList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			MostPopularList
	// 		);
	// 	});

	// // NewReleaseList
	// let NewReleaseList: {
	// 	GameID: string;
	// 	Name: string;
	// 	Moto: string;
	// 	LogoImage: string;
	// 	BigPosterImage: string;
	// 	SmallPosterImage: string;
	// 	OtherImages: string[];
	// 	Genres: string[];
	// 	Feature: string[];
	// 	Description: string;
	// 	Rating: number;
	// 	RatingGivenBy: {
	// 		'PC Gamer': string;
	// 		IGN: string;
	// 		'Game Informer': string;
	// 	};
	// 	Minspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Recomendedspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Discount: number;
	// 	Price: number;
	// 	Developer: string;
	// 	Publisher: string;
	// 	Released: string;
	// 	Platform: string[];
	// 	Players: number;
	// 	Comment: {
	// 		Name: string;
	// 		Rating: number;
	// 		Description: string;
	// 	}[];
	// 	FollowUs: {
	// 		Facebook: string;
	// 		Discord: string;
	// 		Youtube: string;
	// 		Twitter: string;
	// 		Site: string;
	// 	};
	// }[] = [] as {
	// 	GameID: '';
	// 	Name: '';
	// 	Moto: '';
	// 	LogoImage: '';
	// 	BigPosterImage: '';
	// 	SmallPosterImage: '';
	// 	OtherImages: [];
	// 	Genres: [''];
	// 	Feature: [''];
	// 	Description: '';
	// 	FollowUs: {
	// 		Facebook: '';
	// 		Discord: '';
	// 		Youtube: '';
	// 		Twitter: '';
	// 		Site: '';
	// 	};
	// 	Rating: 0;
	// 	RatingGivenBy: {
	// 		'PC Gamer': '';
	// 		IGN: '';
	// 		'Game Informer': '';
	// 	};
	// 	Minspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Recomendedspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Price: 0;
	// 	Discount: 0;
	// 	Developer: '';
	// 	Publisher: '';
	// 	Released: '';
	// 	Platform: ['', '', ''];
	// 	Players: 0;
	// 	Comment: [];
	// }[];
	// await fetch(`http://${process.env.GO_HOST}/game/newrelease`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		NewReleaseList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			NewReleaseList
	// 		);
	// 	});

	// // TopRatedList
	// let TopRatedList: {
	// 	GameID: string;
	// 	Name: string;
	// 	Moto: string;
	// 	LogoImage: string;
	// 	BigPosterImage: string;
	// 	SmallPosterImage: string;
	// 	OtherImages: string[];
	// 	Genres: string[];
	// 	Feature: string[];
	// 	Description: string;
	// 	Rating: number;
	// 	RatingGivenBy: {
	// 		'PC Gamer': string;
	// 		IGN: string;
	// 		'Game Informer': string;
	// 	};
	// 	Minspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Recomendedspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Discount: number;
	// 	Price: number;
	// 	Developer: string;
	// 	Publisher: string;
	// 	Released: string;
	// 	Platform: string[];
	// 	Players: number;
	// 	Comment: {
	// 		Name: string;
	// 		Rating: number;
	// 		Description: string;
	// 	}[];
	// 	FollowUs: {
	// 		Facebook: string;
	// 		Discord: string;
	// 		Youtube: string;
	// 		Twitter: string;
	// 		Site: string;
	// 	};
	// }[] = [] as {
	// 	GameID: '';
	// 	Name: '';
	// 	Moto: '';
	// 	LogoImage: '';
	// 	BigPosterImage: '';
	// 	SmallPosterImage: '';
	// 	OtherImages: [];
	// 	Genres: [''];
	// 	Feature: [''];
	// 	Description: '';
	// 	FollowUs: {
	// 		Facebook: '';
	// 		Discord: '';
	// 		Youtube: '';
	// 		Twitter: '';
	// 		Site: '';
	// 	};
	// 	Rating: 0;
	// 	RatingGivenBy: {
	// 		'PC Gamer': '';
	// 		IGN: '';
	// 		'Game Informer': '';
	// 	};
	// 	Minspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Recomendedspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Price: 0;
	// 	Discount: 0;
	// 	Developer: '';
	// 	Publisher: '';
	// 	Released: '';
	// 	Platform: ['', '', ''];
	// 	Players: 0;
	// 	Comment: [];
	// }[];
	// await fetch(`http://${process.env.GO_HOST}/game/toprated`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		TopRatedList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			TopRatedList
	// 		);
	// 	});
	// // TopSoledList
	// let TopSoledList: {
	// 	GameID: string;
	// 	Name: string;
	// 	Moto: string;
	// 	LogoImage: string;
	// 	BigPosterImage: string;
	// 	SmallPosterImage: string;
	// 	OtherImages: string[];
	// 	Genres: string[];
	// 	Feature: string[];
	// 	Description: string;
	// 	Rating: number;
	// 	RatingGivenBy: {
	// 		'PC Gamer': string;
	// 		IGN: string;
	// 		'Game Informer': string;
	// 	};
	// 	Minspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Recomendedspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Discount: number;
	// 	Price: number;
	// 	Developer: string;
	// 	Publisher: string;
	// 	Released: string;
	// 	Platform: string[];
	// 	Players: number;
	// 	Comment: {
	// 		Name: string;
	// 		Rating: number;
	// 		Description: string;
	// 	}[];
	// 	FollowUs: {
	// 		Facebook: string;
	// 		Discord: string;
	// 		Youtube: string;
	// 		Twitter: string;
	// 		Site: string;
	// 	};
	// }[] = [] as {
	// 	GameID: '';
	// 	Name: '';
	// 	Moto: '';
	// 	LogoImage: '';
	// 	BigPosterImage: '';
	// 	SmallPosterImage: '';
	// 	OtherImages: [];
	// 	Genres: [''];
	// 	Feature: [''];
	// 	Description: '';
	// 	FollowUs: {
	// 		Facebook: '';
	// 		Discord: '';
	// 		Youtube: '';
	// 		Twitter: '';
	// 		Site: '';
	// 	};
	// 	Rating: 0;
	// 	RatingGivenBy: {
	// 		'PC Gamer': '';
	// 		IGN: '';
	// 		'Game Informer': '';
	// 	};
	// 	Minspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Recomendedspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Price: 0;
	// 	Discount: 0;
	// 	Developer: '';
	// 	Publisher: '';
	// 	Released: '';
	// 	Platform: ['', '', ''];
	// 	Players: 0;
	// 	Comment: [];
	// }[];
	// await fetch(`http://${process.env.GO_HOST}/game/topsold`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		TopSoledList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			TopSoledList
	// 		);
	// 	});

	// // TrendingList
	// let TrendingList: {
	// 	GameID: string;
	// 	Name: string;
	// 	Moto: string;
	// 	LogoImage: string;
	// 	BigPosterImage: string;
	// 	SmallPosterImage: string;
	// 	OtherImages: string[];
	// 	Genres: string[];
	// 	Feature: string[];
	// 	Description: string;
	// 	Rating: number;
	// 	RatingGivenBy: {
	// 		'PC Gamer': string;
	// 		IGN: string;
	// 		'Game Informer': string;
	// 	};
	// 	Minspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Recomendedspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Discount: number;
	// 	Price: number;
	// 	Developer: string;
	// 	Publisher: string;
	// 	Released: string;
	// 	Platform: string[];
	// 	Players: number;
	// 	Comment: {
	// 		Name: string;
	// 		Rating: number;
	// 		Description: string;
	// 	}[];
	// 	FollowUs: {
	// 		Facebook: string;
	// 		Discord: string;
	// 		Youtube: string;
	// 		Twitter: string;
	// 		Site: string;
	// 	};
	// }[] = [] as {
	// 	GameID: '';
	// 	Name: '';
	// 	Moto: '';
	// 	LogoImage: '';
	// 	BigPosterImage: '';
	// 	SmallPosterImage: '';
	// 	OtherImages: [];
	// 	Genres: [''];
	// 	Feature: [''];
	// 	Description: '';
	// 	FollowUs: {
	// 		Facebook: '';
	// 		Discord: '';
	// 		Youtube: '';
	// 		Twitter: '';
	// 		Site: '';
	// 	};
	// 	Rating: 0;
	// 	RatingGivenBy: {
	// 		'PC Gamer': '';
	// 		IGN: '';
	// 		'Game Informer': '';
	// 	};
	// 	Minspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Recomendedspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Price: 0;
	// 	Discount: 0;
	// 	Developer: '';
	// 	Publisher: '';
	// 	Released: '';
	// 	Platform: ['', '', ''];
	// 	Players: 0;
	// 	Comment: [];
	// }[];
	// await fetch(`http://${process.env.GO_HOST}/game/trending`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		TrendingList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			TrendingList
	// 		);
	// 	});
	// // CarouselList
	// let CarouselList: {
	// 	GameID: string;
	// 	Name: string;
	// 	Moto: string;
	// 	LogoImage: string;
	// 	BigPosterImage: string;
	// 	SmallPosterImage: string;
	// 	OtherImages: string[];
	// 	Genres: string[];
	// 	Feature: string[];
	// 	Description: string;
	// 	Rating: number;
	// 	RatingGivenBy: {
	// 		'PC Gamer': string;
	// 		IGN: string;
	// 		'Game Informer': string;
	// 	};
	// 	Minspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Recomendedspec: {
	// 		Name: string;
	// 		Value: string;
	// 	}[];
	// 	Discount: number;
	// 	Price: number;
	// 	Developer: string;
	// 	Publisher: string;
	// 	Released: string;
	// 	Platform: string[];
	// 	Players: number;
	// 	Comment: {
	// 		Name: string;
	// 		Rating: number;
	// 		Description: string;
	// 	}[];
	// 	FollowUs: {
	// 		Facebook: string;
	// 		Discord: string;
	// 		Youtube: string;
	// 		Twitter: string;
	// 		Site: string;
	// 	};
	// }[] = [] as {
	// 	GameID: '';
	// 	Name: '';
	// 	Moto: '';
	// 	LogoImage: '';
	// 	BigPosterImage: '';
	// 	SmallPosterImage: '';
	// 	OtherImages: [];
	// 	Genres: [''];
	// 	Feature: [''];
	// 	Description: '';
	// 	FollowUs: {
	// 		Facebook: '';
	// 		Discord: '';
	// 		Youtube: '';
	// 		Twitter: '';
	// 		Site: '';
	// 	};
	// 	Rating: 0;
	// 	RatingGivenBy: {
	// 		'PC Gamer': '';
	// 		IGN: '';
	// 		'Game Informer': '';
	// 	};
	// 	Minspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Recomendedspec: [
	// 		{
	// 			Name: 'OS';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Storage';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'Memory';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'GPU';
	// 			Value: '';
	// 		},
	// 		{
	// 			Name: 'CPU';
	// 			Value: '';
	// 		}
	// 	];
	// 	Price: 0;
	// 	Discount: 0;
	// 	Developer: '';
	// 	Publisher: '';
	// 	Released: '';
	// 	Platform: ['', '', ''];
	// 	Players: 0;
	// 	Comment: [];
	// }[];
	// await fetch(`http://${process.env.GO_HOST}/game/carousel`, {
	// 	method: 'GET'
	// 	//   mode: "no-cors",
	// })
	// 	.then((res) => {
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		CarouselList = data;
	// 		console.log(
	// 			'🚀 ~ file: MostPopular.svelte ~ line 261 ~ GetMostPopularGames ~ MostPopularList',
	// 			CarouselList
	// 		);
	// 	});
	return {
		// Userdata,
		// // MostPopularList,
		// NewReleaseList,
		// TopRatedList,
		// TopSoledList,
		// TrendingList,
		// CarouselList
	};
};
