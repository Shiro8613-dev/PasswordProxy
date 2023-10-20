package databaseSys

import (
	"PasswordProxy/configSys"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBaseStruct struct {
	db *gorm.DB
}

// Connect Connection Database Function
func Connect(config configSys.DatabaseConfig) (DataBaseStruct, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return DataBaseStruct{}, err
	}

	return DataBaseStruct{
		db: db,
	}, err
}

// Init InitDatabase
func (d DataBaseStruct) Init() error {
	return d.db.AutoMigrate(&User{}, &Crypto{})
}

//===User==

// CreateUser create
func (d DataBaseStruct) CreateUser(user User) error {
	err := d.db.Model(&User{}).Where("username = ?", user.Username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = d.db.Create(&user).Error
		if err != nil {
			return err
		}
	}

	return gorm.ErrRegistered
}

// ReadUser read
func (d DataBaseStruct) ReadUser(id int) (User, error) {
	var user User
	err := d.db.First(&user, id).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// ReadUsers reads
func (d DataBaseStruct) ReadUsers(limit int, offset int) ([]User, error) {
	var users []User
	err := d.db.Order("id").Limit(limit).Offset(offset).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// FindUser find
func (d DataBaseStruct) FindUser(name string) (User, error) {
	var user User
	err := d.db.Model(&User{}).Where("username = ? ", name).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, err
	}

	return user, nil

}

// UpdateUser update
func (d DataBaseStruct) UpdateUser(id int, user User) error {
	return d.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

// DeleteUser delete
func (d DataBaseStruct) DeleteUser(id int) error {
	return d.db.Delete(&User{}, id).Error
}

//===User==

//===Crypto==

// CreateCrypto create
func (d DataBaseStruct) CreateCrypto(crypt Crypto) error {
	return d.db.Create(&crypt).Error
}

// ReadCrypto read
func (d DataBaseStruct) ReadCrypto() (Crypto, error) {
	var crypto Crypto
	err := d.db.First(&crypto, 1).Error

	if err != nil {
		return Crypto{}, err
	}

	return crypto, nil
}

// UpdateCrypto update
func (d DataBaseStruct) UpdateCrypto(crypt Crypto) error {
	return d.db.Model(&Crypto{}).Where("id = ?", 1).Updates(crypt).Error
}

//===Crypto==
