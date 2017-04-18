package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type KingSkill struct {
	 Id 	 int32
	 Show 	 int32
	 Name 	 int32
	 NeedRoleLv 	 int32
	 Icon 	 string
	 Icon2 	 string
	 CdTime 	 int32
	 SkillId 	 int32
	 ActionSoundID 	 int32
	 Details 	 int32
	 Notes 	 string

}

var KingSkills []*KingSkill
var KingSkillmap map[int32]*KingSkill

func LoadKingSkill(filepath string) {
    fileName := "KingSkill.json"
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

	KingSkillmap = make(map[int32]*KingSkill)
	for dec.More() {
		var temp KingSkill
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		KingSkills = append(KingSkills, &temp)
		KingSkillmap[temp.Id] = &temp
	}

	LogInfo("Load KingSkill Scheme Success!")
}
