package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Sign struct {
	 Id 	 int32
	 Month 	 int32
	 Index 	 int32
	 VipDouble 	 int32
	 AwardId 	 int32
	 Icon 	 string
	 Details 	 int32
	 Notes 	 string

}

var Signs []*Sign
var Signmap map[int32]*Sign

func LoadSign(filepath string) {
    fileName := "Sign.json"
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

	Signmap = make(map[int32]*Sign)
	for dec.More() {
		var temp Sign
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Signs = append(Signs, &temp)
		Signmap[temp.Id] = &temp
	}

	LogInfo("Load Sign Scheme Success!")
}
