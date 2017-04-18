package role

import (
	"Gameserver/global"
	"Gameserver/logic"
	"Gameserver/logic/rolestate"
	"common/protocol"
	"common/static"
	"errors"
	"fmt"
	. "galaxy"
	"galaxy/define"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
)

func InitRoleModule() {
	init_protocol()
}

const (
	loginTokenKey   = "loginToken:%v"    //%v = token, value = roleId, valid time : tokenExpireTime
	accountTokenKey = "account:%v:token" //%v = roleId ,value = token
	ipKey           = "ipKey:%v"         //%v = role_uid
)

// 注册协议
func init_protocol() {
	GxService().RegisterConnectorLoseLogicCallBack("GateServer", GxService().Config().GetConnectorConfig(0).ConnectServerId(), func() {
		role_list := logic.GetAllRoleByUid()
		for _, role := range role_list {
			role.Offline()
		}

		logic.RemoveAll()
		LogDebug("GateServer lose connect, all role remove")
	})

	global.RegisterMsg(define.MSGCODE_LOSE, func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		logic.RemRoleByUid(role.GetUid())
		logic.RemRoleBySid(sid)

		role.Offline()
		LogDebug("Role[sid:", sid, "] [uid:", role.GetUid(), "] lose connect")
	})

	global.RegisterMsg(int32(protocol.MsgCode_LoginInReq), func(sid int64, msg []byte) {
		LogDebug("into LoginAuthReq:")
		req := &protocol.MsgLoginInReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		LogDebug("Role[sid:", sid, "] [uid:", req.GetTokenKey(), "] Begin LoginIn......")
		role, err := loginByToken(sid, req.GetTokenKey())
		if err != nil {
			LogError(err)
			return
		}
		LogDebug("Role[sid:", sid, "] [uid:", role.GetUid(), "] End LoginIn......")

		ret := &protocol.MsgLoginInRet{}
		ret.SetSystemTime(time.Now().Unix())
		ret.RoleInfo = role.FillRoleBaseInfo()
		ret.ItemListInfo = role.FillItemListInfo()
		ret.HeroCreateInfo = role.FillHeroCreateInfo()
		ret.HeroListInfo = role.FillHeroListInfo()
		ret.AllSoldiers = role.FillAllSoldiersInfo()
		ret.BuildingInfo = role.FillBuildingListInfo()
		ret.DecorationListInfo = role.FillDecorationListInfo()
		ret.KingInfo = role.FillKingInfo()
		ret.MapInfo = role.FillMapInfo()
		ret.PvpHeroList = role.PvpFightHeroList()

		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_LoginInRet), sid, buf)
		LogDebug("Role[sid:", sid, "] [uid:", role.GetUid(), "] Packet LoginIn......")
	})

	global.RegisterMsg(int32(protocol.MsgCode_RoleSetNicknameReq), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		reqMsg := &protocol.MsgRoleSetNicknameReq{}
		err := proto.Unmarshal(msg, reqMsg)
		if err != nil {
			return
		}

		retMsg := &protocol.MsgRoleSetNicknameRet{
			Retcode: proto.Int32(int32(role.SetNickname(reqMsg.GetNickname()))),
		}
		buf, err := proto.Marshal(retMsg)
		if err != nil {
			return
		}
		global.SendMsg(int32(protocol.MsgCode_RoleSetNicknameRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_RoleNewGuideUpdate), func(sid int64, msg []byte) {
		role := logic.GetRoleBySid(sid)
		if role == nil {
			return
		}

		reqMsg := &protocol.MsgRoleNewGuideUpdate{}
		err := proto.Unmarshal(msg, reqMsg)
		if err != nil {
			return
		}

		role.SetNewPlayerGuideStep(reqMsg.GetNewPlayerGuideStep(), true)
	})
}

//登录
func Login(sid int64, role_uid int64) (*Role, error) {
	//查询数据库
	resp, err := GxService().Redis().Cmd("GET", GenRoleCacheKey(role_uid))
	if err != nil {
		LogError(err)
		return nil, err
	}
	var role *Role
	roleBytes, _ := resp.Bytes()
	if roleBytes != nil {
		roleState, err := rolestate.NewRoleState(role_uid)
		if err != nil {
			LogError(err)
			return nil, err
		}

		if roleState.IsFightLock() {
			return nil, errors.New("role is be attacked")
		}

		role = NewRole()
		err = role.LoadAllInfo(roleBytes)
		if err != nil {
			LogError(err)
			return nil, err
		}

		role.sid = sid
		logic.AddRoleBySid(sid, role)
		logic.AddRoleByUid(role_uid, role)
		roleState.SetOnlineServer(global.GetServerId(), true)
		StaticRoleLogin(role)
		LogDebug("Role[", sid, "] LoadDB Success")
	} else {
		roleState, err := rolestate.NewRoleState(role_uid)
		if err != nil {
			LogError(err)
			return nil, err
		}

		role, err = CreateRole(role_uid)
		if err != nil {
			LogError(err)
			return nil, err
		}

		role.sid = sid

		logic.AddRoleBySid(sid, role)
		logic.AddRoleByUid(role_uid, role)
		roleState.SetOnlineServer(global.GetServerId(), true)
		StaticRoleCreate(role)
		LogDebug("Role[", sid, "] Create Success")
	}

	return role, nil
}

//凭据登录
func loginByToken(sid int64, token string) (*Role, error) {
	resp, err := GxService().Redis().Cmd("GET", fmt.Sprintf(loginTokenKey, token))
	if resp.IsNil() || err != nil {
		return nil, errors.New("not found")
	}

	var roleId int64
	buff, _ := resp.Str()
	temp, _ := strconv.Atoi(buff)
	roleId = int64(temp)

	return Login(sid, roleId)
}

func StaticRoleCreate(role *Role) {
	msg := &static.MsgStaticRoleCreate{}
	msg.SetRoleUid(role.GetUid())
	msg.SetLv(role.GetLv())
	msg.SetStone(role.GetStone())
	msg.SetGold(role.GetGold())
	msg.SetFreeGold(role.GetFreeGold())
	msg.SetTrophy(role.GetTrophy())
	msg.SetTotalCharge(0)
	msg.SetLastLoginTime(time.Now().Unix())

	resp, err := GxService().Redis().Cmd("GET", fmt.Sprintf(ipKey, role.GetUid()))
	if err != nil {
		LogError(err)
		msg.SetIp("")
	} else {
		ip, errip := resp.Str()
		if errip != nil {
			LogError(errip)
		}
		msg.SetIp(ip)
	}
	msg.SetCreateTime(time.Now().Unix())

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_RoleCreate), buf)
	LogDebug("StaticRoleCreate msg : ", msg)
}

func StaticRoleLogin(role *Role) {
	msg := &static.MsgStaticRoleLogin{}
	msg.SetRoleUid(role.GetUid())
	msg.SetLv(role.GetLv())
	msg.SetStone(role.GetStone())
	msg.SetGold(role.GetGold())
	msg.SetFreeGold(role.GetFreeGold())
	msg.SetTrophy(role.GetTrophy())
	msg.SetTotalCharge(0)
	msg.SetLastLoginTime(time.Now().Unix())

	resp, err := GxService().Redis().Cmd("GET", fmt.Sprintf(ipKey, role.GetUid()))
	if err != nil {
		LogError(err)
		msg.SetIp("")
	} else {
		ip, errip := resp.Str()
		LogDebug(ip)
		if errip != nil {
			LogError(errip)
		}
		msg.SetIp(ip)
	}
	msg.SetCreateTime(time.Now().Unix())

	buf, err := proto.Marshal(msg)
	if err != nil {
		LogError(err)
		return
	}
	global.SendToStaticServer(int32(static.MsgStaticCode_RoleLogin), buf)
}
