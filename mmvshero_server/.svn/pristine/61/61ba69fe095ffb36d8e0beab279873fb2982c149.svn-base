package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Mall struct {
	 Id 	 int32
	 MallType 	 int32
	 Name 	 int32
	 Details 	 int32
	 Explain 	 int32
	 Args1 	 int32
	 ResourceType 	 int32
	 Price 	 int32
	 AwardId 	 int32
	 OnSale 	 int32
	 OriginalPrice 	 int32
	 Icon 	 string
	 Notes 	 string

}

var Malls []*Mall
var Mallmap map[int32]*Mall

func LoadMall(filepath string) {
    fileName := "Mall.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Malls)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Mallmap = make(map[int32]*Mall)
    for _, v := range Malls {
        Mallmap[v.Id] = v
    }
    LogInfo("Load Mall Scheme Success!")
}
