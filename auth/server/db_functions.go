package main

import (
	"fmt"
	"log"

	"github.com/surrealdb/surrealdb.go"
)

func GetUser(email string) User {
	data, err := db.Query("SELECT * FROM users WHERE email = $userEmail", map[string]interface{}{
		"userEmail": email,
	})

	if err != nil {
		return User{}
	}

	res := make([]User, 1)

	_, err = surrealdb.UnmarshalRaw(data, &res)

	if err != nil {
		log.Printf("Error while unmarshaling data: %v", err)
		return User{}
	}

	return res[0]
}

func CreateUser(user User) error {
	_, err := db.Query("CREATE users:uuid() SET email = $email, passwd = $passwd", map[string]interface{}{
		"email":  user.Email,
		"passwd": user.Passwd,
	})
	if err != nil {
		return fmt.Errorf("Error while creating user: %v", err)
	}
	return nil
}
