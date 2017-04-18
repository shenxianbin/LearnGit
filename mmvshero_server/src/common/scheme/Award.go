package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Award struct {
	 Id 	 int32
	 Reward 	 []*struct {
		Amount int32
		Code int32
		Type int32
		}
	 RandomReward 	 []*struct {
		Data []*struct {
				Amount int32
				Code int32
				Type int32
				Weight int32
				}
		Rate int32
		}
	 SelfRandomReward 	 []*struct {
		Data []*struct {
				Amount int32
				Ban []int32
				Selftype int32
				Weight int32
				}
		Rate int32
		}
	 Notes 	 string

}

var Awards []*Award
var Awardmap map[int32]*Award

func LoadAward(filepath string) {
    fileName := "Award.json"
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

	Awardmap = make(map[int32]*Award)
	for dec.More() {
		var temp Award
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Awards = append(Awards, &temp)
		Awardmap[temp.Id] = &temp
	}

	LogInfo("Load Award Scheme Success!")
}
