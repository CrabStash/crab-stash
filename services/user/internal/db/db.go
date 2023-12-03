package db

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/CrabStash/crab-stash-protofiles/user/proto"
	surrealdb "github.com/surrealdb/surrealdb.go"
)

type Transaction struct {
	Result string `json:"result"`
	Status string `json:"status"`
	Time   string `json:"time"`
}

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

func (h *Handler) GetMeInfo(data *pb.MeInfoRequest) (*pb.MeInfoResponse_Data, error) {
	queryRes, err := h.DB.Select(data.UserID)

	if err != nil {
		return &pb.MeInfoResponse_Data{}, fmt.Errorf("error while querying user info: %v", err)
	}

	res := &pb.MeInfoResponse_Data{
		Data: &pb.MeInfoResponse_Response{},
	}

	err = surrealdb.Unmarshal(queryRes, &res.Data)

	if err != nil {
		return &pb.MeInfoResponse_Data{}, fmt.Errorf("error while unmarshaling data: %v", err)
	}
	return res, nil
}

func (h *Handler) DbUpdateUserInfo(usr *pb.UpdateUserInfoRequest) error {

	_, err := h.DB.Query("UPDATE $userID MERGE $data", map[string]interface{}{
		"userID": usr.UserID,
		"data":   usr.Data,
	})

	if err != nil {
		return fmt.Errorf("error while updating user info: %v", err)
	}

	return nil
}

func (h *Handler) DbGetUserInfo(data *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse_Data, error) {
	queryRes, err := h.DB.Select(data.Id)

	if err != nil {
		return &pb.GetUserInfoResponse_Data{}, fmt.Errorf("error while querying user info: %v", err)
	}

	_, ok := queryRes.(map[string]interface{})

	if !ok {
		return &pb.GetUserInfoResponse_Data{}, fmt.Errorf("user does not exist: %v", err)
	}

	res := &pb.GetUserInfoResponse_Data{
		Data: &pb.GetUserInfoResponse_Response{},
	}

	err = surrealdb.Unmarshal(queryRes, &res.Data)

	if err != nil {
		return &pb.GetUserInfoResponse_Data{}, fmt.Errorf("error while unmarshaling data: %v", err)
	}
	return res, nil
}

func (h *Handler) DbDeleteUser(usr *pb.DeleteUserRequest) error {
	queryRes, err := h.DB.Select(usr.UserID)
	if err != nil {
		return fmt.Errorf("error while checking if user is owner: %v", err)
	}
	isOwner := Ownership{}
	err = surrealdb.Unmarshal(queryRes, &isOwner)
	if err != nil {
		return fmt.Errorf("error while unmarshaling data: %v", err)
	}

	if len(isOwner.Owns) != 0 {
		return fmt.Errorf("cannot delete user if he is an owner")
	}
	_, err = h.DB.Delete(usr.UserID)

	if err != nil {
		return fmt.Errorf("error while deleting user: %v", err)
	}
	return nil
}

func (h *Handler) DbInternalGetUserByEmail(usr *pb.InternalGetUserByEmailRequest) (UserCrucial, error) {
	queryRes, err := h.DB.Select(usr.Email)

	if err != nil {
		return UserCrucial{}, fmt.Errorf("error while querying user: %v", err)
	}
	res := UserCrucial{}
	err = surrealdb.Unmarshal(queryRes, &res)

	if err != nil {
		return UserCrucial{}, fmt.Errorf("error while unmarshaling user data: %v", err)
	}
	return res, nil
}

func (h *Handler) DbGetUserbyUUID(usr *pb.InternalGetUserByUUIDCheck) (*pb.InternalGetUserByUUIDCheck, error) {
	queryRes, err := h.DB.Select(usr.Id)
	if err != nil {
		return &pb.InternalGetUserByUUIDCheck{}, fmt.Errorf("error while querying id: %v ", err)
	}
	userUnmarshal := &pb.InternalGetUserByUUIDCheck{}
	err = surrealdb.Unmarshal(queryRes, &userUnmarshal)

	if err != nil {
		return &pb.InternalGetUserByUUIDCheck{}, fmt.Errorf("error while unmarshaling user data: %v", err)
	}
	return userUnmarshal, nil
}

func (h *Handler) ChangeUserPassword(req *pb.ChangePasswordRequest) *pb.ChangePasswordResponse {

	hashNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
	if err != nil {
		return &pb.ChangePasswordResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while hashing password",
		}
	}

	queryRes, _ := h.DB.Query(`
	BEGIN TRANSACTION; 
	LET $CurrentPassword = (SELECT VALUE passwd FROM ONLY $userID);
	IF crypto::bcrypt::compare($CurrentPassword,$OldPassword) {
		UPDATE $userID SET passwd = $NewPassword;
	} ELSE {
		THROW "Passwords don't match";
	};
	COMMIT TRANSACTION;
	`, map[string]string{
		"OldPassword": req.OldPassword,
		"NewPassword": string(hashNewPassword),
		"userID":      req.UserID,
	})

	transaction := make([]Transaction, 2)

	err = surrealdb.Unmarshal(queryRes, &transaction)
	if err != nil {
		log.Println(err)
		return &pb.ChangePasswordResponse{
			Status:   http.StatusInternalServerError,
			Response: "error while unmarshaling data",
		}
	}

	if transaction[0].Status == "ERR" {
		log.Println(err)

		return &pb.ChangePasswordResponse{
			Status:   http.StatusInternalServerError,
			Response: transaction[1].Result,
		}
	}

	return &pb.ChangePasswordResponse{
		Status:   http.StatusOK,
		Response: "Password Changed",
	}

}
