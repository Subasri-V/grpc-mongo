package rpcService

import (
	"context"
	pro "grpc-mongo/grpc"
	"grpc-mongo/interfaces"
	"grpc-mongo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type RPCServer struct {
	pro.UmimplementedProfileServiceServer
}

var (
	ProfileCollection *mongo.Collection
	ProfileService interfaces.IProfile
)

func (s *RPCServer) CreateProfile(ctx context.Context, req *pro.Profile) (*pro.ProfileResponse, error) {
	dbProfile:=&models.Profile{Name: }
	
	return &pro.ProfileResponse{Name: "Subasri"}, nil

}
