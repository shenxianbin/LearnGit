package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Recharge struct {
	 Id 	 int32
	 Name 	 int32
	 Details 	 int32
	 Price 	 int32
	 GoldNum 	 int32
	 ExtraNum 	 int32
	 Duration 	 int32
	 FirstExtraNum 	 int32
	 FirstDetails 	 int32
	 LimitExtraNum 	 int32
	 LimitDetails 	 int32
	 Icon 	 string
	 Notes 	 string

}

var Recharges []*Recharge
var Rechargemap map[int32]*Recharge

func LoadRecharge(filepath string) {
    fileName := "Recharge.json"
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

	Rechargemap = make(map[int32]*Recharge)
	for dec.More() {
		var temp Recharge
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Recharges = append(Recharges, &temp)
		Rechargemap[temp.Id] = &temp
	}

	LogInfo("Load Recharge Scheme Success!")
}
