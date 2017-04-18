package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Recharges)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Rechargemap = make(map[int32]*Recharge)
    for _, v := range Recharges {
        Rechargemap[v.Id] = v
    }
    LogInfo("Load Recharge Scheme Success!")
}
