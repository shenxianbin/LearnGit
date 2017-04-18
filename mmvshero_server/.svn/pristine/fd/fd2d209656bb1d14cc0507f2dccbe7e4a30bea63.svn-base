package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"io/ioutil"
)

type Award struct {
	Id           int32
	Reward       []*Reward
	RandomReward []*RandomReward
	Notes        string
}

type Reward struct {
	Type   int32
	Code   int32
	Amount int32
}

type RandomReward struct {
	Rate int32
	Data []*RandomRewardData
}

type RandomRewardData struct {
	Type   int32
	Code   int32
	Amount int32
	Weight int32
}

var Awards []*Award
var Awardmap map[int32]*Award

func LoadAward(filepath string) {
	fileName := "Award.json"
	file := fmt.Sprintf("%s/%s", filepath, fileName)
	buff, err := ioutil.ReadFile(file)
	err = json.Unmarshal(buff, &Awards)
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}
	Awardmap = make(map[int32]*Award)
	for _, v := range Awards {
		Awardmap[v.Id] = v
	}
	LogInfo("Load Award Scheme Success!")
}
