package main

import (
	"fmt"
	"galaxy/config"
	"io/ioutil"
)

type ServerConfig struct {
	SchemePath     string
	OpenServerTime string
	RechargePort   string
}

func LoadServerConfig() (*ServerConfig, error) {
	buff, err := ioutil.ReadFile("./config/serverconfig.ini")
	if err != nil {
		fmt.Println("LoadServerConfig Error : ", err)
		return nil, err
	}

	config, err := config.NewConfigData("json", []byte(buff))
	if err != nil {
		fmt.Println("LoadServerConfig json error: ", err)
		return nil, err
	}

	sp := config.String("scheme_path")
	os_time := config.String("open_server_time")
	rp := config.String("recharge_port")

	sc := &ServerConfig{
		SchemePath:     sp,
		OpenServerTime: os_time,
		RechargePort:   rp,
	}

	return sc, nil
}
