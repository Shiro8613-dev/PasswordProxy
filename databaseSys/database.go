package databaseSys

import (
	"PasswordProxy/configSys"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBaseSys struct {
	db *gorm.DB
}

// Connect Connection Database Function
func Connect(config configSys.DatabaseConfig) (DataBaseSys, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return DataBaseSys{}, err
	}

	return DataBaseSys{
		db: db,
	}, err
}

// Init InitDatabase
func (d DataBaseSys) Init() error {
	return d.db.AutoMigrate(&User{}, &Crypto{})
}

//===User==

// CreateUser create
func (d DataBaseSys) CreateUser(user User) error {
	return d.db.Create(&user).Error
}

// ReadUser read
func (d DataBaseSys) ReadUser(id int) (User, error) {
	var user User
	err := d.db.First(&user, id).Error

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// ReadUsers reads
func (d DataBaseSys) ReadUsers(limit int, offset int) ([]User, error) {
	var users []User
	err := d.db.Order("id").Limit(limit).Offset(offset).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser update
func (d DataBaseSys) UpdateUser(id int, user User) error {
	return d.db.Model(&User{}).Where("id = ?", id).Updates(user).Error
}

// DeleteUser delete
func (d DataBaseSys) DeleteUser(id int) error {
	return d.db.Delete(&User{}, id).Error
}

//===User==

//===Crypto==

// CreateCrypto create
func (d DataBaseSys) CreateCrypto(crypt Crypto) error {
	return d.db.Create(&crypt).Error
}

// ReadCrypto read
func (d DataBaseSys) ReadCrypto() (Crypto, error) {
	var crypto Crypto
	err := d.db.First(&crypto, 1).Error

	if err != nil {
		return Crypto{}, err
	}

	return crypto, nil
}

// UpdateCrypto update
func (d DataBaseSys) UpdateCrypto(crypt Crypto) error {
	return d.db.Model(&Crypto{}).Where("id = ?", 1).Updates(crypt).Error
}

//===Crypto==
