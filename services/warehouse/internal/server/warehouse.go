package server

import (
	"context"
	"fmt"
	"log"

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
		log.Println(err)
		return &pb.CreateResponse{}, fmt.Errorf("error while creating warehouse: %v", err)
	}

	return &pb.CreateResponse{WarehouseID: warehouseID}, nil
}

func (s *Server) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	info, err := s.H.GetInfo(req)
	if err != nil {
		log.Println(err)
		return &pb.GetInfoResponse{}, fmt.Errorf("%v", err)
	}
	if info.Owner == "" {
		log.Println(err)
		return &pb.GetInfoResponse{}, fmt.Errorf("warehouse does not exist")
	}
	return info, nil
}

func (s *Server) UpdateWarehouse(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	status, err := s.H.UpdateWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.UpdateResponse{}, fmt.Errorf("%v", err)
	}
	return &status, nil
}

func (s *Server) AddUsersToWarehouse(ctx context.Context, req *pb.AddUsersRequest) (*pb.AddUsersResponse, error) {
	status, err := s.H.AddUserToWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.AddUsersResponse{}, fmt.Errorf("%v", err)
	}
	return &status, nil
}

func (s *Server) RemoveUsersFromWarehouse(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	status, err := s.H.RemoveUserFromWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.RemoveUserResponse{}, fmt.Errorf("%v", err)
	}
	return &status, nil
}

func (s *Server) DeleteWarehouse(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	status, err := s.H.DeleteWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.DeleteResponse{}, fmt.Errorf("%v", err)
	}
	return &status, nil
}
