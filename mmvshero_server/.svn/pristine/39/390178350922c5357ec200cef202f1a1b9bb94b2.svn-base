package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type GLocalization struct {
	 Id 	 int32
	 TipIndex 	 string
	 Text 	 string
	 TextCN 	 string

}

var GLocalizations []*GLocalization
var GLocalizationmap map[int32]*GLocalization

func LoadGLocalization(filepath string) {
    fileName := "GLocalization.json"
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

	GLocalizationmap = make(map[int32]*GLocalization)
	for dec.More() {
		var temp GLocalization
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		GLocalizations = append(GLocalizations, &temp)
		GLocalizationmap[temp.Id] = &temp
	}

	LogInfo("Load GLocalization Scheme Success!")
}
