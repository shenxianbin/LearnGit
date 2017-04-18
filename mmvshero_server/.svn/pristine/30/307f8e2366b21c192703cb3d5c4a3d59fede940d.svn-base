package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type HeroRank struct {
	 Id 	 int32
	 Plan 	 int32
	 Rank 	 int32
	 NeedRankPoint 	 int32
	 BaseRankPoint 	 int32
	 RankParm 	 int32

}

var HeroRanks []*HeroRank
var HeroRankmap map[int32]*HeroRank

func LoadHeroRank(filepath string) {
    fileName := "HeroRank.json"
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

	HeroRankmap = make(map[int32]*HeroRank)
	for dec.More() {
		var temp HeroRank
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		HeroRanks = append(HeroRanks, &temp)
		HeroRankmap[temp.Id] = &temp
	}

	LogInfo("Load HeroRank Scheme Success!")
}
