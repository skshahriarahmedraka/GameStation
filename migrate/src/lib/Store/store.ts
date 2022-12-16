import { writable } from 'svelte/store';

type NavigationState = "loading" | "loaded" | null;

export default writable<NavigationState>(null);


const  UserProData=writable({
    Name: "",
    UserID: "",
    Accounttype: '',

    ProfileImg: '',
    BannerImg: '',

    Email: '',
    Mobile: '',

    Address: '',
    // Region: '',
    Country: '',
    City: '',
    ZipCode: '',

    Coin: 0,

    TransactionHistory: [] as string[],
    WishList: [] as string[],
    ContactAdminMsg: [] as string[],
    ContactDevMsg: [] as string[],
    UserCart: [] as  string[]
})

const UserSettingSelect=writable("General")



export {UserProData,UserSettingSelect}




