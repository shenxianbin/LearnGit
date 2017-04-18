package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type SkillLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 NextLvId 	 int32
	 TriggerType 	 int32
	 TriggerParam 	 string
	 SkillRange 	 int32
	 TargetType 	 string
	 TargetNum 	 string
	 AttackType 	 string
	 AttackParam 	 string
	 BuffType 	 string
	 BuffParam 	 string
	 BuffTime 	 string
	 BuffRange 	 string
	 BuffFlashId 	 string
	 FortressEffect 	 int32
	 Dialogue 	 int32
	 NeedItemId 	 int32
	 NeedItemNum 	 int32
	 DetailsParam 	 string
	 TipMessage 	 string
	 TipParam 	 string

}

var SkillLvUps []*SkillLvUp
var SkillLvUpmap map[int32]*SkillLvUp

func LoadSkillLvUp(filepath string) {
    fileName := "SkillLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &SkillLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    SkillLvUpmap = make(map[int32]*SkillLvUp)
    for _, v := range SkillLvUps {
        SkillLvUpmap[v.Id] = v
    }
    LogInfo("Load SkillLvUp Scheme Success!")
}
