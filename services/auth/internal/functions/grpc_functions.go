package functions

import (
	"context"
	"fmt"
	"log"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.AuthServiceServer
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*emptypb.Empty, error) {

	isUserCreated := GetUser(req.Email)
	if isUserCreated.Email != "" {
		return &emptypb.Empty{}, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("User already exists"),
		)
	}

	pwd, err := utils.HashPassword(req.Passwd)

	if err != nil {
		log.Printf("Error while hashing password: %v", err)
		return &emptypb.Empty{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while hashing password"),
		)
	}
	user := pb.User{
		Email:  req.Email,
		Passwd: pwd,
	}
	err = CreateUser(user)

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while creating user"),
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.User) (*pb.Token, error) {
	user := GetUser(req.Email)

	if user.Email == "" {
		return &pb.Token{}, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Wrong password or email"),
		)
	}

	ok := utils.CheckPasswordHash(req.Passwd, user.Passwd)
	if !ok {
		return &pb.Token{}, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Wrong password or email"),
		)
	}

	token, err := utils.SignJWT(user.Id)
	if err != nil {
		return &pb.Token{}, status.Errorf(
			codes.Internal,
			fmt.Sprintf("%v", err),
		)
	}

	return &pb.Token{JwtToken: token}, nil
}