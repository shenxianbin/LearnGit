package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ActivityLevelBag struct {
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

var ActivityLevelBags []*ActivityLevelBag
var ActivityLevelBagmap map[int32]*ActivityLevelBag

func LoadActivityLevelBag(filepath string) {
    fileName := "ActivityLevelBag.json"
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

	ActivityLevelBagmap = make(map[int32]*ActivityLevelBag)
	for dec.More() {
		var temp ActivityLevelBag
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ActivityLevelBags = append(ActivityLevelBags, &temp)
		ActivityLevelBagmap[temp.Id] = &temp
	}

	LogInfo("Load ActivityLevelBag Scheme Success!")
}
