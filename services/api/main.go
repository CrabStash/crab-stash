package main

import (
	"log"

	pb "github.com/CrabStash/crab-stash/auth/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("crabstash-auth:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to GRPC server: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)

	r := gin.Default()
	AuthRoutes(r, c)

	r.Run(":8080")
}
