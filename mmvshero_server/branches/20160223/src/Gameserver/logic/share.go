package logic

import (
	"common"
	d "common/define"
	"common/scheme"
	"galaxy"
	"github.com/mediocregopher/radix.v2/redis"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//资源转换为骷髅币 时间 魔血 魔魂
func ResourceToCoin(resourceType, resourceAmount int32) int32 {
	var ret, Price1, Price2, Price3, Price4, Param1, Param2, Param3, Param4 float64
	if resourceType == common.RTYPE_TIME {
		Price1 = float64(scheme.Commonmap[d.TimePrice1].Value)
		Price2 = float64(scheme.Commonmap[d.TimePrice2].Value)
		Price3 = float64(scheme.Commonmap[d.TimePrice3].Value)
		Price4 = float64(scheme.Commonmap[d.TimePrice4].Value)
		Param1 = float64(scheme.Commonmap[d.TimeParam1].Value)
		Param2 = float64(scheme.Commonmap[d.TimeParam2].Value)
		Param3 = float64(scheme.Commonmap[d.TimeParam3].Value)
		Param4 = float64(scheme.Commonmap[d.TimeParam4].Value)
	} else if resourceType == common.RTYPE_BLOOD {
		Price1 = float64(scheme.Commonmap[d.BloodPrice1].Value)
		Price2 = float64(scheme.Commonmap[d.BloodPrice2].Value)
		Price3 = float64(scheme.Commonmap[d.BloodPrice3].Value)
		Price4 = float64(scheme.Commonmap[d.BloodPrice4].Value)
		Param1 = float64(scheme.Commonmap[d.BloodParam1].Value)
		Param2 = float64(scheme.Commonmap[d.BloodParam2].Value)
		Param3 = float64(scheme.Commonmap[d.BloodParam3].Value)
		Param4 = float64(scheme.Commonmap[d.BloodParam4].Value)
	} else if resourceType == common.RTYPE_SOUL {
		Price1 = float64(scheme.Commonmap[d.SoulPrice1].Value)
		Price2 = float64(scheme.Commonmap[d.SoulPrice2].Value)
		Price3 = float64(scheme.Commonmap[d.SoulPrice3].Value)
		Price4 = float64(scheme.Commonmap[d.SoulPrice4].Value)
		Param1 = float64(scheme.Commonmap[d.SoulParam1].Value)
		Param2 = float64(scheme.Commonmap[d.SoulParam2].Value)
		Param3 = float64(scheme.Commonmap[d.SoulParam3].Value)
		Param4 = float64(scheme.Commonmap[d.SoulParam4].Value)
	} else {
		return 0
	}

	amount := float64(resourceAmount)
	if amount <= 0 {
		ret = 0
	} else if 0 < amount && amount <= Param1 {
		ret = Price1
	} else if Param1 < amount && amount <= Param2 {
		ret = (Price2-Price1)/(Param2-Param1)*(amount-Param1) + Price1
	} else if Param2 < amount && amount <= Param3 {
		ret = (Price3-Price2)/(Param3-Param2)*(amount-Param2) + Price2
	} else if amount > Param3 {
		ret = (Price4-Price3)/(Param4-Param3)*(amount-Param3) + Price3
	}
	return int32(math.Ceil(ret))
}

//生成随机数
func Rand(min, max int64) int64 {
	if min > max {
		max, min = min, max
	} else if min == max {
		return min
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t := r.Float64()*float64(max-min) + float64(min)
	return int64(math.Floor(t + 0.5))
}

func Min(args ...int32) int32 {
	if len(args) == 0 {
		panic("func Min need args")
		return 0
	}

	var min int32 = args[0]
	for _, k := range args {
		if k < min {
			min = k
		}
	}
	return min
}

func Max(args ...int32) int32 {
	if len(args) == 0 {
		panic("func Max need args")
		return 0
	}

	var max int32 = args[0]
	for _, k := range args {
		if k > max {
			max = k
		}
	}
	return max
}

func Time() int64 {
	return time.Now().Unix()
}

func RefreshTime(hour int32) int64 {
	now := time.Now()
	var day int
	if now.Hour() >= int(hour) {
		day = now.Day()
	} else {
		day = now.Day() - 1
	}
	return time.Date(now.Year(), now.Month(), day, int(hour), 0, 0, 0, time.Local).Unix()
}

func ExtractIntPairMap(str string) map[int32]int32 {
	temp := strings.Split(str, ";")
	ret := make(map[int32]int32)

	for _, v := range temp {
		if v == "" {
			continue
		}
		pair := strings.Split(v, ",")
		key, _ := strconv.Atoi(pair[0])
		value, _ := strconv.Atoi(pair[1])
		ret[int32(key)] = int32(value)
	}

	return ret
}

func ExtractIntMap(str string) map[int32]int32 {
	temp := strings.Split(str, ";")
	ret := make(map[int32]int32, len(temp))

	for i := 0; i < len(temp); i++ {
		t, _ := strconv.Atoi(temp[i])
		ret[int32(i)] = int32(t)
	}
	return ret
}

//根据掉落概率获得真实掉落列表
func GetAwardsByProp(awardsProp map[int32]int32) map[int32]int32 {
	ret := make(map[int32]int32)
	for id, prop := range awardsProp {
		if int32(Rand(0, 100)) <= prop {
			ret[id] = prop
		}
	}
	return ret
}

func RedisCmd(cmd string, args ...interface{}) (*redis.Resp, error) {
	return galaxy.GxService().Redis().Cmd(cmd, args)
}
