package auth

import (
	"log"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	con, err := grpc.Dial("auth-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the auth service: %v", err.Error())
	}

	return pb.NewAuthServiceClient(con)
}
