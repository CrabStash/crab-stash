package user

import (
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.UserServiceClient
	Utils  *utils.Utils
}

func InitServiceClient() pb.UserServiceClient {
	con, err := grpc.Dial(os.Getenv("USER_MS_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the auth service: %v", err.Error())
	}

	return pb.NewUserServiceClient(con)
}
