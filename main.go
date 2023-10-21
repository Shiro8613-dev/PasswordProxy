package main

import (
	"PasswordProxy/configSys"
	"PasswordProxy/databaseSys"
	"PasswordProxy/utils/cryptoSys"
	"PasswordProxy/webSys"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"gorm.io/gorm"
)

var (
	config     configSys.Config
	database   databaseSys.DataBaseStruct
	cryptoSalt databaseSys.Crypto
	store      sessions.Store
)

func init() {
	//===config_load===
	c, err := configSys.Load("config.toml")
	if err != nil {
		panic(err)
	}

	config = c
	//===config_load===

	//===database_connect===
	d, err := databaseSys.Connect(c.Database)
	if err != nil {
		panic(err)
	}

	err = d.Init()
	if err != nil {
		return
	}

	database = d
	//===database_connect===

	//===crypto_salt===
	cryptoSalt, err = d.ReadCrypto()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		salt, err := cryptoSys.GenerateSalt()
		if err != nil {
			panic(err)
		}

		cryptoSalt = databaseSys.Crypto{
			Salt: string(salt),
		}

		err = d.CreateCrypto(cryptoSalt)
		if err != nil {
			panic(err)
		}
	}
	//===crypto_salt===

	//===session===
	salt, err := d.ReadCrypto()
	if err != nil {
		panic(err)
	}

	store, err = redis.NewStore(10, "tcp",
		fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		c.Redis.Password, []byte(salt.Salt))
	if err != nil {
		panic(err)
	}

	//===session===

}

func main() {
	fmt.Println(config)
	server := webSys.NewWebServer(config.Listener, config.Proxy, database, store)
	panic(server.Start())
}
