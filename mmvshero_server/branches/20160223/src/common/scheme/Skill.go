package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Skill struct {
	 Id 	 int32
	 Name 	 int32
	 LvMax 	 int32
	 Icon 	 string
	 Details 	 int32
	 Notes 	 string

}

var Skills []*Skill
var Skillmap map[int32]*Skill

func LoadSkill(filepath string) {
    fileName := "Skill.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Skills)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Skillmap = make(map[int32]*Skill)
    for _, v := range Skills {
        Skillmap[v.Id] = v
    }
    LogInfo("Load Skill Scheme Success!")
}
