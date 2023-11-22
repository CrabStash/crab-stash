package server

import (
	"context"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	"github.com/CrabStash/crab-stash/core/internal/db"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedCoreServiceServer
	H db.Handler
}

func (s *Server) NewCategorySchema(ctx context.Context, req *emptypb.Empty) (*pb.Schema, error) {
	file, err := os.Open("/schemas/category.json")

	if err != nil {
		log.Println(err)
		return &pb.Schema{}, nil
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		return &pb.Schema{}, nil
	}

	return &pb.Schema{
		FileContent: byteValue,
	}, nil

}

func (s *Server) NewFieldSchema(ctx context.Context, req *emptypb.Empty) (*pb.Schema, error) {
	file, err := os.Open("/schemas/field.json")

	if err != nil {
		log.Println(err)
		return &pb.Schema{}, nil
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		log.Println(err)
		return &pb.Schema{}, nil
	}

	return &pb.Schema{
		FileContent: byteValue,
	}, nil
}

func (s *Server) GetCategorySchema(ctx context.Context, req *pb.GenericFetchRequest) (*pb.CategorySchemaResponse, error) {
	res := s.H.GetCategorySchema(req)
	return res, nil
}

// Create functions

func (s *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.GenericCreateResponse, error) {
	res := s.H.CreateCategory(req)
	return res, nil
}

func (s *Server) CreateEntity(ctx context.Context, req *pb.CreateEntityRequest) (*pb.GenericCreateResponse, error) {
	res := s.H.CreateEntity(req)
	return res, nil
}

func (s *Server) CreateField(ctx context.Context, req *pb.CreateFieldRequest) (*pb.GenericCreateResponse, error) {
	res := s.H.CreateField(req)
	return res, nil
}

// Edit functions

func (s *Server) EditCategory(ctx context.Context, req *pb.EditCategoryRequest) (*pb.GenericEditDeleteResponse, error) {
	res := s.H.EditCategory(req)
	return res, nil
}

func (s *Server) EditField(ctx context.Context, req *pb.EditFieldRequest) (*pb.GenericEditDeleteResponse, error) {
	res := s.H.EditField(req)
	return res, nil
}

func (s *Server) EditEntity(ctx context.Context, req *pb.EditEntityRequest) (*pb.GenericEditDeleteResponse, error) {
	res := s.H.EditEntity(req)
	return res, nil
}

// Delete functions

func (s *Server) DeleteCategory(ctx context.Context, req *pb.GenericFetchRequest) (*pb.GenericEditDeleteResponse, error) {
	res := s.H.DeleteCategory(req)
	return res, nil
}

func (s *Server) DeleteField(ctx context.Context, req *pb.GenericFetchRequest) (*pb.GenericEditDeleteResponse, error) {
	res := s.H.DeleteField(req)
	return res, nil
}

func (s *Server) DeleteEntity(ctx context.Context, req *pb.GenericFetchRequest) (*pb.GenericEditDeleteResponse, error) {
	res := s.H.DeleteEntity(req)
	return res, nil
}

// Data Fetch Functions

func (s *Server) GetCategoryData(ctx context.Context, req *pb.GenericFetchRequest) (*pb.GetCategoryDataResponse, error) {
	res := s.H.GetCategoryData(req)
	return res, nil
}

func (s *Server) GetFieldData(ctx context.Context, req *pb.GenericFetchRequest) (*pb.GetFieldDataResponse, error) {
	res := s.H.GetFieldData(req)
	return res, nil
}

func (s *Server) GetEntityDataData(ctx context.Context, req *pb.GenericFetchRequest) (*pb.GetEntityDataResponse, error) {
	res := s.H.GetEntityData(req)
	return res, nil
}

// List
func (s *Server) ListFields(ctx context.Context, req *pb.PaginatedFieldFetchRequest) (*pb.PaginatedFieldsFetchResponse, error) {

	count, err := s.H.Count("", req.WarehouseID, "fields_to_warehouses")

	if err != nil {
		log.Println(err)
		return &pb.PaginatedFieldsFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.PaginatedFieldsFetchResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}

	pages := math.Ceil(float64(count) / float64(req.Limit))

	fields, err := s.H.ListFields(req, int(pages))

	if err != nil {
		log.Println(err)
		return &pb.PaginatedFieldsFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.PaginatedFieldsFetchResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}

	return &pb.PaginatedFieldsFetchResponse{
		Status: http.StatusOK,
		Response: &pb.PaginatedFieldsFetchResponse_Data{
			Data: fields,
		},
	}, nil

}

func (s *Server) ListCategories(ctx context.Context, req *pb.PaginatedCategoriesFetchRequest) (*pb.PaginatedCategoriesFetchResponse, error) {
	categories, err := s.H.ListCategories(req)

	if err != nil {
		log.Println(err)
		return &pb.PaginatedCategoriesFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.PaginatedCategoriesFetchResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}

	return &pb.PaginatedCategoriesFetchResponse{
		Status: http.StatusOK,
		Response: &pb.PaginatedCategoriesFetchResponse_Data{
			Data: categories,
		},
	}, nil

}

func (s *Server) ListEntities(ctx context.Context, req *pb.PaginatedEntitiesFetchRequest) (*pb.PaginatedEntititesFetchResponse, error) {
	count, err := s.H.Count(req.Id, req.WarehouseID, "entities_to_categories")

	if err != nil {
		log.Println(err)
		return &pb.PaginatedEntititesFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.PaginatedEntititesFetchResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}

	pages := math.Ceil(float64(count) / float64(req.Limit))

	fields, err := s.H.ListEntities(req, int(pages))

	if err != nil {
		log.Println(err)
		return &pb.PaginatedEntititesFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.PaginatedEntititesFetchResponse_Error{
				Error: err.Error(),
			},
		}, nil
	}

	return &pb.PaginatedEntititesFetchResponse{
		Status: http.StatusOK,
		Response: &pb.PaginatedEntititesFetchResponse_Data{
			Data: fields,
		},
	}, nil

}

// Misc

func (s *Server) FieldsInheritance(ctx context.Context, req *pb.GenericFetchRequest) (*pb.InheritanceResponse, error) {
	res := s.H.FieldsInheritance(req)
	return res, nil
}

func (s *Server) CoreMiddleware(ctx context.Context, req *pb.CoreMiddlewareRequest) (*pb.CoreMiddlewareResponse, error) {
	res, err := s.H.CoreMiddleware(req)
	if err != nil {
		return &pb.CoreMiddlewareResponse{}, err
	}
	return res, nil
}
