package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &SkillFlashs)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    SkillFlashmap = make(map[int32]*SkillFlash)
    for _, v := range SkillFlashs {
        SkillFlashmap[v.Id] = v
    }
    LogInfo("Load SkillFlash Scheme Success!")
}
