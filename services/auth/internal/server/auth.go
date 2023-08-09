package server

import (
	"context"
	"log"
	"net/http"
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
		log.Println(err)
		return &pb.RegisterResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while querying db",
		}, err
	}
	if user.Email != "" {
		return &pb.RegisterResponse{
			Status:   http.StatusConflict,
			Response: "user already exists",
		}, err
	}

	err = s.H.CreateUser(req)
	if err != nil {
		log.Println(err)
		return &pb.RegisterResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while creating user",
		}, err
	}
	return &pb.RegisterResponse{
		Status:   http.StatusCreated,
		Response: "user created",
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.H.GetUserByEmail(req.Email)
	if err != nil {
		log.Println(err)
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.LoginResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	if user.Email != req.Email {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Response: &pb.LoginResponse_Error{
				Error: "wrong email or password",
			},
		}, err
	}

	ok := utils.CheckPasswordHash(req.Passwd, user.Passwd)

	if !ok {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Response: &pb.LoginResponse_Error{
				Error: "wrong email or password",
			},
		}, err
	}

	token, tokenUUID, err := s.Jwt.SignJWT(user.Id, false)
	if err != nil {
		log.Println(err)
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.LoginResponse_Error{
				Error: "error while signing jwt",
			},
		}, err
	}

	refresh, refreshUUID, err := s.Jwt.SignJWT(user.Id, true)
	if err != nil {
		log.Println(err)
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.LoginResponse_Error{
				Error: "error while signing jwt",
			},
		}, err
	}

	now := time.Now()

	errToken := s.R.Set(ctx, tokenUUID, user.Id, time.Unix(int64(s.Jwt.TokenExp)*int64(time.Hour*24), 0).Sub(now)).Err()
	if errToken != nil {
		log.Println(err)
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.LoginResponse_Error{
				Error: "error while storing token",
			},
		}, err
	}

	errRefresh := s.R.Set(ctx, refreshUUID, user.Id, time.Unix(int64(s.Jwt.RefreshExp)*int64(time.Hour*24), 0).Sub(now)).Err()
	if errRefresh != nil {
		log.Println(err)
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.LoginResponse_Error{
				Error: "error while storing token",
			},
		}, err
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Response: &pb.LoginResponse_Data{
			Data: &pb.LoginResponse_Response{
				Token:   token,
				Refresh: refresh,
			},
		},
	}, nil

}

func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	_, token_uuid, err := s.Jwt.ValidateJWT(req.Token, false)
	if err != nil {
		log.Println(err)
		return &pb.LogoutResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while validating JWT",
		}, err
	}

	_, refresh_uuid, err := s.Jwt.ValidateJWT(req.Refresh, true)
	if err != nil {
		log.Println(err)
		return &pb.LogoutResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while validating jwt",
		}, err
	}

	_, err = s.R.Del(ctx, token_uuid, refresh_uuid).Result()
	if err != nil {
		log.Println(err)
		return &pb.LogoutResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while validating JWT",
		}, err
	}

	return &pb.LogoutResponse{
		Status:   http.StatusOK,
		Response: "logout successful",
	}, nil

}

func (s *Server) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	_, refresh_uuid, err := s.Jwt.ValidateJWT(req.Token, true)
	if err != nil {
		log.Println(err)
		return &pb.RefreshResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.RefreshResponse_Error{
				Error: "error while validating jwt",
			},
		}, err
	}

	userid, err := s.R.Get(ctx, refresh_uuid).Result()
	if err != nil {
		log.Println(err)
		return &pb.RefreshResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.RefreshResponse_Error{
				Error: "error while storing jwt",
			},
		}, err
	}

	user, err := s.H.GetUserByUUID(userid)
	if err != nil {
		log.Println(err)
		return &pb.RefreshResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.RefreshResponse_Error{
				Error: "error while querying db",
			},
		}, err
	}

	if user.Id == "" {
		log.Println(err)
		return &pb.RefreshResponse{
			Status: http.StatusNotFound,
			Response: &pb.RefreshResponse_Error{
				Error: "user does not longer exist",
			},
		}, err
	}

	token, new_token_uuid, err := s.Jwt.SignJWT(user.Id, false)
	if err != nil {
		log.Println(err)
		return &pb.RefreshResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.RefreshResponse_Error{
				Error: "error while signing jwt",
			},
		}, err
	}

	now := time.Now()
	errToken := s.R.Set(ctx, new_token_uuid, user.Id, time.Unix(int64(s.Jwt.TokenExp)*int64(time.Hour*24), 0).Sub(now)).Err()
	if errToken != nil {
		log.Println(err)
		return &pb.RefreshResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.RefreshResponse_Error{
				Error: "error while storing jwt",
			},
		}, err
	}

	return &pb.RefreshResponse{
		Status: http.StatusOK,
		Response: &pb.RefreshResponse_Data{
			Data: &pb.RefreshResponse_Response{
				Token: token,
			},
		},
	}, nil

}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, token_uuid, err := s.Jwt.ValidateJWT(req.Token, false)
	if err != nil {
		log.Println(err)
		return &pb.ValidateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.ValidateResponse_Error{
				Error: "error while validating jwt",
			},
		}, err
	}

	userid, err := s.R.Get(ctx, token_uuid).Result()
	if err != nil {
		log.Println(err)
		return &pb.ValidateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.ValidateResponse_Error{
				Error: "error while storing jwt",
			},
		}, err
	}

	user, err := s.H.GetUserByUUID(userid)
	if err != nil {
		log.Println(err)
		return &pb.ValidateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.ValidateResponse_Error{
				Error: "error while querying db",
			},
		}, err
	}

	if user.Id != userid {
		log.Println(err)
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Response: &pb.ValidateResponse_Error{
				Error: "invalid",
			},
		}, err
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		Response: &pb.ValidateResponse_Data{
			Data: &pb.ValidateResponse_Response{
				Uuid: userid,
			},
		},
	}, nil
}
