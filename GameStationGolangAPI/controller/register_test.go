package controller_test

import (
	"app/config"
	"app/model"
	"encoding/json"
	"fmt"

	// "app/controller"
	"app/database"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	// "reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "go.mongodb.org/mongo-driver/mongo"
)

func init() {
	config.LoadEnvironmentVar()

}
func TestDatabaseCollections_Register(t *testing.T) {
	// SEED FOR gofakeit 
	faker := gofakeit.NewCrypto()
	gofakeit.SetGlobalFaker(faker)

	// type fields struct {
	// 	Mongo *mongo.Database
	// }
	DB2 := database.MongodbConnection()
	H := database.DatabaseInitialization(DB2)

	mockResponse := `{"message":"successfully signed up"}`
	r := SetUpRouter()
	r.POST("/api/register", H.Register())

	regData := model.RegModel{}
	gofakeit.NewCrypto().Rand.Seed(1)
	regData.Name = gofakeit.Name()
	fmt.Println("🚀TEST regData.Name : ", regData.Name)
	regData.Email = gofakeit.Email()
	fmt.Println("🚀TEST regData.Email : ", regData.Email)
	regData.Password = gofakeit.Password(true, true, true, true, true, 10)
	fmt.Println("🚀TEST regData.Password : ", regData.Password)
	
	SentRegData, _ := json.Marshal(regData)
	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(SentRegData))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
