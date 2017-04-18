package main

import (
	. "Loginserver/cache"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	. "galaxy/db/redis"
	. "galaxy/logs"
	"galaxy/utils"
	"github.com/golang/protobuf/proto"
	"github.com/vharitonsky/iniflags"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	roleAutoKey          = "Role:AutoKey"
	anonymousAutoKey     = "anonymous:AutoKey"
	anonymousAccountName = "anonymous:%v"

	accountKey = "account:%v" //%v= roleId, anonymous account

	accountThirdpartTokenKey = "accountToken:%v:%v" //$v = source, token , value = account
	usernameKey              = "accountName:%v"     //%v=username, value = roleId

	loginTokenKey   = "loginToken:%v"    //%v = token, value = roleId, valid time : tokenExpireTime
	accountTokenKey = "account:%v:token" //%v = roleId ,value = tokenExpireTime
	ipKey           = "ipKey:%v"         //%v = role_uid

	tokenExpireTime = 300 //second

	token = "1234567"
)

type response struct {
	Status     int    `json:"status"`
	LoginToken string `json:"loginToken"`
	Server     string `json:"server"`
	Port       int    `json:"port"`
}

type responseAccount struct {
	Status   int    `json:"status"`
	Username string `json:"username"`
}

type errorCode int32

const (
	ERROR_NONE              errorCode = 0
	ERROR_USERNAME_IS_EMPTY errorCode = 17000 + iota
	ERROR_PASSWORD_IS_EMPTY
	ERROR_NOTFOUND
	ERROR_PASSWORD_IS_NOT_MATCH
	ERROR_USERNAME_IS_EXIST
	ERROR_CREATE_ACCOUNT_FAIL
	ERROR_DB
	ERROR_UNKNOWN
)

type Account struct {
	AccountCache
}

func (this *Account) save() error {
	buff, err := proto.Marshal(&this.AccountCache)
	if err != nil {
		return err
	}

	if len(this.SourceToken) != 0 && len(this.Source) != 0 {
		if _, err = gxRedis.Cmd("SET", fmt.Sprintf(accountThirdpartTokenKey, this.Source, this.SourceToken), buff); err != nil {
			return err
		}
	} else {
		if _, err = gxRedis.Cmd("SET", fmt.Sprintf(accountKey, this.RoleId), buff); err != nil {
			return err
		}
	}

	return nil
}

func (this *Account) init() error {
	resp, err := gxRedis.Cmd("INCR", roleAutoKey)
	if err != nil {
		return err
	}

	this.RoleId, _ = resp.Int64()
	this.CreateTime = Time()
	return nil
}

func getAccount(roleId int64) *Account {
	account := &Account{}
	resp, err := gxRedis.Cmd("GET", fmt.Sprintf(accountKey, roleId))
	if resp.IsNil() || err != nil {
		return account
	}

	if buff, _ := resp.Bytes(); buff != nil {
		err = proto.Unmarshal(buff, &account.AccountCache)
		if err != nil {
			return account
		}
	}
	return account
}

func getAccountByToken(source, sourceToken string) *Account {
	account := &Account{}
	resp, err := gxRedis.Cmd("GET", fmt.Sprintf(accountThirdpartTokenKey, source, sourceToken))
	if resp.IsNil() || err != nil {
		return account
	}

	if buff, _ := resp.Bytes(); buff != nil {
		err = proto.Unmarshal(buff, &account.AccountCache)
		if err != nil {
			return account
		}
	}
	return account
}

//用户名登录 device 设备识别号
func login(username, device, ip string) (string, errorCode) {
	if len(username) == 0 {
		return "", ERROR_USERNAME_IS_EMPTY
	}

	var roleId int64
	resp, err := gxRedis.Cmd("GET", fmt.Sprintf(usernameKey, username))
	if resp.IsNil() || err != nil {
		return "", ERROR_NOTFOUND
	}

	buff, _ := resp.Str()
	temp, _ := strconv.Atoi(buff)

	roleId = int64(temp)
	if roleId == 0 {
		return "", ERROR_NOTFOUND
	}

	account := getAccount(roleId)
	if account.RoleId == 0 {
		return "", ERROR_NOTFOUND
	}

	account.LoginTime = Time()
	account.save()

	//返回登录凭据
	loginToken, err1 := account.generateLoginToken(device)
	if err1 != ERROR_NONE {
		return "", err1
	}

	gxRedis.Cmd("SET", fmt.Sprintf(ipKey, roleId), ip)

	return loginToken, ERROR_NONE
}

//数据校验
func verifySign(values url.Values, token string) bool {
	oldSign := values.Get("sign")
	values.Del("sign")

	fmt.Println(values.Encode())

	values.Add("token", token)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(values.Encode())))

	values.Add("sign", sign)
	values.Del("token")
	fmt.Println(values.Encode())
	return sign == oldSign
}

//生成签名
func addSign(values url.Values, token string) {
	values.Add("token", token)
	sign := fmt.Sprintf("%x", md5.Sum([]byte(values.Encode())))
	values.Add("sign", sign)
	values.Del("token")
}

//登录
func thirdpartLogin(source, sourceToken, device, ip string) (string, errorCode) {
	GxLogDebug("login:", source, sourceToken, device, ip)
	account := &Account{}
	account.Source = source
	account.SourceToken = sourceToken

	buff, err := proto.Marshal(&account.AccountCache)
	if err != nil {
		return "", ERROR_DB
	}

	key := fmt.Sprintf(accountThirdpartTokenKey, source, sourceToken)
	resp, err := gxRedis.Cmd("GET", key)
	if err != nil {
		GxLogDebug("err:", err)
		return "", ERROR_DB
	}

	if resp.IsNil() {
		resp, err = gxRedis.Cmd("SET", key, buff)
		if err != nil {
			GxLogDebug("err:", err)
			return "", ERROR_DB
		}

		str, _ := resp.Str()
		if str == "OK" {
			//新建
			if err := account.init(); err != nil {
				return "", ERROR_DB
			}

			account.save()
		}
	} else {
		if buff, _ := resp.Bytes(); buff != nil {
			err = proto.Unmarshal(buff, &account.AccountCache)
			if err != nil {
				return "", ERROR_DB
			}
		}
	}

	//返回登录凭据
	loginToken, err1 := account.generateLoginToken(device)
	if err1 != ERROR_NONE {
		return "", err1
	}

	gxRedis.Cmd("SET", fmt.Sprintf(ipKey, account.RoleId), ip)
	return loginToken, ERROR_NONE
}

func createAnonymousAccount() (string, errorCode) {
	resp, err := gxRedis.Cmd("INCR", anonymousAutoKey)
	if err != nil {
		return "", ERROR_DB
	}

	temp, _ := resp.Int64()
	username := fmt.Sprintf(anonymousAccountName, temp)

	if err := enroll(username); err != ERROR_NONE {
		return "", err
	}

	return username, ERROR_NONE
}

//创建匿名账号
func enroll(username string) errorCode {
	if len(username) == 0 {
		return ERROR_USERNAME_IS_EMPTY
	}

	resp, err := gxRedis.Cmd("SETNX", fmt.Sprintf(usernameKey, username), 0)
	if err != nil {
		return ERROR_UNKNOWN
	}
	if buff, _ := resp.Int(); buff == 0 {
		return ERROR_USERNAME_IS_EXIST
	}

	account := &Account{}
	if err := account.init(); err != nil {
		return ERROR_CREATE_ACCOUNT_FAIL
	}

	account.Username = username
	account.save()

	if _, err := gxRedis.Cmd("SET", fmt.Sprintf(usernameKey, username), account.RoleId); err != nil {
		return ERROR_UNKNOWN
	}

	return ERROR_NONE
}

//生成登录凭据
func (this *Account) generateLoginToken(device string) (string, errorCode) {
	//是否清除之前登录凭据?
	loginToken := MD5(fmt.Sprintf("%v:%s:%s", this, Rand(1, 2176782336), device))

	if _, err := gxRedis.Cmd("SETEX", fmt.Sprintf(loginTokenKey, loginToken), tokenExpireTime, this.RoleId); err != nil {
		return "", ERROR_DB
	}

	if _, err := gxRedis.Cmd("SETEX", fmt.Sprintf(accountTokenKey, this.RoleId), tokenExpireTime, loginToken); err != nil {
		return "", ERROR_DB
	}

	return loginToken, ERROR_NONE
}

func handlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	username, err := createAnonymousAccount()

	res, err1 := json.Marshal(responseAccount{Username: username, Status: int(err)})
	if err1 != nil {
		GxLogError(err1)
	}

	GxLogDebug(string(res))
	fmt.Fprintf(w, string(res))
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	loginToken, err := login(r.PostFormValue("username"), r.PostFormValue("device"), r.RemoteAddr)

	res, err1 := json.Marshal(response{LoginToken: loginToken, Status: int(err), Server: *server, Port: *port})
	if err1 != nil {
		GxLogError(err1)
	}

	GxLogDebug(string(res))
	fmt.Fprintf(w, string(res))
}

func handlerThirdpartLogin(w http.ResponseWriter, r *http.Request) {
	//todo 数据检查
	loginToken, err := thirdpartLogin(r.PostFormValue("source"), r.PostFormValue("sourceToken"), r.PostFormValue("device"), r.RemoteAddr)

	res, err1 := json.Marshal(response{LoginToken: loginToken, Status: int(err), Server: *server, Port: *port})
	if err1 != nil {
		GxLogError(err1)
	}

	GxLogDebug(string(res))
	fmt.Fprintf(w, string(res))
}

var gxRedis *GxRedis = NewGxRedis()
var server *string
var port *int

func main() {
	defer utils.Stack()

	listeningPort := flag.String("listeningPort", "8081", "listening port")
	redisPath := flag.String("redisPath", "192.168.1.220:6380", "redis path")
	redisPassword := flag.String("redisPassword", "", "redis password")
	server = flag.String("server", "192.168.1.220", "gate server ip")
	port = flag.Int("port", 10000, "gate server port")

	iniflags.Parse()
	gxRedis.Run(*redisPath, *redisPassword)

	//启动web服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		GxLogInfo(r.RemoteAddr, r.Method, r.URL, "\""+r.UserAgent()+"\"")

		switch strings.TrimRight(r.URL.Path, "/") {
		case "/anonymous_login":
			handlerLogin(w, r)
		case "/create_anonymous_account":
			handlerCreateAccount(w, r)
		case "/thirdpart_login":
			handlerThirdpartLogin(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	GxLogDebug("listening port: ", *listeningPort)
	GxLogDebug("connect to redis: ", *redisPath)
	GxLogDebug("gate server ip: ", *server)

	GxLogFatal(http.ListenAndServe(fmt.Sprintf(":%s", *listeningPort), nil))
}
