package galaxy

import (
	"errors"
	"fmt"
	"galaxy/config"
)

type ListenerConfig struct {
	listenType    string
	encrypt       bool
	listenHost    string
	listenPort    int
	maxConn       int
	heartInterval int
	recvbufSize   int
	sendbufSize   int
	sendchanSize  int
}

type ConnectorConfig struct {
	connectServerId int
	connectType     string
	encrypt         bool
	connectHost     string
	connectPort     int
	heartInterval   int
	recvbufSize     int
	sendbufSize     int
	sendchanSize    int
}

func (this *ConnectorConfig) ConnectServerId() int {
	return this.connectServerId
}

type GxConfig struct {
	serverId   int
	serverType string
	loglv      int
	logPath    string

	redisHost   string
	redisPort   int
	redisPasswd string

	listenerConfig  []ListenerConfig
	connectorConfig []ConnectorConfig
}

func (this *GxConfig) load(jsonConfig string) error {
	config, err := config.NewConfigData("json", []byte(jsonConfig))
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.serverId, err = config.Int("serverId")
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.serverType = config.String("serverType")

	this.loglv, err = config.Int("loglv")
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.logPath = config.String("logPath")

	this.redisHost = config.String("redis::host")

	this.redisPort, err = config.Int("redis::port")
	if err != nil {
		fmt.Println(err)
		return err
	}

	this.redisPasswd = config.String("redis::passwd")

	//Listener
	listenArray, err := config.DIY("listener")
	if err == nil {
		listenArrayCasted := listenArray.([]interface{})
		if listenArrayCasted == nil {
			fmt.Println("listenArrayCasted nil")
			return errors.New("listenArrayCasted nil")
		} else {
			this.listenerConfig = make([]ListenerConfig, len(listenArrayCasted))
			for index, value := range listenArrayCasted {
				elem := value.(map[string]interface{})
				this.listenerConfig[index].listenType = elem["listenType"].(string)
				encrypt := elem["encrypt"].(string)
				if encrypt == "true" {
					this.listenerConfig[index].encrypt = true
				} else {
					this.listenerConfig[index].encrypt = false
				}
				this.listenerConfig[index].listenHost = elem["listenHost"].(string)
				this.listenerConfig[index].listenPort = int(elem["listenPort"].(float64))
				this.listenerConfig[index].maxConn = int(elem["maxConn"].(float64))
				this.listenerConfig[index].heartInterval = int(elem["heartInterval"].(float64))
				this.listenerConfig[index].recvbufSize = int(elem["recvbufSize"].(float64))
				this.listenerConfig[index].sendbufSize = int(elem["sendbufSize"].(float64))
				this.listenerConfig[index].sendchanSize = int(elem["sendchanSize"].(float64))
			}
		}
	}

	//Connector
	connectArray, err := config.DIY("connector")
	if err == nil {
		connectArrayCasted := connectArray.([]interface{})
		if connectArrayCasted == nil {
			fmt.Println("connectArrayCasted nil")
			return errors.New("connectArrayCasted nil")
		} else {
			this.connectorConfig = make([]ConnectorConfig, len(connectArrayCasted))
			for index, value := range connectArrayCasted {
				elem := value.(map[string]interface{})
				this.connectorConfig[index].connectServerId = int(elem["connectServerId"].(float64))
				this.connectorConfig[index].connectType = elem["connectType"].(string)
				encrypt := elem["encrypt"].(string)
				if encrypt == "true" {
					this.connectorConfig[index].encrypt = true
				} else {
					this.connectorConfig[index].encrypt = false
				}
				this.connectorConfig[index].connectHost = elem["connectHost"].(string)
				this.connectorConfig[index].connectPort = int(elem["connectPort"].(float64))
				this.connectorConfig[index].heartInterval = int(elem["heartInterval"].(float64))
				this.connectorConfig[index].recvbufSize = int(elem["recvbufSize"].(float64))
				this.connectorConfig[index].sendbufSize = int(elem["sendbufSize"].(float64))
				this.connectorConfig[index].sendchanSize = int(elem["sendchanSize"].(float64))
			}
		}
	}

	return nil
}

func (this *GxConfig) GetConnectorConfig(index int) *ConnectorConfig {
	if index >= len(this.connectorConfig) {
		return nil
	}
	return &this.connectorConfig[index]
}

func (this *GxConfig) GetServerId() int32 {
	return int32(this.serverId)
}
