package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type SoldierLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 LvUpKingLv 	 int32
	 Stage 	 int32
	 NeedExp 	 int32
	 ExpCount 	 int32
	 Hp 	 int32
	 Atk 	 int32
	 Spd 	 int32
	 Arg 	 int32
	 Ats 	 int32
	 View 	 int32
	 Text1 	 int32
	 Text2 	 int32
	 Text3 	 int32

}

var SoldierLvUps []*SoldierLvUp
var SoldierLvUpmap map[int32]*SoldierLvUp

func LoadSoldierLvUp(filepath string) {
    fileName := "SoldierLvUp.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &SoldierLvUps)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    SoldierLvUpmap = make(map[int32]*SoldierLvUp)
    for _, v := range SoldierLvUps {
        SoldierLvUpmap[v.Id] = v
    }
    LogInfo("Load SoldierLvUp Scheme Success!")
}
