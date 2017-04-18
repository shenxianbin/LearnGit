package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Item struct {
	 Id 	 int32
	 Name 	 int32
	 ItemType 	 int32
	 HeapLimit 	 int32
	 Lv 	 int32
	 Useage 	 int32
	 Value 	 string
	 GetWay 	 string
	 Price 	 int32
	 Details 	 int32
	 Icon 	 string
	 Icon2 	 string
	 ActionFlashId 	 int32
	 Notes 	 string

}

var Items []*Item
var Itemmap map[int32]*Item

func LoadItem(filepath string) {
    fileName := "Item.json"
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

	Itemmap = make(map[int32]*Item)
	for dec.More() {
		var temp Item
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Items = append(Items, &temp)
		Itemmap[temp.Id] = &temp
	}

	LogInfo("Load Item Scheme Success!")
}
