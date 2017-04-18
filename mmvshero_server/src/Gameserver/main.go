package main

import (
	"Gameserver/logic"
	"Gameserver/logic/achievement"
	"Gameserver/logic/activity"
	"Gameserver/logic/arena"
	"Gameserver/logic/building"
	"Gameserver/logic/challenge"
	"Gameserver/logic/chat"
	"Gameserver/logic/drop"
	"Gameserver/logic/fightreport"
	"Gameserver/logic/friend"
	"Gameserver/logic/gamemap"
	"Gameserver/logic/gm"
	"Gameserver/logic/hero"
	"Gameserver/logic/item"
	"Gameserver/logic/mall"
	"Gameserver/logic/mission"
	"Gameserver/logic/plunder"
	"Gameserver/logic/recharge"
	"Gameserver/logic/role"
	"Gameserver/logic/sign"
	"Gameserver/logic/soldier"
	"Gameserver/logic/stage"
	"common/scheme"
	"fmt"
	"galaxy"
	"galaxy/timer"
	"galaxy/utils"
	"os"
	"os/signal"
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
	galaxy.GxService().Redis().Close()
	galaxy.LogInfo("Service closed.")
	os.Exit(0)
}

func main() {
	defer utils.Stack()
	galaxy.GxService().Run()
	timer.Start(timer.TIMER_TYPE_FIXED, int64(3600*5))
	timer.Start(timer.TIMER_TYPE_FIXED, int64(3600*21))
	timer.Start(timer.TIMER_TYPE_CD)
	go signalHandler()

	serverconfig, err := LoadServerConfig()
	if err != nil {
		return
	}

	scheme.LoadAll(serverconfig.SchemePath)
	scheme.Process()

	logic.Init()
	role.InitRoleModule()
	soldier.InitSoldierModule()
	hero.InitHeroModule()
	item.InitItemModule()
	building.InitBuildingModule()

	stage.InitStageModule()
	mission.InitMissionModule()
	achievement.InitAchievementModule()
	friend.InitFriendModule()

	gamemap.InitMapModule()
	drop.InitDropModule()

	sign.InitSignModule()
	mall.InitMallModule()
	chat.InitChatModule()
	fightreport.InitFightReportModule()
	plunder.InitPlunderModule()
	challenge.InitChallengeModule()
	err = arena.InitArenaModule()
	if err != nil {
		fmt.Println("InitArenaModule : ", err)
		return
	}
	err = activity.InitActivityModule(serverconfig.OpenServerTime)
	if err != nil {
		fmt.Println("InitActivityModule : ", err)
		return
	}
	gm.InitGmModule()
	recharge.InitRechargeModule(serverconfig.RechargePort)

	galaxy.GxService().Wait()
	timer.Wait()
}
