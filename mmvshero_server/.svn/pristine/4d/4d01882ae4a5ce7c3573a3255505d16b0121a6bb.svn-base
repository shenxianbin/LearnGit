package scheme

var heroStageUpExmap map[int64]*HeroStageUp

func HeroStageUpProcess() {
	heroStageUpExmap = make(map[int64]*HeroStageUp)
	for _, value := range HeroStageUpmap {
		var index int64 = int64(value.BaseId)<<32 | int64(value.Stage<<8) | int64(value.Rank)
		heroStageUpExmap[index] = value
	}
}

func HeroStageUpGet(base_id int32, stage int32, rank int32) *HeroStageUp {
	var index int64 = int64(base_id)<<32 | int64(stage<<8) | int64(rank)
	if v, has := heroStageUpExmap[index]; has {
		return v
	}

	return nil
}
