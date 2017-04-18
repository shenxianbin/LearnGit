package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Mission struct {
	 Id 	 int32
	 MissionType 	 int32
	 Once 	 int32
	 NeedRoleLv 	 int32
	 OnlyVip 	 int32
	 TargetNum 	 int32
	 TargetLv 	 int32
	 LvParam 	 int32
	 Name 	 int32
	 Details 	 int32
	 Goto 	 int32
	 AwardId 	 string
	 Icon 	 string
	 Notes 	 string

}

var Missions []*Mission
var Missionmap map[int32]*Mission

func LoadMission(filepath string) {
    fileName := "Mission.json"
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

	Missionmap = make(map[int32]*Mission)
	for dec.More() {
		var temp Mission
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Missions = append(Missions, &temp)
		Missionmap[temp.Id] = &temp
	}

	LogInfo("Load Mission Scheme Success!")
}
