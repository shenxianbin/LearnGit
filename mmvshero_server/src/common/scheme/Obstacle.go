package scheme

import (
	"encoding/json"
	"fmt"
	. "galaxy"
	"os"
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

	Obstaclemap = make(map[int32]*Obstacle)
	for dec.More() {
		var temp Obstacle
		err := dec.Decode(&temp)
		if err != nil {
			panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
		}
		Obstacles = append(Obstacles, &temp)
		Obstaclemap[temp.Id] = &temp
	}

	LogInfo("Load Obstacle Scheme Success!")
}
