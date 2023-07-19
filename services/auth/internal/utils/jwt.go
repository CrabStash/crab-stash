package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type JwtWrapper struct {
	TokenSecret   string
	RefreshSecret string
	TokenPublic   string
	RefreshPublic string
	TokenExp      uint16
	RefreshExp    uint16
}

type jwtClaims struct {
	jwt.StandardClaims
	TokenUUID string
}

func (w *JwtWrapper) SignJWT(userID string, isRefresh bool) (token string, token_uuid string, err error) {
	now := time.Now().Local().Unix()
	var exp int64
	var pk string
	token_uuid = uuid.NewV4().String()
	if isRefresh {
		exp = time.Now().Local().Add(time.Hour * time.Duration(w.RefreshExp)).Unix()
		pk = w.RefreshSecret
	} else {
		exp = time.Now().Local().Add(time.Hour * time.Duration(w.TokenExp)).Unix()
		pk = w.TokenSecret
	}

	claims := &jwtClaims{
		TokenUUID: token_uuid,
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: exp,
			IssuedAt:  now,
			NotBefore: now,
		},
	}

	decodedPK, err := base64.RawStdEncoding.DecodeString(pk)
	if err != nil {
		return "", "", fmt.Errorf("could not decode private key: %v", err.Error())
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPK)
	if err != nil {
		return "", "", fmt.Errorf("could not parse private key: %v", err.Error())
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", "", fmt.Errorf("could not sign token: %v", err.Error())
	}
	return token, token_uuid, nil

}

func (w *JwtWrapper) ValidateJWT(jwtToken string, isRefresh bool) (uuid string, token_uuid string, err error) {
	var pk string
	if isRefresh {
		pk = w.RefreshPublic
	} else {
		pk = w.TokenPublic
	}

	decodedPK, err := base64.RawStdEncoding.DecodeString(pk)
	if err != nil {
		return "", "", fmt.Errorf("could not decode private key: %v", err.Error())
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPK)
	if err != nil {
		return "", "", fmt.Errorf("could not parse private key: %v", err.Error())
	}

	parsedToken, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return "", "", fmt.Errorf("validate error: %v", err.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", "", fmt.Errorf("invalid token")
	}

	return claims["sub"].(string), claims["TokenUUID"].(string), nil

}
