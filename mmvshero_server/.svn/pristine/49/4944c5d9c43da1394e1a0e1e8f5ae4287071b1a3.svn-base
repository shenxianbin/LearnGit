package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Drop struct {
	 Id 	 int32
	 DrawType 	 int32
	 Plan 	 int32
	 Name 	 string
	 NeedGold 	 int32
	 NeedItem 	 int32
	 NeedItemNum 	 int32
	 RandomAwardId 	 string

}

var Drops []*Drop
var Dropmap map[int32]*Drop

func LoadDrop(filepath string) {
    fileName := "Drop.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Drops)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Dropmap = make(map[int32]*Drop)
    for _, v := range Drops {
        Dropmap[v.Id] = v
    }
    LogInfo("Load Drop Scheme Success!")
}
