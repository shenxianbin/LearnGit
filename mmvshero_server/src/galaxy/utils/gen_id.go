package utils

import (
	"sync"
)

var (
	sessionId int64
	mutexSid  sync.Mutex
)

func AllocSid(serverId int) int64 {
	mutexSid.Lock()
	defer mutexSid.Unlock()
	sessionId++
	return int64(int64(serverId<<32) | sessionId)
}
