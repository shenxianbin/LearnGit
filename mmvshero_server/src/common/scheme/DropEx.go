package scheme

import (
	"math/rand"
	"time"
)

var dropExmap map[int32]map[int32]*Drop

func DropProcess() {
	dropExmap = make(map[int32]map[int32]*Drop)
	for _, value := range Dropmap {
		if _, has := dropExmap[value.DrawType]; !has {
			dropExmap[value.DrawType] = make(map[int32]*Drop)
		}
		dropExmap[value.DrawType][value.Plan] = value
	}
}

func DropGet(draw_type int32) *Drop {
	v, has := dropExmap[draw_type]
	if !has {
		return nil
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := random.Int31n(int32(len(v)))
	var index int32
	for _, temp := range v {
		if num == index {
			return temp
		} else {
			index++
		}
	}
	return nil
}
