package scheme

import (
	"math"
)

var stageEliteFirstId int32 = math.MaxInt32

func StageProcess() {
	for _, value := range Stagemap {
		if value.StageType == 0 {
			continue
		}

		if stageEliteFirstId > value.Id {
			stageEliteFirstId = value.Id
		}
	}
}

//获得精英关卡第一关
func StageEliteGetFirst() int32 {
	return stageEliteFirstId
}
