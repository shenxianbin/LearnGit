package main

import (
	"Gameserver/logic"
	"Gameserver/logic/achievement"
	"Gameserver/logic/building"
	"Gameserver/logic/chat"
	"Gameserver/logic/drop"
	"Gameserver/logic/fb"
	"Gameserver/logic/fightreport"
	"Gameserver/logic/friend"
	"Gameserver/logic/gamemap"
	"Gameserver/logic/gm"
	"Gameserver/logic/hero"
	"Gameserver/logic/item"
	"Gameserver/logic/king"
	"Gameserver/logic/mall"
	"Gameserver/logic/mission"
	"Gameserver/logic/pvp"
	"Gameserver/logic/role"
	"Gameserver/logic/sign"
	"Gameserver/logic/soldier"
	"Gameserver/logic/stage"
	"common/scheme"
	"fmt"
	"galaxy"
	"galaxy/timer"
	"galaxy/utils"
)

func main() {
	defer utils.Stack()
	galaxy.GxService().Run()
	timer.Start(timer.TIMER_TYPE_FIXED, int64(3600*5))
	timer.Start(timer.TIMER_TYPE_CD)

	serverconfig, err := LoadServerConfig()
	if err != nil {
		return
	}

	scheme.LoadAll(serverconfig.scheme_path)
	scheme.Process()

	logic.Init()
	role.InitRoleModule()
	soldier.InitSoldierModule()
	hero.InitHeroModule()
	item.InitItemModule()
	building.InitBuildingModule()
	king.InitKingModule()

	stage.InitStageModule()
	mission.InitMissionModule()
	achievement.InitAchievementModule()
	fb.InitFbModule()
	friend.InitFriendModule()

	gamemap.InitMapModule()
	drop.InitDropModule()
	err = pvp.InitPvpModule()
	if err != nil {
		fmt.Println(err)
		return
	}
	sign.InitSignModule()
	mall.InitMallModule()
	chat.InitChatModule()
	fightreport.InitFightReportModule()
	gm.InitGmModule()

	galaxy.GxService().Wait()
	timer.Wait()
}
