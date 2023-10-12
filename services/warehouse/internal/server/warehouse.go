package server

import (
	"context"
	"fmt"
	"log"
	"math"
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
		}, nil
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
				Error: err.Error(),
			},
		}, nil
	}
	if info.Data.Owner == "" {
		log.Println(err)
		return &pb.GetInfoResponse{
			Status: http.StatusNotFound,
			Response: &pb.GetInfoResponse_Error{
				Error: "warehouse does not exist",
			},
		}, nil
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
		}, nil
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
		}, nil
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
		}, nil
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
		}, nil
	}
	return &pb.DeleteResponse{
		Status:   http.StatusOK,
		Response: "warehouse deleted",
	}, nil
}

func (s *Server) ChangeRole(ctx context.Context, req *pb.ChangeRoleRequest) (*pb.ChangeRoleResponse, error) {
	res, err := s.H.ChangeRole(req)
	if err != nil {
		res.Response = err.Error()
		return res, nil
	}
	return res, nil
}

func (s *Server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	count, err := s.H.CountUsers(req)

	if err != nil {
		return &pb.ListUsersResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.ListUsersResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}

	pages := math.Ceil(float64(count) / float64(req.Limit))

	users, err := s.H.ListUsers(req, int(pages))

	if err != nil {
		return &pb.ListUsersResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.ListUsersResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}
	res := &pb.ListUsersResponse{
		Status: http.StatusOK,
		Response: &pb.ListUsersResponse_Data{
			Data: users,
		},
	}

	return res, nil
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

func (s *Server) InternalFetchWarehouseRole(ctx context.Context, req *pb.InternalFetchWarehouseRoleRequest) (*pb.InternalFetchWarehouseRoleResponse, error) {
	role, err := s.H.CheckRole(req)
	if err != nil {
		return &pb.InternalFetchWarehouseRoleResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.InternalFetchWarehouseRoleResponse_Error{
				Error: fmt.Sprintf("error while checking roles: %v", err.Error()),
			},
		}, nil
	}
	return &pb.InternalFetchWarehouseRoleResponse{
		Status: http.StatusOK,
		Response: &pb.InternalFetchWarehouseRoleResponse_Data{
			Data: role,
		},
	}, nil
}
