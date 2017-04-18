package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Hero struct {
	 Id 	 int32
	 Show 	 int32
	 Name 	 string
	 Population 	 int32
	 Corpse 	 int32
	 BattleType 	 int32
	 MagicHeroRankId 	 int32
	 ShowSkillId 	 string
	 Details 	 int32

}

var Heros []*Hero
var Heromap map[int32]*Hero

func LoadHero(filepath string) {
    fileName := "Hero.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Heros)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Heromap = make(map[int32]*Hero)
    for _, v := range Heros {
        Heromap[v.Id] = v
    }
    LogInfo("Load Hero Scheme Success!")
}
