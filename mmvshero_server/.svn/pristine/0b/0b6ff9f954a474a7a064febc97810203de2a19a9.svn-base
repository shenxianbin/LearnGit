package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Achievement struct {
	 Id 	 int32
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Achievements)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Achievementmap = make(map[int32]*Achievement)
    for _, v := range Achievements {
        Achievementmap[v.Id] = v
    }
    LogInfo("Load Achievement Scheme Success!")
}
