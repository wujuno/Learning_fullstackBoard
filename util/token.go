package util

import (
	"fullstackboard/model"

	jwt "github.com/dgrijalva/jwt-go"
)

func makeTokenString(user *model.User) (string error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
}