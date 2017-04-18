package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
)

type StageTask struct {
	 Id 	 int32
	 Details 	 int32
	 Notes 	 string

}

var StageTasks []*StageTask
var StageTaskmap map[int32]*StageTask

func LoadStageTask(filepath string) {
    fileName := "StageTask.json"
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

	StageTaskmap = make(map[int32]*StageTask)
	for dec.More() {
		var temp StageTask
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		StageTasks = append(StageTasks, &temp)
		StageTaskmap[temp.Id] = &temp
	}

	LogInfo("Load StageTask Scheme Success!")
}
