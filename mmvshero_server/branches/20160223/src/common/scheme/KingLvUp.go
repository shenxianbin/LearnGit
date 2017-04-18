package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type KingLvUp struct {
	 Id 	 int32
	 NeedBlood 	 int32
	 NeedSoul 	 int32
	 MapSize 	 string
	 MapStage 	 int32
	 PopLimit 	 int32
	 DigLimit 	 int32
	 EvoSpeedLimit 	 int32
	 PvpLimit 	 int32
	 PvpInterval 	 int32
	 PvpSearchNeedBlood 	 int32
	 MagicHeroLimit 	 int32
	 FortressLimit 	 int32
	 SoldierLimit 	 string
	 SoldierUnlock 	 string
	 SoldierObtain 	 string
	 MagicKingSkillId 	 string

}

var KingLvUps []*KingLvUp
var KingLvUpmap map[int32]*KingLvUp

func LoadKingLvUp(filepath string) {
    fileName := "KingLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &KingLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    KingLvUpmap = make(map[int32]*KingLvUp)
    for _, v := range KingLvUps {
        KingLvUpmap[v.Id] = v
    }
    LogInfo("Load KingLvUp Scheme Success!")
}
