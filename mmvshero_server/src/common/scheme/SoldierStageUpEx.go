package scheme

var soldierStageUpExmap map[int64]*SoldierStageUp

func SoldierStageUpProcess() {
	soldierStageUpExmap = make(map[int64]*SoldierStageUp)
	for _, value := range SoldierStageUpmap {
		var index int64 = int64(value.BaseId)<<16 | int64(value.Stage)
		soldierStageUpExmap[index] = value
	}
}

func SoldierStageUpGet(base_id, stage int32) *SoldierStageUp {
	var index int64 = int64(base_id)<<16 | int64(stage)
	if v, has := soldierStageUpExmap[index]; has {
		return v
	}

	return nil
}
