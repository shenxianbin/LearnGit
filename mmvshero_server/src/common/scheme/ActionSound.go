package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ActionSound struct {
	 Id 	 int32
	 Name 	 string
	 Selected 	 string
	 Die 	 string

}

var ActionSounds []*ActionSound
var ActionSoundmap map[int32]*ActionSound

func LoadActionSound(filepath string) {
    fileName := "ActionSound.json"
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

	ActionSoundmap = make(map[int32]*ActionSound)
	for dec.More() {
		var temp ActionSound
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ActionSounds = append(ActionSounds, &temp)
		ActionSoundmap[temp.Id] = &temp
	}

	LogInfo("Load ActionSound Scheme Success!")
}
