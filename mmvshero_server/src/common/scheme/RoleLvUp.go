package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type RoleLvUp struct {
	 Id 	 int32
	 NeedExp 	 int32
	 TotalExp 	 int32
	 OrderLimit 	 int32
	 OrderAdd 	 int32
	 MapSize 	 string
	 MapStage 	 int32
	 HeroLimit 	 int32
	 DigLimit 	 int32
	 ExDigLimit 	 int32
	 FortressLimit 	 int32
	 SoldierUnlock 	 string
	 SoldierObtain 	 string
	 HeroUnlock 	 string
	 KingSkillId 	 string
	 PlunderSearchNeedSoul 	 int32
	 PlunderTeam 	 string
	 PlunderDailyTimes 	 int32
	 HeroCreateCost 	 int32
	 KingHp 	 int32
	 SoulHour 	 int32
	 SoldierChipHour 	 int32
	 HeroChipHour 	 int32
	 PlunderHeroHpPlus 	 int32
	 PlunderHeroAtkPlus 	 int32

}

var RoleLvUps []*RoleLvUp
var RoleLvUpmap map[int32]*RoleLvUp

func LoadRoleLvUp(filepath string) {
    fileName := "RoleLvUp.json"
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

	RoleLvUpmap = make(map[int32]*RoleLvUp)
	for dec.More() {
		var temp RoleLvUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		RoleLvUps = append(RoleLvUps, &temp)
		RoleLvUpmap[temp.Id] = &temp
	}

	LogInfo("Load RoleLvUp Scheme Success!")
}
