package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Stage struct {
	 Id 	 int32
	 StageType 	 int32
	 StageNo 	 int32
	 NextStageId 	 int32
	 Chapter 	 int32
	 ChapterOpen 	 int32
	 Name 	 int32
	 Notes 	 string
	 DailyTimes 	 int32
	 LeastCostOrder 	 int32
	 VictoryCostOrder 	 int32
	 LeastRoleExp 	 int32
	 VictoryRoleExp 	 int32
	 ShowHead 	 string
	 BattleForce 	 []*struct {
		Npc string
		Time int32
		Wave int32
		}
	 FixedAwardId 	 int32
	 ItemAwardId 	 int32
	 HeroAwardId 	 int32
	 SoldierAwardId 	 int32
	 StageMissionId 	 string
	 Mission1Param 	 string
	 Mission2Param 	 string
	 Mission3Param 	 string
	 SweepNeedItemID 	 int32
	 SweepNeedItemNum 	 int32
	 SweepExtraBonus 	 int32
	 CombatPower 	 int32
	 IsBoss 	 int32
	 PlunderFeatures 	 string
	 StagePosition 	 string
	 LevelLimit 	 int32
	 StoryGroup 	 string

}

var Stages []*Stage
var Stagemap map[int32]*Stage

func LoadStage(filepath string) {
    fileName := "Stage.json"
	file, err := os.Open(fmt.Sprintf("%s/%s", filepath, fileName))
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	_, err = dec.Token()
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}

	Stagemap = make(map[int32]*Stage)
	for dec.More() {
		var temp Stage
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Stages = append(Stages, &temp)
		Stagemap[temp.Id] = &temp
	}

	LogInfo("Load Stage Scheme Success!")
}
