package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"io/ioutil"
)

type Stage struct {
	Id               int32
	Chapter          int32
	StageType        int32
	NextStageId      int32
	Name             int32
	Notes            string
	KeyIcon          string
	NeedItemId       int32
	LeastCostOrder   int32
	VictoryCostOrder int32
	LeastRoleExp     int32
	VictoryRoleExp   int32
	BattleForce      []*BattleForce
	FixedAwardId     int
	ItemAwardId      int
	HeroAwardId      int
	SoldierAwardId   int
	StageMissionId   string
	Mission1Param    string
	Mission2Param    string
	Mission3Param    string
	SweepNeedItemID  int32
	SweepNeedItemNum int32
	SweepExtraBonus  int
}

type BattleForce struct {
	Wave int32
	Npc  string
	Time int32
}

var Stages []*Stage
var Stagemap map[int32]*Stage

func LoadStage(filepath string) {
	fileName := "Stage.json"
	file := fmt.Sprintf("%s/%s", filepath, fileName)
	buff, err := ioutil.ReadFile(file)
	err = json.Unmarshal(buff, &Stages)
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}
	Stagemap = make(map[int32]*Stage)
	for _, v := range Stages {
		Stagemap[v.Id] = v
	}
	LogInfo("Load Stage Scheme Success!")
}
