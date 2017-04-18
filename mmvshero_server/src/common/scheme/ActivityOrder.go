package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ActivityOrder struct {
	Id          int32
	Name        int32
	Icon        string
	Detial      int32
	RuleDetial  int32
	ItemName    string
	ItemContent string
	Award       []*struct {
		Award     int32
		Condition []int32
		Index     int32
	}
	BattleShareAward []*struct {
		Award     int32
		Condition []int32
		Index     int32
	}
}

var ActivityOrders []*ActivityOrder
var ActivityOrdermap map[int32]*ActivityOrder

func LoadActivityOrder(filepath string) {
	fileName := "ActivityOrder.json"
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

	ActivityOrdermap = make(map[int32]*ActivityOrder)
	for dec.More() {
		var temp ActivityOrder
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ActivityOrders = append(ActivityOrders, &temp)
		ActivityOrdermap[temp.Id] = &temp
	}

	LogInfo("Load ActivityOrder Scheme Success!")
}
