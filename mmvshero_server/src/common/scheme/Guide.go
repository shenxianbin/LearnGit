package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type Guide struct {
	 Id 	 int32
	 StepId 	 int32
	 Type 	 int32
	 Skeleton 	 string
	 FlashPath 	 string
	 TextId 	 int32
	 Trigger 	 int32
	 Details 	 string

}

var Guides []*Guide
var Guidemap map[int32]*Guide

func LoadGuide(filepath string) {
    fileName := "Guide.json"
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

	Guidemap = make(map[int32]*Guide)
	for dec.More() {
		var temp Guide
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Guides = append(Guides, &temp)
		Guidemap[temp.Id] = &temp
	}

	LogInfo("Load Guide Scheme Success!")
}