package main

import (
	"Staticserver/ip"
	"Staticserver/mysql"
	"Staticserver/static"
	. "galaxy"
	"galaxy/utils"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	defer utils.Stack()
	GxService().Run()

	serverconfig, err := LoadServerConfig()
	if err != nil {
		LogError(err)
		return
	}

	err = mysql.Init(serverconfig.mysql, int32(serverconfig.mysqlChanSize))
	if err != nil {
		LogError(err)
		return
	}

	err = ip.Init(serverconfig.ipData)
	if err != nil {
		LogError(err)
		return
	}

	static.Init()

	go func() {
		err = http.ListenAndServe(":8090", nil)
		if err != nil {
			LogFatal(err)
			return
		}
	}()

	GxService().Wait()
	ip.Close()
}
