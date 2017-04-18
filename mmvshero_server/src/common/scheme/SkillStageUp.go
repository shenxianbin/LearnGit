package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type SkillStageUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Stage 	 int32
	 SkillFlashId 	 int32
	 SpecialFlashId 	 int32
	 SpecialFlashTime 	 int32
	 SkillSoundId 	 int32

}

var SkillStageUps []*SkillStageUp
var SkillStageUpmap map[int32]*SkillStageUp

func LoadSkillStageUp(filepath string) {
    fileName := "SkillStageUp.json"
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

	SkillStageUpmap = make(map[int32]*SkillStageUp)
	for dec.More() {
		var temp SkillStageUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		SkillStageUps = append(SkillStageUps, &temp)
		SkillStageUpmap[temp.Id] = &temp
	}

	LogInfo("Load SkillStageUp Scheme Success!")
}
