package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type HeroStageUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 int32
	 Stage 	 int32
	 Rank 	 int32
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

var HeroStageUps []*HeroStageUp
var HeroStageUpmap map[int32]*HeroStageUp

func LoadHeroStageUp(filepath string) {
    fileName := "HeroStageUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &HeroStageUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    HeroStageUpmap = make(map[int32]*HeroStageUp)
    for _, v := range HeroStageUps {
        HeroStageUpmap[v.Id] = v
    }
    LogInfo("Load HeroStageUp Scheme Success!")
}
