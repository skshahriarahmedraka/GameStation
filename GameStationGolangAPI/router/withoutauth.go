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
	r.POST("/user/register", H.Register())
	r.POST("/sveltekit/register", H.SveltekitRegister())
	r.POST("/user/login", H.Login())
	r.POST("/sveltekit/login", H.SveltekitLogin())
	r.POST("/admin/addproduct", H.AdminAddProduct())
	r.GET("/game/:gameid", H.GetGameData())
	r.GET("/user/:profileid", H.UserData())

	r.GET("/product/:productid", H.ViewProduct())
	r.GET("/user/search", H.UserSearch())
	r.POST("/imgupload", H.ImgUpload())
	r.POST("/imglink", H.ImgLink())

}
