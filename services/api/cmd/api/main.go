package main

import (
	"log"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/api/internal/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to GRPC server: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)

	r := gin.Default()
	routes.AuthRoutes(r, c)

	r.Run(":8080")
}
