package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Common struct {
	 Id 	 int32
	 Parameter 	 string
	 Value 	 int32
	 Notes 	 string

}

var Commons []*Common
var Commonmap map[int32]*Common

func LoadCommon(filepath string) {
    fileName := "Common.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Commons)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Commonmap = make(map[int32]*Common)
    for _, v := range Commons {
        Commonmap[v.Id] = v
    }
    LogInfo("Load Common Scheme Success!")
}
