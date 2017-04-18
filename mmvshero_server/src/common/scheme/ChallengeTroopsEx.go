package scheme

import (
	"math/rand"
)

var challengeTroopEx map[int32][]*ChallengeTroops

func ChallengeTroopProcess() {
	challengeTroopEx = make(map[int32][]*ChallengeTroops)
	for _, v := range ChallengeTroopsmap {
		key := v.Type<<24 | v.WaveIndex<<12 | v.WaveLv

		if _, has := challengeTroopEx[key]; has {
			challengeTroopEx[key] = append(challengeTroopEx[key], v)
		} else {
			challengeTroopEx[key] = make([]*ChallengeTroops, 1)
			challengeTroopEx[key][0] = v
		}
	}
}

func ChallengeTroopRandGet(Type, WaveIndex, WaveLv int32) *ChallengeTroops {
	key := Type<<24 | WaveIndex<<12 | WaveLv
	troops, has := challengeTroopEx[key]
	if has {
		index := rand.Int31n(int32(len(troops)))
		return troops[index]
	}
	return nil
}
