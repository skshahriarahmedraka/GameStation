package controller

import (
	// "app/LogError"
	// "bytes"
	// "encoding/base64"
	// "app/LogError"
	// "encoding/json"
	// "encoding/base64"
	"fmt"
	// "io/ioutil"
	"log"

	"io"
	// "log"
	"net/http"
	"os"

	// "strings"
	// "path/filepath"

	// "github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (H *DatabaseCollections) ImgUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Request.Header.Set()
		// 	c.Request.Header.Set("Access-Control-Allow-Headers", "Content-Type")
			// c.Request.Header.Set("Content-Type", "application/json; charset=UTF-8")
			c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		
			file, header, err := c.Request.FormFile("img")
            // fmt.Println("🚀 ~ file: imgUpload.go ~ line 35 ~ returnfunc ~ err : ", err)
            // fmt.Println("🚀 ~ file: imgUpload.go ~ line 35 ~ returnfunc ~ header : ", header)
            // fmt.Println("🚀 ~ file: imgUpload.go ~ line 35 ~ returnfunc ~ file : ", file)
			if err != nil {
				c.String(http.StatusBadRequest, "Bad request")
				return
			}
			//  Header calls the filename method to get the file name
			filename := header.Filename
			fmt.Println(file, err, filename)
		
		//  Create a file, the file is named FILENAME, the return value OUT here is also a file pointer
			out, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
		
			defer out.Close()
		
		//  Copy file content to out
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}

		// File saved successfully. Return proper result
        fmt.Println("🚀 ~ file: imgUpload.go ~ line 62 ~ returnfunc ~ Your file has been successfully uploaded : ")
		c.JSON(http.StatusOK, gin.H{"message": "Your file has been successfully uploaded."})
		// c.String(http.StatusOK,"Your file has been successfully uploaded")
		// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! Your file has been successfully uploaded", filename))

	}
}

type My struct {
	Link string `json:"Link"`
}

