package main

import (
	"fmt"
	"galaxy/config"
	"io/ioutil"
)

type ServerConfig struct {
	mysql         string
	mysqlChanSize int
	ipData        string
}

func LoadServerConfig() (*ServerConfig, error) {
	buff, err := ioutil.ReadFile("./config/serverconfig.ini")
	if err != nil {
		fmt.Println("LoadServerConfig Error : ", err)
		return nil, err
	}

	config, err := config.NewConfigData("json", []byte(buff))
	if err != nil {
		fmt.Println("LoadServerConfig Error : ", err)
		return nil, err
	}

	mysql := config.String("mysql")
	mysqlChanSize, _ := config.Int("mysql_chansize")
	ipData := config.String("ip_data")

	sc := &ServerConfig{
		mysql:         mysql,
		mysqlChanSize: mysqlChanSize,
		ipData:        ipData,
	}

	return sc, nil
}
