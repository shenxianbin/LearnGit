package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Building struct {
	 Id 	 int32
	 Show 	 int32
	 Name 	 int32
	 BattleType 	 int32
	 NeedKingLv 	 int32
	 ActionFlashId 	 int32
	 ActionSoundID 	 int32
	 Icon 	 string
	 Icon2 	 string
	 ShowSkillId 	 int32
	 Details 	 int32
	 Notes 	 string

}

var Buildings []*Building
var Buildingmap map[int32]*Building

func LoadBuilding(filepath string) {
    fileName := "Building.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Buildings)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Buildingmap = make(map[int32]*Building)
    for _, v := range Buildings {
        Buildingmap[v.Id] = v
    }
    LogInfo("Load Building Scheme Success!")
}
