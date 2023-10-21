package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JwtGenerate generate
func JwtGenerate(username string, salt string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// JwtVerify verify
func JwtVerify(accessToken string, salt string) (string, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(salt), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	} else {
		return "", err
	}
}
