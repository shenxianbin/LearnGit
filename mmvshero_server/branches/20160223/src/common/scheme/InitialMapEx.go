package scheme

import (
	"errors"
	. "galaxy"
)

func GetMap(x_size int32, y_size int32) [][]*InitalizeMapGrid {
	if x_size/2 == 0 {
		return nil
	}

	LogInfo("x_size : ", x_size, "y_size : ", y_size)

	scheme_x_size := int32(len(WholeMap))
	scheme_y_size := int32(len(WholeMap[0]))
	LogInfo("scheme_x_size : ", scheme_x_size, "scheme_y_size : ", scheme_y_size)

	if x_size > scheme_x_size || y_size > scheme_y_size {
		return nil
	}

	if x_size == scheme_x_size && y_size == scheme_y_size {
		LogDebug("WholeMap")
		return WholeMap
	}

	maps := make([][]*InitalizeMapGrid, x_size)
	fix_x := (scheme_x_size - x_size) / 2
	fix_y := scheme_y_size - y_size
	for index_x, points := range maps {
		if points == nil {
			maps[index_x] = make([]*InitalizeMapGrid, y_size)
		}
		for index_y, _ := range maps[index_x] {
			maps[index_x][index_y] = WholeMap[int32(index_x)+fix_x][int32(index_y)+fix_y]
		}
	}
	return maps
}

func ParseXY(x_before int32, y_before int32, x_size int32, y_size int32) (x_after int32, y_after int32, err error) {
	if x_size/2 == 0 {
		err = errors.New("x_size equals zero")
		return
	}

	LogInfo("x_before : ", x_before, " y_before : ", y_before, " x_size : ", x_size, " y_size : ", y_size)

	scheme_x_size := int32(len(WholeMap))
	scheme_y_size := int32(len(WholeMap[0]))
	LogInfo("scheme_x_size : ", scheme_x_size, " scheme_y_size : ", scheme_y_size)

	if x_size > scheme_x_size || y_size > scheme_y_size {
		err = errors.New("x_size or y_size error")
		return
	}

	if x_size == scheme_x_size && y_size == scheme_y_size {
		x_after = x_before
		y_after = y_before
		return
	}

	fix_x := (scheme_x_size - x_size) / 2
	fix_y := scheme_y_size - y_size
	x_after = x_before - fix_x
	y_after = y_before - fix_y
	LogInfo("x_after : ", x_after, " y_after : ", y_after)
	return
}
