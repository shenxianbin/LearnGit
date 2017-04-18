package scheme

import (
	"strconv"
	"strings"
)

type PlunderRandom struct {
	Num  int32
	Flow []float32
}

type PlunderAwardEx struct {
	SoulAward    []float32
	SoldierAward PlunderRandom
	HeroPool     []int32
	HeroAward    PlunderRandom
}

var PlunderAwardExmap map[int32]*PlunderAwardEx

func PlunderAwardProcess() {
	PlunderAwardExmap = make(map[int32]*PlunderAwardEx)
	for k, v := range PlunderAwardmap {
		plunderAwardEx := new(PlunderAwardEx)

		soul_award := strings.Split(v.SoulAward, ",")
		plunderAwardEx.SoulAward = make([]float32, len(soul_award))
		for index, temp := range soul_award {
			data, _ := strconv.ParseFloat(temp, 32)
			plunderAwardEx.SoulAward[index] = float32(data)
		}

		soldier_award := strings.Split(v.SoldierAward, ";")
		soldier_random := strings.Split(soldier_award[1], ",")
		num, _ := strconv.Atoi(soldier_award[0])
		min, _ := strconv.ParseFloat(soldier_random[0], 32)
		max, _ := strconv.ParseFloat(soldier_random[1], 32)
		plunderAwardEx.SoldierAward.Num = int32(num)
		plunderAwardEx.SoldierAward.Flow = make([]float32, 2)
		plunderAwardEx.SoldierAward.Flow[0] = float32(min)
		plunderAwardEx.SoldierAward.Flow[1] = float32(max)

		hero_pool := strings.Split(v.HeroPool, ",")
		plunderAwardEx.HeroPool = make([]int32, len(hero_pool))
		for index, temp := range hero_pool {
			data, _ := strconv.Atoi(temp)
			plunderAwardEx.HeroPool[index] = int32(data)
		}

		hero_award := strings.Split(v.HeroAward, ";")
		hero_random := strings.Split(hero_award[1], ",")
		num, _ = strconv.Atoi(hero_award[0])
		min, _ = strconv.ParseFloat(hero_random[0], 32)
		max, _ = strconv.ParseFloat(hero_random[1], 32)
		plunderAwardEx.HeroAward.Num = int32(num)
		plunderAwardEx.HeroAward.Flow = make([]float32, 2)
		plunderAwardEx.HeroAward.Flow[0] = float32(min)
		plunderAwardEx.HeroAward.Flow[1] = float32(max)

		PlunderAwardExmap[k] = plunderAwardEx
	}
}
