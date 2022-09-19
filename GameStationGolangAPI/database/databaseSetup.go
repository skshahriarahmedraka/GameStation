package database

import (
	"app/config"
	"app/errors"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Client *mongo.Client= DBSetup() 

func DBSetup()*mongo.Client{
	config.LoadEnvironmentVar()

    fmt.Println("🚀 ~ file: databaseSetup.go ~ line 15 ~ funcDBSetup ~ DBSetup : ")
	mongouri:= fmt.Sprintf("mongodb://%v:%v@%v:%v/?maxPoolSize=%v&w=majority" ,os.Getenv("MONGO_USER"),os.Getenv("MONGO_PASSWORD"),os.Getenv("MONGO_IP"),os.Getenv("MONGO_PORT"),os.Getenv("MONGO_MAXPOOLSIZE"))
    fmt.Println("🚀 ~ file: databaseSetup.go ~ line 13 ~ funcDBSetup ~ mongouri : ", mongouri)
	client ,err :=mongo.NewClient(options.Client().ApplyURI(mongouri))
    errors.ERROR("🚀 ~ file: databaseSetup.go ~ line 17 ~ funcDBSetup ~ err : ",err)
	
	
	ctx,cancel:=context.WithTimeout(context.Background(), 10* time.Second)

	defer cancel()

	err =client.Connect(ctx)
    errors.ERROR("🚀 ~ file: databaseSetup.go ~ line 25 ~ funcDBSetup ~ err : ", err)
	err =client.Ping(context.TODO(),nil)
    errors.ERROR("🚀 ~ file: databaseSetup.go ~ line 27 ~ funcDBSetup ~ err : ", err)
	fmt.Println("sucessfully connected to database")
	return client

}

func UserDatabase(client *mongo.Client,CollectionName string) *mongo.Collection {
	collection :=client.Database("userdb").Collection(CollectionName)
	return collection
}


func ProductDatabase(client *mongo.Client,CollectionName string) *mongo.Collection{
	collection := client.Database("productdb").Collection(CollectionName)
	return collection 
}