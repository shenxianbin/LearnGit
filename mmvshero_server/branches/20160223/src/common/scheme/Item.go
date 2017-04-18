package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Item struct {
	 Id 	 int32
	 Name 	 int32
	 BagShow 	 int32
	 ItemType 	 int32
	 HeapLimit 	 int32
	 Lv 	 int32
	 GetWay 	 string
	 Useage 	 int32
	 Value 	 string
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Items)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Itemmap = make(map[int32]*Item)
    for _, v := range Items {
        Itemmap[v.Id] = v
    }
    LogInfo("Load Item Scheme Success!")
}
