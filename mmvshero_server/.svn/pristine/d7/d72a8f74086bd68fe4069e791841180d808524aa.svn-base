package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type ActionSound struct {
	 Id 	 int32
	 Name 	 string
	 SelectedSound 	 string
	 Attack0Sound 	 string
	 Attack0HitSound 	 string
	 Attack0DurationSound 	 string
	 Attack1Sound 	 string
	 Attack1HitSound 	 string
	 Attack1DurationSound 	 string
	 Attack2Sound 	 string
	 Attack2HitSound 	 string
	 Attack2DurationSound 	 string
	 Attack3Sound 	 string
	 Attack3HitSound 	 string
	 Attack3DurationSound 	 string
	 Attack4Sound 	 string
	 Attack4HitSound 	 string
	 Attack4DurationSound 	 string
	 BeAttackedSound 	 string
	 DieSound 	 string

}

var ActionSounds []*ActionSound
var ActionSoundmap map[int32]*ActionSound

func LoadActionSound(filepath string) {
    fileName := "ActionSound.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &ActionSounds)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    ActionSoundmap = make(map[int32]*ActionSound)
    for _, v := range ActionSounds {
        ActionSoundmap[v.Id] = v
    }
    LogInfo("Load ActionSound Scheme Success!")
}
