package main

import (
	"log"
	"net"

	pb "github.com/CrabStash/crab-stash/auth/proto"
	"github.com/surrealdb/surrealdb.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	addr string = "0.0.0.0:50051"
)

type Server struct {
	pb.AuthServiceServer
}

var db *surrealdb.DB
var err error

func main() {
	db, err = surrealdb.New("ws://localhost:8000/rpc")
	defer db.Close()

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v\n", err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		log.Fatalf("Failed to signin to db: %v\n", err)
	}

	if _, err = db.Use("users", "users"); err != nil {
		log.Fatalf("Failed to use users/users: %v\n", err)
	}

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
