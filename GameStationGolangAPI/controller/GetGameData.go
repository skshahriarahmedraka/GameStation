package controller

import (
	// "app/LogError"
	// "app/model"
	// "context"
	// "fmt"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) GetGameData() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")

		// var myGameData model.Gamedata
		// err := c.Bind(&myGameData)
		// if err != nil {
		// 	LogError.LogError("🚀 ~ file: adminAddProduct.go ~ line 9 ~ returnfunc ~ err : ", err)
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		// }
		// fmt.Println("🚀 ~ file: adminAddProduct.go ~ line 30 ~ returnfunc ~ myGameData : ", myGameData)

		// // Validate
		// err = Validate.Struct(myGameData)
		// if err != nil {
		// 	LogError.LogError("❌ ~ file: signup.go ~ Validate.Struct(u) ~ err : ", err)
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// } else {
		// 	fmt.Println("🚀 ~ file: adminAddProduct.go ~ line 25 ~ Validate ~ Successfull : ")

		// }

		// // insert in mongodb

		// coll := H.Mongo.Collection("GameData")

		// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		// defer cancel()
		// res, err := coll.InsertOne(ctx, myGameData)

		// if err != nil {
		// 	LogError.LogError("🚀 ~ file: adminAddProduct.go ~ line 46 ~ returnfunc ~ err : ", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		// 	return
		// } else {
		// 	fmt.Println("🚀 ~ file: adminAddProduct.go ~ line 44 ~ returnfunc ~ res : ", res)
		// 	fmt.Println("🚀 ~ file: adminAddProduct.go ~ line 44 ~ returnfunc ~ res : ", res.InsertedID)

		// }

		// var req interface{}

		// fmt.Println("🚀 ~ file: adminAddProduct.go ~ line 17 ~ returnfunc ~ req : ", req)

		c.JSON(http.StatusOK, gin.H{"ok": " accepted bro"})

	}
}
