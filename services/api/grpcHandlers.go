package main

import (
	"context"
	"fmt"

	pb "github.com/CrabStash/crab-stash/auth/proto"
	"google.golang.org/grpc/status"
)

func GrpcLogin(c pb.AuthServiceClient, user pb.User) (*pb.Token, error) {
	res, err := c.Login(context.Background(), &user)

	if err != nil {
		s, _ := status.FromError(err)
		return &pb.Token{}, fmt.Errorf("error while logging in: %v", s.Message())
	}

	return res, nil
}

func GrpcCreateUser(c pb.AuthServiceClient, user pb.User) error {
	_, err := c.CreateUser(context.Background(), &user)

	if err != nil {
		s, _ := status.FromError(err)
		return fmt.Errorf("error while logging in: %v", s.Message())
	}

	return nil
}
