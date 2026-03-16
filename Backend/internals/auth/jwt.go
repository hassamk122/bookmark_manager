package auth

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID, username string, secretKey []byte) (string, error) {
	if !validUsernameAndID(userID, username) {
		return "", errors.New("Invalid input")
	}

	claims := NewClaim(userID, username)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseJWT(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func validUsernameAndID(id, username string) bool {
	return removeSpaces(id) != "" && removeSpaces(username) != ""
}
