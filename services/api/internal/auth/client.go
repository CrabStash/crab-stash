package auth

import (
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	con, err := grpc.Dial(os.Getenv("AUTH_MS_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the auth service: %v", err.Error())
	}

	return pb.NewAuthServiceClient(con)
}
