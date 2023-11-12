package server

import (
	"context"
	"io/ioutil"
	"log"
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
