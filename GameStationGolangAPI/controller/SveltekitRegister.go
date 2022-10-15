package controller

import (
	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

//var Validate = validator.New()

func (H *DatabaseCollections) SveltekitRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(http.StatusCreated, "Successfully Signed Up!!")
		//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		//	defer cancel()
		//	var u model.User
		//	err := c.BindJSON(&u)
		//	LogError.LogError("❌ ~ file: signup.go ~ line 22 ~ returnfunc ~ err : ", err)
		//	if err != nil {
		//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//		return
		//	}
		//	err = Validate.Struct(u)
		//	LogError.LogError("❌ ~ file: signup.go ~ Validate.Struct(u) ~ err : ", err)
		//	if err != nil {
		//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//		return
		//	}
		//	// SEARCH EMAIL
		//	count, err := H.Mongo.Collection("usercol").CountDocuments(ctx, bson.M{"email": u.Email})
		//	LogError.LogError("❌ ~ file: signup.go ~ line 36 ~ UserCollection.CountDocuments(ctx, bson.M{\"email\": u.Email}) ~ err : ", err)
		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		//		return
		//	}
		//	if count > 0 {
		//		c.JSON(http.StatusBadRequest, gin.H{"error": "User Already register"})
		//		return
		//	}
		//	//MOBILE
		//	count, err = H.Mongo.Collection("usercol").CountDocuments(ctx, bson.M{"mobile": u.Mobile})
		//	LogError.LogError("❌ ~ file: signup.go ~ line 50 ~ returnfunc ~ err : ", err)
		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		//		return
		//	}
		//	if count > 0 {
		//		c.JSON(http.StatusBadRequest, gin.H{"error": "mobile already in use"})
		//		return
		//	}
		//
		//	hashpass,err:= BcryptHashPass(u.Password)
		//	LogError.LogError("❌ ~ file: signup.go ~ line 61 ~ returnfunc ~ err : ", err)
		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		//		return
		//	}
		//	u.Password =hashpass
		//	u.CreatedAt ,err =time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		//	LogError.LogError("❌ ~ file: signup.go ~ line 68 ~ returnfunc ~ err : ", err)
		//	u.UpdatedAt,err=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		//	LogError.LogError("❌ ~ file: signup.go ~ line 70 ~ returnfunc ~ err : ", err)
		//	u.ID=primitive.NewObjectID()
		//	u.UserID = u.ID.Hex()
		//
		//	jwtToken,refreshToken,err:= token.TokenGenerate(u.Email,u.FirstName,u.LastName, u.UserID)
		//	fmt.Println("🚀 ~ file: register.go ~ line 77 ~ returnfunc ~ refreshToken : ", refreshToken)
		//	fmt.Println("🚀 ~ file: register.go ~ line 77 ~ returnfunc ~ jwtToken : ", jwtToken)
		//	LogError.LogError("❌ ~ file: signup.go ~ line 77 ~ returnfunc ~ err : ", err)
		//
		//	// _=jwtToken
		//	// _=refreshToken
		//	u.UserCart=make([]model.UserProduct, 0)
		//	u.Address=make([]model.Address, 0)
		//	u.OrderStatus=make([]model.Order, 0)
		//	res,err:=H.Mongo.Collection("usercol").InsertOne(ctx,u)
		//	LogError.LogError("❌ ~ file: signup.go ~ line 83 ~ returnfunc ~ err : ", err)
		//	if err!=nil {
		//		c.JSON(http.StatusInternalServerError,gin.H{"error":"not created"})
		//		return
		//	}else {
		//
		//		fmt.Println("🚀 ~ file: signup.go ~ line 83 ~ returnfunc ~ res : ", res)
		//	}
		//
		//	// c.Writer.Header().Set("Auth1", jwtToken)
		//	// c.Writer.Header().Set("Auth1Refresh", refreshToken)
		//	c.SetCookie("Auth1",jwtToken,60*60*24,"/","localhost",false , true)
		//	c.SetCookie("Auth1Refresh",refreshToken,60*60*24,"/","localhost",false , true)
		//	c.JSON(http.StatusOK,"successfully signed up")

		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Request.Header.Set("Content-Type", "application/json")
		c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

		//var resError model.ErrorRes

		var user model.UserData
		//err := json.NewDecoder(r.Body).Decode(&user)

		err := c.BindJSON(&user)
		if err != nil {
			LogError.LogError("❌🔥 error in c.bindjson() ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			//http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("🚀", user)

		user.Accounttype = "normal"
		user.BannerImg = ""
		user.Coin = 0.0
		//user.FrinedListID = []string{}
		user.ProfileImg = ""
		myid := uuid.New()
		user.UserID = myid.String()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		defer cancel()

		// SEARCH EMAIL
		count, err := H.Mongo.Collection("usercol").CountDocuments(ctx, bson.M{"Email": user.Email})
		if err != nil {
			LogError.LogError("❌🔥 error in mongodb connection  ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			LogError.LogError("❌🔥 error in User already registered ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
			//w.WriteHeader(http.StatusBadRequest)
			//resError.ErrorRes = "User already registered"
			//_ = json.NewEncoder(w).Encode(resError)
			return
		}
		//MOBILE
		count, err = H.Mongo.Collection("usercol").CountDocuments(ctx, bson.M{"Mobile": user.Mobile})
		if err != nil {
			LogError.LogError("❌🔥 error in mongodb countDocument mobile connection error ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb countDocument mobile connection error"})


			return

		}
		if count > 0 {
			LogError.LogError("❌🔥 error in mobile number already registered ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "mobile number already registered"})


			// c.JSON(http.StatusBadRequest, gin.H{"error": "mobile already in use"})

			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
		if err != nil {
			log.Println(err)
		}
		//return string(hash)

		res, err := H.Mongo.Collection("usercol").InsertOne(ctx, user)
		LogError.LogError("🚀 ~ file: register.go ~ line 102 ~ func ~ err : ", err)

		if err == nil {
			fmt.Println("🚀 ~ file: register.go ~ line 116 ~ func ~ res : ", res)
		}

		mongoRes, err := H.Mongo.Collection("userdb").InsertOne(ctx, user)
		fmt.Println("🚀 ~ file: register.go ~ line 80 ~ func ~ mongoRes : ", mongoRes)

		if err != nil {
			LogError.LogError("❌🔥 error in mongodb Connection error ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb Connection error"})

			return
		}

		expirationTime := time.Now().Add(time.Minute * 100)

		claims := &model.TokenClaims{
			Name: user.Name,
			Email:    user.Email,
			UserID:   user.UserID,
			// Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			LogError.LogError("❌🔥 error in StatusInternalServerError token generation  ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

			return
		}
		////////////////////////////////
		var SendJwt struct {
			JWT string
		}
		SendJwt.JWT = tokenString

		c.JSON(http.StatusOK, SendJwt)


	}
}

//func BcryptHashPass(s string) (string, error){
//	b,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.DefaultCost)
//	LogError.LogError("❌ ~ file: signup.go ~ line 55 ~ funcBcryptHashPass ~ err : ", err)
//	if err!=nil {
//		return "",err
//	}else {
//		return string(b),nil
//	}
//}
//
//func BcryptVerifyPass(userpass string, Databasepass string) (string,bool){
//	fmt.Println("🚀 ~ file: register.go ~ line 116 ~ funcBcryptVerifyPass ~ Databasepass : ", Databasepass)
//	fmt.Println("🚀 ~ file: register.go ~ line 116 ~ funcBcryptVerifyPass ~ userpass : ", userpass)
//	err:=bcrypt.CompareHashAndPassword([]byte(Databasepass),[]byte(userpass))
//	LogError.LogError("❌ ~ file: register.go ~ line 119 ~ funcBcryptVerifyPass ~ err : ", err)
//	if err!=nil{
//		msg:="Email or Password is incorrect"
//		return msg, false
//	} else {
//		return "",true
//	}
//}
