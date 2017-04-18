package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type RoleLvUp struct {
	 Id 	 int32
	 NeedExp 	 int32
	 OrderLimit 	 int32
	 OrderAdd 	 int32
	 SpiritLimit 	 int32
	 SpiritAdd 	 int32

}

var RoleLvUps []*RoleLvUp
var RoleLvUpmap map[int32]*RoleLvUp

func LoadRoleLvUp(filepath string) {
    fileName := "RoleLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &RoleLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    RoleLvUpmap = make(map[int32]*RoleLvUp)
    for _, v := range RoleLvUps {
        RoleLvUpmap[v.Id] = v
    }
    LogInfo("Load RoleLvUp Scheme Success!")
}
