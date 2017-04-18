package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Missions)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Missionmap = make(map[int32]*Mission)
    for _, v := range Missions {
        Missionmap[v.Id] = v
    }
    LogInfo("Load Mission Scheme Success!")
}
