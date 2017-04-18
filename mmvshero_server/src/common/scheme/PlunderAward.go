package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type PlunderAward struct {
	 Id 	 int32
	 EscortTime 	 int32
	 Quantity 	 int32
	 OnceQuantity 	 int32
	 SoulAward 	 string
	 SoldierAward 	 string
	 HeroPool 	 string
	 HeroAward 	 string
	 Status 	 int32
	 Icon 	 string
	 Detail 	 int32

}

var PlunderAwards []*PlunderAward
var PlunderAwardmap map[int32]*PlunderAward

func LoadPlunderAward(filepath string) {
    fileName := "PlunderAward.json"
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

	PlunderAwardmap = make(map[int32]*PlunderAward)
	for dec.More() {
		var temp PlunderAward
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		PlunderAwards = append(PlunderAwards, &temp)
		PlunderAwardmap[temp.Id] = &temp
	}

	LogInfo("Load PlunderAward Scheme Success!")
}
