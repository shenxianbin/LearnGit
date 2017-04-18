package soldier

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitSoldierModule() {
	initProtocol()
}

func initProtocol() {
	// global.RegisterMsg(int32(protocol.MsgCode_SoldierAllReq), func(sid int64, msg []byte) {
	// 	role := logic.GetRoleBySid(sid)
	// 	if role == nil {
	// 		return
	// 	}

	// 	role.SoldierSendAll()
	// })

	global.RegisterMsg(int32(protocol.MsgCode_SoldierSkillLvReq), func(sid int64, msg []byte) {
		galaxy.LogDebug("enter SoldierSkillLvReq :")
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgSoldierSkillLvReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		retCode := 0
		if !role.SoldierSkillLevelUp(req.GetSoldierId(), req.GetSkillId()) {
			retCode = 1
		}
		ret := &protocol.MsgSoldierSkillLvRet{}
		ret.SetRetCode(int32(retCode))

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		galaxy.LogDebug("end SoldierSkillLvReq :")
		global.SendMsg(int32(protocol.MsgCode_SoldierSkillLvRet), sid, buf)
	})
}
