package router

import (
	"app/controller"
	"app/database"
	"fmt"

	"github.com/gin-gonic/gin"
)


func UserRoutes( incomingRoute *gin.Engine){
    fmt.Println("🚀 ~ file: routers.go ~ line 21 ~ funcUserRoutes ~ database.Client : ", database.Client)
	incomingRoute.POST("/user/Register",controller.Register())
	incomingRoute.POST("/user/login",controller.Login())
	incomingRoute.POST("/admin/addproduct",controller.AdminAddProduct())
	incomingRoute.GET("/user/productview",controller.ViewProduct())
	incomingRoute.GET("/user/search",controller.UserSearch())


}
