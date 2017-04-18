package chat

import (
	. "Gameserver/logic"
	"common"
	"common/define"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	cache_chat_key_t   = "Role:%v:Chat"
	cache_chat_world_t = "ChatWorld"
)

const (
	content_limit_per = 4
)

func genChatCacheKey(role_uid int64) string {
	return fmt.Sprintf(cache_chat_key_t, role_uid)
}

type ChatSys struct {
	owner            IRole
	cache_info_key   string
	query_cd_private int64
	query_cd_world   int64
	chat_cd          int64
}

func (this *ChatSys) Init(owner IRole) {
	this.owner = owner
	this.cache_info_key = genChatCacheKey(this.owner.GetUid())
}

func (this *ChatSys) Check() {
	_, err := GxService().Redis().Cmd("LTRIM", this.cache_info_key, 0, scheme.Commonmap[define.ChatPrivateShowNum].Value)
	if err != nil {
		LogFatal(err)
	}
}

func (this *ChatSys) ChatQuery(chat_type protocol.ChatType) (common.RetCode, []*protocol.ChatInfo) {
	switch chat_type {
	case protocol.ChatType_Private:
		now := time.Now().Unix()
		if this.query_cd_private > now {
			return common.RetCode_CoolDown, nil
		}
		this.query_cd_private = now + int64(scheme.Commonmap[define.ChatQueryLimitT].Value)

		resp, err := GxService().Redis().Cmd("LRANGE", this.cache_info_key, 0, scheme.Commonmap[define.ChatPrivateShowNum].Value)
		if err != nil {
			LogFatal(err)
			return common.RetCode_Fail, nil
		}

		chat_list, err := resp.List()
		if err != nil {
			LogFatal(err)
			return common.RetCode_Fail, nil
		}

		chat_info := make([]*protocol.ChatInfo, len(chat_list))
		for index, v := range chat_list {
			temp := new(protocol.ChatInfo)
			err := proto.Unmarshal([]byte(v), temp)
			if err != nil {
				LogError(err)
				return common.RetCode_Fail, nil
			}
			chat_info[index] = temp
		}
		return common.RetCode_Success, chat_info

	case protocol.ChatType_World:
		now := time.Now().Unix()
		if this.query_cd_world > now {
			return common.RetCode_CoolDown, nil
		}
		this.query_cd_world = now + int64(scheme.Commonmap[define.ChatQueryLimitT].Value)

		resp, err := GxService().Redis().Cmd("LRANGE", cache_chat_world_t, 0, scheme.Commonmap[define.ChatWorldShowNum].Value)
		if err != nil {
			LogFatal(err)
			return common.RetCode_Fail, nil
		}

		chat_list, err := resp.List()
		if err != nil {
			LogFatal(err)
			return common.RetCode_Fail, nil
		}

		chat_info := make([]*protocol.ChatInfo, len(chat_list))
		for index, v := range chat_list {
			temp := new(protocol.ChatInfo)
			err := proto.Unmarshal([]byte(v), temp)
			if err != nil {
				LogError(err)
				return common.RetCode_Fail, nil
			}
			chat_info[index] = temp
		}
		return common.RetCode_Success, chat_info
	}

	return common.RetCode_Fail, nil
}

func (this *ChatSys) Chat(chat_type protocol.ChatType, role_uid int64, content []byte) (common.RetCode, *protocol.ChatInfo) {
	now := time.Now().Unix()
	if this.chat_cd > now {
		return common.RetCode_CoolDown, nil
	}

	if len(content) > int(scheme.Commonmap[define.ChatWordsLimit].Value*content_limit_per) {
		return common.RetCode_Unable, nil
	}

	switch chat_type {
	case protocol.ChatType_Private:
		if role_uid <= 0 {
			return common.RetCode_Fail, nil
		}

		info := &protocol.ChatInfo{
			TimeStamp: proto.Int64(now),
			Content:   content,
		}

		buf, err := proto.Marshal(info)
		if err != nil {
			LogError(err)
			return common.RetCode_Fail, nil
		}

		_, err = GxService().Redis().Cmd("LPUSH", this.cache_info_key, buf)
		if err != nil {
			LogError(err)
			return common.RetCode_Fail, nil
		}

		_, err = GxService().Redis().Cmd("LPUSH", genChatCacheKey(role_uid), buf)
		if err != nil {
			LogError(err)
			return common.RetCode_Fail, nil
		}

		this.chat_cd = now + int64(scheme.Commonmap[define.ChatInterval].Value)
		return common.RetCode_Success, info

	case protocol.ChatType_World:
		if !this.owner.CostChatFreeTime(true, true) {
			if !this.owner.IsEnoughGold(scheme.Commonmap[define.ChatNeedGold].Value) {
				return common.RetCode_Unable, nil
			}
			this.owner.CostGold(scheme.Commonmap[define.ChatNeedGold].Value, true, true)
		}

		info := &protocol.ChatInfo{
			TimeStamp: proto.Int64(now),
			Content:   content,
		}

		buf, err := proto.Marshal(info)
		if err != nil {
			LogError(err)
			return common.RetCode_Fail, nil
		}

		_, err = GxService().Redis().Cmd("LPUSH", cache_chat_world_t, buf)
		if err != nil {
			LogError(err)
			return common.RetCode_Fail, nil
		}

		this.chat_cd = now + int64(scheme.Commonmap[define.ChatInterval].Value)
		return common.RetCode_Success, info
	}

	return common.RetCode_Fail, nil
}
