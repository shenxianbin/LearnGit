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

		retcode := role.HeroSkillLvUp(req_msg.GetHeroId(), req_msg.GetSkillId(), true)
		ret_msg := &protocol.MsgHeroSkillLvUpRet{
			Retcode: proto.Int32(int32(retcode)),
		}
		buf, err := proto.Marshal(ret_msg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_HeroSkillLvUpRet), sid, buf)
	})
}
