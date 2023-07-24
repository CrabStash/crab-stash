package db

import (
	"fmt"
	"log"
	"os"

	pb "github.com/CrabStash/crab-stash-protofiles/auth/proto"
	"github.com/CrabStash/crab-stash/auth/internal/utils"
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

	if _, err = db.Use("users", "users"); err != nil {
		log.Fatalf("Failed to use users/users: %v\n", err.Error())
	}
	return Handler{db}
}

func (h *Handler) GetUserByEmail(email string) (UserCrucial, error) {
	data, err := h.DB.Query("SELECT email, passwd, id FROM users WHERE email = $userEmail", map[string]interface{}{
		"userEmail": email,
	})

	if err != nil {
		return UserCrucial{}, fmt.Errorf("error while querying db: %v", err.Error())
	}

	res := make([]UserCrucial, 1)

	_, err = surrealdb.UnmarshalRaw(data, &res)

	if err != nil {
		return UserCrucial{}, fmt.Errorf("error while unmarshaling data: %v", err.Error())
	}

	return res[0], nil

}

func (h *Handler) GetUserByUUID(uuid string) (UserCrucial, error) {
	data, err := h.DB.Query("SELECT id FROM users WHERE id = $id", map[string]interface{}{
		"id": uuid,
	})

	if err != nil {
		return UserCrucial{}, fmt.Errorf("error while querying db: %v", err.Error())
	}

	res := make([]UserCrucial, 1)

	_, err = surrealdb.UnmarshalRaw(data, &res)

	if err != nil {
		return UserCrucial{}, fmt.Errorf("error while unmarshaling data: %v", err.Error())
	}

	if res[0].Id == "" {
		return UserCrucial{}, fmt.Errorf("user don't exist")
	}

	return res[0], nil

}

func (h *Handler) CreateUser(user *pb.RegisterRequest) error {
	hash, err := utils.HashPassword(user.Passwd)
	if err != nil {
		return fmt.Errorf("error while hashing passwd: %v", err.Error())
	}

	_, err = h.DB.Query("CREATE users:uuid() SET email = $email, passwd = $passwd, firstName = $firstName, lastName = $lastName", map[string]interface{}{
		"email":     user.Email,
		"passwd":    hash,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	})
	if err != nil {
		return fmt.Errorf("error while creating user: %v", err)
	}
	return nil
}
