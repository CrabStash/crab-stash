package db

import (
	"log"
	"net/http"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type Handler struct {
	DB *surrealdb.DB
}

func Init() Handler {
	db, err := surrealdb.New(os.Getenv("SURREALDB_ADDR"))

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v\n", err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": os.Getenv("SURREAL_USER"),
		"pass": os.Getenv("SURREAL_PASSWD"),
	}); err != nil {
		log.Fatalf("Failed to signin to db: %v\n", err)
	}

	if _, err = db.Use("crabstash", "data"); err != nil {
		log.Fatalf("Failed to use crabstash/data: %v\n", err.Error())
	}
	return Handler{db}
}

func (h *Handler) FieldsInheritance(data *pb.InheritanceRequest) *pb.InheritanceResponse {
	queryRes, err := h.DB.Query("SELECT title, id, properties[*].name as fieldNames FROM (SELECT VALUE parents FROM ONLY $categoryID);", map[string]string{
		"categoryID": data.CategoryID,
	})

	if err != nil {
		return &pb.InheritanceResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.InheritanceResponse_Error{
				Error: err.Error(),
			},
		}
	}

	res, err := surrealdb.SmartUnmarshal[[]*pb.InheritanceResponse_Parent](queryRes, nil)

	if err != nil {
		return &pb.InheritanceResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.InheritanceResponse_Error{
				Error: err.Error(),
			},
		}
	}

	return &pb.InheritanceResponse{
		Status: http.StatusOK,
		Response: &pb.InheritanceResponse_Data{
			Data: &pb.InheritanceResponse_Response{
				Parents: res,
			},
		},
	}

}

func (h *Handler) GetCategory(data *pb.ServeCategoryRequest) *pb.ServeCategoryResponse {
	queryRes, err := h.DB.Query("SELECT description, title, array::flatten(array::union(properties[*].*, parents.properties[*].*)) as properties FROM ONLY $categoryID;", map[string]string{
		"categoryID": data.CategoryID,
	})

	if err != nil {
		return &pb.ServeCategoryResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.ServeCategoryResponse_Error{
				Error: err.Error(),
			},
		}
	}

	log.Println(queryRes)

	return &pb.ServeCategoryResponse{
		Status: http.StatusOK,
		Response: &pb.ServeCategoryResponse_Data{
			Data: &pb.ServeCategoryResponse_Response{
				Schema: &pb.Category{
					Title:       "test",
					Description: "test",
					Type:        "object",
					Properties:  make(map[string]*pb.Field),
				},
			},
		},
	}

}

// func (h *Handler) CoreMiddleware(data *pb.CoreMiddlewareRequest) *pb.CoreMiddlewareResponse {
// 	queryRes, err := h.DB.Query("SELECT * FROM $target WHERE in = $entityID AND out = $warehouseID", map[string]string{
// 		"target":      strings.Split(data.Target, ":")[0],
// 		"entityID":    data.Target,
// 		"warehouseID": data.WarehouseID,
// 	})

// 	if err != nil {
// 		return &pb.CoreMiddlewareResponse{
// 			DoesItBelong: false,
// 		}
// 	}

// }
