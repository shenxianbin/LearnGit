package main

import (
	"Gateserver/client"
	"galaxy"
	"galaxy/utils"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	//"runtime/pprof"
	"syscall"
)

func signalHandler() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-sc
	galaxy.GxService().Stop()
	galaxy.LogInfo("Service closed.")
	os.Exit(0)
}

func main() {
	defer utils.Stack()

	galaxy.GxService().Run()
	client.ClientManager()

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	//	f, err := os.OpenFile("./tmp/cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	//	if err != nil {
	//		return
	//	}
	//	defer f.Close()
	//	pprof.StartCPUProfile(f)
	//	defer pprof.StopCPUProfile()

	galaxy.GxService().Wait()
}
