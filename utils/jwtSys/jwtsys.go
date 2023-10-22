package jwtSys

import (
	"PasswordProxy/databaseSys"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

var salt string

// SetSalt set
func SetSalt(s string) {
	salt = s
}

// JwtGenerate generate
func JwtGenerate(username string, database databaseSys.DataBaseStruct) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	crypt, err := database.ReadCrypto()
	if err != nil {
		return "", err
	}

	salt = crypt.Salt

	accessToken, err := token.SignedString([]byte(crypt.Salt))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// JwtVerify verify
func JwtVerify(accessToken string) (string, error) {

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
