package include

import (
	"gopkg.in/ini.v1"
)

type configStruct struct {
	Port  int  `ini:"port"`
	Debug bool `ini:"debug"`
}

var Config *configStruct

func init() {
	LoadConfig()
}

func LoadConfig() {
	Config = &configStruct{}
	configPath := "config.ini"
	if ConfigPath != "" {
		configPath = ConfigPath
	}
	cfg, err := ini.InsensitiveLoad(configPath)
	if err != nil {
		panic(err)
	}
	err = cfg.MapTo(Config)
	if err != nil {
		panic(err)
	}
}
