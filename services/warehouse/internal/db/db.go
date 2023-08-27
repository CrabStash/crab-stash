package db

import (
	"fmt"
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	surrealdb "github.com/surrealdb/surrealdb.go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	DB *surrealdb.DB
}

type Transaction struct {
	Result []map[string]interface{} `json:"result"`
	Status string                   `json:"status"`
	Time   string                   `json:"time"`
}

func Init() Handler {
	db, err := surrealdb.New(os.Getenv("SURREALDB_ADDR"))

	if err != nil {
		log.Fatalf("Failed to connect to DB: %v\n", err.Error())
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": os.Getenv("SURREAL_USER"),
		"pass": os.Getenv("SURREAL_PASSWD"),
	}); err != nil {
		log.Fatalf("Failed to signin to db: %v\n", err.Error())
	}

	if _, err = db.Use("crabstash", "data"); err != nil {
		log.Fatalf("Failed to use crabstash/data: %v\n", err.Error())
	}
	return Handler{db}
}

func (h *Handler) CreateWarehouse(data *pb.CreateRequest) (string, error) {
	res, err := h.DB.Query(`
	BEGIN TRANSACTION;
	LET $warehouse = type::thing("warehouse", rand::uuid());
	CREATE $warehouse CONTENT {
		name: $name,
		desc: $desc,
		logo: $logo,
		owner: $userID,
		isPhysical: $isPhysical,
		capacity: $capacity
	} RETURN id;
	UPDATE $userID SET owns += $warehouse RETURN NONE;
	RELATE $userID -> manages -> $warehouse SET roles = ["owner"] RETURN NONE;
	COMMIT TRANSACTION;
	`, map[string]interface{}{
		"userID":     data.OwnerID,
		"name":       data.Name,
		"desc":       data.Desc,
		"logo":       data.Logo,
		"isPhysical": data.IsPhysical,
		"capacity":   data.Capacity,
	})

	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error while creating warehouse: %v", err.Error())
	}

	var finalRes []Transaction

	err = surrealdb.Unmarshal(res, &finalRes)

	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("error while unmarshaling data")
	}

	warehouseID, ok := finalRes[1].Result[0]["id"].(string)
	if !ok {
		log.Println(err)
		return "", fmt.Errorf("error while asserting type")
	}

	return warehouseID, nil
}

func (h *Handler) GetInfo(data *pb.GetInfoRequest) (*pb.GetInfoResponse_Data, error) {
	queryRes, err := h.DB.Select(data.WarehouseID)

	if err != nil {
		log.Println(err)
		return &pb.GetInfoResponse_Data{}, fmt.Errorf("error while querying db: %v", err)
	}

	res := &pb.GetInfoResponse_Data{
		Data: &pb.GetInfoResponse_Response{},
	}
	err = surrealdb.Unmarshal(queryRes, res.Data)
	if err != nil {
		log.Println(err)
		return &pb.GetInfoResponse_Data{}, fmt.Errorf("error while unmarshalling data: %v", err)
	}

	return res, nil
}

func (h *Handler) UpdateWarehouse(data *pb.UpdateRequest) error {
	_, err := h.DB.Change(data.WarehouseID, data)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while updating record: %v", err)
	}
	return nil
}

func (h *Handler) DeleteWarehouse(data *pb.DeleteRequest) error {
	_, err := h.DB.Delete(data.WarehouseID)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while deleting warehouse: %v", err)
	}
	return nil
}

func (h *Handler) AddUserToWarehouse(data *pb.AddUsersRequest) error {
	_, err := h.DB.Query("RELATE $userID -> manages -> $warehouse SET roles = [];", map[string]string{
		"warehouse": data.WarehouseID,
		"userID":    data.UserIds,
	})
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while adding user to warehouse: %v", err)
	}
	return nil
}

func (h *Handler) RemoveUserFromWarehouse(data *pb.RemoveUserRequest) error {
	_, err := h.DB.Query("DELETE $userID -> manages WHERE out=$warehouse;", map[string]string{
		"warehouse": data.WarehouseID,
		"userID":    data.UserIds,
	})
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while adding user to warehouse: %v", err)
	}
	return nil
}

func (h *Handler) FetchWarehouses(data *pb.InternalFetchWarehousesRequest) (*pb.InternalFetchWarehousesResponse, error) {
	queryRes, err := h.DB.Query("SELECT out as warehouseID, roles FROM manages WHERE in = $userID", map[string]string{
		"userID": data.UserID,
	})

	if err != nil {
		log.Println(err)
		return &pb.InternalFetchWarehousesResponse{}, fmt.Errorf("error while querying db: %v", err)
	}

	res := &pb.InternalFetchWarehousesResponse{}
	err = surrealdb.Unmarshal(queryRes, res)
	if err != nil {
		log.Println(err)
		return &pb.InternalFetchWarehousesResponse{}, fmt.Errorf("error while unmarshaling data: %v", err)
	}

	return res, nil
}

func (h *Handler) DeleteAccount(data *pb.InternalDeleteAccRequest) (*emptypb.Empty, error) {
	_, err := h.DB.Query("DELETE $userID->manages", map[string]string{
		"userID": data.UserID,
	})

	if err != nil {
		log.Println(err)
		return &emptypb.Empty{}, fmt.Errorf("error while deleting user from warehouses: %v", err)
	}

	return &emptypb.Empty{}, nil

}