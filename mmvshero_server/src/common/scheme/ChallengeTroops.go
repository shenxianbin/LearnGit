package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ChallengeTroops struct {
	 Id 	 int32
	 Type 	 int32
	 WaveIndex 	 int32
	 WaveLv 	 int32
	 WaveHpPlus 	 int32
	 WaveAtkPlus 	 int32
	 NPC 	 string
	 FeatureID 	 string
	 Head 	 string
	 Interval 	 int32
	 PlunderName 	 int32
	 PlunderAward 	 int32

}

var ChallengeTroopss []*ChallengeTroops
var ChallengeTroopsmap map[int32]*ChallengeTroops

func LoadChallengeTroops(filepath string) {
    fileName := "ChallengeTroops.json"
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

	ChallengeTroopsmap = make(map[int32]*ChallengeTroops)
	for dec.More() {
		var temp ChallengeTroops
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ChallengeTroopss = append(ChallengeTroopss, &temp)
		ChallengeTroopsmap[temp.Id] = &temp
	}

	LogInfo("Load ChallengeTroops Scheme Success!")
}
