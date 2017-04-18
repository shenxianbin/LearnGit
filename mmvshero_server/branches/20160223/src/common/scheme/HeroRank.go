package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type HeroRank struct {
	 Id 	 int32
	 Plan 	 int32
	 Rank 	 int32
	 NeedRankPoint 	 int32
	 BaseRankPoint 	 int32

}

var HeroRanks []*HeroRank
var HeroRankmap map[int32]*HeroRank

func LoadHeroRank(filepath string) {
    fileName := "HeroRank.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &HeroRanks)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    HeroRankmap = make(map[int32]*HeroRank)
    for _, v := range HeroRanks {
        HeroRankmap[v.Id] = v
    }
    LogInfo("Load HeroRank Scheme Success!")
}
