package db

import (
	"fmt"
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/warehouse/proto"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type Handler struct {
	DB *surrealdb.DB
}

type UserCrucial struct {
	Email  string `json:"email"`
	Id     string `json:"id"`
	Passwd string `json:"passwd"`
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
	};
	UPDATE $userID SET owns += $warehouse;
	RELATE $userID -> manages -> $warehouse SET roles = ["owner"];
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
		return "", fmt.Errorf("error while creating warehouse: %v", err)
	}

	fmt.Printf("%v RESPONSE", res)

	return "ok", nil
}

func (h *Handler) GetInfo(data *pb.GetInfoRequest) (pb.GetInfoResponse, error) {
	queryRes, err := h.DB.Query("SELECT * FROM $warehouse;", map[string]string{
		"warehouse": data.WarehouseID,
	})
	if err != nil {
		return pb.GetInfoResponse{}, fmt.Errorf("error while querying db: %v", err.Error())
	}

	var res pb.GetInfoResponse
	err = surrealdb.Unmarshal(queryRes, &res)
	if err != nil {
		return pb.GetInfoResponse{}, fmt.Errorf("error while unmarshalling data: %v", err.Error())
	}

	return res, nil
}

func (h *Handler) UpdateWarehouse(data *pb.UpdateRequest) (pb.UpdateResponse, error) {
	_, err := h.DB.Update(data.WarehouseID, data)
	if err != nil {
		return pb.UpdateResponse{}, fmt.Errorf("error while updating record: %v", err.Error())
	}
	return pb.UpdateResponse{Status: "ok", Response: "record updated"}, nil
}

func (h *Handler) DeleteWarehouse(data *pb.DeleteRequest) (pb.DeleteResponse, error) {
	_, err := h.DB.Delete(data.WarehouseID)
	if err != nil {
		return pb.DeleteResponse{}, fmt.Errorf("error while deleting warehouse: %v", err.Error())
	}
	return pb.DeleteResponse{Status: "ok", Response: "warehouse deleted"}, nil
}

func (h *Handler) AddUserToWarehouse(data *pb.AddUsersRequest) (pb.AddUsersResponse, error) {
	_, err := h.DB.Query("RELATE $userID -> manages -> $warehouse SET roles = ['display'];", map[string]string{
		"warehouse": data.WarehouseID,
		"userID":    data.UserIds,
	})
	if err != nil {
		return pb.AddUsersResponse{}, fmt.Errorf("error while adding user to warehouse: %v", err.Error())
	}
	return pb.AddUsersResponse{Status: "ok", Response: "user added to warehouse"}, nil
}

func (h *Handler) RemoveUserFromWarehouse(data *pb.RemoveUserRequest) (pb.RemoveUserResponse, error) {
	_, err := h.DB.Query("DELETE $userID -> manages WHERE out=$warehouse;", map[string]string{
		"warehouse": data.WarehouseID,
		"userID":    data.UserIds,
	})
	if err != nil {
		return pb.RemoveUserResponse{}, fmt.Errorf("error while adding user to warehouse: %v", err.Error())
	}
	return pb.RemoveUserResponse{Status: "ok", Response: "user added to warehouse"}, nil
}
