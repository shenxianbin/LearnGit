package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type SoldierStageUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 test 	 string
	 Name 	 int32
	 Stage 	 int32
	 LvLimit 	 int32
	 NextStageId 	 int32
	 SkillId 	 string
	 AttackList 	 string
	 ActionFlashId 	 int32
	 ActionSoundId 	 int32
	 EvoNeedTime 	 int32
	 EvoNeedItemId 	 string
	 EvoNeedItemNum 	 string
	 EvoNeedMagicHeroId 	 int32
	 EvoNeedMagicHeroStage 	 int32
	 EvoNeedMagicHeroRank 	 int32
	 EvoNeedMagicHeroLv 	 int32
	 EvoNeedMagicHeroNum 	 int32
	 Icon 	 string
	 Icon2 	 string

}

var SoldierStageUps []*SoldierStageUp
var SoldierStageUpmap map[int32]*SoldierStageUp

func LoadSoldierStageUp(filepath string) {
    fileName := "SoldierStageUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &SoldierStageUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    SoldierStageUpmap = make(map[int32]*SoldierStageUp)
    for _, v := range SoldierStageUps {
        SoldierStageUpmap[v.Id] = v
    }
    LogInfo("Load SoldierStageUp Scheme Success!")
}
