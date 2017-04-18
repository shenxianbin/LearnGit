package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
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
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &StageTasks)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    StageTaskmap = make(map[int32]*StageTask)
    for _, v := range StageTasks {
        StageTaskmap[v.Id] = v
    }
    LogInfo("Load StageTask Scheme Success!")
}
