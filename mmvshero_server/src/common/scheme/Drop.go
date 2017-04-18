package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Drop struct {
	 Id 	 int32
	 DrawType 	 int32
	 Plan 	 int32
	 Name 	 string
	 RandomAwardId 	 int32

}

var Drops []*Drop
var Dropmap map[int32]*Drop

func LoadDrop(filepath string) {
    fileName := "Drop.json"
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

	Dropmap = make(map[int32]*Drop)
	for dec.More() {
		var temp Drop
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Drops = append(Drops, &temp)
		Dropmap[temp.Id] = &temp
	}

	LogInfo("Load Drop Scheme Success!")
}
