package main

import (
	"log"
	"net"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/auth/internal/functions"
	surrealdb "github.com/surrealdb/surrealdb.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	addr string = "auth:50051"
)

var err error

func main() {
	functions.DB, err = surrealdb.New("ws://surrealdb:8000/rpc")
	defer functions.DB.Close()

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v\n", err)
	}

	if _, err = functions.DB.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		log.Fatalf("Failed to signin to db: %v\n", err)
	}

	if _, err = functions.DB.Use("users", "users"); err != nil {
		log.Fatalf("Failed to use users/users: %v\n", err)
	}

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", addr)

	s := grpc.NewServer()
	serverStruct := functions.Server{}
	pb.RegisterAuthServiceServer(s, &serverStruct)
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
