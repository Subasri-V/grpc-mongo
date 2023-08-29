package main

import (
	"context"
	"fmt"
	"net"

	"grpc-mongo/config"
	"grpc-mongo/constants"
	pro "grpc-mongo/grpc/profile"

	rpc "grpc-mongo/controllers"
	"grpc-mongo/services"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	rpc.ProfileCollection = config.GetCollection(client, "bankdb", "profiles")
	rpc.ProfileService = services.NewProfileServiceInit(rpc.ProfileCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterProfileServiceServer(s, &rpc.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
