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
