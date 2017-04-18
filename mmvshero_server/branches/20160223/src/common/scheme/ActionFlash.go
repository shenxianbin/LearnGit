package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type ActionFlash struct {
	 Id 	 int32
	 Name 	 string
	 SkeletonPath 	 string
	 TexturePath 	 string
	 DragonBones 	 string
	 Skeleton 	 string
	 Hit 	 string
	 Die 	 string
	 Walk 	 string
	 Grap 	 string
	 Standby 	 string
	 Standby2 	 string
	 HeadPosition 	 string
	 ChestPosition 	 string

}

var ActionFlashs []*ActionFlash
var ActionFlashmap map[int32]*ActionFlash

func LoadActionFlash(filepath string) {
    fileName := "ActionFlash.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &ActionFlashs)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    ActionFlashmap = make(map[int32]*ActionFlash)
    for _, v := range ActionFlashs {
        ActionFlashmap[v.Id] = v
    }
    LogInfo("Load ActionFlash Scheme Success!")
}
