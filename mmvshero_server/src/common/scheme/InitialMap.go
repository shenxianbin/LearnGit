package scheme

import (
	"common"
	"encoding/json"
	"fmt"
	. "galaxy"
	"io/ioutil"
	"strconv"
	"strings"
)

type MapJson struct {
	Name    string
	X       int32
	Y       int32
	GridArr []string
	JdArr   []string
}

type InitalizeMapPoint struct {
	PointId    int32
	NeedKingLv int32
	RType      int32
	Value      int32
	IsLock     bool //0,已解锁 1.未解锁
}

type InitalizeMapPointBuilding struct {
	BuildingId int32
	IsWay      int32
	X          int32
	Y          int32
}

type InitalizeMapGrid struct {
	PointId   int32
	IsWay     int32
	PointType int32
	ObjType   int32
	SchemeId  int32
}

var mapJson *MapJson
var WholeMap [][]*InitalizeMapGrid
var PointMap map[int32]*InitalizeMapPoint
var PointMapBuilding map[int32][]*InitalizeMapPointBuilding

func LoadInitialMap(filepath string) {
	fileName := "InitialMap.json"
	file := fmt.Sprintf("%s/%s", filepath, fileName)
	buff, err := ioutil.ReadFile(file)
	err = json.Unmarshal(buff, &mapJson)
	if err != nil {
		panic(fmt.Sprintf("Read [file ：%s]occurs error: %s", fileName, err))
	}

	total_size := mapJson.X * mapJson.Y
	if total_size != int32(len(mapJson.GridArr)) {
		panic("InitialMap Data size error")
	}

	WholeMap = make([][]*InitalizeMapGrid, mapJson.X)
	PointMapBuilding = make(map[int32][]*InitalizeMapPointBuilding)
	for index_x, v_x := range WholeMap {
		if v_x == nil {
			WholeMap[index_x] = make([]*InitalizeMapGrid, mapJson.Y)
		}

		for index_y, _ := range WholeMap[index_x] {
			s := strings.Split(mapJson.GridArr[index_y*int(mapJson.X)+index_x], "-")
			pointId, _ := strconv.Atoi(s[0])
			isWay, _ := strconv.Atoi(s[1])
			pointType, _ := strconv.Atoi(s[2])
			objType, _ := strconv.Atoi(s[3])
			schemeId, _ := strconv.Atoi(s[4])

			WholeMap[index_x][index_y] = &InitalizeMapGrid{
				PointId:   int32(pointId),
				IsWay:     int32(isWay),
				PointType: int32(pointType),
				ObjType:   int32(objType),
				SchemeId:  int32(schemeId),
			}

			if objType == common.MAP_OBJ_BUILDING && pointId != 0 {
				if _, has := PointMapBuilding[int32(pointId)]; has {
					PointMapBuilding[int32(pointId)] = append(PointMapBuilding[int32(pointId)], &InitalizeMapPointBuilding{
						BuildingId: int32(schemeId),
						IsWay:      int32(isWay),
						X:          int32(index_x),
						Y:          int32(index_y),
					})
				} else {
					PointMapBuilding[int32(pointId)] = make([]*InitalizeMapPointBuilding, 1)
					PointMapBuilding[int32(pointId)][0] = &InitalizeMapPointBuilding{
						BuildingId: int32(schemeId),
						IsWay:      int32(isWay),
						X:          int32(index_x),
						Y:          int32(index_y),
					}
				}
			}
		}
	}

	PointMap = make(map[int32]*InitalizeMapPoint)
	for _, v := range mapJson.JdArr {
		s := strings.Split(v, "-")
		pointId, _ := strconv.Atoi(s[0])
		needKingLv, _ := strconv.Atoi(s[1])
		rType, _ := strconv.Atoi(s[2])
		value, _ := strconv.Atoi(s[3])
		isLock, _ := strconv.Atoi(s[4])

		PointMap[int32(pointId)] = &InitalizeMapPoint{
			PointId:    int32(pointId),
			NeedKingLv: int32(needKingLv),
			RType:      int32(rType),
			Value:      int32(value),
			IsLock:     isLock == 1,
		}
	}

	LogInfo("Load InitialMap Scheme Success!")
}
