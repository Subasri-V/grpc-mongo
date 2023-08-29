package server

import (
	"context"
	"grpc-mongo/config"
	"grpc-mongo/services"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDatabase(client *mongo.Client) {
	rpc.ProfileCollection = config.GetCollection(client,"bankdb","profiles");
	rpc.ProfileService=services.NewProfileServiceInit(rpc.ProfileCollection,context.Background())
}

func main(){
	mongoClient,err:=config.ConnectDataBase()
	defer mongoClient.Disconnect()
	
}
