package scheme

func AwardProcess() {
	for _, list := range Awardmap {
		for _, v := range list.RandomReward {
			var sum int32
			for _, data := range v.Data {
				data.Weight += sum
				sum = data.Weight
			}
		}
	}
}
