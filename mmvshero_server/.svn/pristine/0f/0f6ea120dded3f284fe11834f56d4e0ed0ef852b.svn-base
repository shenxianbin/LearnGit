package scheme

var skillLvUpExmap map[int32]*SkillLvUp

func SkillLvUpProcess() {
	skillLvUpExmap = make(map[int32]*SkillLvUp)
	for _, value := range SkillLvUpmap {
		index := value.BaseId<<8 | value.Lv
		skillLvUpExmap[index] = value
	}
}

func SkillLvUpGet(base_id int32, lv int32) *SkillLvUp {
	index := base_id<<8 | lv
	if v, has := skillLvUpExmap[index]; has {
		return v
	}

	return nil
}
