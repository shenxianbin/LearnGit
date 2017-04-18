package scheme

var arenaRankEx map[int32]int32
var lastRank int32
var lastAward int32

func ArenaRankProcess() {
	arenaRankEx = make(map[int32]int32)
	for _, v := range ArenaRanks {
		for i := v.RankUplimit; i <= v.RankDownlimit; i++ {
			arenaRankEx[i] = v.RankAward
		}
		if v.RankDownlimit == -1 {
			lastRank = v.RankUplimit
			lastAward = v.RankAward
		}
	}
}

func ArenaRankGet(rank int32) int32 {
	award, has := arenaRankEx[rank]
	if has {
		return award
	}

	if rank >= lastRank {
		return lastAward
	}
	return 0
}
