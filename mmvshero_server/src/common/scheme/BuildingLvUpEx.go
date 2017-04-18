package scheme

var buildingLvUpExmap map[int32]*BuildingLvUp

func BuildingLvUpProcess() {
	buildingLvUpExmap = make(map[int32]*BuildingLvUp)
	for _, value := range BuildingLvUpmap {
		index := value.BaseId<<8 | value.Lv
		buildingLvUpExmap[index] = value
	}
}

func BuildingLvUpGet(base_id int32, lv int32) *BuildingLvUp {
	index := base_id<<8 | lv
	if v, has := buildingLvUpExmap[index]; has {
		return v
	}

	return nil
}
