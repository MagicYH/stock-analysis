package tools

import "github.com/koding/multiconfig"

type ConfigMysql struct {
	Name     string `required:"true"`
	Host     string `required:"true"`
	Port     int    `required:"true"`
	Passwd   string `required:"true"`
	User     string `required:"true"`
	Database string `required:"true"`
}

type ConfigDb struct {
	Mysql []ConfigMysql
}

type GlobalConfig struct {
	Database ConfigDb `toml:"database"`
}

var _globalConfig *GlobalConfig

func init() {
	_globalConfig = new(GlobalConfig)
}

func LocdGlobalConfig(path string) {
	m := multiconfig.NewWithPath(path) // supports TOML, JSON and YAML
	err := m.Load(_globalConfig)       // Check for error
	if err != nil {
		panic(err)
	}
	m.MustLoad(_globalConfig) // Panic's if there is any error
}

func GetGlobalConfig() GlobalConfig {
	return *_globalConfig
}
