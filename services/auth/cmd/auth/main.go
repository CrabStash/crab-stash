package main

import (
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/auth/internal/db"
	"github.com/CrabStash/crab-stash/auth/internal/redis"
	"github.com/CrabStash/crab-stash/auth/internal/server"
	"github.com/CrabStash/crab-stash/auth/internal/utils"
	"google.golang.org/grpc"
)

var err error

func main() {
	h := db.Init()
	r := redis.Init()

	defer h.DB.Close()

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %v\n", ":50051")

	tokenExp, err := strconv.ParseInt(os.Getenv("TOKEN_EXP"), 10, 16)
	if err != nil {
		log.Fatal("error parsing token_exp")
	}

	refreshExp, err := strconv.ParseInt(os.Getenv("REFRESH_EXP"), 10, 16)
	if err != nil {
		log.Fatal("error parsing token_exp")
	}

	jwt := utils.JwtWrapper{
		TokenSecret:   os.Getenv("TOKEN_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
		TokenPublic:   os.Getenv("TOKEN_PUBLIC"),
		RefreshPublic: os.Getenv("REFRESH_PUBLIC"),
		TokenExp:      uint16(tokenExp),
		RefreshExp:    uint16(refreshExp),
	}

	s := server.Server{
		H:   h,
		Jwt: jwt,
		R:   r,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}

}
