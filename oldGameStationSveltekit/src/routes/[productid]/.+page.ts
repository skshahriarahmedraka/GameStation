import { json } from '@sveltejs/kit';
/** @type {import('./$types').PageLoad} */
export async  function load({ params }) {
	// console.log(params.productid);
    // let mydata:any

    // const res = await fetch(
    //     `/${params.productid}`,{
        
    //     }
    //   );
    // mydata = await res.json();



	// await fetch(`http://localhost:8001/game/${params}`, {
	// 	method: 'GET',
	// 	mode: 'no-cors',
	// 	// headers: { 'Content-Type': 'application/json', Accept: 'application/json' },
    //     // body :JSON.stringify({"GameId":params})

	// }).then((res)=>{
    //     res.json()
    // }).then((d) => {
    //     // console.log("🚀 ~ file: +page.ts ~ line 14 ~ load ~ d", d)
    //     mydata=d        
    // }).catch((err)=> {
    //     console.log("🚀 ~ file: +page.ts ~ line 14 ~ load ~ err", err)
        
    // })

    return {
        
    }

}
