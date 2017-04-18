package tcp

import (
	"fmt"
	. "galaxy/event"
	. "galaxy/logs"
	. "galaxy/nets/packet"
	"galaxy/utils"
	"net"
	"sync"
)

type Listener struct {
	listenAddr *net.TCPAddr
	listener   *net.TCPListener

	name      string
	isEncrypt bool

	maxConn       int
	recvBufSize   int
	sendBufSize   int
	sendChanSize  int
	heartInterval int64

	connectedSession map[int64]*Session
	sessionCount     chan int64
	exitChan         chan struct{}
	mutex            sync.Mutex

	genSidCb      func() int64
	loseLogicCb   func(sid int64)
	sessionRecvCb func(s *Session, packet *GxPacket)

	waitGroup sync.WaitGroup
	runStatus bool
	runMutex  sync.Mutex
}

func NewListener(name string, isEncrypt bool,
	maxConn int, recvBufSize int, sendBufSize int,
	sendChanSize int, heartInterval int64,
	genSidCb func() int64,
	sessionRecvCb func(s *Session, packet *GxPacket)) *Listener {

	l := new(Listener)
	l.name = name
	l.isEncrypt = isEncrypt
	l.maxConn = maxConn
	l.recvBufSize = recvBufSize
	l.sendBufSize = sendBufSize
	l.sendChanSize = sendChanSize
	l.heartInterval = heartInterval

	l.genSidCb = genSidCb
	if l.genSidCb == nil {
		panic(fmt.Sprintln("NewListener[", name, "] genSidCb is nil"))
	}

	l.sessionRecvCb = sessionRecvCb
	if l.sessionRecvCb == nil {
		panic(fmt.Sprintln("NewListener[", name, "] sessionRecvCb is nil"))
	}

	return l
}

func (this *Listener) RegisterLoseLogicCallBack(f func(sid int64)) {
	this.loseLogicCb = f
}

func (this *Listener) initSession() {
	this.sessionCount = make(chan int64, this.maxConn)
	for i := 0; i < this.maxConn; i++ {
		sid := this.genSidCb()
		this.sessionCount <- sid
	}
}

func (this *Listener) Listening(host string, port int) (err error) {
	this.runMutex.Lock()
	defer this.runMutex.Unlock()
	if !this.runStatus {
		ip := net.ParseIP(host)
		this.listenAddr = &net.TCPAddr{IP: ip, Port: port}
		this.listener, err = net.ListenTCP("tcp", this.listenAddr)
		if err != nil {
			this.listener.Close()
			GxLogError(err.Error())
			return
		}
		this.resetSession()
		this.initSession()

		GxLogDebug("Listener [", this.name, "] IP: ", this.listener.Addr(), " create success, now begin to Listen......")
		this.waitGroup.Add(1)
		go func() {
			defer utils.Stack()
			GxLogDebug("Listener [", this.name, "] IP: ", this.listener.Addr(), " begin to Listen......")
		Exit:
			for {
				select {
				case <-this.exitChan:
					break Exit
				case sid := <-this.sessionCount:
					GxLogDebug("Listener [", this.name, "] IP: ", this.listener.Addr(), " accepting")
					conn, err := this.listener.AcceptTCP()
					if err == nil {
						GxLogDebug("Listener [", this.name, "] IP: ", this.listener.Addr(), " accept conn IP: ", conn.RemoteAddr())
						session := NewSession(sid, this.isEncrypt, false, this.heartInterval,
							func(s *Session) {
								this.delSession(s.Sid())
								if this.loseLogicCb != nil {
									GxEvent().Execute(func(args ...interface{}) {
										this.loseLogicCb(s.Sid())
									})
								}
								this.sessionCount <- s.Sid()
							},
							this.sessionRecvCb)
						session.Start(conn, this.recvBufSize, this.sendBufSize, this.sendChanSize)
						this.addSession(session)
					} else {
						GxLogDebug("Listener [", this.name, "] IP: ", this.listener.Addr(), " accept failed")
					}
				}
			}
			GxLogDebug("Listener [", this.name, "] IP: ", this.listener.Addr(), " end to Listen......")
			this.waitGroup.Done()
		}()

		this.runStatus = true
	}

	return
}

func (this *Listener) SendPacket(sid int64, p *GxPacket) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if s, has := this.connectedSession[sid]; has {
		s.SendPacket(p)
	}
}

func (this *Listener) Remove(sid int64) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	session, has := this.connectedSession[sid]
	if has && session != nil {
		session.Stop()
		delete(this.connectedSession, sid)
		GxLogDebug("Remove sid : ", sid, " and num : ", len(this.connectedSession))
	}
}

func (this *Listener) RemoveAll() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	GxLogDebug("RemoveAll Start")
	for _, session := range this.connectedSession {
		if session != nil {
			session.Stop()
			session.Wait()
		}
	}
	this.connectedSession = make(map[int64]*Session)
	GxLogDebug("RemoveAll End")
}

func (this *Listener) Stop() {
	this.runMutex.Lock()
	defer this.runMutex.Unlock()
	if this.runStatus {
		this.mutex.Lock()
		defer this.mutex.Unlock()
		for _, session := range this.connectedSession {
			if session != nil {
				session.Stop()
				session.Wait()
			}
		}
		close(this.exitChan)
		this.listener.Close()
		this.waitGroup.Wait()
		this.runStatus = false
	}
}

func (this *Listener) Wait() {
	this.waitGroup.Wait()
}

func (this *Listener) LocalAddr() net.Addr {
	return this.listenAddr
}

func (this *Listener) addSession(session *Session) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.connectedSession[session.sid] = session
	GxLogDebug("Listener [", this.name, "] addSession and num : ", len(this.connectedSession))
}

func (this *Listener) delSession(sid int64) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	delete(this.connectedSession, sid)
	GxLogDebug("Listener [", this.name, "] DelSession and num : ", len(this.connectedSession))
}

func (this *Listener) resetSession() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.connectedSession = make(map[int64]*Session)
	GxLogDebug("Listener [", this.name, "] ResetSession !")
}
