package database

import (
	// "app/controller"
	"app/LogError"
	"context"
	"fmt"
	"time"

	"os"

	// "log"

	// "go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbConnection() *mongo.Database {

	// mongoURI:= fmt.Sprintf("mongodb://%v:%v@%v:%v/?maxPoolSize=%v&w=majority" ,os.Getenv("MONGO_USER"),os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_IP"),os.Getenv("MONGO_PORT"),os.Getenv("MONGO_MAXPOOLSIZE"))

	// // mongoURI := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=" + os.Getenv("MONGO_MAXPOOLSIZE") + "&w=" + os.Getenv("MONGO_W")

	// client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	log.Println("❌ failed to NewCient Client", err)
	// }
	// // ctx := context.Background()
	// ctx,cancel:=context.WithTimeout(context.Background(), 10* time.Second)

	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal("❌ Failed  to connect to mongodb db", err)
	// }
	// // defer client.Disconnect(ctx)
	// Mydb := client.Database("userdb")
	// fmt.Println("✨🥰 ~ file: mongodb.go ~ line 32 ~ funcMongodbConnection ~ Mydb : ", Mydb)

	// // Mycol := Mydb.Collection("stackdb")


	// return Mydb
	mongouri:= fmt.Sprintf("mongodb://%v:%v@%v:%v/?maxPoolSize=%v&w=majority" ,os.Getenv("MONGO_USER"),os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_IP"),os.Getenv("MONGO_PORT"),os.Getenv("MONGO_MAXPOOLSIZE"))
    fmt.Println("🚀 ~ file: databaseSetup.go ~ line 13 ~ funcDBSetup ~ mongouri : ", mongouri)
	client ,err :=mongo.NewClient(options.Client().ApplyURI(mongouri))
    LogError.LogError("🚀 ~ file: databaseSetup.go ~ line 17 ~ funcDBSetup ~ err : ",err)
	
	
	ctx,cancel:=context.WithTimeout(context.Background(), 10* time.Second)

	defer cancel()

	err =client.Connect(ctx)
    LogError.LogError("❌ ~ file: databaseSetup.go ~ line 25 ~ funcDBSetup ~ err : ", err)
	err =client.Ping(context.TODO(),nil)
    LogError.LogError("❌ ~ file: databaseSetup.go ~ line 27 ~ funcDBSetup ~ err : ", err)
	Mydb := client.Database("userdb")
	if err== nil {

		fmt.Println("⚡😍 sucessfully connected to database")
	}
	return Mydb

}


