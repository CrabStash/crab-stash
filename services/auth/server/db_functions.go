package main

import (
	"fmt"
	"log"

	pb "github.com/CrabStash/crab-stash/auth/proto"
	"github.com/surrealdb/surrealdb.go"
)

func GetUser(email string) pb.User {
	data, err := db.Query("SELECT * FROM users WHERE email = $userEmail", map[string]interface{}{
		"userEmail": email,
	})

	if err != nil {
		return pb.User{}
	}

	res := make([]pb.User, 1)

	_, err = surrealdb.UnmarshalRaw(data, &res)

	if err != nil {
		log.Printf("Error while unmarshaling data: %v", err)
		return pb.User{}
	}

	return res[0]
}

func CreateUser(user pb.User) error {
	_, err := db.Query("CREATE users:uuid() SET email = $email, passwd = $passwd", map[string]interface{}{
		"email":  user.Email,
		"passwd": user.Passwd,
	})
	if err != nil {
		return fmt.Errorf("Error while creating user: %v", err)
	}
	return nil
}
