package friend

import (
	"Gameserver/global"
	"Gameserver/logic"
	"common/protocol"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

func InitFriendModule() {
	initProtocol()
}

func initProtocol() {
	global.RegisterMsg(int32(protocol.MsgCode_FriendAllReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendAllReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.FriendAll())
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_FriendAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendSearchReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			galaxy.LogDebug("role == nil")
			return
		}

		req := &protocol.MsgFriendSearchReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		friend, retCode := role.FriendSearch(req.GetAlias())

		ret := &protocol.MsgFriendSearchRet{}
		ret.SetRetCode(int32(retCode))
		ret.SetFriend(friend)

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_FriendSearchRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendDeleteReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendDeleteReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgFriendDeleteRet{}
		retCode := role.FriendDelete(req.GetFriendIds())

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendDeleteRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendSendExcitationReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendSendExcitationReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgFriendSendExcitationRet{}
		affectedRows := role.FriendSendExcitation(req.GetFriendIds())

		ret.SetAffectedRows(affectedRows)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendSendExcitationRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendUseExcitationReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendUseExcitationReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgFriendUseExcitationRet{}
		total, affectedRows := role.FriendUseExcitation(req.GetFriendIds())

		ret.SetTotal(total)
		ret.SetAffectedRows(affectedRows)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendUseExcitationRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendSavePvpResultReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendSavePvpResultReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgFriendSavePvpResultRet{}
		record := string(req.GetRecord())
		retCode := role.FriendSavePvpResult(req.GetFriendId(), req.GetAttackerWin(), record)

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendSavePvpResultRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendRequestAllReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendRequestAllReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		buf, err := proto.Marshal(role.FriendRequestAll())
		if err != nil {
			galaxy.LogError(err)
			return
		}
		global.SendMsg(int32(protocol.MsgCode_FriendRequestAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendRequestAddReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendRequestAddReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgFriendRequestAddRet{}
		retCode := role.FriendRequestAdd(req.GetFriendId())

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendRequestAddRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendRequestDealWithReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendRequestDealWithReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		ret := &protocol.MsgFriendRequestDealWithRet{}
		retCode := role.FriendRequestDealWith(req.GetFriendId(), req.GetIsAgreed())

		ret.SetRetCode(int32(retCode))
		ret.SetIsAgreed(req.GetIsAgreed())
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendRequestDealWithRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendInviteAddIdReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgFriendInviteAddIdReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		ret := &protocol.MsgFriendInviteAddIdRet{}
		retCode := role.FriendInviteAddId(req.GetInviteId())

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_FriendInviteAddIdRet), sid, buf)
	})
}
