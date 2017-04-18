package item

import (
	"Gameserver/global"
	. "Gameserver/logic"
	// . "Gameserver/logic/award"
	"common"
	. "common/cache"
	"common/protocol"
	"common/scheme"
	"errors"
	"fmt"
	. "galaxy"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
)

const (
	cache_item_autokey_t = "Role:%v:ItemAutoKey"
	cache_itemlist_key_t = "Role:%v:Item"
	cache_itemobj_key_t  = "Role:%v:Item:%v"
)

func GenItemListKey(role_uid int64) string {
	return fmt.Sprintf(cache_itemlist_key_t, role_uid)
}

func GenItemCacheKey(role_uid int64, item_uid int64) string {
	return fmt.Sprintf(cache_itemobj_key_t, role_uid, item_uid)
}

func genItemAutoKey(role_uid int64) string {
	return fmt.Sprintf(cache_item_autokey_t, role_uid)
}

type Item struct {
	ItemCache
	scheme_data *scheme.Item
}

func NewItem(scheme_id int32, num int32, role_uid int64) (*Item, error) {
	item_scheme, has := scheme.Itemmap[scheme_id]
	if !has {
		return nil, errors.New("NewItem Scheme Error")
	}

	resp, err := GxService().Redis().Cmd("INCR", genItemAutoKey(role_uid))
	if err != nil {
		return nil, err
	}

	obj := new(Item)
	uid, _ := resp.Int64()
	obj.ItemCache.SetUid(uid)
	obj.ItemCache.SetSchemeId(scheme_id)
	obj.ItemCache.SetNum(num)
	obj.scheme_data = item_scheme
	return obj, nil
}

func LoadItem(buf []byte) (*Item, error) {
	obj := new(Item)
	err := proto.Unmarshal(buf, &obj.ItemCache)
	if err != nil {
		return nil, err
	}

	item_scheme, has := scheme.Itemmap[obj.GetSchemeId()]
	if !has {
		return nil, errors.New("LoadItem Scheme Error")
	}
	obj.scheme_data = item_scheme
	return obj, nil
}

func (this *Item) GetType() int32 {
	return this.scheme_data.ItemType
}

func (this *Item) GetHeapLimit() int32 {
	return this.scheme_data.HeapLimit
}

func (this *Item) GetLv() int32 {
	return this.scheme_data.Lv
}

func (this *Item) GetUseage() int32 {
	return this.scheme_data.Useage
}

func (this *Item) GetValue() []int32 {
	value_str := strings.Split(this.scheme_data.Value, ";")
	values := make([]int32, len(value_str))
	for index, v := range value_str {
		value, _ := strconv.Atoi(v)
		values[index] = int32(value)
	}
	return values
}

func (this *Item) FillItemInfo() *protocol.ItemInfo {
	msg := new(protocol.ItemInfo)
	msg.SetUid(this.ItemCache.GetUid())
	msg.SetSchemeId(this.ItemCache.GetSchemeId())
	msg.SetNum(this.ItemCache.GetNum())
	return msg
}

type ItemSys struct {
	owner          IRole
	item_list      map[int64]*Item
	cache_list_key string
}

func (this *ItemSys) Init(owner IRole) {
	this.owner = owner
	this.item_list = make(map[int64]*Item)
	this.cache_list_key = fmt.Sprintf(cache_itemlist_key_t, this.owner.GetUid())
}

func (this *ItemSys) Load() error {
	resp, err := GxService().Redis().Cmd("SMEMBERS", this.cache_list_key)
	if err != nil {
		return err
	}

	cacheKeys, err := resp.List()
	if err != nil {
		return err
	}

	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", this.cache_list_key, key)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			item, err := LoadItem(buf)
			if err != nil {
				LogFatal(err)
				continue
			}
			this.item_list[item.GetUid()] = item
		}
	}

	return nil
}

func (this *ItemSys) Save(item *Item) {
	buf, err := proto.Marshal(&item.ItemCache)
	if err != nil {
		LogFatal(err)
		return
	}

	key := GenItemCacheKey(this.owner.GetUid(), item.ItemCache.GetUid())
	if _, err := GxService().Redis().Cmd("SET", key, buf); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SADD", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
}

func (this *ItemSys) Del(item_uid int64) {
	key := GenItemCacheKey(this.owner.GetUid(), item_uid)
	if _, err := GxService().Redis().Cmd("DEL", key); err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SREM", this.cache_list_key, key); err != nil {
		LogFatal(err)
		return
	}
	delete(this.item_list, item_uid)
}

func (this *ItemSys) FillItemListInfo() *protocol.ItemListInfo {
	msg := new(protocol.ItemListInfo)
	msg.ItemList = make([]*protocol.ItemInfo, len(this.item_list))
	index := 0
	for _, v := range this.item_list {
		msg.ItemList[index] = v.FillItemInfo()
		index++
	}
	return msg
}

//考虑是否发包
func (this *ItemSys) ItemAdd(scheme_id int32, num int32, is_notify bool) common.RetCode {
	if num <= 0 {
		return common.RetCode_ItemNumArgsError
	}

	item_scheme, has := scheme.Itemmap[scheme_id]
	if !has {
		return common.RetCode_SchemeData_Error
	}

	var total_num int32
	var found_item *Item
	for _, found_item = range this.item_list {
		if found_item.ItemCache.GetSchemeId() == scheme_id {
			total_num += found_item.ItemCache.GetNum()
			break
		}
	}

	if total_num == 0 {
		item, err := NewItem(scheme_id, num, this.owner.GetUid())
		if err != nil {
			LogFatal(err)
			return common.RetCode_Failed
		}
		this.item_list[item.GetUid()] = item
		this.Save(item)
		if is_notify {
			this.send_updateinfo(item.FillItemInfo())
		}
	} else {
		found_item.SetNum(found_item.GetNum() + num)
		if found_item.GetNum() > item_scheme.HeapLimit {
			found_item.SetNum(item_scheme.HeapLimit)
		}
		this.Save(found_item)
		if is_notify {
			this.send_updateinfo(found_item.FillItemInfo())
		}
	}

	return common.RetCode_Success
}

func (this *ItemSys) ItemAddByUid(uid int64, num int32, is_notify bool) common.RetCode {
	if num <= 0 {
		return common.RetCode_ItemNumArgsError
	}

	if item, has := this.item_list[uid]; !has {
		return common.RetCode_ItemUidError
	} else {
		item.SetNum(item.GetNum() + num)
		if item.GetNum() > item.scheme_data.HeapLimit {
			item.SetNum(item.scheme_data.HeapLimit)
		}
		this.Save(item)
		if is_notify {
			this.send_updateinfo(item.FillItemInfo())
		}
	}
	return common.RetCode_Success
}

func (this *ItemSys) ItemCost(scheme_id int32, num int32, is_notify bool) common.RetCode {
	if num <= 0 {
		return common.RetCode_ItemNumArgsError
	}

	var total_num int32
	var found_item *Item
	for _, found_item = range this.item_list {
		if found_item.ItemCache.GetSchemeId() == scheme_id {
			total_num += found_item.ItemCache.GetNum()
			break
		}
	}

	if total_num < num {
		return common.RetCode_ItemNotEnough
	}

	found_item.SetNum(found_item.GetNum() - num)
	if found_item.GetNum() == 0 {
		this.Del(found_item.GetUid())
	} else {
		this.Save(found_item)
	}

	if is_notify {
		this.send_updateinfo(found_item.FillItemInfo())
	}

	return common.RetCode_Success
}

func (this *ItemSys) ItemCostByUid(uid int64, num int32, is_notify bool) common.RetCode {
	if num <= 0 {
		return common.RetCode_Success
	}

	if item, has := this.item_list[uid]; !has {
		return common.RetCode_ItemUidError
	} else {
		if item.GetNum() < num {
			return common.RetCode_ItemNotEnough
		}

		item.SetNum(item.GetNum() - num)
		if item.GetNum() == 0 {
			this.Del(item.GetUid())
		} else {
			this.Save(item)
		}
		if is_notify {
			this.send_updateinfo(item.FillItemInfo())
		}
	}
	return common.RetCode_Success
}

func (this *ItemSys) ItemFixNum(uid int64, num int32, is_notify bool) common.RetCode {
	if num <= 0 {
		return common.RetCode_Success
	}

	if item, has := this.item_list[uid]; !has {
		return common.RetCode_ItemUidError
	} else {
		item.SetNum(num)
		if item.GetNum() == 0 {
			this.Del(item.GetUid())
		} else {
			this.Save(item)
		}
		if is_notify {
			this.send_updateinfo(item.FillItemInfo())
		}
	}
	return common.RetCode_Success
}

func (this *ItemSys) ItemIsEnough(scheme_id int32, num int32) bool {
	if num <= 0 {
		return true
	}

	var total_num int32
	var found_item *Item
	for _, found_item = range this.item_list {
		if found_item.ItemCache.GetSchemeId() == scheme_id {
			total_num += found_item.ItemCache.GetNum()
			break
		}
	}

	return total_num >= num
}

func (this *ItemSys) ItemIsEnoughByUid(uid int64, num int32) bool {
	if num <= 0 {
		return true
	}

	if item, has := this.item_list[uid]; !has {
		return false
	} else {
		return item.GetNum() >= num
	}
}

func (this *ItemSys) ItemGet(item_uid int64) IItem {
	if item, has := this.item_list[item_uid]; has {
		return item
	}
	return nil
}

func (this *ItemSys) ItemUse(items []*protocol.ItemUseInfo, user_type protocol.ItemUserType, user_id int64) common.RetCode {
	LogDebug("Item:", items, " user_type:", user_type, " user_id", user_id)
	LogDebug(this.item_list)

	if len(items) == 0 {
		return common.RetCode_ItemUseArgsLenNull
	}

	for _, v := range items {
		item, has := this.item_list[v.GetUid()]
		if !has || item.GetNum() < v.GetNum() || item.GetUseage() == -1 {
			LogError("ItemUse args valid error")
			return common.RetCode_ItemUseArgsError
		}
	}

	switch user_type {
	case protocol.ItemUserType_HERO:
		if !this.owner.HeroFind(int32(user_id)) {
			LogError("ItemUse Hero Null")
			return common.RetCode_ItemUseHeroIdError
		}
	case protocol.ItemUserType_SOLDIER:
		if this.owner.SoldierGetInCamp(int32(user_id)) == nil {
			LogError("ItemUse Soldier Null")
			return common.RetCode_ItemUseSoldierIdError
		}
	}

	for _, v := range items {
		item, _ := this.item_list[v.GetUid()]
		LogDebug("ItemUse item_uid (", item.GetUid(), ") item_scheme_id (", item.GetSchemeId(), ") item_useage (", item.GetUseage(), ")")
		num, ret := ItemUseageInstance().Invoker(item.scheme_data.Useage, v.GetNum(), item, this.owner, user_type, user_id)
		if ret == common.RetCode_Success {
			this.ItemCostByUid(v.GetUid(), num, true)
		} else {
			LogError("ItemUse Error")
			return ret
		}
	}

	return common.RetCode_Success
}

func (this *ItemSys) ItemSell(uid int64, num int32, is_notify bool) common.RetCode {
	if num <= 0 {
		return common.RetCode_Success
	}

	if item, has := this.item_list[uid]; !has {
		return common.RetCode_ItemUidError
	} else {
		if item.GetNum() < num {
			return common.RetCode_ItemNotEnough
		}

		item.SetNum(item.GetNum() - num)
		if item.GetNum() == 0 {
			this.Del(item.GetUid())
		} else {
			this.Save(item)
		}

		this.owner.AddSoul(item.scheme_data.Price*num, true, true)

		if is_notify {
			this.send_updateinfo(item.FillItemInfo())
		}
	}
	return common.RetCode_Success
}

func (this *ItemSys) send_updateinfo(info *protocol.ItemInfo) {
	msg := &protocol.MsgItemInfoUpdateNotify{}
	msg.SetInfos(info)

	buf, err := proto.Marshal(msg)
	if err != nil {
		return
	}
	global.SendMsg(int32(protocol.MsgCode_ItemInfoUpdateNotify), this.owner.GetSid(), buf)
}

var item_useage *ItemUseage

func ItemUseageInstance() *ItemUseage {
	if item_useage == nil {
		item_useage = new(ItemUseage)
		item_useage.Init()
	}
	return item_useage
}

type ItemFunc func(use_num int32, item *Item, role IRole, user_type protocol.ItemUserType, user_id int64) (curCost int32, retcode common.RetCode)

type ItemUseage struct {
	useage_list map[int32]ItemFunc
}

//调注册方法
func (this *ItemUseage) Init() {
	this.useage_list = make(map[int32]ItemFunc)
	this.Register(common.ITEM_USEAGE_HEROCHANGE, func(use_num int32, item *Item, role IRole, user_type protocol.ItemUserType, user_id int64) (curCost int32, retcode common.RetCode) {
		value_str := strings.Split(item.scheme_data.Value, ";")
		if len(value_str) != 1 {
			LogDebug("Item_args len error")
			return 0, common.RetCode_SchemeData_Error
		}

		exp, err := strconv.Atoi(value_str[0])
		if err != nil {
			LogDebug("Item_args content error : ", value_str)
			return 0, common.RetCode_SchemeData_Error
		}

		if ret := role.HeroAddRank(int32(user_id), int32(exp)*use_num, true); ret != common.RetCode_Success {
			return 0, ret
		}

		return use_num, common.RetCode_Success
	})

	this.Register(common.ITEM_USEAGE_LVEXP, func(use_num int32, item *Item, role IRole, user_type protocol.ItemUserType, user_id int64) (curCost int32, retcode common.RetCode) {
		value_str := strings.Split(item.scheme_data.Value, ";")
		if len(value_str) != 1 {
			return 0, common.RetCode_SchemeData_Error
		}

		exp, err := strconv.Atoi(value_str[0])
		if err != nil {
			return 0, common.RetCode_SchemeData_Error
		}
		switch user_type {
		case protocol.ItemUserType_HERO:
			if ret := role.HeroLvUp(int32(user_id), item.GetSchemeId(), int32(exp)*use_num, true); ret != common.RetCode_Success {
				return 0, ret
			}
		case protocol.ItemUserType_SOLDIER:
			if !role.SoldierLevelUp(int32(user_id), item.GetSchemeId(), int32(exp)*use_num) {
				return 0, common.RetCode_Failed
			}
		default:
			return 0, common.RetCode_Failed
		}
		return use_num, common.RetCode_Success
	})
}

func (this *ItemUseage) Register(itemUseageType int32, f ItemFunc) {
	this.useage_list[itemUseageType] = f
}

func (this *ItemUseage) Invoker(itemUseageType int32, use_num int32, item *Item, role IRole, user_type protocol.ItemUserType, user_id int64) (curCost int32, retcode common.RetCode) {
	if function, has := this.useage_list[itemUseageType]; has {
		return function(use_num, item, role, user_type, user_id)
	}
	return 0, common.RetCode_ItemUseageTypeError
}
