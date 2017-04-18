package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type KingSkill struct {
	 Id 	 int32
	 Show 	 int32
	 Name 	 int32
	 NeedKingLv 	 int32
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &KingSkills)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    KingSkillmap = make(map[int32]*KingSkill)
    for _, v := range KingSkills {
        KingSkillmap[v.Id] = v
    }
    LogInfo("Load KingSkill Scheme Success!")
}
