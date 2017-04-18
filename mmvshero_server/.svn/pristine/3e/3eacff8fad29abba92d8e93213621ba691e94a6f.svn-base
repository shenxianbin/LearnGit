package tcp

import (
	"fmt"
	. "galaxy/define"
	. "galaxy/logs"
	. "galaxy/nets/packet"
	"galaxy/utils"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Session struct {
	sid           int64
	isEncrpty     bool
	isHeartPump   bool
	heartInterval int64

	loseCb func(s *Session)
	recvCb func(s *Session, packet *GxPacket)

	conn       *net.TCPConn
	recvCache  []byte
	recvBuffer *utils.BinaryBuffer
	sendChan   chan *GxPacket
	exitChan   chan bool

	heartbeat_time_remote int64
	heartbeat_time_local  int64
	isLegal               bool

	waitGroup sync.WaitGroup
	runStatus bool
	runMutex  sync.Mutex
}

// 构造器
func NewSession(sid int64, isEncrpty bool,
	isHeartPump bool, heartInterval int64,
	loseCb func(s *Session),
	recvCb func(s *Session, packet *GxPacket)) *Session {
	s := new(Session)
	s.sid = sid
	s.isEncrpty = isEncrpty
	s.isHeartPump = isHeartPump
	s.heartInterval = heartInterval
	if s.heartInterval < HEART_MIN_INTERVAL {
		s.heartInterval = HEART_MIN_INTERVAL
	}
	s.loseCb = loseCb
	if s.loseCb == nil {
		panic(fmt.Sprintln("NewSession[", sid, "] loseCb is nil"))
	}

	s.recvCb = recvCb
	if s.recvCb == nil {
		panic(fmt.Sprintln("NewSession[", sid, "] recvCb is nil"))
	}

	return s
}

func (this *Session) LocalAddr() net.Addr {
	if this.conn != nil {
		return this.conn.LocalAddr()
	}
	return nil
}

func (this *Session) RemoteAddr() net.Addr {
	if this.conn != nil {
		return this.conn.RemoteAddr()
	}
	return nil
}

func (this *Session) Sid() int64 {
	return this.sid
}

func (this *Session) Start(conn *net.TCPConn, recvbufSize int, sendbufSize int, sendChanSize int) {
	this.runMutex.Lock()
	defer this.runMutex.Unlock()
	GxLogDebug("TcpSession[", this.sid, "] begin start!")
	if !this.runStatus {
		this.conn = conn
		this.conn.SetReadBuffer(recvbufSize)
		this.conn.SetWriteBuffer(sendbufSize)

		this.recvCache = make([]byte, recvbufSize)
		this.recvBuffer = utils.NewBinaryBuffer(recvbufSize)
		this.sendChan = make(chan *GxPacket, sendChanSize)
		this.exitChan = make(chan bool)
		this.isLegal = false
		GxLogDebug("TcpSession[", this.sid, "] connected ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())

		this.startIo()
		this.runStatus = true
	}
}

func (this *Session) Stop() {
	this.runMutex.Lock()
	defer this.runMutex.Unlock()
	if this.runStatus {
		GxLogDebug("TcpSession[", this.sid, "] Stop ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		close(this.exitChan)
		close(this.sendChan)
		this.conn.Close()
		this.runStatus = false
	}
}

func (this *Session) Wait() {
	this.waitGroup.Wait()
}

func (this *Session) SendPacket(p *GxPacket) {
	if this.runStatus {
		this.sendChan <- p
		// GxLogDebug("TcpSession[", this.sid, "] SendPacket remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
	}
}

func (this *Session) SetLegal() {
	this.isLegal = true
	GxLogDebug("TcpSession[", this.sid, "] is legal now ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
}

func (this *Session) IsLegal() bool {
	return this.isLegal
}

func (this *Session) startIo() {
	this.local_idle_reset()
	this.remote_heartbeat_active()

	this.waitGroup.Add(3)
	//recv
	go func() {
		defer utils.Stack()
		GxLogDebug("TcpSession[", this.sid, "] recv_thread work ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		var contentLen uint32
		var packet *GxPacket
		for {
			recviced, err := this.conn.Read(this.recvCache)
			if err != nil {
				GxLogDebug("TcpSession[", this.sid, "] read_err :", err.Error(), "! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
				this.Stop()
				break
			}
			this.recvBuffer.Write(this.recvCache[:recviced])
			for {
				if this.recvBuffer.Len() == 0 {
					break
				}
				if contentLen == 0 {
					if this.recvBuffer.Len() < PACKET_HEAD_LEN {
						break
					}
					packet = NewRecvPacket(this.recvBuffer)
					contentLen = packet.Length()
				}
				if contentLen > uint32(this.recvBuffer.Len()) {
					break
				}
				if contentLen != 0 {
					content := make([]byte, contentLen)
					this.recvBuffer.Read(content)
					packet.SetContent(content)
					contentLen = 0
				}
				this.recvCb(this, packet) //回调函数 recvCb
				this.remote_heartbeat_active()
			}
		}
		GxLogDebug("TcpSession[", this.sid, "] recv_thread free ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		this.waitGroup.Done()
	}()
	//send
	go func() {
		defer utils.Stack()
		GxLogDebug("TcpSession[", this.sid, "] send_thread work ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
	Exit:
		for {
			select {
			case <-this.exitChan:
				break Exit
			case p := <-this.sendChan:
				if p != nil {
					//fmt.Println("send [ ", this.sid, " ]: ", p)
					_, err := this.conn.Write(p.ToBytes())
					if err != nil {
						//GxLogDebug("TcpSession[", this.sid, "] send_err :", err.Error(), "! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
						this.Stop()
						break Exit
					}
					this.local_idle_reset()
				}
			}
		}
		GxLogDebug("TcpSession[", this.sid, "] send_thread free ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		this.waitGroup.Done()
	}()
	//heart
	go func() {
		defer utils.Stack()
		GxLogDebug("TcpSession[", this.sid, "] heart_thread work ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		timer := time.NewTicker(time.Second)
	Exit:
		for {
			select {
			case <-timer.C:
				this.checkHeartAndLegal()
			case <-this.exitChan:
				break Exit
			}
		}
		GxLogDebug("TcpSession[", this.sid, "] heart_thread free ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		this.waitGroup.Done()
	}()
	//wait
	go func() {
		defer utils.Stack()
		this.Wait()
		this.loseCb(this)
		GxLogDebug("TcpSession[", this.sid, "] all_thread free ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
	}()
}

func (this *Session) checkHeartAndLegal() {
	if this.remote_heartbeat_check() {
		this.local_idle_check()
	}
}

func (this *Session) remote_heartbeat_active() {
	atomic.StoreInt64(&this.heartbeat_time_remote, time.Now().Unix())
}

func (this *Session) remote_heartbeat_check() bool {
	is_die := (time.Now().Unix() - atomic.LoadInt64(&this.heartbeat_time_remote)) > (this.heartInterval + HEART_FIX_INTERVAL)
	if is_die {
		GxLogError("TcpSession[", this.sid, "] heart death ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
		this.Stop()
	}
	return !is_die
}

func (this *Session) local_idle_reset() {
	atomic.StoreInt64(&this.heartbeat_time_local, time.Now().Unix())
}

func (this *Session) local_idle_check() {
	max_idle := (time.Now().Unix()-atomic.LoadInt64(&this.heartbeat_time_local) >= this.heartInterval)
	if max_idle {
		if !this.isLegal {
			GxLogDebug("TcpSession[", this.sid, "] checkLegal failed and kick out ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
			this.Stop()
			return
		}
		this.heartPump()
	}
}

func (this *Session) heartPump() {
	// GxLogInfo("TcpSession[", this.sid, "] heartPump ! remoteIP: ", this.RemoteAddr(), " localIP: ", this.LocalAddr())
	packet := NewPacket(this.sid, MSGCODE_HEART)
	this.SendPacket(packet)
}
