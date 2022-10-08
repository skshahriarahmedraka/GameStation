package main

import (
	"fmt"
	"log"
	"os"

	"app/config"
	// "app/controller"
	"app/middleware"
	"app/router"

	"github.com/gin-gonic/gin"
	// "github.com/gobuffalo/packr/v2"
	// "github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func init(){

	// LOAD ENVIRONMENT VARIABLES
	config.LoadEnvironmentVar()

}

func main() {
	//  SET MODE TO RELEASE
	gin.SetMode(gin.ReleaseMode)
	
	//  fmt.Println("environment var :  ", os.Getenv("NAME"))
	//  fmt.Println("environment var :  ", os.Getenv("NAME")=="SK SHAHRIAR AHMED RAKA")
	// fmt.Println("environment var :", os.Getenv("POSTGRES_TIMEZONE"))
	// fmt.Println("environment var :", os.Getenv("POSTGRES_TIMEZONE")=="Asia/Dhaka")
	fmt.Println("🚀✨ Api is started")

	r := gin.New()
	r.Use(gin.Logger())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true
  
	r.Use(cors.New(config))
	// ROUTE WITHOUT AUTHENTICATION
	router.RouteWithoutAuth(r)

	// AUTHENTICATION
	r.Use(middleware.Authentication())

	// ROUTE ACCESSABLE AFTER AUTHENTICATION
	router.RouteWithAuth(r)
	// r.GET("/addtocart", controller.AddtoCart())
	// r.GET("/removeitem", controller.Removeitem())
	// r.GET("/listcart", controller.GetItemFromCart())
	// r.POST("/addaddress", controller.Addaddress())
	// r.POST("/edithomeaddress", controller.Edithomeaddress())
	// r.POST("/editworkaddresss", controller.Edithomeaddress())
	// r.GET("/deleteaddress", controller.Deleteaddress())
	// r.GET("/cartchackout", controller.Cartcheckout())
	// r.GET("/instantbuy", controller.Instantbuy())

	log.Println("Server is started in PORT 8001 ...👨‍💻 ")
	if e := r.Run(os.Getenv("HOST") + ":" + os.Getenv("PORT")); e != nil {
		log.Fatalln("❌ ERROR when Server is start   👨‍💻 : ", e)
	}

}
