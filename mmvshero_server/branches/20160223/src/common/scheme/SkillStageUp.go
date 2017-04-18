package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type SkillStageUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Stage 	 int32
	 SkillFlashId 	 int32
	 SpecialFlashId 	 int32
	 SpecialFlashTime 	 int32

}

var SkillStageUps []*SkillStageUp
var SkillStageUpmap map[int32]*SkillStageUp

func LoadSkillStageUp(filepath string) {
    fileName := "SkillStageUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &SkillStageUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    SkillStageUpmap = make(map[int32]*SkillStageUp)
    for _, v := range SkillStageUps {
        SkillStageUpmap[v.Id] = v
    }
    LogInfo("Load SkillStageUp Scheme Success!")
}
