package models

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"net/url"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type midasConfig struct {
	AppId             string
	AppKey            string
	Format            string
	GameKey           string
	IsSandbox         bool
	SandboxServerName string
	ProductServerName string
	Protocol          string
	GameNotify        string
}

var MidasConfig *midasConfig

func init() {
	configfilepath := filepath.Join("conf", "midas.conf")
	iniconf, err := config.NewConfig("ini", configfilepath)
	if err != nil {
		logs.Critical(err)
	}
	MidasConfig = new(midasConfig)
	MidasConfig.AppId = iniconf.String("appId")
	MidasConfig.AppKey = iniconf.String("appKey")
	MidasConfig.Format = iniconf.String("format")
	MidasConfig.GameKey = iniconf.String("gameKey")
	MidasConfig.IsSandbox, err = iniconf.Bool("isSandbox")
	if err != nil {
		logs.Critical(err)
	}
	MidasConfig.ProductServerName = iniconf.String("productServerName")
	MidasConfig.SandboxServerName = iniconf.String("sandboxServerName")
	MidasConfig.GameNotify = iniconf.String("GameNotify")
	MidasConfig.Protocol = iniconf.String("protocol")
}
func MakeSign(method string, url string, params url.Values, secret string) string {
	data := MakeSource(method, url, params)
	logs.Debug("data=", data)
	key := []byte(secret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	logs.Debug("sign=", sign)
	return sign
}

func MakeSource(method string, urlpath string, params url.Values) string {
	var buf bytes.Buffer
	buf.WriteString(strings.ToUpper(method))
	buf.WriteString("&")
	buf.WriteString(url.QueryEscape(urlpath))
	buf.WriteString("&")
	queryString := EncodeValues(params)
	queryString = strings.Replace(queryString, "~", "%7E", -1)
	buf.WriteString(queryString)
	return buf.String()
}

func EncodeValues(params url.Values) string {
	if params == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		prefix := k + "="
		buf.WriteString(prefix)
		buf.WriteString(params.Get(k))
	}
	return url.QueryEscape(buf.String())
}
func VerifySig(method string, scriptName string, params url.Values, secret string, sig string) bool {
	params.Del("sig")
	encodeParams := url.Values{}
	for key, _ := range params {
		encodeParams.Add(key, encodeValue(params.Get(key)))
	}
	sign := MakeSign(method, scriptName, encodeParams, secret)
	decode, err := url.QueryUnescape(sig)
	if err != nil {
		logs.Error(err)
	}
	return sign == decode
}

func encodeValue(value string) string {
	var buf bytes.Buffer
	for i := 0; i < len(value); i++ {
		c := value[i]
		flag, err := regexp.Match(`[0-9A-Za-z\!\*\(\)]{1,1}`, []byte{c})
		if err != nil {
			fmt.Println(err)
			return ""
		}
		if flag {
			buf.WriteByte(c)
		} else {
			buf.WriteString("%")
			buf.WriteString(fmt.Sprintf("%02X", int(c)))
		}
	}
	return buf.String()
}
