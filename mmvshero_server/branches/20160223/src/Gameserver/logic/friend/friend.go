package friend

import (
	. "Gameserver/logic"
	. "Gameserver/logic/award"
	"common"
	d "common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	"galaxy"
	"strconv"
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
}

func (this *UserFriends) Init(user IRole) {
	this.user = user
	this.Friends = make(map[int64]*Friend)
	this.Requests = make(map[int64]*Request)
	this.Invitees = make(map[int64]*Invitee)
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

func (this *UserFriends) FriendInviteAddId(inviteId int64) common.RetCode {
	if this.InviteId == 0 && inviteId != 0 {
		this.InviteId = inviteId
		this.saveInviteId()
		return common.RetCode_Success
	}
	return common.RetCode_Fail
}

//添加好友请求 done
func (this *UserFriends) FriendRequestAdd(friendId int64) common.RetCode {
	//是否已经添加有好友请求
	if _, ok := this.Requests[friendId]; ok {
		return common.RetCode_Fail
	}

	//判断是否已经是好友
	if _, ok := this.Friends[friendId]; ok {
		return common.RetCode_Fail
	}

	//检查是否存在这个用户
	if this.user.OfflineRoleBase(friendId) == nil {
		galaxy.LogError("role not exist.", friendId)
		return common.RetCode_Fail
	}

	//在对方请求列表中添加好友请求
	request := &Request{}
	request.id = this.user.GetUid()
	request.selfId = friendId
	request.timestamp = Time()
	// request.friendInfo = &FriendInfo{}
	request.friendInfo = this.newFriendInfo(this.user)
	request.save()
	return common.RetCode_Success
}

//处理好友请求 done
func (this *UserFriends) FriendRequestDealWith(friendId int64, isAgreed bool) common.RetCode {
	if request, ok := this.Requests[friendId]; ok {
		if false == isAgreed {
			delete(this.Requests, friendId)
			request.delete()
			return common.RetCode_Success
		}

		//判断是否大于各自好友上限
		limit := scheme.Commonmap[d.FriendNumLimit].Value
		if int32(len(this.Friends)) >= limit || this.getCount(friendId) >= limit {
			galaxy.LogDebug("int32(len(this.Friends)) >= limit || this.getCount(friendId) >= limit :")
			galaxy.LogDebug(int32(len(this.Friends)), limit, this.getCount(friendId))
			return common.RetCode_Fail
		}

		//添加好友
		myFriend := &Friend{}
		myFriend.friendInfo = request.friendInfo.copy()
		myFriend.id = friendId
		myFriend.selfId = this.user.GetUid()
		// galaxy.LogDebug("myFriend:", myFriend.id, myFriend.selfId, request.id)
		if myFriend.save() == false {
			galaxy.LogError("myFriend.save() == false ")
		}
		this.Friends[friendId] = myFriend

		//对方添加
		asFriend := &Friend{}
		asFriend.id = this.user.GetUid()
		asFriend.selfId = friendId
		asFriend.friendInfo = this.newFriendInfo(this.user)
		// galaxy.LogDebug("asFriend:", asFriend.id, asFriend.selfId)
		if asFriend.save() == false {
			galaxy.LogError("asFriend.save() == false ")
		}

		//删除好友请求
		delete(this.Requests, friendId)
		request.delete()

		return common.RetCode_Success
	}
	// galaxy.LogDebug("not found friendId:", friendId, this.Requests)
	return common.RetCode_Fail
}

//删除好友 双向删除 done
func (this *UserFriends) FriendDelete(friendIds []int64) common.RetCode {
	if len(friendIds) == 0 {
		return common.RetCode_Fail
	}

	//待删除好友列表
	myFriendDelete := []string{}
	asFriendDelete := []string{}
	i := 0
	for _, friendId := range friendIds {
		// myFriendDelete[i] = fmt.Sprintf(friendkey, this.user.GetUid(), friendId)
		myFriendDelete = append(myFriendDelete, fmt.Sprintf(friendkey, this.user.GetUid(), friendId))
		// asFriendDelete[i] = fmt.Sprintf(friendkey, friendId, this.user.GetUid())
		asFriendDelete = append(asFriendDelete, fmt.Sprintf(friendkey, friendId, this.user.GetUid()))
		i++
	}
	//删除自己的
	this.delete(myFriendDelete)

	//遍历删除对方的好友列表
	for _, friendId := range asFriendDelete {
		asFriend := &Friend{}
		asFriend.id = this.user.GetUid()
		temp, _ := strconv.Atoi(friendId)
		asFriend.selfId = int64(temp)
		asFriend.delete()
	}

	return common.RetCode_Success
}

//发送奖励 返回获得激励的好友个数 done
func (this *UserFriends) FriendSendExcitation(friendIds []int64) int32 {
	var count int32 = 0
	for _, friendId := range friendIds {
		myFriend, ok := this.Friends[friendId]
		if !ok {
			galaxy.LogError("can not send excitation non-friends:", this.user.GetUid(), friendId)
			continue
		}

		//判断是否发送过
		if myFriend.sendExcitationTime >= RefreshTime(5) {
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

		count++
	}
	return count
}

//接收激励 done
func (this *UserFriends) FriendUseExcitation(friendIds []int64) int32 {
	var count int32 = 0
	for _, friendId := range friendIds {
		if friend, ok := this.Friends[friendId]; ok {
			//是否大于上限
			limit := scheme.Commonmap[d.FriendReceiveTime].Value
			if count >= limit {
				galaxy.LogDebug("count can not large than limit")
				return 0
			}

			//过期的
			if friend.receiveExcitationTime < RefreshTime(5) {
				continue
			}

			if friend.useExcitationTime != 0 {
				continue
			}

			//接收奖励
			Award(scheme.Commonmap[d.FriendGiveAward].Value, this.user, true)
			friend.useExcitationTime = Time()
			friend.save()

			count++
		}
	}
	return count
}

//保存好友战斗结果
func (this *UserFriends) FriendSavePvpResult(friendId int64, attackerWin bool, record string) common.RetCode {
	//保存自己的
	if friend, ok := this.Friends[friendId]; ok {
		self := &Friend{}
		self.id = this.user.GetUid()
		self.selfId = friendId
		if err := self.read(); err != nil {
			galaxy.LogError(err)
			return common.RetCode_Fail
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

		return common.RetCode_Success
	}

	return common.RetCode_Fail
}

func (this *UserFriends) FriendSearch(alias string) (*protocol.Friend, common.RetCode) {
	roleId := this.user.GetRoleIdByAlias(alias)
	if roleId == 0 {
		return nil, common.RetCode_Fail
	}

	// galaxy.LogDebug("roleId:", roleId)

	role := this.user.OfflineRoleBase(roleId)
	myFriend := &Friend{}
	myFriend.friendInfo = &FriendInfo{}
	myFriend.friendInfo = this.newFriendInfo(role)
	myFriend.id = roleId
	return myFriend.toProtocol(), common.RetCode_Success
}

func (this *UserFriends) FriendAll() *protocol.MsgFriendAllRet {
	this.reloadFriends()
	this.updateInviteesLv()
	// galaxy.LogDebug("this.Friends:", this.Friends)

	list1 := make([]*protocol.Friend, len(this.Friends))

	var i int32 = 0
	for _, v := range this.Friends {
		list1[i] = v.toProtocol()
		i++
	}

	m := new(protocol.MsgFriendAllRet)
	m.SetFriends(list1)
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
