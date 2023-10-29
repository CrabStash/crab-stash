package main

import (
	"log"
	"net"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	"github.com/CrabStash/crab-stash/core/internal/db"
	"github.com/CrabStash/crab-stash/core/internal/server"

	"google.golang.org/grpc"
)

var err error

func main() {
	h := db.Init()

	defer h.DB.Close()

	lis, err := net.Listen("tcp", ":50055")

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", ":50055")

	s := server.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoreServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}

}
