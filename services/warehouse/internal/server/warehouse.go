package server

import (
	"context"
	"fmt"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	"github.com/CrabStash/crab-stash/warehouse/internal/db"
)

type Server struct {
	pb.UnimplementedWarehouseServiceServer
	H db.Handler
}

func (s *Server) CreateWarehouse(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	warehouseID, err := s.H.CreateWarehouse(req)
	if err != nil {
		return &pb.CreateResponse{}, fmt.Errorf("Error while creating warehouse: %v", err.Error())
	}

	return &pb.CreateResponse{WarehouseID: warehouseID}, nil
}

func (s *Server) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	info, err := s.H.GetInfo(req)
	if err != nil {
		return &pb.GetInfoResponse{}, fmt.Errorf("%v", err.Error())
	}
	if info.Owner == "" {
		return &pb.GetInfoResponse{}, fmt.Errorf("warehouse does not exist")
	}
	return &info, nil
}

func (s *Server) UpdateWarehouse(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	status, err := s.H.UpdateWarehouse(req)
	if err != nil {
		return &pb.UpdateResponse{}, fmt.Errorf("%v", err.Error())
	}
	return &status, nil
}

func (s *Server) AddUsersToWarehouse(ctx context.Context, req *pb.AddUsersRequest) (*pb.AddUsersResponse, error) {
	status, err := s.H.AddUserToWarehouse(req)
	if err != nil {
		return &pb.AddUsersResponse{}, fmt.Errorf("%v", err.Error())
	}
	return &status, nil
}

func (s *Server) RemoveUsersFromWarehouse(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	status, err := s.H.RemoveUserFromWarehouse(req)
	if err != nil {
		return &pb.RemoveUserResponse{}, fmt.Errorf("%v", err.Error())
	}
	return &status, nil
}

func (s *Server) DeleteWarehouse(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	status, err := s.H.DeleteWarehouse(req)
	if err != nil {
		return &pb.DeleteResponse{}, fmt.Errorf("%v", err.Error())
	}
	return &status, nil
}
