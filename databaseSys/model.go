package databaseSys

import (
	"PasswordProxy/utils/cryptoSys"
	"gorm.io/gorm"
	"strconv"
)

type Crypto struct {
	gorm.Model
	Salt string
}

type User struct {
	gorm.Model
	Username  string
	Password1 string
	Password2 string
	Password3 string
	PinCode   string
	Admin     bool
}

// UserCreate create
func UserCreate(
	username string, password1 string,
	password2 string, password3 string,
	pinCode int, admin bool) User {

	return User{
		Username:  username,
		Password1: cryptoSys.GenerateHashedPassword(password1),
		Password2: cryptoSys.GenerateHashedPassword(password2),
		Password3: cryptoSys.GenerateHashedPassword(password3),
		PinCode:   cryptoSys.GenerateHashedPassword(strconv.Itoa(pinCode)),
		Admin:     admin,
	}
}

// CryptoCreate create
func CryptoCreate(salt string) Crypto {
	return Crypto{
		Salt: salt,
	}
}
