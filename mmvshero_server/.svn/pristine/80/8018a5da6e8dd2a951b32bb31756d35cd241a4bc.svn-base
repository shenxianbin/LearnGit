package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type GLocalization struct {
	 Id 	 int32
	 TipIndex 	 string
	 TextJP 	 string
	 Text 	 string

}

var GLocalizations []*GLocalization
var GLocalizationmap map[int32]*GLocalization

func LoadGLocalization(filepath string) {
    fileName := "GLocalization.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &GLocalizations)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    GLocalizationmap = make(map[int32]*GLocalization)
    for _, v := range GLocalizations {
        GLocalizationmap[v.Id] = v
    }
    LogInfo("Load GLocalization Scheme Success!")
}
