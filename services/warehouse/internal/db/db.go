package db

import (
	"fmt"
	"log"
	"net/http"
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

type AddUserTransaction struct {
	Result string `json:"result"`
	Status string `json:"status"`
	Time   string `json:"time"`
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
	RELATE $userID -> manages -> $warehouse SET role = $role RETURN NONE;
	COMMIT TRANSACTION;
	`, map[string]interface{}{
		"userID":     data.OwnerID,
		"name":       data.Name,
		"desc":       data.Desc,
		"logo":       data.Logo,
		"isPhysical": data.IsPhysical,
		"capacity":   data.Capacity,
		"role":       pb.Roles_ROLE_OWNER,
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
	queryRes, err := h.DB.Query("SELECT *, (SELECT role FROM ONLY manages WHERE out = $warehouseID AND in = $userID).role as role FROM $warehouseID", map[string]string{
		"warehouseID": data.WarehouseID,
		"userID":      data.UserID,
	})

	if err != nil {
		log.Println(err)
		return &pb.GetInfoResponse_Data{}, fmt.Errorf("error while querying db: %v", err)
	}

	res, err := surrealdb.SmartUnmarshal[[]pb.GetInfoResponse_Response](queryRes, nil)

	if err != nil {
		log.Println(err)
		return &pb.GetInfoResponse_Data{}, fmt.Errorf("error while unmarshalling data: %v", err)
	}

	return &pb.GetInfoResponse_Data{
		Data: &res[0],
	}, nil
}

func (h *Handler) UpdateWarehouse(data *pb.UpdateRequest) error {
	_, err := h.DB.Query("UPDATE $warehouseID MERGE $data", map[string]interface{}{
		"warehouseID": data.WarehouseID,
		"data":        data,
	})
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
	queryRes, err := h.DB.Query(`
	IF (SELECT VALUE email FROM user WHERE email == $email) IS NOT [] {
		RELATE (SELECT VALUE id FROM user WHERE email == $email) -> manages -> $warehouse SET role = $role;
	} ELSE {
		THROW "User does not exist!"
	};`, map[string]interface{}{
		"warehouse": data.WarehouseID,
		"email":     data.Email,
		"role":      pb.Roles_ROLE_VIEWER,
	})

	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while adding user to warehouse: %v", err)
	}

	res := make([]AddUserTransaction, 1)

	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while unmarshaling data: %v", err)
	}

	if res[0].Status == "ERR" {
		return fmt.Errorf("%s", res[0].Result)
	}

	return nil
}

func (h *Handler) RemoveUserFromWarehouse(data *pb.RemoveUserRequest) error {
	_, err := h.DB.Query("DELETE manages WHERE in=$userID AND out=$warehouse;", map[string]string{
		"warehouse": data.WarehouseID,
		"userID":    data.UserIds,
	})
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error while adding user to warehouse: %v", err)
	}
	return nil
}

func (h *Handler) ListWarehouses(data *pb.ListWarehousesRequest, pageCount int) (*pb.ListWarehousesResponse_Response, error) {
	var page int32

	if data.Page > int32(pageCount) {
		page = int32(pageCount)
	} else {
		page = int32(data.Page) - 1
	}

	queryRes, err := h.DB.Query("SELECT out.* as warehouse, role FROM manages WHERE in = $uuid ORDER BY role DESC LIMIT $limit START $page", map[string]interface{}{
		"uuid":  data.Uuid,
		"limit": data.Limit,
		"page":  data.Limit * page,
	})

	if err != nil {
		log.Println(err)
		return &pb.ListWarehousesResponse_Response{}, fmt.Errorf("error while querying db: %v", err.Error())
	}

	res := &pb.ListWarehousesResponse_Response{
		Pagination: &pb.ListWarehousesResponsePagination{
			Limit: data.Limit,
			Page:  data.Page,
			Total: int32(pageCount),
		},
	}

	list, err := surrealdb.SmartUnmarshal[[]*pb.ListWarehousesResponseList](queryRes, nil)

	if err != nil {
		log.Println(err)
		return &pb.ListWarehousesResponse_Response{}, fmt.Errorf("error while unmarshaling data: %v", err)
	}

	res.List = list

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

func (h *Handler) ChangeRole(data *pb.ChangeRoleRequest) (*pb.ChangeRoleResponse, error) {
	role, err := h.CheckRole(&pb.InternalFetchWarehouseRoleRequest{
		WarehouseID: data.WarehouseID,
		UserID:      data.Uuid,
	})

	if err != nil {
		log.Println(err)
		return &pb.ChangeRoleResponse{
			Status: http.StatusInternalServerError,
		}, fmt.Errorf("error while checking user role")
	}

	if data.NewRole >= role.Role {
		return &pb.ChangeRoleResponse{
			Status: http.StatusUnauthorized,
		}, fmt.Errorf("user's role cannot be equal or higher than your current role")
	}

	_, err = h.DB.Query("UPDATE manages SET role = $newRole WHERE in = $userID AND out = $warehouseID", map[string]interface{}{
		"newRole":     data.NewRole,
		"userID":      data.TargetUserID,
		"warehouseID": data.WarehouseID,
	})

	if err != nil {
		return &pb.ChangeRoleResponse{
			Status: http.StatusInternalServerError,
		}, fmt.Errorf("error while updating users role: %v", err.Error())
	}

	return &pb.ChangeRoleResponse{
		Status:   http.StatusOK,
		Response: "updated user role",
	}, nil
}

func (h *Handler) CheckRole(data *pb.InternalFetchWarehouseRoleRequest) (*pb.InternalFetchWarehouseRoleResponse_Response, error) {
	queryRes, err := h.DB.Query("SELECT role FROM manages WHERE in = $userID AND out = $warehouseID LIMIT 1", map[string]string{
		"userID":      data.UserID,
		"warehouseID": data.WarehouseID,
	})

	if err != nil {
		log.Println(err)
		return &pb.InternalFetchWarehouseRoleResponse_Response{}, fmt.Errorf("error while querying db: %v", err.Error())
	}
	res := make([]Transaction, 1)

	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		log.Println(err)
		return &pb.InternalFetchWarehouseRoleResponse_Response{}, fmt.Errorf("error while unmarshaling data: %v", err)
	}

	if len(res[0].Result) == 0 {
		return &pb.InternalFetchWarehouseRoleResponse_Response{}, fmt.Errorf("user does not belong to requested warehouse")
	}

	return &pb.InternalFetchWarehouseRoleResponse_Response{
		Role: pb.Roles(res[0].Result[0]["role"].(float64)),
	}, nil

}

func (h *Handler) ListUsers(data *pb.ListUsersRequest, pageCount int) (*pb.ListUsersResponse_Response, error) {
	var page int32

	if data.Page > int32(pageCount) {
		page = int32(pageCount)
	} else {
		page = int32(data.Page) - 1
	}

	queryRes, err := h.DB.Query("SELECT in.* as user, role FROM manages WHERE out = $warehouseID ORDER BY role DESC LIMIT $limit START $page", map[string]interface{}{
		"warehouseID": data.WarehouseID,
		"limit":       data.Limit,
		"page":        data.Limit * page,
	})

	if err != nil {
		log.Println(err)
		return &pb.ListUsersResponse_Response{}, fmt.Errorf("error while querying db: %v", err.Error())
	}

	res := &pb.ListUsersResponse_Response{
		Pagination: &pb.ListUsersResponsePagination{
			Limit: data.Limit,
			Page:  data.Page,
			Total: int32(pageCount),
		},
	}

	list, err := surrealdb.SmartUnmarshal[[]*pb.ListUsersResponseList](queryRes, nil)

	if err != nil {
		log.Println(err)
		return &pb.ListUsersResponse_Response{}, fmt.Errorf("error while unmarshaling data: %v", err)
	}

	res.List = list

	return res, nil

}

func (h *Handler) CountUsers(data *pb.ListUsersRequest) (int, error) {

	queryRes, err := h.DB.Query("SELECT count() FROM manages WHERE out = $warehouseID GROUP ALL", map[string]string{
		"warehouseID": data.WarehouseID,
	})

	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error while counting users: %v", err.Error())
	}

	res := make([]Transaction, 1)

	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error while unmarshaling data:%v", err)
	}

	if len(res[0].Result) < 1 {
		return 0, nil
	}

	return int(res[0].Result[0]["count"].(float64)), nil
}

func (h *Handler) CountWarehouses(data *pb.ListWarehousesRequest) (int, error) {

	queryRes, err := h.DB.Query("SELECT count() FROM manages WHERE in = $uuid GROUP ALL", map[string]string{
		"uuid": data.Uuid,
	})

	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error while counting users: %v", err.Error())
	}

	res := make([]Transaction, 1)

	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error while unmarshaling data:%v", err)
	}

	if len(res[0].Result) < 1 {
		return 0, nil
	}

	return int(res[0].Result[0]["count"].(float64)), nil

}
