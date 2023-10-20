package main

import (
	"PasswordProxy/configSys"
	"fmt"
)

var config configSys.Config

func init() {
	//===config_load===
	c, err := configSys.Load("config.toml")
	if err != nil {
		panic(err)
	}
	config = c
	//===config_load===
}

func main() {
	fmt.Println(config)
}
