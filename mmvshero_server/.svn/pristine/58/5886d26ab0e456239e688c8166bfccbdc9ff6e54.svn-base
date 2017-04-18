package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Pvp struct {
	 Id 	 int32
	 Name 	 int32
	 LvUpTrophy 	 int32
	 LvDownTrophy 	 int32
	 RewardStoneWin 	 int32
	 RewardStoneLose 	 int32

}

var Pvps []*Pvp
var Pvpmap map[int32]*Pvp

func LoadPvp(filepath string) {
    fileName := "Pvp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Pvps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Pvpmap = make(map[int32]*Pvp)
    for _, v := range Pvps {
        Pvpmap[v.Id] = v
    }
    LogInfo("Load Pvp Scheme Success!")
}
