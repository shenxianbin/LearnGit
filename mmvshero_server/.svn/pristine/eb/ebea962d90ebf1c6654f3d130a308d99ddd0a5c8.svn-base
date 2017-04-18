package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Soldier struct {
	 Id 	 int32
	 Show 	 int32
	 Name 	 string
	 Population 	 int32
	 Corpse 	 int32
	 BattleType 	 int32
	 ShowSkillId 	 string
	 Details 	 int32

}

var Soldiers []*Soldier
var Soldiermap map[int32]*Soldier

func LoadSoldier(filepath string) {
    fileName := "Soldier.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Soldiers)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Soldiermap = make(map[int32]*Soldier)
    for _, v := range Soldiers {
        Soldiermap[v.Id] = v
    }
    LogInfo("Load Soldier Scheme Success!")
}
