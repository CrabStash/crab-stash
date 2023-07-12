package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/CrabStash/crab-stash/auth/config"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", fmt.Errorf("Error while hashing password: %v", err)
	}

	return string(bytes), nil
}

func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func SignJWT(userID string) (string, error) {
	sampleSecret := []byte("LoffciamF1")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(config.Cfg.GetJWTDuration()) * time.Minute)
	claims["id"] = userID

	tokenString, err := token.SignedString(sampleSecret)
	if err != nil {
		return "", fmt.Errorf("Error while signing jwt: %v", err)
	}

	return tokenString, nil
}

func ValidateJWT(jwtToken string) bool {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("error while validating token")
		}
		return "", nil
	})

	if err != nil {
		log.Printf("Error while Validating JWT: %v", err)
		return false
	}

	log.Printf("%v", token.Valid)
	return token.Valid
}
