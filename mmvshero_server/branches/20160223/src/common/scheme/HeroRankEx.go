package scheme

var heroRankExmap map[int32]*HeroRank

func HeroRankProcess() {
	heroRankExmap = make(map[int32]*HeroRank)
	for _, value := range HeroRankmap {
		index := value.Plan<<16 | value.Rank
		heroRankExmap[index] = value
	}
}

func HeroRankGet(plan int32, rank int32) *HeroRank {
	index := plan<<16 | rank
	if v, has := heroRankExmap[index]; has {
		return v
	}

	return nil
}
