package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type KingSkillLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 LvUpBlood 	 int32
	 LvUpSoul 	 int32
	 LvUpTime 	 int32
	 LvUpKingLv 	 int32
	 Xeffect 	 int32
	 SkillId 	 int32

}

var KingSkillLvUps []*KingSkillLvUp
var KingSkillLvUpmap map[int32]*KingSkillLvUp

func LoadKingSkillLvUp(filepath string) {
    fileName := "KingSkillLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &KingSkillLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    KingSkillLvUpmap = make(map[int32]*KingSkillLvUp)
    for _, v := range KingSkillLvUps {
        KingSkillLvUpmap[v.Id] = v
    }
    LogInfo("Load KingSkillLvUp Scheme Success!")
}
