package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type BuffFlash struct {
	 Id 	 int32
	 BuffType 	 int32
	 Name 	 string
	 SkeletonPath 	 string
	 TexturePath 	 string
	 DragonBones 	 string
	 Skeleton 	 string
	 BuffStart 	 string
	 BuffCycle 	 string
	 BuffEnd 	 string
	 ShadeType 	 int32

}

var BuffFlashs []*BuffFlash
var BuffFlashmap map[int32]*BuffFlash

func LoadBuffFlash(filepath string) {
    fileName := "BuffFlash.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &BuffFlashs)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    BuffFlashmap = make(map[int32]*BuffFlash)
    for _, v := range BuffFlashs {
        BuffFlashmap[v.Id] = v
    }
    LogInfo("Load BuffFlash Scheme Success!")
}
