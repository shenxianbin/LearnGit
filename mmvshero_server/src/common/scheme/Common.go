package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Common struct {
	 Id 	 int32
	 Parameter 	 string
	 Value 	 int32
	 Notes 	 string

}

var Commons []*Common
var Commonmap map[int32]*Common

func LoadCommon(filepath string) {
    fileName := "Common.json"
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

	Commonmap = make(map[int32]*Common)
	for dec.More() {
		var temp Common
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Commons = append(Commons, &temp)
		Commonmap[temp.Id] = &temp
	}

	LogInfo("Load Common Scheme Success!")
}
