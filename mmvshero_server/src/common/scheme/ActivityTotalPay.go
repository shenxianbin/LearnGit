package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ActivityTotalPay struct {
	 Id 	 int32
	 Name 	 int32
	 Icon 	 string
	 Detial 	 int32
	 RuleDetial 	 int32
	 ItemContent 	 string
	 Award 	 []*struct {
		Award int32
		Condition int32
		Index int32
		}

}

var ActivityTotalPays []*ActivityTotalPay
var ActivityTotalPaymap map[int32]*ActivityTotalPay

func LoadActivityTotalPay(filepath string) {
    fileName := "ActivityTotalPay.json"
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

	ActivityTotalPaymap = make(map[int32]*ActivityTotalPay)
	for dec.More() {
		var temp ActivityTotalPay
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ActivityTotalPays = append(ActivityTotalPays, &temp)
		ActivityTotalPaymap[temp.Id] = &temp
	}

	LogInfo("Load ActivityTotalPay Scheme Success!")
}
