package config

import (
	"fmt"
)

type ConfigContainer interface {
	Set(key, val string) error
	String(key string) string
	Strings(key string) []string
	Int(key string) (int, error)
	Int64(key string) (int64, error)
	Bool(key string) (bool, error)
	Float(key string) (float64, error)
	DefaultString(key string, defaultval string) string
	DefaultStrings(key string, defaultval []string) []string
	DefaultInt(key string, defaultval int) int
	DefaultInt64(key string, defaultval int64) int64
	DefaultBool(key string, defaultval bool) bool
	DefaultFloat(key string, defaultval float64) float64
	DIY(key string) (interface{}, error)
	GetSection(section string) (map[string]string, error)
	SaveConfigFile(filename string) error
}

type Config interface {
	Parse(key string) (ConfigContainer, error)
	ParseData(data []byte) (ConfigContainer, error)
}

var adapters = make(map[string]Config)

func Register(name string, adapter Config) {
	if adapter == nil {
		panic("config: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("config: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewConfig(adapterName, fileaname string) (ConfigContainer, error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("config: unknown adaptername %q (forgotten import?)", adapterName)
	}
	return adapter.Parse(fileaname)
}

func NewConfigData(adapterName string, data []byte) (ConfigContainer, error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("config: unknown adaptername %q (forgotten import?)", adapterName)
	}
	return adapter.ParseData(data)
}
