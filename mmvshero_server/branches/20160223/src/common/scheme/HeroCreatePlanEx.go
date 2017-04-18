package scheme

import (
	"math/rand"
	"time"
)

var maxIndex int32

func HeroCreatePlanProcess() {
	for key, _ := range HeroCreatePlanmap {
		if maxIndex < key {
			maxIndex = key
		}
	}
}

func HeroCreateRandPlanId() int32 {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	random := r.Int31n(maxIndex)
	return int32(random)
}
