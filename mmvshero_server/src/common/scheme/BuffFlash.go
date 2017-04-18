package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
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
	file, err := os.Open(fmt.Sprintf("%s/%s", filepath, fileName))
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	_, err = dec.Token()
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}

	BuffFlashmap = make(map[int32]*BuffFlash)
	for dec.More() {
		var temp BuffFlash
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		BuffFlashs = append(BuffFlashs, &temp)
		BuffFlashmap[temp.Id] = &temp
	}

	LogInfo("Load BuffFlash Scheme Success!")
}
