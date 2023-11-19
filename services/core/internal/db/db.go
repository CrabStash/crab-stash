package db

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/core/proto"
	surrealdb "github.com/surrealdb/surrealdb.go"
	"google.golang.org/protobuf/types/known/anypb"
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

// Create

func (h *Handler) CreateField(data *pb.CreateFieldRequest) *pb.GenericCreateResponse {
	queryRes, err := h.DB.Query(`
		BEGIN TRANSACTION;
		LET $field = type::thing("fields", rand::uuid());
		CREATE $field CONTENT $data RETURN id;
		RELATE $field -> fields_to_warehouses -> $warehouse RETURN NONE;
		COMMIT TRANSACTION;
	`, map[string]interface{}{
		"data":      data.FormData,
		"warehouse": data.WarehouseID,
	})

	if err != nil {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while creating field: %s", err.Error()).Error(),
			},
		}
	}

	var finalRes []Transaction

	err = surrealdb.Unmarshal(queryRes, &finalRes)

	if err != nil {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while unmarshalling data: %s", err.Error()).Error(),
			},
		}
	}

	fieldID, ok := finalRes[1].Result[0]["id"].(string)
	if !ok {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while asserting type: %s", err.Error()).Error(),
			},
		}
	}

	return &pb.GenericCreateResponse{
		Status: http.StatusInternalServerError,
		Response: &pb.GenericCreateResponse_Data{
			Data: &pb.GenericCreateResponse_Response{
				Id: fieldID,
			},
		},
	}
}

func (h *Handler) CreateCategory(data *pb.CreateCategoryRequest) *pb.GenericCreateResponse {
	queryRes, err := h.DB.Query(`
		BEGIN TRANSACTION;
		LET $categories = type::thing("categories", rand::uuid());
		CREATE $categories CONTENT $data RETURN id;
		RELATE $categories -> categories_to_warehouses -> $warehouse RETURN NONE;
		COMMIT TRANSACTION;
	`, map[string]interface{}{
		"data":      data.FormData,
		"warehouse": data.WarehouseID,
	})

	if err != nil {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while creating field: %s", err.Error()).Error(),
			},
		}
	}

	var finalRes []Transaction

	err = surrealdb.Unmarshal(queryRes, &finalRes)

	if err != nil {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while unmarshalling data: %s", err.Error()).Error(),
			},
		}
	}

	categoryID, ok := finalRes[1].Result[0]["id"].(string)
	if !ok {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while asserting type: %s", err.Error()).Error(),
			},
		}
	}

	return &pb.GenericCreateResponse{
		Status: http.StatusInternalServerError,
		Response: &pb.GenericCreateResponse_Data{
			Data: &pb.GenericCreateResponse_Response{
				Id: categoryID,
			},
		},
	}
}

func (h *Handler) CreateEntity(data *pb.CreateEntityRequest) *pb.GenericCreateResponse {
	queryRes, err := h.DB.Query(`
		BEGIN TRANSACTION;
		LET $entities = type::thing("entities", rand::uuid());
		CREATE $entities CONTENT $data RETURN id;
		RELATE $entities -> entities_to_categories -> $category RETURN NONE;
		COMMIT TRANSACTION;
	`, map[string]interface{}{
		"data":     data.FormData,
		"category": data.CategoryID,
	})

	if err != nil {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while creating field: %s", err.Error()).Error(),
			},
		}
	}

	var finalRes []Transaction

	err = surrealdb.Unmarshal(queryRes, &finalRes)

	if err != nil {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while unmarshalling data: %s", err.Error()).Error(),
			},
		}
	}

	entityID, ok := finalRes[1].Result[0]["id"].(string)
	if !ok {
		log.Println(err)
		return &pb.GenericCreateResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GenericCreateResponse_Error{
				Error: fmt.Errorf("error while asserting type: %s", err.Error()).Error(),
			},
		}
	}

	return &pb.GenericCreateResponse{
		Status: http.StatusInternalServerError,
		Response: &pb.GenericCreateResponse_Data{
			Data: &pb.GenericCreateResponse_Response{
				Id: entityID,
			},
		},
	}
}

// Editing
func (h *Handler) EditField(data *pb.EditFieldRequest) *pb.GenericEditDeleteResponse {
	_, err := h.DB.Query("UPDATE $field MERGE $data", map[string]interface{}{
		"field": data.FieldID,
		"data":  data.FormData,
	})

	if err != nil {
		return &pb.GenericEditDeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: fmt.Errorf("error while updating field: %s", err.Error()).Error(),
		}
	}

	return &pb.GenericEditDeleteResponse{
		Status:   http.StatusOK,
		Response: "field edited",
	}

}

func (h *Handler) EditEntity(data *pb.EditEntityRequest) *pb.GenericEditDeleteResponse {
	_, err := h.DB.Query("UPDATE $entity MERGE $data", map[string]interface{}{
		"entity": data.EntityID,
		"data":   data.FormData,
	})

	if err != nil {
		return &pb.GenericEditDeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: fmt.Errorf("error while updating entity: %s", err.Error()).Error(),
		}
	}

	return &pb.GenericEditDeleteResponse{
		Status:   http.StatusOK,
		Response: "entity edited",
	}

}

func (h *Handler) EditCategory(data *pb.EditCategoryRequest) *pb.GenericEditDeleteResponse {
	_, err := h.DB.Query("UPDATE $category MERGE $data", map[string]interface{}{
		"category": data.CategoryID,
		"data":     data.FormData,
	})

	if err != nil {
		return &pb.GenericEditDeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: fmt.Errorf("error while updating category: %s", err.Error()).Error(),
		}
	}

	return &pb.GenericEditDeleteResponse{
		Status:   http.StatusOK,
		Response: "category edited",
	}

}

// Delete

func (h *Handler) DeleteCategory(data *pb.GenericFetchRequest) *pb.GenericEditDeleteResponse {
	_, err := h.DB.Query(`
		BEGIN TRANSACTION;
		LET $categories = SELECT VALUE id FROM categories WHERE parents CONTAINS $categoryID;
		FOR $category IN $categories {
			DELETE entities WHERE ->entities_to_categories->out = $category;
			DELETE $category;
		};
		DELETE entities WHERE ->entities_to_categories->out = $categoryID;
		DELETE $categoryID;
		COMMIT TRANSACTION;
	`, map[string]string{
		"categoryID": data.EntityID,
	})

	if err != nil {
		return &pb.GenericEditDeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: fmt.Errorf("error while deleting category: %s", err.Error()).Error(),
		}
	}
	return &pb.GenericEditDeleteResponse{
		Status:   http.StatusOK,
		Response: "category deleted",
	}
}

func (h *Handler) DeleteField(data *pb.GenericFetchRequest) *pb.GenericEditDeleteResponse {
	_, err := h.DB.Delete(data.EntityID)
	if err != nil {
		return &pb.GenericEditDeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: fmt.Errorf("error while deleting field: %s", err.Error()).Error(),
		}
	}
	return &pb.GenericEditDeleteResponse{
		Status:   http.StatusOK,
		Response: "field deleted",
	}
}

func (h *Handler) DeleteEntity(data *pb.GenericFetchRequest) *pb.GenericEditDeleteResponse {
	_, err := h.DB.Delete(data.EntityID)
	if err != nil {
		return &pb.GenericEditDeleteResponse{
			Status:   http.StatusInternalServerError,
			Response: fmt.Errorf("error while deleting entity: %s", err.Error()).Error(),
		}
	}
	return &pb.GenericEditDeleteResponse{
		Status:   http.StatusOK,
		Response: "entity deleted",
	}
}

// Fetching
func (h *Handler) GetCategorySchema(data *pb.GenericFetchRequest) *pb.CategorySchemaResponse {
	queryRes, err := h.DB.Query("SELECT description, title, array::flatten(array::union(properties[*].*, parents.properties[*].*)) as properties FROM ONLY $categoryID;", map[string]string{
		"categoryID": data.EntityID,
	})

	if err != nil {
		return &pb.CategorySchemaResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.CategorySchemaResponse_Error{
				Error: err.Error(),
			},
		}
	}

	res := make([]ServeSchema, 1)
	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		return &pb.CategorySchemaResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.CategorySchemaResponse_Error{
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

	return &pb.CategorySchemaResponse{
		Status: http.StatusOK,
		Response: &pb.CategorySchemaResponse_Data{
			Data: &pb.CategorySchemaResponse_Response{
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

func (h *Handler) GetCategoryData(data *pb.GenericFetchRequest) *pb.GetCategoryDataResponse {
	queryRes, err := h.DB.Query("SELECT title, description, parents, properties FROM $category", map[string]string{"category": data.EntityID})

	if err != nil {
		return &pb.GetCategoryDataResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetCategoryDataResponse_Error{
				Error: fmt.Errorf("error while getting category data: %s", err.Error()).Error(),
			},
		}
	}

	res, err := surrealdb.SmartUnmarshal[[]*pb.Category](queryRes, nil)

	log.Println(res, queryRes)

	if err != nil {
		return &pb.GetCategoryDataResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetCategoryDataResponse_Error{
				Error: fmt.Errorf("error while unmarshalling data: %s", err.Error()).Error(),
			},
		}
	}

	return &pb.GetCategoryDataResponse{
		Status: http.StatusOK,
		Response: &pb.GetCategoryDataResponse_Data{
			Data: &pb.GetCategoryDataResponse_Response{
				FormData: res[0],
			},
		},
	}
}

func (h *Handler) GetFieldData(data *pb.GenericFetchRequest) *pb.GetFieldDataResponse {
	queryRes, err := h.DB.Query("SELECT title, type, help FROM $field", map[string]string{"field": data.EntityID})

	if err != nil {
		return &pb.GetFieldDataResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetFieldDataResponse_Error{
				Error: fmt.Errorf("error while getting Field data: %s", err.Error()).Error(),
			},
		}
	}

	res, err := surrealdb.SmartUnmarshal[[]*pb.Field](queryRes, nil)

	log.Println(res, queryRes)

	if err != nil {
		return &pb.GetFieldDataResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetFieldDataResponse_Error{
				Error: fmt.Errorf("error while unmarshalling data: %s", err.Error()).Error(),
			},
		}
	}

	return &pb.GetFieldDataResponse{
		Status: http.StatusOK,
		Response: &pb.GetFieldDataResponse_Data{
			Data: &pb.GetFieldDataResponse_Response{
				FormData: res[0],
			},
		},
	}
}

func (h *Handler) GetEntityData(data *pb.GenericFetchRequest) *pb.GetEntityDataResponse {
	queryRes, err := h.DB.Query("SELECT * FROM $entity", map[string]string{"entity": data.EntityID})

	if err != nil {
		return &pb.GetEntityDataResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetEntityDataResponse_Error{
				Error: fmt.Errorf("error while getting Entity data: %s", err.Error()).Error(),
			},
		}
	}

	res := make([]Transaction, 1)

	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		return &pb.GetEntityDataResponse{
			Status: http.StatusInternalServerError,
			Response: &pb.GetEntityDataResponse_Error{
				Error: fmt.Errorf("error while getting Entity data: %s", err.Error()).Error(),
			},
		}
	}

	log.Println(queryRes, res[0].Result)
	return &pb.GetEntityDataResponse{
		Status: http.StatusOK,
		Response: &pb.GetEntityDataResponse_Data{
			Data: &pb.GetEntityDataResponse_Response{
				FormData: make(map[string]*anypb.Any),
			},
		},
	}
}

// Misc

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

func (h *Handler) CoreMiddleware(data *pb.CoreMiddlewareRequest) (*pb.CoreMiddlewareResponse, error) {
	queryRes, err := h.DB.Query("SELECT * FROM type::table($target) WHERE in = $entityID AND out = $warehouseID", map[string]string{
		"target":      data.Type,
		"entityID":    data.In,
		"warehouseID": data.Out,
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
