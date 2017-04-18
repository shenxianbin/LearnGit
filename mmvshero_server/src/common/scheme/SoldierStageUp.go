package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type SoldierStageUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Notes 	 string
	 Name 	 int32
	 Stage 	 int32
	 SkillId 	 string
	 AttackList 	 string
	 ActionFlashId 	 int32
	 ActionSoundId 	 int32
	 Icon 	 string
	 Icon2 	 string

}

var SoldierStageUps []*SoldierStageUp
var SoldierStageUpmap map[int32]*SoldierStageUp

func LoadSoldierStageUp(filepath string) {
    fileName := "SoldierStageUp.json"
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

	SoldierStageUpmap = make(map[int32]*SoldierStageUp)
	for dec.More() {
		var temp SoldierStageUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		SoldierStageUps = append(SoldierStageUps, &temp)
		SoldierStageUpmap[temp.Id] = &temp
	}

	LogInfo("Load SoldierStageUp Scheme Success!")
}
