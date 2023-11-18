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

type Transaction struct {
	Result []map[string]interface{} `json:"result"`
	Status string                   `json:"status"`
	Time   string                   `json:"time"`
}

type SchemaProperties struct {
	Title string `json:"title"`
	Help  string `json:"help"`
	Type  string `json:"type"`
	Id    string `json:"id"`
}
type SchemaResult struct {
	Description string             `json:"description"`
	Title       string             `json:"title"`
	Properties  []SchemaProperties `json:"properties"`
}

type ServeSchema struct {
	Result SchemaResult `json:"result"`
	Status string       `json:"status"`
	Time   string       `json:"time"`
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

func (h *Handler) FieldsInheritance(data *pb.GenericFetchRequest) *pb.InheritanceResponse {
	queryRes, err := h.DB.Query("SELECT title, id, properties[*].title as fieldNames FROM (SELECT VALUE parents FROM ONLY $categoryID);", map[string]string{
		"categoryID": data.EntityID,
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

func (h *Handler) GetCategory(data *pb.GenericFetchRequest) *pb.GenericFetchResponse {
	queryRes, err := h.DB.Query("SELECT description, title, array::flatten(array::union(properties[*].*, parents.properties[*].*)) as properties FROM ONLY $categoryID;", map[string]string{
		"categoryID": data.EntityID,
	})

	if err != nil {
		return &pb.GenericFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericFetchResponse_Error{
				Error: err.Error(),
			},
		}
	}

	res := make([]ServeSchema, 1)
	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		return &pb.GenericFetchResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericFetchResponse_Error{
				Error: err.Error(),
			},
		}
	}

	properties := make(map[string]*pb.Field)

	for i := 0; i < len(res[0].Result.Properties); i++ {
		field := res[0].Result.Properties[i]
		properties[field.Id] = &pb.Field{
			Title: field.Title,
			Type:  field.Type,
			Help:  field.Help,
		}
	}

	return &pb.GenericFetchResponse{
		Status: http.StatusOK,
		Response: &pb.GenericFetchResponse_Data{
			Data: &pb.GenericFetchResponse_Response{
				Schema: &pb.GenericSchema{
					Title:       res[0].Result.Title,
					Description: res[0].Result.Description,
					Type:        "object",
					Properties:  properties,
				},
			},
		},
	}

}

func (h *Handler) CoreMiddleware(data *pb.GenericFetchRequest) (*pb.CoreMiddlewareResponse, error) {
	queryRes, err := h.DB.Query("SELECT * FROM type::table($target) WHERE in = $entityID AND out = $warehouseID", map[string]string{
		"target":      data.Type,
		"entityID":    data.EntityID,
		"warehouseID": data.WarehouseID,
	})

	if err != nil {
		return &pb.CoreMiddlewareResponse{}, err
	}

	res := make([]Transaction, 1)
	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		return &pb.CoreMiddlewareResponse{}, err
	}

	if len(res[0].Result) == 0 {
		return &pb.CoreMiddlewareResponse{
			DoesItBelong: false,
		}, nil
	} else {
		return &pb.CoreMiddlewareResponse{
			DoesItBelong: true,
		}, nil
	}
}
