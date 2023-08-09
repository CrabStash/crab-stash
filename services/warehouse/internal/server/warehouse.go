package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	"github.com/CrabStash/crab-stash/warehouse/internal/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedWarehouseServiceServer
	H db.Handler
}

func (s *Server) CreateWarehouse(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	warehouseID, err := s.H.CreateWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.CreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.CreateResponse_Error{
				Error: err.Error(),
			},
		}, fmt.Errorf("error while creating warehouse: %v", err)
	}

	return &pb.CreateResponse{
		Status: http.StatusCreated,
		Response: &pb.CreateResponse_Data{
			Data: &pb.CreateResponse_Response{
				WarehouseID: warehouseID,
			},
		},
	}, nil
}

func (s *Server) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	info, err := s.H.GetInfo(req)
	if err != nil {
		log.Println(err)
		return &pb.GetInfoResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetInfoResponse_Error{
				Error: "error while querying db",
			},
		}, fmt.Errorf("%v", err)
	}
	if info.Data.Owner == "" {
		log.Println(err)
		return &pb.GetInfoResponse{
			Status: http.StatusNotFound,
			Response: &pb.GetInfoResponse_Error{
				Error: "error while querying db",
			},
		}, fmt.Errorf("warehouse does not exist")
	}
	return &pb.GetInfoResponse{
		Status:   http.StatusOK,
		Response: info,
	}, nil
}

func (s *Server) UpdateWarehouse(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	err := s.H.UpdateWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.UpdateResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while creating warehouse",
		}, fmt.Errorf("%v", err)
	}
	return &pb.UpdateResponse{
		Status:   http.StatusOK,
		Response: "record updated",
	}, nil
}

func (s *Server) AddUsersToWarehouse(ctx context.Context, req *pb.AddUsersRequest) (*pb.AddUsersResponse, error) {
	err := s.H.AddUserToWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.AddUsersResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while adding user to warehouse",
		}, fmt.Errorf("%v", err)
	}
	return &pb.AddUsersResponse{
		Status:   http.StatusOK,
		Response: "user added to warehouse",
	}, nil
}

func (s *Server) RemoveUserFromWarehouse(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	err := s.H.RemoveUserFromWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.RemoveUserResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while removing user from warehouse",
		}, fmt.Errorf("%v", err)
	}
	return &pb.RemoveUserResponse{
		Status:   http.StatusOK,
		Response: "user removed from warehouse",
	}, nil
}

func (s *Server) DeleteWarehouse(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := s.H.DeleteWarehouse(req)
	if err != nil {
		log.Println(err)
		return &pb.DeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while deleting warehouse",
		}, fmt.Errorf("%v", err)
	}
	return &pb.DeleteResponse{
		Status:   http.StatusOK,
		Response: "warehouse deleted",
	}, nil
}

func (s *Server) InternalFetchWarehouses(ctx context.Context, req *pb.InternalFetchWarehousesRequest) (*pb.InternalFetchWarehousesResponse, error) {
	warehouses, err := s.H.FetchWarehouses(req)
	if err != nil {
		log.Println(err)
		return &pb.InternalFetchWarehousesResponse{}, fmt.Errorf("%v", err)
	}
	return warehouses, nil
}

func (s *Server) InternalDeleteAcc(ctx context.Context, req *pb.InternalDeleteAccRequest) (*emptypb.Empty, error) {
	_, err := s.H.DeleteAccount(req)
	if err != nil {
		log.Println(err)
		return &emptypb.Empty{}, fmt.Errorf("%v", err)
	}
	return &emptypb.Empty{}, nil
}
