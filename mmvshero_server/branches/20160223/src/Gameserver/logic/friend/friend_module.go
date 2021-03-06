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
		// galaxy.LogDebug("enter MsgCode_FriendAllReq -------")
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

		ret := role.FriendAll()
		// galaxy.LogDebug("FriendAll: ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		// galaxy.LogDebug("end MsgCode_FriendAllReq -------")
		global.SendMsg(int32(protocol.MsgCode_FriendAllRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_FriendSearchReq), func(sid int64, msg []byte) {
		// galaxy.LogDebug("enter FriendSearchReq----------")
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

		// galaxy.LogDebug("req:", req.GetAlias())
		friend, retCode := role.FriendSearch(req.GetAlias())

		// galaxy.LogDebug("retCode:", retCode)
		ret := &protocol.MsgFriendSearchRet{}
		ret.SetRetCode(int32(retCode))
		ret.SetFriend(friend)

		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		// galaxy.LogDebug("end FriendSearchReq----------")
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
		affectedRows := role.FriendUseExcitation(req.GetFriendIds())

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
		// galaxy.LogDebug("enter MsgCode_FriendRequestAllReq -------")
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

		ret := role.FriendRequestAll()
		// galaxy.LogDebug("FriendRequestAll: ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		// galaxy.LogDebug("end MsgCode_FriendRequestAllReq -------")
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
		//galaxy.LogDebug("enter MsgCode_FriendRequestDealWithReq -------")

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

		//galaxy.LogDebug("req.GetFriendId(), req.GetIsAgreed():", req.GetFriendId(), req.GetIsAgreed())
		ret := &protocol.MsgFriendRequestDealWithRet{}
		retCode := role.FriendRequestDealWith(req.GetFriendId(), req.GetIsAgreed())

		ret.SetRetCode(int32(retCode))
		buf, err := proto.Marshal(ret)
		if err != nil {
			galaxy.LogError(err)
			return
		}
		//galaxy.LogDebug("retCode:", retCode)
		//galaxy.LogDebug("end MsgCode_FriendRequestDealWithReq -------")

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
