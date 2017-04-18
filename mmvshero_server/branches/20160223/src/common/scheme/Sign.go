package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Signs)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Signmap = make(map[int32]*Sign)
    for _, v := range Signs {
        Signmap[v.Id] = v
    }
    LogInfo("Load Sign Scheme Success!")
}
