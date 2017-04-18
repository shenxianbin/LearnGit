package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type JsonConfig struct {
}

func (js *JsonConfig) Parse(filename string) (ConfigContainer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return js.ParseData(content)
}

func (js *JsonConfig) ParseData(data []byte) (ConfigContainer, error) {
	x := &JsonConfigContainer{
		data: make(map[string]interface{}),
	}
	err := json.Unmarshal(data, &x.data)
	if err != nil {
		var wrappingArray []interface{}
		err2 := json.Unmarshal(data, &wrappingArray)
		if err2 != nil {
			return nil, err
		}
		x.data["rootArray"] = wrappingArray
	}
	return x, nil
}

type JsonConfigContainer struct {
	data map[string]interface{}
	sync.RWMutex
}

func (c *JsonConfigContainer) Bool(key string) (bool, error) {
	val := c.getData(key)
	if val != nil {
		if v, ok := val.(bool); ok {
			return v, nil
		}
		return false, errors.New("not bool value")
	}
	return false, errors.New("not exist key:" + key)
}

func (c *JsonConfigContainer) DefaultBool(key string, defaultval bool) bool {
	if v, err := c.Bool(key); err == nil {
		return v
	}
	return defaultval
}

func (c *JsonConfigContainer) Int(key string) (int, error) {
	val := c.getData(key)
	if val != nil {
		if v, ok := val.(float64); ok {
			return int(v), nil
		}
		return 0, errors.New("not int value")
	}
	return 0, errors.New("not exist key:" + key)
}

func (c *JsonConfigContainer) DefaultInt(key string, defaultval int) int {
	if v, err := c.Int(key); err == nil {
		return v
	}
	return defaultval
}

func (c *JsonConfigContainer) Int64(key string) (int64, error) {
	val := c.getData(key)
	if val != nil {
		if v, ok := val.(float64); ok {
			return int64(v), nil
		}
		return 0, errors.New("not int64 value")
	}
	return 0, errors.New("not exist key:" + key)
}

func (c *JsonConfigContainer) DefaultInt64(key string, defaultval int64) int64 {
	if v, err := c.Int64(key); err == nil {
		return v
	}
	return defaultval
}

func (c *JsonConfigContainer) Float(key string) (float64, error) {
	val := c.getData(key)
	if val != nil {
		if v, ok := val.(float64); ok {
			return v, nil
		}
		return 0.0, errors.New("not float64 value")
	}
	return 0.0, errors.New("not exist key:" + key)
}

func (c *JsonConfigContainer) DefaultFloat(key string, defaultval float64) float64 {
	if v, err := c.Float(key); err == nil {
		return v
	}
	return defaultval
}

func (c *JsonConfigContainer) String(key string) string {
	val := c.getData(key)
	if val != nil {
		if v, ok := val.(string); ok {
			return v
		}
	}
	return ""
}

func (c *JsonConfigContainer) DefaultString(key string, defaultval string) string {
	// TODO FIXME should not use "" to replace non existance
	if v := c.String(key); v != "" {
		return v
	}
	return defaultval
}

func (c *JsonConfigContainer) Strings(key string) []string {
	stringVal := c.String(key)
	if stringVal == "" {
		return []string{}
	}
	return strings.Split(c.String(key), ";")
}

func (c *JsonConfigContainer) DefaultStrings(key string, defaultval []string) []string {
	if v := c.Strings(key); len(v) > 0 {
		return v
	}
	return defaultval
}

func (c *JsonConfigContainer) GetSection(section string) (map[string]string, error) {
	if v, ok := c.data[section]; ok {
		return v.(map[string]string), nil
	}
	return nil, errors.New("nonexist section " + section)
}

func (c *JsonConfigContainer) SaveConfigFile(filename string) (err error) {
	// Write configuration file by filename.
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.MarshalIndent(c.data, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}

func (c *JsonConfigContainer) Set(key, val string) error {
	c.Lock()
	defer c.Unlock()
	c.data[key] = val
	return nil
}

func (c *JsonConfigContainer) DIY(key string) (v interface{}, err error) {
	val := c.getData(key)
	if val != nil {
		return val, nil
	}
	return nil, errors.New("not exist key")
}

func (c *JsonConfigContainer) getData(key string) interface{} {
	if len(key) == 0 {
		return nil
	}

	c.RLock()
	defer c.RUnlock()

	sectionKeys := strings.Split(key, "::")
	if len(sectionKeys) >= 2 {
		curValue, ok := c.data[sectionKeys[0]]
		if !ok {
			return nil
		}
		for _, key := range sectionKeys[1:] {
			if v, ok := curValue.(map[string]interface{}); ok {
				if curValue, ok = v[key]; !ok {
					return nil
				}
			}
		}
		return curValue
	}
	if v, ok := c.data[key]; ok {
		return v
	}
	return nil
}

func init() {
	Register("json", &JsonConfig{})
}
