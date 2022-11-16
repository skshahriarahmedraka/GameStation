// package router

// import (
// 	"app/controller"
// 	"app/database"
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// func RoutesWithoutAuth( incomingRoute *gin.Engine){
//     fmt.Println("🚀 ~ file: routers.go ~ line 21 ~ funcUserRoutes ~ database.Client : ", database.Client)
// 	incomingRoute.POST("/user/register",controller.Register())
// 	incomingRoute.POST("/user/login",controller.Login())
// 	incomingRoute.POST("/admin/addproduct",controller.AdminAddProduct())
// 	incomingRoute.GET("/user/productview",controller.ViewProduct())
// 	incomingRoute.GET("/user/search",controller.UserSearch())

// }
package router

import (
	// "app/controller"
	"app/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RouteWithoutAuth(r *gin.Engine) {
	// fmt.Println("🚀 ~ file: routers.go ~ line 21 ~ funcUserRoutes ~ database.Client : ", database.Client)
	DB2 := database.MongodbConnection()
	H := database.DatabaseInitialization(DB2)
	fmt.Println("🚀 ~ file: WithoutAuth.go ~ line 13 ~ funcRouteWithoutAuth ~ RouteWithoutAuth  ", H)
	r.POST("/sveltekit/login", H.SveltekitLogin())
	r.POST("/sveltekit/register", H.SveltekitRegister())
	r.POST("/user/login", H.Login())
	r.POST("/user/register", H.Register())
	r.POST("/user/profileupdate", H.ProfileUpdate())
	r.POST("/user/cart", H.ProfileCartData())// not complete
	r.POST("/user/transactiondata", H.ProfileTransactionData())// not complete
	// r.GET("/user/moneyreq", H.ProfileMoneyReq())
	r.GET("/user/:profileid", H.UserData())

	r.POST("/admin/addproduct", H.AdminAddProduct())

	r.GET("/game/:gameid", H.GetGameData())
	r.POST("/game/addcomment", H.AddGameComment())
	//r.GET("/game/:gameid", H.GetGameData())
	r.POST("/news/addnews", H.AddNews())
	r.GET("/news/:newsid", H.GetNewsByID())
	r.GET("/news", H.NewsList())
	
	r.POST("/user/moneytokenreq", H.ProfileTokenReq())         // completed: user send token request
	r.POST("/user/moneytokenreqlist", H.ProfileTokenReqList()) // complete: user get token request list
	r.POST("/user/rechargewallet", H.ProfileRechargeWallet())  //
	// r.GET("/user/moneyreq", H.ProfileMoneyReq())
	
	r.GET("/admin/adminmoneymanagement", H.AdminMoneyManagement()) // complete: admin money management
	r.POST("/admin/moneyreqaccept", H.AdminMoneyReqAccept())       // complete: admin money request accept
	
	r.GET("/user/browse", H.BrowseGames())
	
	// default filters
	r.GET("/game/mostpopular", H.MostPopular())
	r.GET("/game/trending", H.Trending())
	r.GET("/game/newrelease", H.NewRelease())
	r.GET("/game/topsold", H.TopSold())
	r.GET("/game/toprated", H.TopRated())

	r.GET("/game/carousel", H.GetCarousel())
	// r.GET("/game/carousel", H.GetCarousel())



	//r.GET("/product/:productid", H.ViewProduct())
	r.GET("/user/search", H.UserSearch())
	r.POST("/imgupload", H.ImgUpload())
	r.POST("/imglink", H.ImgLink())
	r.GET("/ws", H.Ws())
}
