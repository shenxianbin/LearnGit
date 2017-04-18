package scheme

import (
	"strconv"
	"strings"
)

type RoleLvUpEx struct {
	MapSize       []int32
	SoldierUnlock []int32
	SoldierObtain []int32
	HeroUnlock    []int32
	KingSkillId   []int32
	PlunderTeam   []int32
}

var RoleLvUpExmap map[int32]*RoleLvUpEx

func RoleLvUpProcess() {
	RoleLvUpExmap = make(map[int32]*RoleLvUpEx)
	for k, v := range RoleLvUpmap {
		roleLvUpEx := new(RoleLvUpEx)
		map_size := strings.Split(v.MapSize, ";")
		roleLvUpEx.MapSize = make([]int32, len(map_size))
		for index, temp := range map_size {
			data, _ := strconv.Atoi(temp)
			roleLvUpEx.MapSize[index] = int32(data)
		}

		soldier_unlock := strings.Split(v.SoldierUnlock, ";")
		roleLvUpEx.SoldierUnlock = make([]int32, len(soldier_unlock))
		for index, temp := range soldier_unlock {
			data, _ := strconv.Atoi(temp)
			roleLvUpEx.SoldierUnlock[index] = int32(data)
		}

		soldier_obtain := strings.Split(v.SoldierObtain, ";")
		roleLvUpEx.SoldierObtain = make([]int32, len(soldier_obtain))
		for index, temp := range soldier_obtain {
			data, _ := strconv.Atoi(temp)
			roleLvUpEx.SoldierObtain[index] = int32(data)
		}

		hero_unlock := strings.Split(v.HeroUnlock, ";")
		roleLvUpEx.HeroUnlock = make([]int32, len(hero_unlock))
		for index, temp := range hero_unlock {
			data, _ := strconv.Atoi(temp)
			roleLvUpEx.HeroUnlock[index] = int32(data)
		}

		king_skillid := strings.Split(v.KingSkillId, ";")
		roleLvUpEx.KingSkillId = make([]int32, len(king_skillid))
		for index, temp := range king_skillid {
			data, _ := strconv.Atoi(temp)
			roleLvUpEx.KingSkillId[index] = int32(data)
		}

		plunder_team := strings.Split(v.PlunderTeam, ";")
		roleLvUpEx.PlunderTeam = make([]int32, len(plunder_team))
		for index, temp := range plunder_team {
			data, _ := strconv.Atoi(temp)
			roleLvUpEx.PlunderTeam[index] = int32(data)
		}

		RoleLvUpExmap[k] = roleLvUpEx
	}
}
