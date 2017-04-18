package tcp

import (
	"fmt"
	. "galaxy/logs"
	. "galaxy/nets/packet"
	"galaxy/utils"
	"net"
	"sync"
	"time"
)

type Connector struct {
	connectAddr    *net.TCPAddr
	connectSession *Session

	name      string
	isEncrypt bool

	recvBufSize   int
	sendBufSize   int
	sendChanSize  int
	heartInterval int64

	exitChan chan bool

	genSidCb      func() int64
	helloCb       func(s *Session)
	loseLogicCb   func()
	sessionRecvCb func(s *Session, packet *GxPacket)

	waitGroup sync.WaitGroup
	runStatus bool
	runMutex  sync.Mutex
}

func NewConnector(name string, isEncrypt bool,
	recvBufSize int, sendBufSize int,
	sendChanSize int, heartInterval int64,
	genSidCb func() int64,
	helloCb func(s *Session),
	sessionRecvCb func(s *Session, packet *GxPacket)) *Connector {

	c := new(Connector)
	c.name = name
	c.isEncrypt = isEncrypt
	c.recvBufSize = recvBufSize
	c.sendBufSize = sendBufSize
	c.sendChanSize = sendChanSize
	c.heartInterval = heartInterval

	c.genSidCb = genSidCb
	if c.genSidCb == nil {
		panic(fmt.Sprintln("NewConnector[", name, "] genSidCb is nil"))
	}

	c.helloCb = helloCb
	if c.helloCb == nil {
		panic(fmt.Sprintln("NewConnector[", name, "] helloCb is nil"))
	}

	c.sessionRecvCb = sessionRecvCb
	if c.sessionRecvCb == nil {
		panic(fmt.Sprintln("NewConnector[", name, "] sessionRecvCb is nil"))
	}

	return c
}

func (this *Connector) RegisterLoseLogicCallBack(f func()) {
	this.loseLogicCb = f
}

func (this *Connector) connect() error {
	conn, err := net.DialTCP("tcp", nil, this.connectAddr)
	if err != nil {
		return err
	}
	this.connectSession = NewSession(this.genSidCb(), this.isEncrypt, true, this.heartInterval,
		func(s *Session) {
			if this.loseLogicCb != nil {
				this.loseLogicCb() //回调函数　关联上层逻辑
			}
		},
		this.sessionRecvCb)
	this.connectSession.Start(conn, this.recvBufSize, this.sendBufSize, this.sendChanSize)
	this.connectSession.SetLegal()
	if this.helloCb != nil {
		this.helloCb(this.connectSession)
	}
	GxLogDebug("Connector [", this.name, "] IP: ", this.connectAddr, " connect success ")
	this.connectSession.Wait()

	return nil
}

func (this *Connector) Connect(host string, port int) {
	this.runMutex.Lock()
	defer this.runMutex.Unlock()
	if !this.runStatus {
		ip := net.ParseIP(host)
		this.connectAddr = &net.TCPAddr{IP: ip, Port: port}
		GxLogDebug("Connector [", this.name, "] IP: ", this.connectAddr, " init session success ")
		this.waitGroup.Add(1)
		go func() {
			defer utils.Stack()
		Exit:
			for {
				select {
				case <-this.exitChan:
					break Exit
				default:
					GxLogDebug("Connector [", this.name, "] IP: ", this.connectAddr, " start connect... ")
					err := this.connect()
					if err != nil {
						GxLogError("Connector [", this.name, "] IP: ", this.connectAddr, "error : ", err.Error())
					}
					time.Sleep(time.Second)
					GxLogDebug("Connector [", this.name, "] IP: ", this.connectAddr, " maybe reconnect... ")
				}
			}
			GxLogDebug("Connector [", this.name, "] IP: ", this.connectAddr, " end to connect")
			this.waitGroup.Done()
		}()

		this.runStatus = true
	}

	return
}

func (this *Connector) SendPacket(p *GxPacket) {
	if this.runStatus {
		this.connectSession.SendPacket(p)
	}
}

func (this *Connector) Stop() {
	this.runMutex.Lock()
	defer this.runMutex.Unlock()
	if this.runStatus {
		close(this.exitChan)
		this.connectSession.Stop()
		this.connectSession.Wait()
		this.waitGroup.Wait()
		this.runStatus = false
	}
}

func (this *Connector) Wait() {
	this.waitGroup.Wait()
}

func (this *Connector) DialAddr() net.Addr {
	return this.connectAddr
}
