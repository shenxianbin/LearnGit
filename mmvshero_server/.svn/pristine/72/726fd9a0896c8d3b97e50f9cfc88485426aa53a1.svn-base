package main

import (
	"fmt"
	"galaxy/config"
	"io/ioutil"
)

type ServerConfig struct {
	scheme_path string
}

func LoadServerConfig() (*ServerConfig, error) {
	buff, err := ioutil.ReadFile("./config/serverconfig.ini")
	if err != nil {
		fmt.Println("LoadServerConfig Error : ", err)
		return nil, err
	}

	config, err := config.NewConfigData("json", []byte(buff))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sp := config.String("scheme_path")

	sc := &ServerConfig{
		scheme_path: sp,
	}

	return sc, nil
}
