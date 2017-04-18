package alias

import (
	. "Gameserver/logic"
	"fmt"
	"galaxy"
	"strconv"
	"strings"
)

const (
	roleKey  = "Role:%v:Alias" // value alias string
	aliasKey = "Alias:%v:Role" //value Role id int64
)

type Alias struct {
	user  IRole
	alias string
}

func (this *Alias) Init(user IRole) {
	this.user = user
	this.Load()
}

func (this *Alias) Load() error {
	this.alias = this.GetAliasByRoleId(this.user.GetUid())

	if this.alias == "" {
		this.GetAlias()
	}

	return nil
}

func (this *Alias) GetAlias() string {
	if this.alias != "" {
		return this.alias
	}

	for i := 0; i < 10; i++ {
		aliasString := generate()
		resp, err := RedisCmd("setnx", fmt.Sprintf(aliasKey, aliasString), this.user.GetUid())
		if err != nil {
			galaxy.LogError(err)
			return "" //没有
		}
		if buf, _ := resp.Int(); buf == 0 {
			continue
		}

		RedisCmd("set", fmt.Sprintf(roleKey, this.user.GetUid()), aliasString)
		this.alias = aliasString
		break
	}

	return this.alias
}

func (this *Alias) GetAliasByRoleId(roleId int64) string {
	resp, _ := RedisCmd("get", fmt.Sprintf(roleKey, roleId))
	str, _ := resp.Str()
	return str
}

func (this *Alias) GetRoleIdByAlias(alias string) int64 {
	resp, _ := RedisCmd("get", fmt.Sprintf(aliasKey, strings.ToUpper(alias)))
	roleId, _ := resp.Int64()
	return roleId
}

func generate() string {
	return strings.ToUpper(fmt.Sprintf("%06s", strconv.FormatInt(Rand(1, 2176782336), 36)))
}
