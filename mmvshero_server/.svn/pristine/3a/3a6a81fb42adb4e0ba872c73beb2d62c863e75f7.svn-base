package galaxy

import (
	"fmt"
	. "galaxy/db/redis"
	"galaxy/define"
	. "galaxy/logs"
	. "galaxy/nets/packet"
	. "galaxy/nets/tcp"
	. "galaxy/utils"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

var gxService *Service

type Service struct {
	gxConfig    GxConfig
	listeners   map[string]*Listener
	dispatchers map[string]*Dispatcher
	connectors  map[string]map[int]*Connector
	redisDb     *GxRedis
}

func (this *Service) Run() error {
	//读取本地配置
	buff, err := ioutil.ReadFile("./config/gxconfig.ini")
	if err != nil {
		fmt.Println("LoadGxConfig Error : ", err)
	}
	this.gxConfig.load(string(buff))

	//读取网络配置

	//连接数据库
	if this.gxConfig.redisHost != "" {
		this.redisDb = NewGxRedis()
		this.redisDb.Run(fmt.Sprintf("%v:%v", this.gxConfig.redisHost, this.gxConfig.redisPort), this.gxConfig.redisPasswd)
	}

	//设置日志
	LogLevel(this.gxConfig.loglv)
	LogSetLogger("console", "")
	LogSetLogger("file", fmt.Sprintf(`{"filename":"%v"}`, this.gxConfig.logPath))

	//创建网络对象
	this.listeners = make(map[string]*Listener)
	this.connectors = make(map[string]map[int]*Connector)
	this.dispatchers = make(map[string]*Dispatcher)
	//监听
	for _, config := range this.gxConfig.listenerConfig {
		if _, has := this.listeners[config.listenType]; has {
			panic(fmt.Sprintf("ListenerType[%v] Repeat", config.listenType))
		}

		dispatcher, has := this.dispatchers[config.listenType]
		if !has {
			dispatcher = NewDispatcher(config.listenType)
			this.dispatchers[config.listenType] = dispatcher
		}

		listen := NewListener(config.listenType, config.encrypt, config.maxConn,
			config.recvbufSize, config.sendbufSize, config.sendchanSize, int64(config.heartInterval),
			func() int64 {
				return AllocSid(this.gxConfig.serverId)
			},
			func(s *Session, p *GxPacket) {
				dispatcher.Process(s, p)
			})
		err := listen.Listening(config.listenHost, config.listenPort)
		if err != nil {
			panic(fmt.Sprintf("ListenerType[%v:%v] Run Failed %v", config.listenType, config.listenPort, err.Error()))
		}
		this.listeners[config.listenType] = listen
	}

	//连接
	for _, config := range this.gxConfig.connectorConfig {
		dispatcher, has := this.dispatchers[config.connectType]
		if !has {
			dispatcher = NewDispatcher(config.connectType)
			this.dispatchers[config.connectType] = dispatcher
		}

		connect := NewConnector(config.connectType, config.encrypt,
			config.recvbufSize, config.sendbufSize, config.sendchanSize, int64(config.heartInterval),
			func() int64 {
				return AllocSid(this.gxConfig.serverId)
			},
			func(s *Session) {
				packet := NewPacket(s.Sid(), define.MSGCODE_HELLO_REQ)
				msg := &define.MsgHelloReq{
					ServerId:   proto.Int32(int32(this.gxConfig.serverId)),
					ServerType: proto.String(this.gxConfig.serverType)}

				data, err := proto.Marshal(msg)
				if err != nil {
					GxLogError("MSG_HELLO_REQ Marshal error")
					return
				}

				packet.SetContent(data)
				s.SendPacket(packet)
			},
			func(s *Session, p *GxPacket) {
				dispatcher.Process(s, p)
			})
		connect.Connect(config.connectHost, config.connectPort)
		if v, has := this.connectors[config.connectType]; has {
			v[config.connectServerId] = connect
		} else {
			this.connectors[config.connectType] = make(map[int]*Connector)
			this.connectors[config.connectType][config.connectServerId] = connect
		}
	}

	return nil
}

func (this *Service) RegisterMsg(serverType string, msgCode int32, f func(s *Session, packet *GxPacket)) {
	if d, has := this.dispatchers[serverType]; has {
		d.RegisterMsg(msgCode, f)
	} else {
		panic(fmt.Sprintf("registerMsg serverType [%v] nil", serverType))
	}
}

func (this *Service) RegisterUnknowMsgCb(serverType string, f func(s *Session, packet *GxPacket)) {
	if d, has := this.dispatchers[serverType]; has {
		d.RegisterUnknowMsgCb(f)
	} else {
		panic(fmt.Sprintf("registerUnknowMsgCb serverType [%v] nil", serverType))
	}
}

func (this *Service) RegisterListenerLoseLogicCallBack(listenType string, f func(sid int64)) {
	if l, has := this.listeners[listenType]; has {
		l.RegisterLoseLogicCallBack(f)
	} else {
		panic(fmt.Sprintf("registerListenerLoseLogicCallBack listenType [%v] nil", listenType))
	}
}

func (this *Service) RegisterConnectorLoseLogicCallBack(connectType string, connectServerId int, f func()) {
	if c, has := this.connectors[connectType][connectServerId]; has {
		c.RegisterLoseLogicCallBack(f)
	} else {
		panic(fmt.Sprintf("registerConnectorLoseLogicCallBack connectType [%v] connectServerId [%v] nil", connectType, connectServerId))
	}
}

func (this *Service) SendToConnector(connectType string, connectServerId int, p *GxPacket) {
	if c, has := this.connectors[connectType][connectServerId]; has {
		c.SendPacket(p)
	}
}

func (this *Service) Redis() *GxRedis {
	return this.redisDb
}

func (this *Service) RemoveListenerSession(listenType string, sid int64) {
	if l, has := this.listeners[listenType]; has {
		l.Remove(sid)
	}
}

func (this *Service) RemoveAllListenerSession(listenType string) {
	if l, has := this.listeners[listenType]; has {
		LogDebug("RemoveAllListenerSession : ", listenType)
		l.RemoveAll()
	}
}

func (this *Service) Stop() {
	for _, l := range this.listeners {
		l.Stop()
	}

	for _, cMap := range this.connectors {
		for _, c := range cMap {
			c.Stop()
		}
	}
}

func (this *Service) Wait() {
	for _, l := range this.listeners {
		l.Wait()
	}

	for _, cMap := range this.connectors {
		for _, c := range cMap {
			c.Wait()
		}
	}
}

func (this *Service) Config() *GxConfig {
	return &this.gxConfig
}

func GxService() *Service {
	return gxService
}

func init() {
	gxService = new(Service)
}
