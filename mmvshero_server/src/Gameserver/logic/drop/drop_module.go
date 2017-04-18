package drop

import (
	"Gameserver/global"
	"Gameserver/logic"
	"Gameserver/logic/award"
	"common"
	"common/define"
	"common/protocol"
	"common/scheme"
	. "galaxy"

	"github.com/golang/protobuf/proto"
)

func InitDropModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_DropReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req_msg := &protocol.MsgDropReq{}
		err := proto.Unmarshal(msg, req_msg)
		if err != nil {
			LogError(err)
			return
		}

		var ret common.RetCode
		var can bool
		var needItemId int32
		var needItemNum int32
		var needGold int32
		var awardinfo []*protocol.AwardInfo
		drop := scheme.DropGet(req_msg.GetType())
		if drop == nil {
			ret = common.RetCode_SchemeData_Error
		} else {
			if req_msg.GetType() == common.DROP_ONE {
				needItemId = scheme.Commonmap[define.DropOneNeedItemId].Value
				needItemNum = scheme.Commonmap[define.DropOneNeedItemNum].Value
				needGold = scheme.Commonmap[define.DropOneNeedGold].Value
			} else {
				needItemId = scheme.Commonmap[define.DropTenNeedItemId].Value
				needItemNum = scheme.Commonmap[define.DropTenNeedItemNum].Value
				needGold = scheme.Commonmap[define.DropTenNeedGold].Value
			}

			if role.ItemIsEnough(needItemId, needItemNum) {
				role.ItemCost(needItemId, needItemNum, true)
				can = true
			} else {
				if role.IsEnoughGold(needGold) {
					role.CostGold(needGold, true, true)
					can = true
				}
			}

			if can {
				awardinfo, ret = award.Award(int32(drop.RandomAwardId), role, true)
				LogDebug(awardinfo)
			} else {
				ret = common.RetCode_RoleNotEnoughGold
			}
		}

		var ret_msg *protocol.MsgDropRet
		if len(awardinfo) == 0 {
			ret_msg = &protocol.MsgDropRet{
				Retcode: proto.Int32(int32(ret)),
			}
		} else {
			ret_msg = &protocol.MsgDropRet{
				Retcode: proto.Int32(int32(ret)),
				Infos:   awardinfo,
			}
		}

		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_DropRet), sid, buf)
	})
}
