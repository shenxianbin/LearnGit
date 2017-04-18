package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type SkillLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 NextLvId 	 int32
	 NeedSoul 	 int32
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
	 DetailValue 	 string
	 DetailType 	 string
	 TipMessage 	 string
	 TipParam 	 string

}

var SkillLvUps []*SkillLvUp
var SkillLvUpmap map[int32]*SkillLvUp

func LoadSkillLvUp(filepath string) {
    fileName := "SkillLvUp.json"
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

	SkillLvUpmap = make(map[int32]*SkillLvUp)
	for dec.More() {
		var temp SkillLvUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		SkillLvUps = append(SkillLvUps, &temp)
		SkillLvUpmap[temp.Id] = &temp
	}

	LogInfo("Load SkillLvUp Scheme Success!")
}
