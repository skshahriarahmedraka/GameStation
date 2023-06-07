// import { error } from '@sveltejs/kit';
import type { GameDataType, UserDataType } from '$lib/Store/types';
import type { PageLoad } from './$types';



export const load:PageLoad = async ({ fetch }) => {

    let MostPopularList: GameDataType[] = [] as GameDataType[];
    
    console.log("🚀 +page.ts MostPopularList: ","http://${import.meta.env.VITE_GO_HOST}/api/game/mostpopular", MostPopularList)
	await fetch(`http://${import.meta.env.VITE_GO_HOST}/api/game/mostpopular`, {
        method: 'GET'
		//   mode: "no-cors",
	})
		.then((res) => {
			return res.json();
		})
		.then((data) => {
            MostPopularList = data;
			console.log(
                '🚀 ~ MostPopularList',
				MostPopularList
                );
            });
            
            console.log("🚀 +page.ts MostPopularList:", MostPopularList)
    
			const NewReleaseList: GameDataType[] = [] as GameDataType[];
			const TopRatedList: GameDataType[] = [] as GameDataType[];
			const TopSoledList: GameDataType[] = [] as GameDataType[];
			const TrendingList: GameDataType[] = [] as GameDataType[];
			const CarouselList: GameDataType[] = [] as GameDataType[];

			const Userdata: UserDataType = {} as UserDataType;

    return {
		Userdata,
        MostPopularList,
		NewReleaseList,
		TopRatedList,
		TopSoledList,
		TrendingList,
		CarouselList
    }
 
//   throw error(404, 'Not found');
}
