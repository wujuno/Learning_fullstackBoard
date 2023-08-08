package util

import (
	"fmt"
	"fullstackboard/model"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func makeTokenString(user *model.User) (string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	
	claims["userId"] = user.Id
	claims["username"] = user.Name
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	jwtKey := os.Getenv("JWT_SECRET_KEY")
	t, err := accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		return "", fmt.Errorf("Failed to make token string. %s", err)
	}
	return t, nil
}