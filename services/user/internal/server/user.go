package server

import (
	"context"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	"github.com/CrabStash/crab-stash/user/internal/db"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	H db.Handler
}

func (s *Server) MeInfo(ctx context.Context, req *pb.MeInfoRequest) (*pb.MeInfoResponse, error) {
	me, err := s.H.GetMeInfo(req)
	if err != nil {
		return &pb.MeInfoResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.MeInfoResponse_Error{
				Error: "could not get me info",
			},
		}, nil
	}
	return &pb.MeInfoResponse{
		Status:   http.StatusOK,
		Response: me,
	}, nil

}

func (s *Server) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	err := s.H.DbUpdateUserInfo(req)
	if err != nil {
		return &pb.UpdateUserInfoResponse{
			Status:   http.StatusInternalServerError,
			Response: "could not update user info",
		}, nil
	}
	return &pb.UpdateUserInfoResponse{
		Status:   http.StatusOK,
		Response: "user updated",
	}, nil
}

func (s *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	usrInfo, err := s.H.DbGetUserInfo(req)
	if err != nil {
		return &pb.GetUserInfoResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetUserInfoResponse_Error{
				Error: "could not get user info",
			},
		}, nil
	}
	return &pb.GetUserInfoResponse{
		Status:   http.StatusOK,
		Response: usrInfo,
	}, nil

}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.H.DbDeleteUser(req)
	if err != nil {
		log.Println(err)
		return &pb.DeleteUserResponse{
			Status:   http.StatusInternalServerError,
			Response: err.Error(),
		}, nil
	}
	return &pb.DeleteUserResponse{
		Status:   http.StatusOK,
		Response: "user deleted",
	}, nil
}

func (s *Server) InternalGetUserByEmailAuth(ctx context.Context, req *pb.InternalGetUserByEmailRequest) (*pb.InternalGetUserByEmailAuthResponse, error) {
	usrInfo, err := s.H.DbInternalGetUserByEmail(req)
	if err != nil {
		return &pb.InternalGetUserByEmailAuthResponse{}, nil
	}
	res := &pb.InternalGetUserByEmailAuthResponse{
		Id:     usrInfo.Id,
		Passwd: usrInfo.Passwd,
		Email:  usrInfo.Email,
	}
	return res, nil

}

func (s *Server) InternalGetUserByEmailWarehouse(ctx context.Context, req *pb.InternalGetUserByEmailRequest) (*pb.InternalGetUserByEmailWarehouseResponse, error) {
	usrInfo, err := s.H.DbInternalGetUserByEmail(req)
	if err != nil {
		return &pb.InternalGetUserByEmailWarehouseResponse{}, nil
	}
	res := &pb.InternalGetUserByEmailWarehouseResponse{
		Id: usrInfo.Id,
	}
	return res, nil

}

func (s *Server) InternalGetUserByUUIDCheck(ctx context.Context, req *pb.InternalGetUserByUUIDCheck) (*pb.InternalGetUserByUUIDCheck, error) {
	usrID, err := s.H.DbGetUserbyUUID(req)
	if err != nil {
		return &pb.InternalGetUserByUUIDCheck{}, nil
	}
	return usrID, nil
}
