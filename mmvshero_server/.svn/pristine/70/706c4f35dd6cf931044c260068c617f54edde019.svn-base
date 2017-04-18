package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ArenaRank struct {
	 Id 	 int32
	 RankUplimit 	 int32
	 RankDownlimit 	 int32
	 RankAward 	 int32

}

var ArenaRanks []*ArenaRank
var ArenaRankmap map[int32]*ArenaRank

func LoadArenaRank(filepath string) {
    fileName := "ArenaRank.json"
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

	ArenaRankmap = make(map[int32]*ArenaRank)
	for dec.More() {
		var temp ArenaRank
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ArenaRanks = append(ArenaRanks, &temp)
		ArenaRankmap[temp.Id] = &temp
	}

	LogInfo("Load ArenaRank Scheme Success!")
}
