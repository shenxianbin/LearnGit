package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type ModuleOpen struct {
	 Id 	 int32
	 NeedRoleLv 	 int32
	 Notes 	 string

}

var ModuleOpens []*ModuleOpen
var ModuleOpenmap map[int32]*ModuleOpen

func LoadModuleOpen(filepath string) {
    fileName := "ModuleOpen.json"
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

	ModuleOpenmap = make(map[int32]*ModuleOpen)
	for dec.More() {
		var temp ModuleOpen
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		ModuleOpens = append(ModuleOpens, &temp)
		ModuleOpenmap[temp.Id] = &temp
	}

	LogInfo("Load ModuleOpen Scheme Success!")
}