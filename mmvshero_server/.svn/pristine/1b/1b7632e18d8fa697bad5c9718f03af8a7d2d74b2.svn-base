package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type BuildingLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 Hp 	 int32
	 Arg 	 int32
	 Ats 	 int32
	 BloodMax 	 int32
	 BloodOutput 	 int32
	 BloodOutputInterval 	 int32
	 BloodOutputMax 	 int32
	 SoulMax 	 int32
	 SoulOutput 	 int32
	 SoulOutputInterval 	 int32
	 SoulOutputMax 	 int32
	 LvUpBlood 	 int32
	 LvUpSoul 	 int32
	 LvUpTime 	 int32
	 LvUpKingLv 	 int32
	 Xeffect 	 int32
	 SkillId 	 string
	 AttackList 	 string
	 Text1 	 int32
	 Text2 	 int32
	 Text3 	 int32

}

var BuildingLvUps []*BuildingLvUp
var BuildingLvUpmap map[int32]*BuildingLvUp

func LoadBuildingLvUp(filepath string) {
    fileName := "BuildingLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &BuildingLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    BuildingLvUpmap = make(map[int32]*BuildingLvUp)
    for _, v := range BuildingLvUps {
        BuildingLvUpmap[v.Id] = v
    }
    LogInfo("Load BuildingLvUp Scheme Success!")
}
