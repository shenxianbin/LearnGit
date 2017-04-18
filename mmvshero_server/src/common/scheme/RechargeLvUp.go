package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type RechargeLvUp struct {
	 Id 	 int32
	 NeedExp 	 int32
	 AwardId 	 int32
	 Icon 	 string
	 Notes 	 string

}

var RechargeLvUps []*RechargeLvUp
var RechargeLvUpmap map[int32]*RechargeLvUp

func LoadRechargeLvUp(filepath string) {
    fileName := "RechargeLvUp.json"
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

	RechargeLvUpmap = make(map[int32]*RechargeLvUp)
	for dec.More() {
		var temp RechargeLvUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		RechargeLvUps = append(RechargeLvUps, &temp)
		RechargeLvUpmap[temp.Id] = &temp
	}

	LogInfo("Load RechargeLvUp Scheme Success!")
}
