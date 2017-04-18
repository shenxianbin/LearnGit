package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type HeroLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 LvUpKingLv 	 int32
	 Stage 	 int32
	 Rank 	 int32
	 NeedExp 	 int32
	 ExpCount 	 int32
	 Hp 	 int32
	 Atk 	 int32
	 Spd 	 int32
	 Arg 	 int32
	 Ats 	 int32
	 View 	 int32
	 SkillEnergy 	 int32
	 Text1 	 int32
	 Text2 	 int32
	 Text3 	 int32

}

var HeroLvUps []*HeroLvUp
var HeroLvUpmap map[int32]*HeroLvUp

func LoadHeroLvUp(filepath string) {
    fileName := "HeroLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &HeroLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    HeroLvUpmap = make(map[int32]*HeroLvUp)
    for _, v := range HeroLvUps {
        HeroLvUpmap[v.Id] = v
    }
    LogInfo("Load HeroLvUp Scheme Success!")
}
