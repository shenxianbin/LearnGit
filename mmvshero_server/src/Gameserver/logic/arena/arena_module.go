package arena

import (
	"Gameserver/global"
	. "Gameserver/logic"
	. "Gameserver/logic/achievement"
	"Gameserver/logic/award"
	"common"
	. "common/cache"
	"common/define"
	"common/protocol"
	"common/scheme"
	"errors"
	"fmt"
	. "galaxy"
	. "galaxy/event"
	"galaxy/timer"
	"time"

	"github.com/golang/protobuf/proto"
)

const (
	cache_arena_sys_key       = "Arena"
	cache_arena_rank_key      = "{ArenaRank}"
	cache_arena_rank_copy_key = "{ArenaRank}Copy"
)

var arena_manager *ArenaManager

type ArenaManager struct {
	ArenaSystemCache
}

func (this *ArenaManager) Init() {
	this.ArenaSystemCache.BossAward1 = make([]*ArenaAward, 0)
	this.ArenaSystemCache.BossAward2 = make([]*ArenaAward, 0)
	this.ArenaSystemCache.BossAward3 = make([]*ArenaAward, 0)
	this.ArenaSystemCache.ShopAward = make([]*ArenaAward, 0)
}

func (this *ArenaManager) String() string {
	return fmt.Sprintf("BossVer : %v, BossId : %v, BossAward : %v,%v,%v, ShopVer : %v, ShopPanel : %v, ShopAward : %v",
		this.GetBossVersion(), this.GetBossId(), this.GetBossAward1(), this.GetBossAward2(), this.GetBossAward3(), this.GetShopVersion(), this.GetShopPanel(), this.GetShopAward())
}

func (this *ArenaManager) Load() error {
	LogDebug("Enter ArenaManager Load")
	resp, err := GxService().Redis().Cmd("GET", cache_arena_sys_key)
	if err != nil {
		LogError(err)
		return err
	}

	if buf, _ := resp.Bytes(); buf != nil {
		LogDebug("ArenaManager Load refresh")
		err := proto.Unmarshal(buf, &this.ArenaSystemCache)
		if err != nil {
			LogError(err)
			return err
		}
		LogDebug("ArenaManager Load refresh before : ", this)
		err = this.refreshBoss()
		if err != nil {
			LogError(err)
			return err
		}
		err = this.refreshShop()
		if err != nil {
			LogError(err)
			return err
		}
		LogDebug("ArenaManager Load refresh after : ", this)
	} else {
		LogDebug("ArenaManager Load init")
		err := this.init()
		if err != nil {
			LogError(err)
			return err
		}
	}

	return nil
}

func (this *ArenaManager) Save() {
	buf, err := proto.Marshal(&this.ArenaSystemCache)
	if err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SET", cache_arena_sys_key, buf); err != nil {
		LogFatal(err)
		return
	}
}

func (this *ArenaManager) genAward(awardId int32) []*ArenaAward {
	arenaAwards := make([]*ArenaAward, 0)
	for _, v := range award.AwardGen(awardId) {
		arenaAward := new(ArenaAward)
		arenaAward.SetType(v.GetType())
		arenaAward.SetCode(v.GetCode())
		arenaAward.SetAmount(v.GetAmount())
		arenaAwards = append(arenaAwards, arenaAward)
	}
	return arenaAwards
}

func (this *ArenaManager) genBossAward(boss_id int32) ([]*ArenaAward, []*ArenaAward, []*ArenaAward, error) {
	boss_scheme, has := scheme.ArenaBossmap[boss_id]
	if !has {
		return nil, nil, nil, errors.New("ArenaBoss Error")
	}

	return this.genAward(boss_scheme.Award1), this.genAward(boss_scheme.Award2), this.genAward(boss_scheme.Award3), nil
}

func (this *ArenaManager) genShopAward(panel int32) ([]*ArenaAward, error) {
	shop := scheme.ArenaShopGet(panel)
	if shop == nil {
		return nil, errors.New("ArenaShop Error")
	}
	shopAwards := make([]*ArenaAward, 0)
	for _, v := range shop {
		shopAward := this.genAward(v.ExchangeAward)
		shopAwards = append(shopAwards, shopAward...)
	}
	return shopAwards, nil
}

func (this *ArenaManager) clearRank() {
	_, err := GxService().Redis().Cmd("DEL", cache_arena_rank_key)
	if err != nil {
		LogError(err)
		return
	}
}

func (this *ArenaManager) init() error {
	boss_version := RefreshTime(21)
	this.ArenaSystemCache.SetBossVersion(boss_version)
	arenaBossAward1, arenaBossAward2, arenaBossAward3, err := this.genBossAward(1)
	if err != nil {
		return err
	}
	this.ArenaSystemCache.SetBossAward1(arenaBossAward1)
	this.ArenaSystemCache.SetBossAward2(arenaBossAward2)
	this.ArenaSystemCache.SetBossAward3(arenaBossAward3)
	this.ArenaSystemCache.SetBossId(1)
	this.clearRank()

	shop_version := RefreshTime(5)
	this.ArenaSystemCache.SetShopVersion(shop_version)
	panel := scheme.ArenaShopRandomPanel()
	shopAwards, err := this.genShopAward(panel)
	if err != nil {
		return err
	}
	this.ArenaSystemCache.SetShopAward(shopAwards)
	this.ArenaSystemCache.SetShopPanel(panel)
	this.Save()
	return nil
}

func (this *ArenaManager) copyRankAndAward() error {
	_, err := GxService().Redis().Cmd("ZUNIONSTORE", cache_arena_rank_copy_key, 1, cache_arena_rank_key)
	if err != nil {
		LogError(err)
		return err
	}

	go func() {
		resp, err := GxService().Redis().Cmd("ZREVRANGE", cache_arena_rank_copy_key, 0, -1, "WITHSCORES")
		if err != nil {
			LogError(err)
			return
		}

		top_rank_str, err := resp.List()
		if err != nil {
			LogError(err)
			return
		}

		if top_rank_str != nil && len(top_rank_str) != 0 {
			var rank int32
			for i := 0; i < len(top_rank_str); i = i + 2 {
				rank++
				uid, _ := parse_sortkey_to_id_name(top_rank_str[i])
				award_id := scheme.ArenaRankGet(rank)
				GxEvent().Execute(func(args ...interface{}) {
					LogDebug("Execute ArenaRankAward Uid : ", uid)
					if role := GetRoleByUid(uid); role != nil {
						award.Award(award_id, role, true)
					} else {
						resp, err := GxService().Redis().Cmd("GET", fmt.Sprintf(cache_arena_key_t, uid))
						if err != nil {
							LogError(err)
							return
						}

						if buf, _ := resp.Bytes(); buf != nil {
							arena := new(ArenaCache)
							err := proto.Unmarshal(buf, arena)
							if err != nil {
								return
							}
							awards := award.AwardGen(award_id)
							for _, v := range awards {
								if v.GetType() == common.RTYPE_ARENAPOINT {
									arena.SetPoint(arena.GetPoint() + v.GetAmount())
								}
							}

							newbuf, err := proto.Marshal(arena)
							if err != nil {
								LogFatal(err)
								return
							}

							if _, err := GxService().Redis().Cmd("SET", fmt.Sprintf(cache_arena_key_t, uid), newbuf); err != nil {
								LogFatal(err)
								return
							}
						}
					}

					//添加成就
					AchievementAddNumByUid(uid, 15, rank, true)
				})
				time.Sleep(time.Second)
			}
		}
		_, err = GxService().Redis().Cmd("DEL", cache_arena_rank_copy_key)
		if err != nil {
			LogError(err)
			return
		}
	}()

	return nil
}

func (this *ArenaManager) refreshBoss() error {
	change := false
	now := time.Now().Unix()
	if now-this.ArenaSystemCache.GetBossVersion() > int64(scheme.Commonmap[define.ArenaChangeBossTime].Value*86400-5) {
		version := RefreshTime(21)
		this.ArenaSystemCache.SetBossVersion(version)
		arenaBossAward1, arenaBossAward2, arenaBossAward3, err := this.genBossAward(this.ArenaSystemCache.GetBossId() + 1)
		if err != nil {
			arenaBossAward1, arenaBossAward2, arenaBossAward3, err = this.genBossAward(1)
			if err != nil {
				return err
			} else {
				this.ArenaSystemCache.SetBossId(1)
			}
		} else {
			this.ArenaSystemCache.SetBossId(this.ArenaSystemCache.GetBossId() + 1)
		}
		this.ArenaSystemCache.SetBossAward1(arenaBossAward1)
		this.ArenaSystemCache.SetBossAward2(arenaBossAward2)
		this.ArenaSystemCache.SetBossAward3(arenaBossAward3)
		this.clearRank()
		change = true
	}

	if change {
		this.Save()
	}
	return nil
}

func (this *ArenaManager) refreshShop() error {
	change := false
	now := time.Now().Unix()
	if now-this.ArenaSystemCache.GetShopVersion() > int64(scheme.Commonmap[define.ArenaChangeShopTime].Value*86400-5) {
		version := RefreshTime(5)
		this.ArenaSystemCache.SetShopVersion(version)
		panel := scheme.ArenaShopRandomPanel()
		shopAwards, err := this.genShopAward(panel)
		if err != nil {
			return err
		}
		this.ArenaSystemCache.SetShopAward(shopAwards)
		this.ArenaSystemCache.SetShopPanel(panel)
		change = true
	}

	if change {
		this.Save()
	}
	return nil
}

func InitArenaModule() error {
	arena_manager = new(ArenaManager)
	arena_manager.Init()
	err := arena_manager.Load()
	if err != nil {
		return err
	}

	init_protocol()

	timer.AddFixedTimerEvent(timer.FIXED_TIMER_TYPE_DAY, 3600*5, func() {
		err := arena_manager.refreshShop()
		if err != nil {
			LogError(err)
		}
		LogDebug("Refresh ArenaShop")
	})

	timer.AddFixedTimerEvent(timer.FIXED_TIMER_TYPE_DAY, 3600*21, func() {
		err := arena_manager.copyRankAndAward()
		if err != nil {
			LogError(err)
		}
		err = arena_manager.refreshBoss()
		if err != nil {
			LogError(err)
		}
		LogDebug("Refresh ArenaBoss")
	})
	return nil
}

func init_protocol() {
	global.RegisterMsg(int32(protocol.MsgCode_ArenaQueryReq), func(sid int64, msg []byte) {
		LogDebug("MsgCode_ArenaQueryReq")
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		ret := role.ArenaQuery()
		LogDebug("MsgCode_ArenaQueryRet :: uid:", role.GetUid(), " ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_ArenaQueryRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ArenaFightReq), func(sid int64, msg []byte) {
		LogDebug("MsgCode_ArenaFightReq")
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgArenaFightReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		ret := role.ArenaFight(req.GetScore(), req.GetIsCostOrder())
		LogDebug("MsgCode_ArenaFightRet :: uid:", role.GetUid(), " ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_ArenaFightRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ArenaShopQueryReq), func(sid int64, msg []byte) {
		LogDebug("MsgCode_ArenaShopQueryReq")
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		ret := role.ArenaShopQuery()
		LogDebug("MsgCode_ArenaShopQueryRet :: uid:", role.GetUid(), " ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_ArenaShopQueryRet), sid, buf)
	})

	global.RegisterMsg(int32(protocol.MsgCode_ArenaShopBuyReq), func(sid int64, msg []byte) {
		LogDebug("MsgCode_ArenaShopBuyReq")
		role := GetRoleBySid(sid)
		if role == nil {
			return
		}

		req := &protocol.MsgArenaShopBuyReq{}
		err := proto.Unmarshal(msg, req)
		if err != nil {
			LogError(err)
			return
		}

		ret := role.ArenaShopBuy(req.GetPos())
		LogDebug("MsgCode_ArenaShopBuyRet :: uid:", role.GetUid(), " ", ret)
		buf, err := proto.Marshal(ret)
		if err != nil {
			LogError(err)
			return
		}

		global.SendMsg(int32(protocol.MsgCode_ArenaShopBuyRet), sid, buf)
	})
}
