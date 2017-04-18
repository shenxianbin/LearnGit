package hero

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"

	"github.com/golang/protobuf/proto"
)

func InitHeroModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_HeroCreateFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode := role.HeroCreateFinish(true)
		ret_msg := &protocol.MsgHeroCreateFinishRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroCreateFinishRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroCreateShockReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode := role.HeroCreateShock(true)
		ret_msg := &protocol.MsgHeroCreateShockRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroCreateShockRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroCreateGiveUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode := role.HeroCreateGiveUp(true)
		ret_msg := &protocol.MsgHeroCreateGiveUpRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroCreateGiveUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroSkillLvUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req_msg := &protocol.MsgHeroSkillLvUpReq{}
		err := proto.Unmarshal(msg, req_msg)
		if err != nil {
			return
		}

		retcode := role.HeroSkillLvUp(req_msg.GetHeroUid(), req_msg.GetSkillId(), true)
		ret_msg := &protocol.MsgHeroSkillLvUpRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroSkillLvUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroEvoReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req_msg := &protocol.MsgHeroEvoReq{}
		err := proto.Unmarshal(msg, req_msg)
		if err != nil {
			return
		}

		retcode := role.HeroEvoStart(req_msg.GetHeroUid(), req_msg.GetNeedHeroUid(), req_msg.GetUseMoney() != 0, true)
		ret_msg := &protocol.MsgHeroEvoRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroEvoRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroEvoFinishReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req_msg := &protocol.MsgHeroEvoFinishReq{}
		err := proto.Unmarshal(msg, req_msg)
		if err != nil {
			return
		}

		retcode := role.HeroEvoFinish(req_msg.GetUseMoney() != 0, true)
		ret_msg := &protocol.MsgHeroEvoFinishRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroEvoFinishRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroEvoSpeedUpReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode := role.HeroEvoSpeedUp(true)
		ret_msg := &protocol.MsgHeroEvoSpeedUpRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroEvoSpeedUpRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_HeroMixReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req_msg := &protocol.MsgHeroMixReq{}
		err := proto.Unmarshal(msg, req_msg)
		if err != nil {
			return
		}

		retcode := role.HeroMix(req_msg.GetTargetUid(), req_msg.GetUids(), true)
		ret_msg := &protocol.MsgHeroMixRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroMixRet), sid, buf)
	})
}
