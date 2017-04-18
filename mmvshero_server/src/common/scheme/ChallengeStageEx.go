package scheme

import (
	"math/rand"
)

var challengeStageEx map[int32][]*ChallengeStage

func ChallengeStageProcess() {
	challengeStageEx = make(map[int32][]*ChallengeStage)
	for _, v := range ChallengeStagemap {
		if _, has := challengeStageEx[v.Stage]; has {
			challengeStageEx[int32(v.Stage)] = append(challengeStageEx[int32(v.Stage)], v)
		} else {
			challengeStageEx[int32(v.Stage)] = make([]*ChallengeStage, 1)
			challengeStageEx[int32(v.Stage)][0] = v
		}
	}
}

func ChallengeStageRandGet(stage int32) *ChallengeStage {
	stages, has := challengeStageEx[stage]
	if has {
		index := rand.Int31n(int32(len(stages)))
		return stages[index]
	}
	return nil
}
