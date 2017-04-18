package recharge

import (
	"Gameserver/logic"
	"Gameserver/logic/role"
	"bytes"
	"common/scheme"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	. "galaxy"
	"galaxy/event"
	"net/http"
	"strconv"
)

const (
	key = "zxpoi098&"
)

type ResultMessage struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

func handlerRecharge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	r.ParseForm()

	orderIdStr := r.Form.Get("orderId")
	roleIdStr := r.Form.Get("roleId")
	itemIdStr := r.Form.Get("itemId")
	//serverIdStr := r.Form.Get("serverId")
	//platformStr := r.Form.Get("platform")
	md5Str := r.Form.Get("md5")

	result := new(ResultMessage)
	result.Ret = 1

	// 验证MD5
	var buf bytes.Buffer
	buf.WriteString(orderIdStr)
	buf.WriteString(roleIdStr)
	buf.WriteString(itemIdStr)
	buf.WriteString(key)
	h := md5.New()
	h.Write(buf.Bytes())
	cipherStr := h.Sum(nil)
	check_md5Str := hex.EncodeToString(cipherStr)

	defer func() {
		content, err := json.Marshal(result)
		if err != nil {
			LogError(err)
		}
		w.Write(content)
	}()

	if md5Str != check_md5Str {
		result.Msg = "MD5 Check Error"
		return
	}

	//检测订单号重复
	resp, err := GxService().Redis().Cmd("SETNX", orderIdStr, "")
	if err != nil {
		result.Msg = "Redis SETNX Error"
		return
	}

	ret, err := resp.Int()
	if err != nil {
		result.Msg = "Redis SETNX Ret Error"
		return
	}

	if ret == 0 {
		result.Msg = "OrderId Exists"
		return
	}

	//操作玩家
	roleUid, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		result.Msg = "RoleUid Parse Error"
		return
	}

	itemId, err := strconv.ParseInt(itemIdStr, 10, 32)
	if err != nil {
		result.Msg = "ItemId Parse Error"
		return
	}

	LogDebug("RoleUid : ", roleUid, " itemId : ", itemId)
	recharge_scheme, has := scheme.Rechargemap[int32(itemId)]
	if !has {
		result.Msg = "Recharge Scheme Error"
		return
	}

	event.GxEvent().Execute(func(args ...interface{}) {
		online_role := logic.GetRoleByUid(roleUid)
		if online_role != nil {
			count := online_role.GetRechargeRecord(int32(itemId))
			if count != 0 {
				online_role.AddGold(recharge_scheme.GoldNum, true, true)
			} else {
				online_role.AddGold(recharge_scheme.GoldNum+recharge_scheme.FirstExtraNum, true, true)
			}
			if recharge_scheme.Duration != 0 {
				online_role.AddVip(int64(recharge_scheme.Duration*24*3600), true, true)
			}
			online_role.AddRechargeRecord(int32(itemId))
		} else {
			temp := role.NewRole()
			offline_role := temp.OfflineRoleBase(roleUid)
			if offline_role != nil {
				result.Msg = "Load OfflineRole Error"
				return
			}

			count := offline_role.GetRechargeRecord(int32(itemId))
			if count != 0 {
				offline_role.AddGold(recharge_scheme.GoldNum, false, true)
			} else {
				offline_role.AddGold(recharge_scheme.GoldNum+recharge_scheme.FirstExtraNum, false, true)
			}
			if recharge_scheme.Duration != 0 {
				offline_role.AddVip(int64(recharge_scheme.Duration*24*3600), false, true)
			}
			offline_role.AddRechargeRecord(int32(itemId))
		}
	})

	result.Ret = 0
	result.Msg = "Success"
}

func InitRechargeModule(addr string) {
	http.HandleFunc("/recharge", handlerRecharge)
	http.NotFoundHandler()

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			LogFatal(err)
		}
	}()
}
