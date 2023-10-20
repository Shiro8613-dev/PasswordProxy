package cryptoSys

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"io"
)

// GenerateSalt generate
func GenerateSalt() ([]byte, error) {
	var r = rand.Reader
	b := make([]byte, 64)

	_, err := io.ReadFull(r, b)

	if err != nil {
		return nil, err
	}

	return []byte(hex.EncodeToString(b)), nil
}

// GenerateHashedPassword generator
func GenerateHashedPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 15)

	return string(hashed)
}

// VerifyPassword verify
func VerifyPassword(rawPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))

	if err == nil {
		return true
	}

	return false
}
