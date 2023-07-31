package server

import (
	"context"
	"fmt"
	"time"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/auth/internal/db"
	"github.com/CrabStash/crab-stash/auth/internal/redis"
	"github.com/CrabStash/crab-stash/auth/internal/utils"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	H   db.Handler
	R   redis.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user, err := s.H.GetUserByEmail(req.Email)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}
	if user.Email != "" {
		return &pb.RegisterResponse{}, fmt.Errorf("user already exists")
	}

	err = s.H.CreateUser(req)
	if err != nil {
		return &pb.RegisterResponse{}, err
	}
	return &pb.RegisterResponse{
		Status:   "ok",
		Response: "user created",
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.H.GetUserByEmail(req.Email)
	if err != nil {
		return &pb.LoginResponse{}, err
	}

	if user.Email != req.Email {
		return &pb.LoginResponse{}, fmt.Errorf("wrong email or password")
	}

	ok := utils.CheckPasswordHash(req.Passwd, user.Passwd)

	if !ok {
		return &pb.LoginResponse{}, fmt.Errorf("wrong email or password")
	}

	token, tokenUUID, err := s.Jwt.SignJWT(user.Id, false)
	if err != nil {
		return &pb.LoginResponse{}, fmt.Errorf("%v", err.Error())
	}

	refresh, refreshUUID, err := s.Jwt.SignJWT(user.Id, true)
	if err != nil {
		return &pb.LoginResponse{}, fmt.Errorf("%v", err.Error())
	}

	now := time.Now()

	errToken := s.R.Set(ctx, tokenUUID, user.Id, time.Unix(int64(s.Jwt.TokenExp)*int64(time.Hour*24), 0).Sub(now)).Err()
	if errToken != nil {
		return &pb.LoginResponse{}, fmt.Errorf("%v", err.Error())
	}

	errRefresh := s.R.Set(ctx, refreshUUID, user.Id, time.Unix(int64(s.Jwt.RefreshExp)*int64(time.Hour*24), 0).Sub(now)).Err()
	if errRefresh != nil {
		return &pb.LoginResponse{}, fmt.Errorf("%v", err.Error())
	}

	return &pb.LoginResponse{
		Token:   token,
		Refresh: refresh,
	}, nil

}

func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	_, token_uuid, err := s.Jwt.ValidateJWT(req.Token, false)
	if err != nil {
		return &pb.LogoutResponse{}, fmt.Errorf("error validating token: %v", err.Error())
	}

	_, refresh_uuid, err := s.Jwt.ValidateJWT(req.Refresh, true)
	if err != nil {
		return &pb.LogoutResponse{}, fmt.Errorf("error validating refresh: %v", err.Error())
	}

	_, err = s.R.Del(ctx, token_uuid, refresh_uuid).Result()
	if err != nil {
		return &pb.LogoutResponse{}, fmt.Errorf("error deleting tokens: %v", err.Error())
	}

	return &pb.LogoutResponse{
		Status:   "ok",
		Response: "logout successful",
	}, nil

}

func (s *Server) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	_, refresh_uuid, err := s.Jwt.ValidateJWT(req.Token, true)
	if err != nil {
		return &pb.RefreshResponse{}, fmt.Errorf("error validating token: %v", err.Error())
	}

	userid, err := s.R.Get(ctx, refresh_uuid).Result()
	if err != nil {
		return &pb.RefreshResponse{}, fmt.Errorf("error while getting token: %v", err.Error())
	}

	user, err := s.H.GetUserByUUID(userid)
	if err != nil {
		return &pb.RefreshResponse{}, fmt.Errorf("%v", err.Error())
	}

	if user.Id == "" {
		return &pb.RefreshResponse{}, fmt.Errorf("user belonging to this token does not exist", err.Error())
	}

	token, new_token_uuid, err := s.Jwt.SignJWT(user.Id, false)
	if err != nil {
		return &pb.RefreshResponse{}, fmt.Errorf("error while signing token: %v", err.Error())
	}

	now := time.Now()
	errToken := s.R.Set(ctx, new_token_uuid, user.Id, time.Unix(int64(s.Jwt.TokenExp)*int64(time.Hour*24), 0).Sub(now)).Err()
	if errToken != nil {
		return &pb.RefreshResponse{}, fmt.Errorf("%v", err.Error())
	}

	return &pb.RefreshResponse{
		Token: token,
	}, nil

}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, token_uuid, err := s.Jwt.ValidateJWT(req.Token, false)
	if err != nil {
		return &pb.ValidateResponse{}, fmt.Errorf("error validating token: %v", err.Error())
	}

	userid, err := s.R.Get(ctx, token_uuid).Result()
	if err != nil {
		return &pb.ValidateResponse{}, fmt.Errorf("token is invalid or session has expired")
	}

	user, err := s.H.GetUserByUUID(userid)
	if err != nil {
		return &pb.ValidateResponse{}, fmt.Errorf("%v", err.Error())
	}

	if user.Id != userid {
		return &pb.ValidateResponse{}, fmt.Errorf("user belonging to this token does longer exist")
	}

	return &pb.ValidateResponse{
		Uuid: userid,
	}, nil
}
