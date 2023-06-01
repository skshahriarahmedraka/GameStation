import { writable } from 'svelte/store';


type NavigationState = "loading" | "loaded" | null;

export default writable<NavigationState>(null);

const syn1 = new URL('../../lib/images/assassinscreedsyndicate/1.jpg', import.meta.url).href;
const syn2 = new URL('../../lib/images/assassinscreedsyndicate/2.jpg', import.meta.url).href;
const syn3 = new URL('../../lib/images/assassinscreedsyndicate/3.jpg', import.meta.url).href;
const syn4 = new URL('../../lib/images/assassinscreedsyndicate/4.jpg', import.meta.url).href;
const syn5 = new URL('../../lib/images/assassinscreedsyndicate/5.jpg', import.meta.url).href;
// const syn6 = new URL('../../lib/images/assassinscreedsyndicate/6.jpg', import.meta.url).href;
const syn7 = new URL('../../lib/images/assassinscreedsyndicate/7.jpg', import.meta.url).href;

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


const SampleGameData=writable( {
    GameID: '1234',
    Name: "Assassin's Creed Syndicate",
    Moto: 'London, 1868. In the heart of the Industrial Revolution, play as Jacob Frye - a brash and charismatic Assassin.',

    LogoImage:
        'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACV_UCS12483_EGST_Logo_Color_1000x407-1000x407-48fad9206ee895c599365e2b7c5bebcab099bd26.png?h=270&resize=1&w=480',
    BigPosterImage:
        'https://cdn2.unrealengine.com/Diesel%2Fproductv2%2Fassassins-creed-syndicate%2Fhome%2FACS-STD-2560x1440-635b7b6c86f18730071426375e7c4fe0bd831ddd.jpg',
    SmallPosterImage:
        'https://cdn1.epicgames.com/epic/offer/AC_Syndicate_Portrait-1280x1420-b74c2aa94670d9e97cc6ddab0a5d4dd0.png?h=854&resize=1&w=640',

    OtherImages: [syn1, syn2, syn3, syn4, syn5, syn7],
    Genres: ['Action', 'RPG', 'Adventure'],
    Feature: ['Single Player', 'Achivement', 'Control', 'Support'],
    Description: `As Jacob Frye, a young and reckless Assassin, use your skills to help those trampled by the march of progress.
    Travel the city at the height of the Industrial Revolution and meet iconic historical figures. From Westminster to Whitechapel, you will come across Darwin, Dickens, Queen Victoria… and many more.
    As a gang leader, strengthen your stronghold and rally rival gang members to your cause, in order to take back the capital from the Templars’ hold.`,
    FollowUs: {
        Facebook: 'https://www.facebook.com/epicgames',
        Discord: 'https://discord.com/app',
        Youtube: 'https://www.youtube.com/channel/UC5Qk8mWBwtMyEj7iQQYRk1A',
        Twitter: 'https://twitter.com/epicgames',
        Site: 'https://store.epicgames.com/en-US/p/assassins-creed-syndicate'
    },
    Rating: 9,
    RatingGivenBy: {
        'PC Gamer':
            'The Assassin’s Creed series has been running for eight years, and that kind of longevity doesn’t happen without taking a few risks. To combat fatigue, each entry attempts to punctuate the familiar with new elements. Not all of these experiments pay off, ',
        IGN: 'Syndicate smartly negotiates this internal conflict by dramatising it in the form of its twin playable characters, Evie and Jacob Frye. The former is a devout Assassin, intent on stopping the Templars by tracking down the remaining pieces of Eden. Her brother Jacob, however, is a pragmatist – a social reformer who rails against the Assassin’s burdensome legacy',
        'Game Informer':
            "Syndicate moves Assassin's Creed forward with a solid new adventure, a beautiful London playground and a renewed sense of fun"
    },
    Minspec: {
        OS: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)',
        CPU: 'Intel Core i5 2400s @ 2.5 GHz / AMD FX 6350 @ 3.9 GHz',
        GPU: 'NVIDIA GeForce GTX 660 / AMD Radeon R9 270 (2GB VRAM with Shader Model 5.0)',
        Memory: '6 GB RAM',
        Storage: '100 GB',
        'Sound Card': 'DirectX Compatible Sound Card with latest drivers'
    },
    Recomendedspec: {
        OS: 'Windows 7 SP1 or Windows 8.1 or Windows 10 (64bit versions)',
        CPU: 'Intel Core i7 3770 @ 3.5 GHz / AMD FX 8350 @ 4.0 GHz',
        GPU: 'NVIDIA GeForce GTX 760 (4GB) or the newer GTX 970 (4GB) / AMD Radeon R9 280X (3GB of VRAM) or better',
        Memory: '8 GB RAM',
        Storage: '200 GB',
        'Sound Card': 'DirectX Compatible Sound Card with latest drivers'
    },
    Price: 34.55,
    Discount: 25,
    Developer: 'Ubisoft',
    Publisher: 'Ubisoft',
    Released: '02/20/20',
    Platform: ['windows', 'linux', 'mac'],
    Players: 23478238
})



export {UserProData,UserSettingSelect,SampleGameData}




