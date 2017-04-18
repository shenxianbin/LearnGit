package challenge

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitChallengeModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_ChallengeQueryReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgChallengeQueryReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.ChallengeQuery())
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ChallengeQueryRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ChallengeStartFightReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgChallengeStartFightReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.ChallengeStartFight(req.GetLayer()))
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ChallengeStartFightRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ChallengeFightResultReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgChallengeFightResultReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.ChallengeFightResult(req.GetLayer(), req.GetIsSuccess()))
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ChallengeFightResultRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ChallengeResetReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		buf, err := proto.Marshal(role.ChallengeReset())
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_ChallengeResetRet), sid, buf)
	})
}
