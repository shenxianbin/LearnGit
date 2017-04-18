package gm

import (
	. "Gameserver/cache"
	. "Gameserver/logic"
	"common"
	"common/gm"
	. "galaxy"

	"github.com/golang/protobuf/proto"
)

type GmSys struct {
	owner IRole
}

func (this *GmSys) Init(owner IRole) {
	this.owner = owner
}

func (this *GmSys) GmProcess(is_notify bool) {
	resp, err := GxService().Redis().Cmd("SMEMBERS", gm.GenGmOrderListKey(this.owner.GetUid()))
	if err != nil {
		LogError(err)
		return
	}

	cacheKeys, err := resp.List()
	if err != nil {
		LogError(err)
		return
	}

	for _, key := range cacheKeys {
		resp, err := GxService().Redis().Cmd("GET", key)
		if err != nil {
			GxService().Redis().Cmd("SREM", gm.GenGmOrderListKey(this.owner.GetUid()), key)
			LogError(err)
			continue
		}

		if buf, _ := resp.Bytes(); buf != nil {
			gmOrder := new(GmCommandOrder)
			err := proto.Unmarshal(buf, gmOrder)
			if err != nil {
				continue
			}

			if gmOrder.GetOrderStatus() != int32(gm.OrderStatus_NoProcess) {
				continue
			}

			this.process(gmOrder, is_notify)

			newBuf, err := proto.Marshal(gmOrder)
			if err != nil {
				LogError(err)
				continue
			}

			if _, err := GxService().Redis().Cmd("SET", key, newBuf); err != nil {
				LogError(err)
				continue
			}
		}
	}
}

func (this *GmSys) process(order *GmCommandOrder, is_notify bool) {
	module := order.GetCommandModule()
	switch gm.CommandModule(module) {
	case gm.CommandModule_Role:
		this.processRole(order, is_notify)
	case gm.CommandModule_King:
		this.processKing(order, is_notify)
	case gm.CommandModule_Item:
		this.processItem(order, is_notify)
	case gm.CommandModule_Hero:
		this.processHero(order, is_notify)
	case gm.CommandModule_Soldier:
		this.processSoldier(order, is_notify)
	case gm.CommandModule_Building:
		this.processBuilding(order, is_notify)
	}
}

// 体力、精力、造兵点数，增加和减少（注：包括减少，减少类测试不够时的提示）
// 骷髅币数量，增加和减少
// 魔血数量调整，增加和减少
// 魔魂数量调整，增加和减少
// 玩家等级(经验)，增加
func (this *GmSys) processRole(order *GmCommandOrder, is_notify bool) {
	property := order.GetCommandProperty()
	if property == int32(gm.CommandProperty_Fix) {
		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := content.GetValue()

			switch gm.CommandOption(option) {
			case gm.CommandOption_Exp:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddExp(int32(value), is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Order:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddOrder(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditOrder(int32(value), is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Spirit:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddSpirit(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditSpirit(int32(value), is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Gold:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddGold(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditGold(int32(value), is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_FreeGold:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddFreeGold(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditFreeGold(int32(value), is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Blood:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddBlood(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditBlood(value, is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Soul:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddSoul(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditSoul(value, is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Stone:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					this.owner.AddStone(int32(value), is_notify, true)
				case gm.CommandOperation_Edit:
					this.owner.EditStone(int32(value), is_notify, true)
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}
		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	}
	order.SetOrderStatus(int32(gm.OrderStatus_Fail))
}

// 魔王等级，增加
// 魔王技能等级，增加
func (this *GmSys) processKing(order *GmCommandOrder, is_notify bool) {
	property := order.GetCommandProperty()
	if property == int32(gm.CommandProperty_Fix) {
		for _, content := range order.Content {
			value := content.GetValue()
			if value <= 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := content.GetValue()

			switch gm.CommandOption(option) {
			case gm.CommandOption_Lv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					if this.owner.KingAddLvGm(int32(value)) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_SkillLv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if this.owner.KingEditSkillLv(int32(value/1000), int32(value%1000)) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}
		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	}
	order.SetOrderStatus(int32(gm.OrderStatus_Fail))
}

// 扫荡券，数量增加
// 经验宝宝，数量增加
// 魔物蛋，数量增加
// 魔物、魔使进化套装部件增加
func (this *GmSys) processItem(order *GmCommandOrder, is_notify bool) {
	property := order.GetCommandProperty()
	if property == int32(gm.CommandProperty_New) {
		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		schemeId := int32(order.GetUidOrSchemeId())
		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := content.GetValue()

			switch gm.CommandOption(option) {
			case gm.CommandOption_Num:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if this.owner.ItemAdd(schemeId, int32(value), is_notify) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	} else if property == int32(gm.CommandProperty_Fix) {
		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		uid := order.GetUidOrSchemeId()
		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := content.GetValue()

			switch gm.CommandOption(option) {
			case gm.CommandOption_Num:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Add:
					if this.owner.ItemAddByUid(uid, int32(value), is_notify) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				case gm.CommandOperation_Edit:
					if this.owner.ItemFixNum(uid, int32(value), is_notify) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}
		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	}
	order.SetOrderStatus(int32(gm.OrderStatus_Fail))
}

// 添加魔使
// 修改魔使技能等级
func (this *GmSys) processHero(order *GmCommandOrder, is_notify bool) {
	property := order.GetCommandProperty()
	if property == int32(gm.CommandProperty_New) {
		schemeId := int32(order.GetUidOrSchemeId())
		lv := int32(1)
		rank := int32(1)

		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := content.GetValue()

			switch gm.CommandOption(option) {
			case gm.CommandOption_Lv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if value != 0 {
						lv = int32(value)
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Rank:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if value != 0 {
						rank = int32(value)
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		if this.owner.HeroObtain(schemeId, lv, rank, is_notify) != common.UID_FAILED {
			order.SetOrderStatus(int32(gm.OrderStatus_Success))
			return
		}
	} else if property == int32(gm.CommandProperty_Fix) {
		heroUid := order.GetUidOrSchemeId()
		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := int32(content.GetValue())

			switch gm.CommandOption(option) {
			case gm.CommandOption_Lv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if this.owner.HeroEditLv(heroUid, int32(value), is_notify) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Rank:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if this.owner.HeroEditRank(heroUid, int32(value), is_notify) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_SkillLv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					LogDebug("heroUid=", heroUid, " skillId=", value/1000, "lv=", value%1000)
					if this.owner.HeroEditSkillLv(heroUid, value/1000, value%1000, is_notify) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))

						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}
		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	}
	order.SetOrderStatus(int32(gm.OrderStatus_Fail))
}

// 添加魔物
// 修改魔物技能等级
func (this *GmSys) processSoldier(order *GmCommandOrder, is_notify bool) {
	property := order.GetCommandProperty()
	if property == int32(gm.CommandProperty_Fix) {
		id := int32(order.GetUidOrSchemeId())
		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := int32(content.GetValue())

			switch gm.CommandOption(option) {
			case gm.CommandOption_Lv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if !this.owner.SoldierEditLv(id, value) {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_SkillLv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if !this.owner.SoldierEditSkillLv(id, value/1000, value%1000) {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			case gm.CommandOption_Num:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if !this.owner.SoldierEditNum(id, value) {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}
		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	}
	order.SetOrderStatus(int32(gm.OrderStatus_Fail))
}

// 建筑等级，增加
func (this *GmSys) processBuilding(order *GmCommandOrder, is_notify bool) {
	property := order.GetCommandProperty()
	if property == int32(gm.CommandProperty_Fix) {
		buildingUid := order.GetUidOrSchemeId()
		for _, content := range order.Content {
			value := content.GetValue()
			if value < 0 {
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		for _, content := range order.Content {
			option := content.GetOption()
			operation := content.GetOperation()
			value := int32(content.GetValue())

			switch gm.CommandOption(option) {
			case gm.CommandOption_Lv:
				switch gm.CommandOperation(operation) {
				case gm.CommandOperation_Edit:
					if this.owner.BuildingEditLv(buildingUid, value, true) != common.RetCode_Success {
						order.SetOrderStatus(int32(gm.OrderStatus_Fail))
						return
					}
				default:
					order.SetOrderStatus(int32(gm.OrderStatus_Fail))
					return
				}
			default:
				order.SetOrderStatus(int32(gm.OrderStatus_Fail))
				return
			}
		}

		order.SetOrderStatus(int32(gm.OrderStatus_Success))
		return
	}
	order.SetOrderStatus(int32(gm.OrderStatus_Fail))
}