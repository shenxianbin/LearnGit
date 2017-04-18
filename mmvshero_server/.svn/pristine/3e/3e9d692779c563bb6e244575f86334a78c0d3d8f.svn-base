package gm

import (
	"fmt"
)

const (
	cache_gm_orderlist_t = "Role:%v:GmOrderList"
	cache_gm_order_t     = "Role:%v:GmOrder:%v"
)

func GenGmOrderListKey(roleUid int64) string {
	return fmt.Sprintf(cache_gm_orderlist_t, roleUid)
}

func GenGmOrderKey(roleUid int64, timeStamp int64) string {
	return fmt.Sprintf(cache_gm_order_t, roleUid, timeStamp)
}
