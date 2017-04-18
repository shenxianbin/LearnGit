package friend

import (
	"Gameserver/global"
	. "Gameserver/logic"
	. "Gameserver/logic/achievement"
	. "Gameserver/logic/award"
	. "common"
	. "common/cache"
	d "common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	"galaxy"

	"github.com/golang/protobuf/proto"
)

const (
	friendkey  = "{Role:%v}:Friend:%v"
	inviteeKey = "{Role:%v}:Invitee:%v"
	requestkey = "{Role:%v}:Request:%v"

	//set
	friendskey  = "Role:%v:Friends"
	inviteesKey = "Role:%v:Invitees"
	requestskey = "Role:%v:Requests"

	inviteIdKey = "Role:%v:InviteId"

	receiveExcitationKey = "ReceiveExcitation:%v"
)

type Friend struct {
	id                    int64 //好友id
	selfId                int64 //自身id
	winTimes              int32
	loseTimes             int32
	sendExcitationTime    int64 //发送激励的时间
	receiveExcitationTime int64 //接收激励的时间
	useExcitationTime     int64 //使用激励的时间
	friendInfo            *FriendInfo
}

//cache the under data every 30m
type FriendInfo struct {
	id          int64
	loginTime   int64 //last login time
	level       int32
	leagueLevel int32
	trophy      int32
	timestamp   int64 //last refresh time
	nickname    string
}

type Request struct {
	id         int64 //user id
	selfId     int64 //自身id
	timestamp  int64 //create time
	friendInfo *FriendInfo
}

type Invitee struct {
	id     int64 //user id
	selfId int64 //自身id
	level  int32
}

type UserFriends struct {
	user     IRole
	Friends  map[int64]*Friend //key user Id
	Requests map[int64]*Request
	Invitees map[int64]*Invitee
	InviteId int64
	*ReceiveExcitationCache
}

func (this *UserFriends) Init(user IRole) {
	this.user = user
	this.Friends = make(map[int64]*Friend)
	this.Requests = make(map[int64]*Request)
	this.Invitees = make(map[int64]*Invitee)
	this.ReceiveExcitationCache = &ReceiveExcitationCache{}
}

func (this *UserFriends) Load() error {
	err := this.reloadFriends()
	if err != nil {
		galaxy.LogError(err)
		return err
	}

	err = this.reloadRequests()
	if err != nil {
		galaxy.LogError(err)
		return err
	}

	err = this.reloadInvitees()
	if err != nil {
		galaxy.LogError(err)
		return err
	}

	err = this.updateInviteesLv()
	if err != nil {
		galaxy.LogError(err)
		return err
	}

	this.loadInviteId()
	this.loadReceiveExcitationTimes()
	this.checkReceiveExcitationTimes()

	return nil
}

func (this *UserFriends) getCount(friendId int64) int32 {
	//scard
	resp, err := galaxy.GxService().Redis().Cmd("SCARD", fmt.Sprintf(friendskey, friendId))
	if err != nil {
		galaxy.LogError(err)
		return 0
	}
	count, _ := resp.Int()
	return int32(count)
}

func (this *UserFriends) FriendInviteAddId(inviteId int64) RetCode {
	if this.InviteId != 0 {
		return RetCode_FriendInviteIdHasExist
	}

	//todo 检查inviteId合法性？

	if inviteId != 0 {
		this.InviteId = inviteId
		this.saveInviteId()
		return RetCode_Success
	}
	return RetCode_FriendInviteIdError
}

//添加好友请求 done
func (this *UserFriends) FriendRequestAdd(friendId int64) RetCode {
	//是否已经添加有好友请求
	if _, ok := this.Requests[friendId]; ok {
		return RetCode_FriendRequestHasExist
	}

	if this.checkRequestIsExist(friendId, this.user.GetUid()) {
		return RetCode_FriendRequestHasExist
	}

	//判断是否已经是好友
	if _, ok := this.Friends[friendId]; ok {
		return RetCode_FriendHasBeenGot
	}

	//检查是否存在这个用户
	if this.user.OfflineRoleBase(friendId) == nil {
		return RetCode_FriendSearchNotFound
	}

	//在对方请求列表中添加好友请求
	request := &Request{}
	request.id = this.user.GetUid()
	request.selfId = friendId
	request.timestamp = Time()
	request.friendInfo = this.newFriendInfo(this.user)
	request.save()

	if role := GetRoleByUid(friendId); role != nil {
		role.FriendRequestNotify(request.toProtocol())
	}

	return RetCode_Success
}

//处理好友请求 done
func (this *UserFriends) FriendRequestDealWith(friendId int64, isAgreed bool) RetCode {
	if request, ok := this.Requests[friendId]; ok {
		if false == isAgreed {
			delete(this.Requests, friendId)
			request.delete()
			return RetCode_Success
		}

		//判断是否大于各自好友上限
		limit := scheme.Commonmap[d.FriendNumLimit].Value
		if int32(len(this.Friends)) >= limit || this.getCount(friendId) >= limit {
			return RetCode_FriendHasReachedTheMaximumNumberOfFriends
		}

		//添加好友
		myFriend := &Friend{}
		myFriend.friendInfo = request.friendInfo.copy()
		myFriend.id = friendId
		myFriend.selfId = this.user.GetUid()
		myFriend.save()
		this.Friends[friendId] = myFriend

		//对方添加
		asFriend := &Friend{}
		asFriend.id = this.user.GetUid()
		asFriend.selfId = friendId
		asFriend.friendInfo = this.newFriendInfo(this.user)
		asFriend.save()

		//删除好友请求
		delete(this.Requests, friendId)
		request.delete()

		//添加成就
		//自身
		this.user.AchievementAddNum(17, int32(len(this.Friends)), true)
		//对方
		AchievementAddNumByUid(friendId, 17, this.getCount(friendId), true)

		if role := GetRoleByUid(friendId); role != nil {
			role.FriendReload(friendId)
			role.FriendAddNotify(asFriend.toProtocol())
		}

		return RetCode_Success
	}
	return RetCode_FriendRequestNotFound
}

//删除好友 双向删除
func (this *UserFriends) FriendDelete(friendIds []int64) RetCode {
	if len(friendIds) == 0 {
		return RetCode_FriendNotFound
	}

	//待删除好友列表
	myFriendDelete := []string{}
	asFriendDelete := []int64{}
	i := 0
	for _, friendId := range friendIds {
		myFriendDelete = append(myFriendDelete, fmt.Sprintf(friendkey, this.user.GetUid(), friendId))
		asFriendDelete = append(asFriendDelete, friendId)
		i++
	}
	//删除自己的
	this.delete(myFriendDelete)

	//遍历删除对方的好友列表
	for _, friendId := range asFriendDelete {
		asFriend := &Friend{}
		asFriend.id = this.user.GetUid()
		asFriend.selfId = friendId
		galaxy.LogDebug("friendId:", friendId)
		asFriend.delete()

		if role := GetRoleByUid(friendId); role != nil {
			role.FriendReload(friendId)
			role.FriendDeleteNotify(friendId)
		}
	}

	return RetCode_Success
}

//发送奖励 返回获得激励的好友个数
func (this *UserFriends) FriendSendExcitation(friendIds []int64) int32 {
	var count int32 = 0
	for _, friendId := range friendIds {
		myFriend, ok := this.Friends[friendId]
		if !ok {
			galaxy.LogError("can not send excitation non-friends:", this.user.GetUid(), friendId)
			continue
		}

		//判断是否发送过
		if myFriend.sendExcitationTime >= RefreshTime(d.SysResetTime) {
			continue
		}

		myFriend.sendExcitationTime = Time()
		myFriend.save()

		asFriend := &Friend{}
		asFriend.id = this.user.GetUid()
		asFriend.selfId = friendId
		if asFriend.read() != nil {
			galaxy.LogDebug("not found:", this.user.GetUid(), friendId)
			continue
		}

		asFriend.receiveExcitationTime = Time()
		asFriend.useExcitationTime = 0
		asFriend.save()

		if role := GetRoleByUid(friendId); role != nil {
			role.FriendReload(friendId)
			role.FriendExcitationNotify(friendId)
		}

		count++
	}
	return count
}

//接收激励 总数，影响数
func (this *UserFriends) FriendUseExcitation(friendIds []int64) (int32, int32) {
	this.checkReceiveExcitationTimes()
	var times int32
	for _, friendId := range friendIds {
		if friend, ok := this.Friends[friendId]; ok {
			//是否大于上限
			limit := scheme.Commonmap[d.FriendReceiveTime].Value
			if this.ReceiveExcitationCache.GetTimes() >= limit {
				galaxy.LogDebug("count can not large than limit:", this.GetTimes(), limit)
				return this.ReceiveExcitationCache.GetTimes(), 0
			}

			//过期的
			if friend.receiveExcitationTime < RefreshTime(d.SysResetTime) {
				continue
			}

			if friend.useExcitationTime != 0 {
				continue
			}

			//接收奖励
			Award(scheme.Commonmap[d.FriendGiveAward].Value, this.user, true)
			friend.useExcitationTime = Time()
			friend.save()

			this.ReceiveExcitationCache.SetTimes(this.ReceiveExcitationCache.GetTimes() + 1)
			galaxy.LogDebug("this.GetTimes():", this.ReceiveExcitationCache.GetTimes())

			times++
		}
	}

	if times > 0 {
		this.saveReceiveExcitationTimes()
	}

	galaxy.LogDebug("this.GetTimes()", this.ReceiveExcitationCache.GetTimes())
	return this.ReceiveExcitationCache.GetTimes(), times
}

//保存好友战斗结果
func (this *UserFriends) FriendSavePvpResult(friendId int64, attackerWin bool, record string) RetCode {
	//保存自己的
	if friend, ok := this.Friends[friendId]; ok {
		self := &Friend{}
		self.id = this.user.GetUid()
		self.selfId = friendId
		if err := self.read(); err != nil {
			galaxy.LogError(err)
			return RetCode_Failed
		}
		if attackerWin {
			self.winTimes++
			friend.loseTimes++
		} else {
			self.loseTimes++
			friend.winTimes++
		}
		self.save()
		friend.save()

		//todo 保存战斗录像

		return RetCode_Success
	}

	return RetCode_FriendNotFound
}

func (this *UserFriends) FriendSearch(alias string) (*protocol.Friend, RetCode) {
	roleId := this.user.GetRoleIdByAlias(alias)
	if roleId == 0 {
		return nil, RetCode_FriendSearchNotFound
	}

	role := this.user.OfflineRoleBase(roleId)
	myFriend := &Friend{}
	myFriend.friendInfo = &FriendInfo{}
	myFriend.friendInfo = this.newFriendInfo(role)
	myFriend.id = roleId
	return myFriend.toProtocol(), RetCode_Success
}

func (this *UserFriends) FriendAll() *protocol.MsgFriendAllRet {
	this.reloadFriends()
	this.updateInviteesLv()

	list1 := make([]*protocol.Friend, len(this.Friends))

	var i int32 = 0
	for _, v := range this.Friends {
		list1[i] = v.toProtocol()
		i++
	}

	m := new(protocol.MsgFriendAllRet)
	m.SetFriends(list1)

	this.checkReceiveExcitationTimes()
	m.SetReceiveExcitationTimes(this.GetTimes())
	m.SetReceiveExcitationTimestamp(this.GetTimestamp())
	return m
}

func (this *UserFriends) FriendRequestAll() *protocol.MsgFriendRequestAllRet {
	this.reloadRequests()
	// galaxy.LogDebug("this.Request:", this.Requests)

	list1 := make([]*protocol.FriendRequest, len(this.Requests))

	var i int32 = 0
	for _, v := range this.Requests {
		list1[i] = v.toProtocol()
		i++
	}

	m := new(protocol.MsgFriendRequestAllRet)
	m.SetRequests(list1)
	return m
}

//好友请求通知
func (this *UserFriends) FriendRequestNotify(request *protocol.FriendRequest) {
	msg := &protocol.MsgFriendRequestNotify{}

	msg.SetRequest(request)
	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	global.SendMsg(int32(protocol.MsgCode_FriendRequestNotify), this.user.GetSid(), buf)
}

//添加好友通知
func (this *UserFriends) FriendAddNotify(friend *protocol.Friend) {
	msg := &protocol.MsgFriendAddNotify{}
	msg.SetFriend(friend)

	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	global.SendMsg(int32(protocol.MsgCode_FriendAddNotify), this.user.GetSid(), buf)
}

//删除好友通知
func (this *UserFriends) FriendDeleteNotify(friendUid int64) {
	msg := &protocol.MsgFriendDeleteNotify{}

	msg.SetFriendId(friendUid)
	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	global.SendMsg(int32(protocol.MsgCode_FriendDeleteNotify), this.user.GetSid(), buf)
}

//收到好友激励通知
func (this *UserFriends) FriendExcitationNotify(friendUid int64) {
	msg := &protocol.MsgFriendExcitationNotify{}

	msg.SetFriendId(friendUid)
	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	global.SendMsg(int32(protocol.MsgCode_FriendExcitationNotify), this.user.GetSid(), buf)
}

//重新载入好友信息
func (this *UserFriends) FriendReload(friendUid int64) {
	key := fmt.Sprintf(friendkey, this.user.GetUid(), friendUid)
	resp, err := RedisCmd("GET", key)
	if err != nil {
		galaxy.LogError(err)
		return
	}
	if buff, _ := resp.Bytes(); buff != nil {
		cache := &FriendCache{}
		err = proto.Unmarshal(buff, cache)
		friend := &Friend{}
		friend.readCache(cache)
		this.Friends[cache.GetId()] = friend
	} else {
		RedisCmd("SREM", fmt.Sprintf(friendskey, this.user.GetUid()), key)
	}
}
