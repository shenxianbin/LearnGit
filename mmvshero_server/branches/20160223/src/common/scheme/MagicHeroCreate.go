package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type MagicHeroCreate struct {
	 Id 	 int32
	 Name 	 string
	 MagicValue 	 int32
	 Drop 	 int32
	 ActionFlashId 	 int32

}

var MagicHeroCreates []*MagicHeroCreate
var MagicHeroCreatemap map[int32]*MagicHeroCreate

func LoadMagicHeroCreate(filepath string) {
    fileName := "MagicHeroCreate.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &MagicHeroCreates)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    MagicHeroCreatemap = make(map[int32]*MagicHeroCreate)
    for _, v := range MagicHeroCreates {
        MagicHeroCreatemap[v.Id] = v
    }
    LogInfo("Load MagicHeroCreate Scheme Success!")
}
