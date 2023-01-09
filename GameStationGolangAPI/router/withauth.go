package router

import (
	// "app/H"

	"app/database"

	"github.com/gin-gonic/gin"
)

func RouteWithAuth(r *gin.Engine) {
	DB2 := database.MongodbConnection()
	H := database.DatabaseInitialization(DB2)
	// these routes need authentication before use
	r.GET("/home2", H.Home())
	//r.GET("/seecartlist", H.AddtoCart())
	//r.GET("/removeitem", H.Removeitem())
	//r.GET("/listcart", H.GetItemFromCart())
	//r.POST("/addaddress", H.Addaddress())
	//
	//r.POST("/edithomeaddress", H.Edithomeaddress())
	//r.POST("/editworkaddresss", H.Edithomeaddress())
	//r.GET("/deleteaddress", H.Deleteaddress())
	//r.GET("/cartchackout", H.Cartcheckout())
	//r.GET("/instantbuy", H.Instantbuy())

}
