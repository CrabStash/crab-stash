package main

import (
	"log"
	"net"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	"github.com/CrabStash/crab-stash/user/internal/db"
	"github.com/CrabStash/crab-stash/user/internal/server"

	"google.golang.org/grpc"
)

var err error

func main() {
	h := db.Init()

	defer h.DB.Close()

	lis, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", ":50054")

	s := server.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}

}
