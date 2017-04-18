package scheme

import (
    "encoding/json"
    "fmt"
    . "galaxy"
    "io/ioutil"
)

type Obstacle struct {
	 Id 	 int32
	 Name 	 int32
	 RType 	 int32
	 Num 	 int32
	 Size 	 string
	 Icon 	 string
	 ObstacleFlashId 	 int32

}

var Obstacles []*Obstacle
var Obstaclemap map[int32]*Obstacle

func LoadObstacle(filepath string) {
    fileName := "Obstacle.json"
    file := fmt.Sprintf("%s/%s", filepath, fileName)
    buff, err := ioutil.ReadFile(file)
    err = json.Unmarshal(buff, &Obstacles)
    if err != nil {
        panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
    }
    Obstaclemap = make(map[int32]*Obstacle)
    for _, v := range Obstacles {
        Obstaclemap[v.Id] = v
    }
    LogInfo("Load Obstacle Scheme Success!")
}
