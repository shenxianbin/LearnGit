package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type KingSkillCost struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 UseNum 	 int32
	 UseEnergy 	 int32

}

var KingSkillCosts []*KingSkillCost
var KingSkillCostmap map[int32]*KingSkillCost

func LoadKingSkillCost(filepath string) {
    fileName := "KingSkillCost.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &KingSkillCosts)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    KingSkillCostmap = make(map[int32]*KingSkillCost)
    for _, v := range KingSkillCosts {
        KingSkillCostmap[v.Id] = v
    }
    LogInfo("Load KingSkillCost Scheme Success!")
}
