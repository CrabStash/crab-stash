package handlers

import (
	"context"
	"fmt"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"google.golang.org/grpc/status"
)

func GrpcLogin(client pb.AuthServiceClient, user pb.User) (*pb.Token, error) {
	res, err := client.Login(context.Background(), &user)

	if err != nil {
		s, _ := status.FromError(err)
		return &pb.Token{}, fmt.Errorf("error while logging in: %v", s.Message())
	}

	return res, nil
}

func GrpcCreateUser(client pb.AuthServiceClient, user pb.User) error {
	_, err := client.CreateUser(context.Background(), &user)

	if err != nil {
		s, _ := status.FromError(err)
		return fmt.Errorf("error while registering: %v", s.Message())
	}

	return nil
}
