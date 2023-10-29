package core

import (
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CoreServiceClient
}

func InitServiceClient() pb.CoreServiceClient {
	con, err := grpc.Dial(os.Getenv("CORE_MS_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the core service: %v", err.Error())
	}

	return pb.NewCoreServiceClient(con)
}
