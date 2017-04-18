package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Activity struct {
	 Id 	 int32
	 ActType 	 int32
	 ActId 	 int32
	 TimeType 	 int32
	 StartTime 	 string
	 EndTime 	 string
	 RoleLv 	 int32
	 IsOpen 	 int32
	 IsClean 	 int32
	 IsReceiveClose 	 int32
	 ActName 	 int32
	 ActIcon 	 string

}

var Activitys []*Activity
var Activitymap map[int32]*Activity

func LoadActivity(filepath string) {
    fileName := "Activity.json"
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

	Activitymap = make(map[int32]*Activity)
	for dec.More() {
		var temp Activity
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Activitys = append(Activitys, &temp)
		Activitymap[temp.Id] = &temp
	}

	LogInfo("Load Activity Scheme Success!")
}
