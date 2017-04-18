package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ActivityGrowFund struct {
	 Id 	 int32
	 Name 	 int32
	 Icon 	 string
	 Detial 	 int32
	 RuleDetial 	 int32
	 Buy 	 []*struct {
		Cost int32
		Index int32
		}
	 ItemContent 	 string
	 Award 	 []*struct {
		Award int32
		Condition int32
		Index int32
		}

}

var ActivityGrowFunds []*ActivityGrowFund
var ActivityGrowFundmap map[int32]*ActivityGrowFund

func LoadActivityGrowFund(filepath string) {
    fileName := "ActivityGrowFund.json"
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

	ActivityGrowFundmap = make(map[int32]*ActivityGrowFund)
	for dec.More() {
		var temp ActivityGrowFund
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ActivityGrowFunds = append(ActivityGrowFunds, &temp)
		ActivityGrowFundmap[temp.Id] = &temp
	}

	LogInfo("Load ActivityGrowFund Scheme Success!")
}
