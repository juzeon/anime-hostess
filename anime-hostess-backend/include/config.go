package include

import (
	"gopkg.in/ini.v1"
	"log"
)

type configStruct struct {
	Port       int    `ini:"port"`
	Debug      bool   `ini:"debug"`
	VideoRoot  string `ini:"video-root"`
	VideoTypes string `ini:"video-types"`
}

var Config *configStruct

func init() {
	LoadConfig()
}

func LoadConfig() {
	// set default config
	Config = &configStruct{
		Port:       9777,
		Debug:      true,
		VideoRoot:  ".",
		VideoTypes: "mp4,flv",
	}
	configPath := "config.ini"
	if ConfigPath != "" {
		configPath = ConfigPath
	}
	cfg, err := ini.InsensitiveLoad(configPath)
	if err != nil {
		log.Println("Use default config")
	} else {
		err = cfg.MapTo(Config)
		if err != nil {
			panic(err)
		}
	}
}
