package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
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

	Skillmap = make(map[int32]*Skill)
	for dec.More() {
		var temp Skill
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Skills = append(Skills, &temp)
		Skillmap[temp.Id] = &temp
	}

	LogInfo("Load Skill Scheme Success!")
}
