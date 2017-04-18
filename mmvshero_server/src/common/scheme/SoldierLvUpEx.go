package scheme

var soldierLvUpExmap map[int64]*SoldierLvUp

func SoldierLvUpProcess() {
	soldierLvUpExmap = make(map[int64]*SoldierLvUp)
	for _, value := range SoldierLvUpmap {
		var index int64 = int64(value.BaseId)<<16 | int64(value.Lv)
		soldierLvUpExmap[index] = value
	}
}

func SoldierLvUpGet(base_id int32, lv int32) *SoldierLvUp {
	var index int64 = int64(base_id)<<16 | int64(lv)
	if v, has := soldierLvUpExmap[index]; has {
		return v
	}

	return nil
}
