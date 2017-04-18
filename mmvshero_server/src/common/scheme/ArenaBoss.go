package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ArenaBoss struct {
	 Id 	 int32
	 BossNpc 	 int32
	 Line1 	 int32
	 Award1 	 int32
	 Line2 	 int32
	 Award2 	 int32
	 Line3 	 int32
	 Award3 	 int32

}

var ArenaBosss []*ArenaBoss
var ArenaBossmap map[int32]*ArenaBoss

func LoadArenaBoss(filepath string) {
    fileName := "ArenaBoss.json"
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

	ArenaBossmap = make(map[int32]*ArenaBoss)
	for dec.More() {
		var temp ArenaBoss
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ArenaBosss = append(ArenaBosss, &temp)
		ArenaBossmap[temp.Id] = &temp
	}

	LogInfo("Load ArenaBoss Scheme Success!")
}
