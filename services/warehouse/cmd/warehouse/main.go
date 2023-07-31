package main

import (
	"log"
	"net"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	"github.com/CrabStash/crab-stash/warehouse/internal/db"
	"github.com/CrabStash/crab-stash/warehouse/internal/server"
	"google.golang.org/grpc"
)

var err error

func main() {
	h := db.Init()

	defer h.DB.Close()

	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", ":50052")

	s := server.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterWarehouseServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}

}
