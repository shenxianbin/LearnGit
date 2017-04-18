package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type SkillFlash struct {
	 Id 	 int32
	 Name 	 string
	 SkeletonPath 	 string
	 TexturePath 	 string
	 DragonBones 	 string
	 Skeleton 	 string
	 ActionFlashName 	 string
	 FlySpd 	 int32
	 FlyOffSet 	 string
	 SelfEffect 	 int32
	 SelfEffectStart 	 int32
	 BulletEffect 	 int32
	 BulletEffectStart 	 int32
	 HitEffect 	 int32
	 HitEffectStart 	 int32

}

var SkillFlashs []*SkillFlash
var SkillFlashmap map[int32]*SkillFlash

func LoadSkillFlash(filepath string) {
    fileName := "SkillFlash.json"
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

	SkillFlashmap = make(map[int32]*SkillFlash)
	for dec.More() {
		var temp SkillFlash
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		SkillFlashs = append(SkillFlashs, &temp)
		SkillFlashmap[temp.Id] = &temp
	}

	LogInfo("Load SkillFlash Scheme Success!")
}
