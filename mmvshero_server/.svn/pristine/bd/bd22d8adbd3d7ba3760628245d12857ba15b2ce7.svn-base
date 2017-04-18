package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type SoldierLvUp struct {
	 Id 	 int32
	 BaseId 	 int32
	 Name 	 string
	 Lv 	 int32
	 LvUpRoleLv 	 int32
	 Stage 	 int32
	 NeedExp 	 int32
	 ExpCount 	 int32
	 AddRoleExp 	 int32
	 Hp 	 int32
	 Atk 	 int32
	 Spd 	 int32
	 Arg 	 int32
	 Ats 	 int32
	 View 	 int32
	 Text1 	 int32
	 Text2 	 int32
	 Power 	 int32
	 AtsText 	 string
	 ArgText 	 string

}

var SoldierLvUps []*SoldierLvUp
var SoldierLvUpmap map[int32]*SoldierLvUp

func LoadSoldierLvUp(filepath string) {
    fileName := "SoldierLvUp.json"
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

	SoldierLvUpmap = make(map[int32]*SoldierLvUp)
	for dec.More() {
		var temp SoldierLvUp
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		SoldierLvUps = append(SoldierLvUps, &temp)
		SoldierLvUpmap[temp.Id] = &temp
	}

	LogInfo("Load SoldierLvUp Scheme Success!")
}
