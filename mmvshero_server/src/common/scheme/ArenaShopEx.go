package scheme

import (
	"math/rand"
)

var arenaShopEx map[int32][]*ArenaShop
var arenaShopPanel []int32

func ArenaShopProcess() {
	arenaShopEx = make(map[int32][]*ArenaShop)
	arenaShopPanel = make([]int32, 0)
	for _, v := range ArenaShops {
		if _, has := arenaShopEx[v.Panel]; has {
			arenaShopEx[int32(v.Panel)] = append(arenaShopEx[int32(v.Panel)], v)
		} else {
			arenaShopEx[int32(v.Panel)] = make([]*ArenaShop, 1)
			arenaShopEx[int32(v.Panel)][0] = v
		}
		arenaShopPanel = append(arenaShopPanel, v.Panel)
	}
}

func ArenaShopRandomPanel() int32 {
	random := rand.Int31n(int32(len(arenaShopPanel)))
	return arenaShopPanel[random]
}

func ArenaShopGet(panel int32) []*ArenaShop {
	shop, has := arenaShopEx[panel]
	if has {
		return shop
	}
	return nil
}
