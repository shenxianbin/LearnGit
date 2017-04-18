package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Mall struct {
	 Id 	 int32
	 GroupId 	 int32
	 MallType 	 int32
	 LimitCount 	 int32
	 CD 	 int32
	 ConArgsType 	 int32
	 ConArgsValue 	 int32
	 NextId 	 int32
	 LastId 	 int32
	 CostType 	 int32
	 CostValue 	 int32
	 AwardId 	 int32
	 OnSale 	 int32
	 OriginalPrice 	 int32
	 Name 	 int32
	 Details 	 int32
	 Explain 	 int32
	 Icon 	 string
	 Notes 	 string

}

var Malls []*Mall
var Mallmap map[int32]*Mall

func LoadMall(filepath string) {
    fileName := "Mall.json"
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

	Mallmap = make(map[int32]*Mall)
	for dec.More() {
		var temp Mall
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Malls = append(Malls, &temp)
		Mallmap[temp.Id] = &temp
	}

	LogInfo("Load Mall Scheme Success!")
}
