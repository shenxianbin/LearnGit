package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type BuildingLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 LvUpRoleLv 	 int32
	 LvUpSoul 	 int32
	 Hp 	 int32
	 Arg 	 int32
	 Ats 	 int32
	 SkillId 	 string
	 AttackList 	 string

}

var BuildingLvUps []*BuildingLvUp
var BuildingLvUpmap map[int32]*BuildingLvUp

func LoadBuildingLvUp(filepath string) {
    fileName := "BuildingLvUp.json"
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

	BuildingLvUpmap = make(map[int32]*BuildingLvUp)
	for dec.More() {
		var temp BuildingLvUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		BuildingLvUps = append(BuildingLvUps, &temp)
		BuildingLvUpmap[temp.Id] = &temp
	}

	LogInfo("Load BuildingLvUp Scheme Success!")
}
