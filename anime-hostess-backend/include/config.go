package include

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

type configStruct struct {
	Listen        string `ini:"listen"`
	Debug         bool   `ini:"debug"`
	VideoRoot     string `ini:"video-root"`
	VideoTypes    string `ini:"video-types"`
	AsGRPC        bool   `ini:"as-grpc"`
	GRPCServer    string `ini:"grpc-server"`
	RedisServer   string `ini:"redis-server"`
	RedisPassword string `ini:"redis-password"`
}

var Config *configStruct

func LoadConfig() {
	// set default config
	Config = &configStruct{
		Listen:        "0.0.0.0:9777",
		Debug:         true,
		VideoRoot:     ".",
		VideoTypes:    "mp4,flv,mov,mkv",
		AsGRPC:        false,
		GRPCServer:    "",
		RedisServer:   "127.0.0.1:6379",
		RedisPassword: "",
	}
	configPath := "config.ini"
	if *ConfigPath != "" {
		configPath = *ConfigPath
	}
	cfg, err := ini.InsensitiveLoad(configPath)
	if err != nil {
		log.Println("Use default config")
	} else {
		err = cfg.MapTo(Config)
		if err != nil {
			panic(err)
		}
		fmt.Printf("LoadConfig: %#v\n", Config)
	}
}
