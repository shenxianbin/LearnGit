package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type RechargeLvUp struct {
	 Id 	 int32
	 NeedExp 	 int32
	 AwardId 	 int32
	 Icon 	 string
	 Notes 	 string

}

var RechargeLvUps []*RechargeLvUp
var RechargeLvUpmap map[int32]*RechargeLvUp

func LoadRechargeLvUp(filepath string) {
    fileName := "RechargeLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &RechargeLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    RechargeLvUpmap = make(map[int32]*RechargeLvUp)
    for _, v := range RechargeLvUps {
        RechargeLvUpmap[v.Id] = v
    }
    LogInfo("Load RechargeLvUp Scheme Success!")
}
