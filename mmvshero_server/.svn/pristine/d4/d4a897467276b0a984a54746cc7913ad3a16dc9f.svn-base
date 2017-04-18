package sign

import (
	. "Gameserver/cache"
	. "Gameserver/logic"
	"Gameserver/logic/award"
	"common"
	"common/protocol"
	"common/scheme"
	"fmt"
	. "galaxy"
	"github.com/golang/protobuf/proto"
	"time"
)

const (
	cache_sign_key_t = "Role:%v:Sign"
)

func genSignCacheKey(role_uid int64) string {
	return fmt.Sprintf(cache_sign_key_t, role_uid)
}

type SignSys struct {
	owner IRole
	SignCache
	cache_key string
}

func (this *SignSys) Init(owner IRole) {
	this.owner = owner
	this.cache_key = genSignCacheKey(this.owner.GetUid())
}

func (this *SignSys) Load() error {
	resp, err := GxService().Redis().Cmd("GET", this.cache_key)
	if err != nil {
		return err
	}

	if buf, _ := resp.Bytes(); buf != nil {
		err = proto.Unmarshal(buf, &this.SignCache)
		if err != nil {
			LogError(err)
			return err
		}
	}

	return nil
}

func (this *SignSys) Save() {
	buf, err := proto.Marshal(&this.SignCache)
	if err != nil {
		LogFatal(err)
		return
	}

	if _, err := GxService().Redis().Cmd("SET", this.cache_key, buf); err != nil {
		LogFatal(err)
		return
	}
}

func (this *SignSys) fillSignInfo(now_day int32) *protocol.SignInfo {
	msg := new(protocol.SignInfo)
	msg.SetYearMonth(this.SignCache.GetYearMonth())
	msg.SetCount(this.SignCache.GetCount())
	if now_day > this.SignCache.GetDayIndex() {
		msg.SetDayIndex(this.SignCache.GetCount() + 1)
	} else {
		msg.SetDayIndex(0)
	}

	return msg
}

func (this *SignSys) SignInit() *protocol.SignInfo {
	now := time.Now()
	year := int32(now.Year())
	month := int32(now.Month())
	day := int32(now.Day())

	year_month := int32(year*100 + month)
	if this.SignCache.GetYearMonth() != year_month {
		this.SignCache.SetYearMonth(year_month)
		this.SignCache.SetCount(0)
		this.SignCache.SetDayIndex(0)
		this.Save()
	}

	return this.fillSignInfo(day)
}

func (this *SignSys) SignIn() (common.RetCode, *protocol.SignInfo) {
	now := time.Now()
	year := int32(now.Year())
	month := int32(now.Month())
	day := int32(now.Day())

	year_month := int32(year*100 + month)
	if this.SignCache.GetYearMonth() != year_month {
		this.SignCache.SetYearMonth(year_month)
		this.SignCache.SetCount(0)
		this.SignCache.SetDayIndex(0)
		this.Save()
		return common.RetCode_TimeOut, this.fillSignInfo(day)
	} else {
		if this.SignCache.GetDayIndex() >= day {
			LogDebug(fmt.Sprintf("Sign Already YearMonth : %v, Count : %v, DayIndex : %v", this.SignCache.GetYearMonth(), this.SignCache.GetCount(), this.SignCache.GetDayIndex()))
			return common.RetCode_Unable, this.fillSignInfo(day)
		}

		this.SignCache.SetCount(this.SignCache.GetCount() + 1)
		this.SignCache.SetDayIndex(day)
	}

	sign_award := scheme.SignGet(month, this.SignCache.GetCount())
	_, ret := award.Award(sign_award.AwardId, this.owner, true)
	if ret != common.RetCode_Success {
		LogError("Award Fail : ", sign_award.AwardId)
	}
	if this.owner.IsVip() && sign_award.VipDouble == 1 {
		_, ret := award.Award(sign_award.AwardId, this.owner, true)
		if ret != common.RetCode_Success {
			LogError("Award Fail : ", sign_award.AwardId)
		}
	}
	this.Save()
	return common.RetCode_Success, this.fillSignInfo(day)
}

func (this *SignSys) FixVipSignIn() {
	sign_award := scheme.SignGet(this.SignCache.GetYearMonth()/100, this.SignCache.GetCount())
	if this.owner.IsVip() && sign_award.VipDouble == 1 {
		_, ret := award.Award(sign_award.AwardId, this.owner, true)
		if ret != common.RetCode_Success {
			LogError("Award Fail : ", sign_award.AwardId)
		}
	}
}
