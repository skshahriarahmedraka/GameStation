<script lang="ts">
	import type { PageData } from './$types';
	import { UserProData } from '$lib/Store/store';

	export let data: PageData;

	let { Userdata, CartGameList } = data;
	console.log('🚀 ~ file: +page.svelte ~ line 8 ~ CartGameList', CartGameList);
	console.log('🚀 ~ file: +page.svelte ~ line 10 ~ Userdata', Userdata);
	UserProData.update((d) => (d = Userdata));

	import Apple from '$lib/svgs/apple.svelte';
	import Linux from '$lib/svgs/linux.svelte';
	import Windows from '$lib/svgs/windows.svelte';
	import General from '$lib/settingsCom/General.svelte';
	import { UserSettingSelect } from '$lib/Store/store';
	import { goto } from '$app/navigation';

	// import  {UserProData} from '$lib/Store/store';
	// import { goto } from '$app/navigation';

	let GeneralData = {
		UserID: $UserProData.UserID,
		Name: $UserProData.Name,
		Accounttype: $UserProData.Accounttype,

		ProfileImg: $UserProData.ProfileImg,
		BannerImg: $UserProData.BannerImg,

		Email: $UserProData.Email,
		Mobile: $UserProData.Mobile,

		Address: $UserProData.Address,
		// Region: '',
		Country: $UserProData.Country,
		City: $UserProData.City,
		ZipCode: $UserProData.ZipCode,

		Coin: $UserProData.Coin,

		TransactionHistory: $UserProData.TransactionHistory || [],
		WishList: $UserProData.WishList || [],
		ContactAdminMsg: $UserProData.ContactAdminMsg || [],
		ContactDevMsg: $UserProData.ContactDevMsg || [],
		UserCart: $UserProData.UserCart || []
	};
	// 	let li: {
	//     name: string;
	//     midCoverImage: string;
	//     feature: string[];
	//     price: number;
	//     discount: number;
	//     discountEnd: string;
	//     supportedPlatform: string[];
	//     rating: number;
	// }[] = [

	//     {
	// 			name: 'Cyberpunk 2077',
	// 			midCoverImage:
	// 				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
	// 			feature: ['Action', 'RPG', 'Explore'],
	//             price:29.45,
	//             discount:50,
	//             discountEnd: "9/25/2022",
	//             supportedPlatform : ["mac", "windows", "linux"],
	//             rating:9.4,
	// 		},
	//         {
	// 			name: 'Cyberpunk 2077',
	// 			midCoverImage:
	// 				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
	// 			feature: ['Action', 'RPG', 'Explore'],
	//             price:29.45,
	//             discount:50,rating:9.4,
	//             discountEnd: "9/25/2022",
	//             supportedPlatform : ["mac", "windows", "linux"]
	// 		},
	//         {
	// 			name: 'Cyberpunk 2077',
	// 			midCoverImage:
	// 				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
	// 			feature: ['Action', 'RPG', 'Explore'],
	//             price:29.45,rating:9.4,
	//             discount:50,
	//             discountEnd: "9/25/2022",
	//             supportedPlatform : ["mac", "windows", "linux"]
	// 		},
	//         {
	// 			name: 'Cyberpunk 2077',
	// 			midCoverImage:
	// 				'https://cdn1.epicgames.com/offer/77f2b98e2cef40c8a7437518bf420e47/EGS_Cyberpunk2077_CDPROJEKTRED_S2_03_1200x1600-b1847981214ac013383111fc457eb9c5?h=1280&resize=1&w=960',
	// 			feature: ['Action', 'RPG', 'Explore'],
	//             price:29.45,rating:9.4,
	//             discount:50,
	//             discountEnd: "9/25/2022",
	//             supportedPlatform : ["mac", "windows", "linux"]
	// 		},
	// 	];

	let total = {
		price: 0,
		discount: 0
	};
	let MoneyError = [false, 'Not Enough Money'];
	// let GrandTotal: { TotalPrice: number; TotalDiscount: number } = {} as {
	// 	TotalPrice: number;
	// 	TotalDiscount: number;
	// };
	// (GrandTotal.TotalPrice = Math.round((total.price - total.discount) * 100) / 100),
	// 	(GrandTotal.TotalDiscount = Math.round(total.discount * 100) / 100);

	function CountTotal() {
		total.price = 0;
		total.discount = 0;
		if (CartGameList != null) {
			for (let i = 0; i < CartGameList.length; i++) {
				total.price += CartGameList[i].Price;
				total.discount += (CartGameList[i].Price * CartGameList[i].Discount) / 100;
			}
		}
	}
	CountTotal();

	async function GetCardGameList() {
		const reqStruct = {
			UserID: $UserProData.UserID
		};

		await fetch(`/api/profile/cart`, {
			method: 'POST',
			mode: 'no-cors',
			body: JSON.stringify(reqStruct)
		})
			.then((res) => {
				return res.json();
			})
			.then((da) => {
				// console.log('🚀 ~ file: +page.server.ts ~ line 106 ~ myfetch ~ da', da);

				CartGameList = da;
			});
	}

	function BuyAllCartNow() {
		let TotalPrice = Math.round((total.price - total.discount) * 100) / 100;
		console.log('🚀 ~ file: +page.svelte ~ line 145 ~ BuyAllCartNow ~ TotalPrice', TotalPrice);

		if (TotalPrice > $UserProData.Coin) {
			MoneyError[0] = true;
			MoneyError[1] = 'Not Enough Money';
			return;
		}
		GeneralData.Coin=$UserProData.Coin-TotalPrice
		GeneralData.TransactionHistory.push(...GeneralData.UserCart)
		GeneralData.UserCart=[]

		UpdateProfileData()

		UserSettingSelect.set('Transaction');
		goto(`/profile/settings`);
	}

	// function RemoveFromCar(id: number) {
	// 	console.log('remove from cart');
	// 	console.log(id);
	// }
	async function UpdateProfileData() {
		console.log(
			'🚀 ~ file: gameProfile.svelte ~ line 49 ~ UpdateProfileData ~ GeneralData',
			GeneralData
		);
		await fetch('/api/profile/general', {
			// credentials: 'same-origin',
			method: 'POST',
			// mode: 'cors',
			body: JSON.stringify(GeneralData)
		})
			.then((res) => {
				return res.json();
			})
			.then((resdata) => {
				console.log('🚀 ~ file: General.svelte ~ line 148 ~ .then ~ resdata', resdata);
				UserProData.update((n) => {
					return (n = resdata);
				});
				// GeneralData = resdata;
				console.log('🚀 ~ file: General.svelte ~ line 152 ~ .then ~ UserProData', $UserProData);
			});
	}
	function RemoveFromCart(MyGameID: string) {
		if ($UserProData.UserCart != null) {
			let temp: string[] = [];
			for (let i = 0; i < $UserProData.UserCart.length; i++) {
				if ($UserProData.UserCart[i] != MyGameID) {
					temp.push($UserProData.UserCart[i]);
				}
			}
			GeneralData.UserCart = temp;
			UpdateProfileData();
			GetCardGameList();

			// (GrandTotal.TotalPrice = Math.round((total.price - total.discount) * 100) / 100),
			// (GrandTotal.TotalDiscount = Math.round(total.discount * 100) / 100);
		}
		CountTotal();

		// GeneralData.UserCart=$UserProData.UserCart;
		// UpdateProfileData()
		// $UserProData.UserCart.push(GameID);
		// console.log("🚀 ~ file: gameProfile.svelte ~ line 20 ~ AddToCart ~ UserProData", $UserProData)
		// console.log('🚀 ~ file: +page.svelte ~ line 182 ~ RemoveFromCart ~ GrandTotal', GrandTotal);
	}
</script>

<div class="flex max-h-fit min-h-[90%] w-full flex-col justify-start lg:px-48  ">
	<p class="  font-Poppins text-4xl text-slate-200 ">My Cart</p>
	<div
		class="flex w-full flex-row font-Poppins text-slate-200 sm:m-5 sm:flex-col-reverse sm:gap-3 md:m-5 md:flex-col-reverse md:gap-3 xs:m-5 xs:flex-col-reverse xs:gap-3"
	>
		<!-- CART GAMES -->
		<div class="flex w-full flex-col gap-2 ">
			{#if CartGameList != null}
				{#each CartGameList as i}
					<div
						class=" flex h-56  w-[1110px] flex-row items-center rounded-md  bg-[#202020] pl-3 hover:bg-slate-800 sm:w-full md:w-full xs:w-full "
					>
						<img src={i.SmallPosterImage} alt={i.Name} class=" h-[200px] w-[150px] " />
						<div class=" ml-4 mr-2 flex h-full w-full flex-row">
							<div class="flex h-full w-fit flex-col">
								<p class=" mt-4 text-2xl">{i.Name}</p>
								<div class="flex flex-row gap-2">
									{#each i.Feature as j}
										<p class=" text-lg font-thin text-blue-500">{j}</p>
									{/each}
								</div>
								<p class=" font-semibold">Rating {i.Rating}/10</p>
								<div class="grow" />
								<div class=" mb-4 flex flex-row gap-2">
									{#each i.Platform as j}
										{#if j == 'mac'}
											<Apple class="h-6 w-6 fill-slate-400" />
										{:else if j == 'linux'}
											<Linux class="h-6 w-6 fill-slate-400" />
										{:else}
											<Windows class="h-6 w-6 fill-slate-400" />
										{/if}
									{/each}
								</div>
							</div>
							<div class="grow" />
							<div class=" mt-4 mb-3 flex h-full w-fit flex-col">
								<del class=" text-gray-400">${i.Price} </del>
								<p class="text-xl font-bold">${i.Price - (i.Price * i.Discount) / 100}</p>
								<p class="">sales Ends soon</p>
								<div class="grow" />
								<button
									on:click={() => {
										RemoveFromCart(i.GameID);
									}}
									class=" mb-8 text-gray-500 underline hover:text-gray-200 ">Remove</button
								>
							</div>
						</div>
					</div>
				{/each}
			{:else}
				<!-- else content here -->
				<div class=" flex h-full w-full  items-center justify-center">
					<p class=" text-2xl font-thin">No Game in Cart</p>
				</div>
			{/if}
		</div>
		<!-- SUMMARY -->
		<div class=" flex h-fit w-80 flex-col gap-4">
			<p class=" mb-6 text-4xl">Games and Apps Summary</p>
			<div class="flex flex-row">
				<p class="">Price</p>
				<div class=" grow " />
				<p class="">${total.price}</p>
			</div>
			<div class="flex flex-row">
				<p class="">Sale Discount</p>
				<div class=" grow " />
				<p class="">-${Math.round((total.discount) * 100) / 100}</p>
			</div>
			<div class="flex flex-row">
				<p class="">Taxes</p>
				<div class=" grow " />
				<p class=" text-gray-500">Calculated at Checkout</p>
			</div>

			<div class="mt-2 flex flex-row border-t-[1px] border-gray-700 pt-2 ">
				<p class=" text-lg font-semibold">SubTotal</p>
				<div class=" grow " />
				<p class=" text-lg font-semibold">
					${Math.round((total.price - total.discount) * 100) / 100}
				</p>
			</div>
			{#if MoneyError[0]}
				<!-- content here -->

				<p class="text-red-500">{MoneyError[1]}</p>
			{/if}

			<button
				on:click={() => {
					BuyAllCartNow();
				}}
				class=" h-[50px] w-[296px] bg-blue-500 transition-all duration-150 ease-linear hover:bg-sky-500 hover:font-semibold"
			>
				Buy Now
			</button>
		</div>
	</div>
</div>

<style>
	/* your styles go here */
</style>
