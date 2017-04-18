package scheme

var heroLvUpExmap map[int64]*HeroLvUp

func HeroLvUpProcess() {
	heroLvUpExmap = make(map[int64]*HeroLvUp)
	for _, value := range HeroLvUpmap {
		var index int64 = int64(value.BaseId)<<32 | int64(value.Lv<<16) | int64(value.Rank)
		heroLvUpExmap[index] = value
	}
}

func HeroLvUpGet(base_id int32, lv int32, rank int32) *HeroLvUp {
	var index int64 = int64(base_id)<<32 | int64(lv<<16) | int64(rank)
	if v, has := heroLvUpExmap[index]; has {
		return v
	}

	return nil
}
