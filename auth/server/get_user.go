package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/CrabStash/crab-stash/auth/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type User struct {
	Email  string `json:"email,omitempty"`
	Passwd string `json:"passwd"`
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*emptypb.Empty, error) {
	log.Printf("%v %v\n", req.Email, req.Passwd)
	user := User{
		Email:  req.Email,
		Passwd: req.Passwd,
	}
	_, err := db.Create("user", user)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal server error :c"),
		)
	}
	return &emptypb.Empty{}, nil
}
