package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ArenaShop struct {
	 Id 	 int32
	 Panel 	 int32
	 ExchangeTimes 	 int32
	 ExchangeAward 	 int32
	 CostPoint 	 int32

}

var ArenaShops []*ArenaShop
var ArenaShopmap map[int32]*ArenaShop

func LoadArenaShop(filepath string) {
    fileName := "ArenaShop.json"
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

	ArenaShopmap = make(map[int32]*ArenaShop)
	for dec.More() {
		var temp ArenaShop
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ArenaShops = append(ArenaShops, &temp)
		ArenaShopmap[temp.Id] = &temp
	}

	LogInfo("Load ArenaShop Scheme Success!")
}
