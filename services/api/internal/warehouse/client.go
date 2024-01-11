package warehouse

import (
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	"github.com/CrabStash/crab-stash/api/internal/utils"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.WarehouseServiceClient
	Utils  *utils.Utils
}

func InitServiceClient() pb.WarehouseServiceClient {
	con, err := grpc.Dial(os.Getenv("WAREHOUSE_MS_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to the warehouse service: %v", err.Error())
	}

	return pb.NewWarehouseServiceClient(con)
}
