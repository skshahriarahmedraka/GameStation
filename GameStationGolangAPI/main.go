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
	// "github.com/gin-contrib/cors"
)

func init() {

	// LOAD ENVIRONMENT VARIABLES
	config.LoadEnvironmentVar()

}

func main() {
	//  SET MODE TO RELEASE
	gin.SetMode(gin.ReleaseMode)


	fmt.Println("🚀✨ Api is started")
	// r := gin.Default()

	r := gin.New()
	
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())

	

	// r.Use(cors.New(conf))
	// ROUTE WITHOUT AUTHENTICATION
	router.RouteWithoutAuth(r)

	// AUTHENTICATION
	r.Use(middleware.Authentication())

	// ROUTE ACCESSABLE AFTER AUTHENTICATION
	router.RouteWithAuth(r)
	

	log.Println("Server is started in PORT ",os.Getenv("HOST_ADDR")," ...👨‍💻 ")
	if e := r.Run(os.Getenv("HOST_ADDR")); e != nil {
		log.Fatalln("❌ ERROR when Server is start   👨‍💻 : ", e)
	}

}
