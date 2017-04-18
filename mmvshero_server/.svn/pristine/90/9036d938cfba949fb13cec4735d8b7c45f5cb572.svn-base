package sign

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	. "galaxy"

	"github.com/golang/protobuf/proto"
)

func InitSignModule() {
	init_protocol()
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_SignInitReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		ret := &protocol.MsgSignInitRet{
			Info: role.SignInit(),
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_SignInitRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_SignInReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		retcode, info := role.SignIn()
		ret := &protocol.MsgSignInRet{
			Retcode: proto.Int32(int32(retcode)),
			Info:    info,
		}

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_SignInRet), sid, buf)
	})
}
