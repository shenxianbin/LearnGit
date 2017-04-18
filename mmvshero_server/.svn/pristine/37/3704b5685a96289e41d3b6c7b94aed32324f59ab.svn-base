package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Achievement struct {
	 Id 	 int32
	 DrawOrder 	 int32
	 TargetNum 	 string
	 Name 	 string
	 Details 	 string
	 AwardId 	 string
	 Notes 	 string

}

var Achievements []*Achievement
var Achievementmap map[int32]*Achievement

func LoadAchievement(filepath string) {
    fileName := "Achievement.json"
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

	Achievementmap = make(map[int32]*Achievement)
	for dec.More() {
		var temp Achievement
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Achievements = append(Achievements, &temp)
		Achievementmap[temp.Id] = &temp
	}

	LogInfo("Load Achievement Scheme Success!")
}
